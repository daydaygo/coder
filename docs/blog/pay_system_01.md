# coder| 支付系统0x01: 基础设施 & 初版架构

一周紧张的开发后, 支付系统的大致雏形已经出来的, 需要完善的地方还有很多(每天都是长长的 todolist), 这里先讲一讲基础设施和初版架构.

# 基础设施篇

## docker

既然是微服务设计, 当然跑不了 docker 的使用. 这里不赘述, 单从工具之美的角度聊一聊.

核心概念:

- repository 仓库: 镜像仓库
- image 镜像: 定义程序运行的环境
- container 容器: 实例化的镜像, 程序的运行环境, 通常一个 container 代表一个 service(服务)

**备注**: `container == service`, 之后不再重申这个概念

没错, 核心概念就这么多, 下面再来看看 **周边**, 我喜欢简单归纳到工具里面:

- docker hub, 类似 github, 你也可以类比 repository/image 的概念和你的 git 项目来理解, 我一般使用国内镜像源(aliyun)替代
- dockerfile: 一系列指令来构建一个 image
- dockerd: docker server 端程序
- docker: docker client 端程序, 用来和 docker server 通信
- docker-compose: container 编排程序, 适合服务较少场景(比如开发环境)下使用

好了, 差不多知道这么多概念就可以动手了:

- docker 安装: 直接使用 [阿里云 - 容器服务 - 镜像加速器](https://cr.console.aliyun.com/#/accelerator), 里面有详细的文档
- docker-compose 安装: 直接使用 [daocloud 软件中心](https://download.daocloud.io/Docker_Mirror/Docker_Compose); 我通常会配置 `alias doc='docker-compose'` 来加速.

**备注**: 使用 `docker-for-window` 一站式解决 docker 安装, 之后用好 docker, 完全可以把 window 变成很好的开发平台.

好了, 所了这么多, 直接实战, 你就会发现这个到底有多简单了, 先来看 docker:

```bash
docker help # 这个不用多说吧
docker info # 查看 docker 信息, 一般用来确认 docker 安装启动是否正常

docker pull image:tag # 熟悉 git 对 tag 就不陌生了吧
docker images # 参看本地镜像
docker rmi $(docker images -f "dangling=true" -q) # 高级点的应用, 删除所有镜像

docker build -t image:tag xxxDockerfile # 根据 dockerfile 构建 image

docker ps # 查看运行中的容易, -a 查看所有
docker rm `docker ps -a -q` # 删除所有容器, -q 只显示 container id

# run 运行容器, 可以说是最复杂的命令了, 这里直说常用的
docker run --name xxx-app -d -p 8080:80 image:tag # 运行容器 + 命名(--name) + daemon(-d) 运行 + 绑定端口(-p)
docker run -ti --rm centos:7.4 bash # 我常用这个来开一个 linux 环境, -ti 如同终端一样操作, --rm 退出 container 后自动删除
```

看到这里, 你可能还没有感受到 docker 的强大, 但是请你注意这里:

`docker run -ti --rm centos:7.4 bash`, 你就执行一下这条命令, 我就有了一个 centos 7.4 的环境了随便你折腾了, 对的, 就一条命令!

**备注**: 实际使用中 docker 命令其实也不怎么用, 主要还是使用 doc(`alias doc='docker-compose'` 后面不再提示)

然后再来看 dockerfile, 首先为什么会有 dockerfile 呢? 我的理解是:

> 将软件的运行环境 文本化/指令化, 变成程序源码这种容易分发的方式.

dockerfile 作为一系列指令(用来构建运行环境)的集合, 实在是简单的无须多讲, 这里就简单列举一下这次支付系统使用到的例子:

```dockerfile
# mysql.Dockerfile
FROM mysql:8 # 官方镜像
LABEL maintainer="1252409767@qq.com" # 镜像维护者
ADD my.cnf /etc/mysql/conf.d/my.cnf # 添加配置文件
RUN chmod -R 644 /etc/mysql/conf.d/*
CMD ["mysqld"] # 运行 mysqld 服务
EXPOSE 3306 # 开启 3306 端口

# redis.Dockerfile
FROM redis:4.0-alpine
LABEL maintainer="1252409767@qq.com"
COPY redis.conf /etc/redis/redis.conf
CMD redis-server /etc/redis/redis.conf --appendonly yes
```

这样, 我系统里就有了随时可用的 mysql 和 redis 服务了.

这里说一下 dockerfile 的相关的学习方法:

- 大致看一下 dockerfile 指令, 一些容易混淆的指令注意一下(这个教程里都会突出的)
- 使用官方镜像, **务必**大致浏览官方提供的 `readme`, 常见用法基本都可以在这里找到
- **务必** 看一下官方镜像的 dockerfile, 这样可以帮你有效解决 环境依赖(比如 php 镜像使用 pecl 安装插件需要安装哪些软件)/环境变量(有哪些变量, 怎么使用) 相关的问题

好了, 到 doc 登场了, 按照上面的 dockerfile, 我们建好了一个又一个服务, 怎么愉快的玩耍起来呢? 直接看 「栗子」:

- 首先, 你要有一个 `docker-compose.yml` 的配置文件, 基于 yaml 语法(和 json 一样简单, 出现原因大概是 xml 实在是太长了)

```yml
version '3'
services: # 服务编排
    nginx:
        build: # 需要 dockerfile 来构建镜像
            context: ./server
            dockerfile: nginx.Dockerfile
        volumes:
            - ../:/var/www # 挂载项目目录
            - ./logs/nginx/:/var/log/nginx # 挂载 nginx 日志
        links:
            - fpm # 需要使用 fpm service
        ports: # 端口绑定
           - "80:80"
           - "443:443"
    fpm:
        build:
            context: ./server
            dockerfile: fpm.Dockerfile
        volumes:
            - ../:/var/www
            - ./logs/fpm/:/var/log/php7
        links:
            - mysql
            - redis
    mysql:
        build:
            context: ./server
            dockerfile: mysql.Dockerfile
        volumes:
            - ./data/mysql:/var/lib/mysql # 挂载数据目录, 实现数据持久化
        ports:
            - "3306:3306"
        environment: # 设置环境变量
            # MYSQL_DATABASE: test
            # MYSQL_USER: test
            # MYSQL_PASSWORD: test
            MYSQL_ROOT_PASSWORD: root
    redis:
        build:
            context: ./server
            dockerfile: redis.Dockerfile
        volumes:
            - ./data/redis:/data
            - ./logs/redis:/var/log/redis
```

有了 `docker-compose.yml` 后, 运行起来就好了:

```bash
doc up -d nginx # daemon 状态运行 nginx, 因为 nginx 需要 fpm, fpm 需要 mysql redis, 所有 4 个服务都会运行起来
doc up -d --build nginx # 重新构建 image 并运行

doc stop # stop service
doc logs xxxService # 查看 service 日志, 在 service 启动失败时会有报错信息, 如果报错信息不足以帮助解决问题, 就可以使用这个查看一下日志来获取更多信息
```

到这里你可能会说怎么弄个 lnmp 环境怎么这么麻烦, 七七八八的折腾了这么久, 工具一个接着一个, 但是请关注一下结果, `doc up -d nginx`, 只要执行一下这个, 你的 lnmp 好了, 而且你想用什么版本就用什么版本, 更主要的是:

> 快, 这才是老司机的本质.

作为一个使用 docker 2年+ 的老司机, 继续安利一波, docker 当做工具来使用真的非常非常简单, 大致了解完基础后, 在去看看几个开源项目怎么使用的就可以用起来了, 比如下面这几个:

- [laradock](https://github.com/laradock/laradock): 这里顺便安利一下 laravel 框架, 不一定要用它, 但是从他丰富的 **周边** 真的可以学到很多
- [swoole-docker](https://github.com/tmtbe/swoole-docker): swoole distribution 的 docker 运行环境, 作为 php 程序员, 去了解一下 swoole 吧, 你会发现很多很有意思的人, 而不是 「php是最好的语言」 或者 「php 真 low」

> 想深入的话推荐这本书: [<容器与容器云ed2>](http://www.ituring.com.cn/book/1899)

### alpine linux

这里再补充一点 alpine linux 的知识. 各种官方镜像, 实际也是在 linux 系统上, 通过安装软件的方式构建而成, 通常我们把这个 linux 系统叫做 **基础镜像**. 现在比较流行的是 2 个 `Debian` 和 `alpine`. `Debian` 我不多说了, 同系的 `Ubuntu` 以及 redhat 系的 `centos` 都很常见, 这里重点推荐一下 `alpine` linux.

- 首先是镜像容量对比: 几M vs 接近百M
- 然后是 alpine 的 **工具** 属性: 从 busybox 这个继承了上百个常用 linux 命令的工具集升级而来. 在我眼里 alpine = linux kernel + tool(包管理其实也是为了增减工具)

使用这么久下来, 通过笔记([wiki - os - linux 发行版](http://wiki.dayddaygo.top))来看, 实在是不要太简单:

```
# apk 包管理
echo -e "http://mirrors.aliyun.com/alpine/v3.4/main\nhttp://mirrors.aliyun.com/alpine/v3.4/community" > /etc/apk/repositories
apk add xxx
apk search -v xxx
apk info -a xxx
apk info

# network
apk add iproute2 # ss vs netstat
ss -ptl
```

## git

关于 git, 能说的太多了, 比如 [<pro git>](https://git-scm.com/book/en/v2) 这本公版书. 这里简单提四个点.

1. git 是 Linus 发明的. Linus 也是是 linux 的发明(就是这么牛逼). 而 git 之所以叫 git, **是因为我就是一个自私混点, 做的东西都喜欢用自己来命名, linux 是, git 也是**. 另一个经典语录是一次公开场合下 **I am your god**.

2. 当然绕不开 git 和 svn 的对比. 本质是 **分布式系统 vs 集中式系统**, 更简单一点是, 你没有一个远程仓库, 你依旧可以可以开心的和 git 玩耍, 远程仓库帮助你分发和协作. 总结一下: **珍爱生命, 原理 svn**. (发邮件出文件更新列表来上线的日子, 我在三年前刚工作的时候也经历过, 没想到现在还能碰到, **活久见**)

3. git flow 工作流, 这种方式比较复杂, 但是非常推荐你去了解一下, 当你见识到了 **终极复杂** 的方式, 按需精简成需要的工作流就简单了. 这里给几个我使用下来的经验: master(主) 分支是 **产品** 分支, 要保证一直可用, 所以 merge(合并) 到 master 分支需要具有一定能力(权限); 至少保留 dev 分支作为测试使用; 开发时间较长的需求务必使用 feature(特性/功能) 分支; 上线使用 tag(标签)

4. 当然还是 git 常用功能小结了:

```
git status # 查看发生改动的文件列表
gitk # 图形化工具查看文件改动; 安装 git 后自带的图形化界面, 可以直接命令行调起, 我本人强烈推荐, 我不安装其他 git 图形化工具
# 使用 gitk 确认文件修改无误 -> 一定要自己 review 一下代码, 你自己都不愿意 code review, 就不要憧憬那些有 code review 的牛逼团队了
# 在 review 之前, 一定要自测一遍自己的代码, 不要没事就给测试妹子刷存在感
git checkout . # 取消所有修改
git checkout xxxFile # 取消某个文件的修改
git add -A # 添加所有修改
git commit -m 'commit message' # 提交修改

git push # 提价代码到远程分支, 这里注意, 远程仓库一般使用 origin 作为名称, 这里省略了
git pull # 从远程仓库获取最新代码

git fetch # 从远程获取 branch/tag 信息
git branch # 查看本地分支, -r 查看远程分支, -a 查看本地+远程
git checkout dev # 切换到 dev 分支, -b 新建并切换到分支
git merge dev # 将 dev 分支 merge 到当前分支

git tag # 查看 tag
git tag -a v1.0 -m 'tag message' # 添加 tag
git push --tags # 同时推送 tag 到远程仓库

git stash # 暂存本地修改
# do something else
git stash pop # 恢复暂存的修改

git remote set-url origin git://new.url.here # 修改远程仓库地址, 我经常会用到这个
git help xxxCmd # 查看帮助文件, 这里会用浏览器打开 html 文件, 非常方便
git config --global user.name "daydaygo" # git 设置, --global 全局生效

# pull request & code review
# 1. fork(复制) 原项目到自己的仓库
git clone origin-git-url # 克隆自己的仓库
// changge - commit
# 2. 在 gitub/gitlab 提交 pr(pull request), 原项目上就可以看到, 之后提交后会自动提 pr
git remote add upstream upstream-git-url # 添加原仓库地址; 取名为 upstream 和 origin 一样，习惯而已
git pull upstream master # 从原仓库获取更新
```

还有一个重要概念是 **冲突**, 其实很简单, 你的版本和其他人的版本**无法合并**, 比如你写了 `a=1`, 别人写了 `a=2`, 所以, 和当事人确认一下代码, 改一下就好.

**周边** 还要很多, 比如 `.gitignore / .gitkeep` 文件, 慢慢 `get` 就好.

> [<只是为了好玩 - Linus自传>](https://book.douban.com/subject/25930025/), 强烈推荐这本书, 也许可以境界提升.

## gitlab

讲完 git, 就会发现, 少一个 **远程仓库**, 对, gitlab 就是干这的. 当然, 他的功能远不止于此:

- web 操作界面
- 仓库的权限管理: 用户 / 用户组 /  https & ssh & deploy key
- pull request & code review
- CI 持续集成

因为使用时间并不长(还有一个原因是规模并不大), 探索目前只到这个阶段, 这里晒一下 CI 的成果:

![](http://7xksek.com1.z0.glb.clouddn.com/gitlab-ci.PNG)

过程有点小坎坷, 不过还是套用之前 blog 里提到的观点: [唯手生尔](http://www.jianshu.com/p/8991cb09fe0d), 这里简单叙述一下:

- gitlab ci 的工作原理: 安装 `gitlab-runner` 服务并执行 `gitlab-runner register` 和 gitlab 上的项目关联
- 安装 `gitlab-runner` 按照官网的文档来就好了, 测试了 `shell` 版和 `docker` 版, 最终还是选用了 docker 版, 并且使用了 阿里云镜像仓库
- master分支/tag 变更触发 CI, gitlab 会按照 `.gitlab-ci.yml` 中定义的任务创建一个 job, job 分发给关联的 `gitlab-runner`, runner 会先跟新项目, 然后运行 job.

使用 docker 版的 `gitlab-runner` 也非常简单, 还是使用 doc 进行编排:

```yaml
services:
    ci:
        image: gitlab/gitlab-runner:alpine
        volumes:
            - ./data/gitlab-runner/config:/etc/gitlab-runner
            - /var/run/docker.sock:/var/run/docker.sock:Z
```

然后执行:

```bash
doc up -d ci # 启动 gitlab-runner 服务
doc exec ci gitlab-runner register # 执行注册管理, 可以多次执行, 定义多个 runner

```

`.gitlab-ci.yml` 定义非常简单:

```yaml
before_script:
  - cd pay-support

phpunit: # 类似这样定义一个又一个任务就好了
  script:
    - phpunit --coverage-text --colors=never
```

基础设施到这里先告一段落, 还有 监控报警/日志收集分析 等, 下次再做分解.

# 初版架构篇

其实在上一篇 [支付系统0X00: 支付系统预研](http://www.jianshu.com/p/8991cb09fe0d) 就已经将架构方面的内容侃得七七八八了, 多是在实施的过程中继续参考 [凤凰牌老熊 - 现代支付系统设计](http://blog.lixf.cn/) 中各章节, 在细节上继续下功夫.

> 一周时间毕竟有限, 还需要细细打磨

## 初版概况
> 流程图 / 项目结构图: https://www.processon.com/view/link/5a0baa75e4b049e7f4fd90f4

- 核心业务流程: 用户 -> 商户 -> 支付接口(验参验签) -> 支付路由 -> 支付方式 -> 收单 -> 支付成功
- 架构分层: 产品服务 核心系统(支付核心 + 支付服务) 支撑系统

系统边界:

- 支付网关: api路由 -> 聚合支付; 接口安全
- 支付产品: 风控 支付路由 参数校验 支付流程(交易记录, 支付渠道, 同步/异步通知)
- 支付渠道: 和支付渠道对接, 按照支付产品预定格式化统一化结果输出

项目结构:

- pay-gateway: 支付网关, 使用 api 和商户交互
- pay-product: 封装支付产品
- pay-channel: 支付渠道 sdk, 和支付渠道交互
- pay-lib: 依赖库
- pay-support: 支撑系统, 目前只包含 api doc 和 phpunit

重要实现细节:

- 商户模式: 通过支持商户模式, 让支付系统具有平台属性, 方便业务拓展
- 请求参数设计: 公共参数 + 业务参数 + 业务拓展参数, 保持接口的灵活
- 请求全异步化机制: 接到请求后立刻返回请求是否成功受理 + 异步处理耗时任务 + 异步通知商户

## 基于 swoole 的任务队列
> 感谢开源作品: [swoole-jobs](https://github.com/kcloze/swoole-jobs)

从上面的实现细节 **请求全异步化机制** 可知, 支付系统引入了任务队列, 而且相当重要. 原来项目基于 `yii2 console + redis queue + yii2 mutex` 实现了一套任务机制, 最终基于 crontab 运行. 核心实现细节如下:

```php
// mutex 使用
$lock = \yii::$app->mutex->acquire( $lock_name );
\register_shutdown_function(function() use($lock_name) {
    return \yii::$app->mutex->release( $lock_name );
});

// 基于 redis list 实现的任务队列
public static function push($params = []) {
    $redis = self::getRedis($params);
    return $redis->executeCommand('RPUSH', $params);
}
public static function pop($params = []) {
    $redis = self::getRedis($params);
    return $redis->executeCommand('LPOP', $params);
}

// 执行任务
$lock = CommonHelper::lock(); // 封装 mutex 的使用
if (!$lock) {
    return self::EXIT_CODE_ERROR;
}
$tag_file = '/tmp/xxx.tag'.$id; // yii2 console 应用 bug, 检测到 tag 后关闭脚本
$content = RedisQueue::pop([RedisQueue::LIST_CHANNEL_FEEDBACK]); # 执行任务
```

当然, 这样的结构完全可以胜任 **任务队列** 这一定义, 并且经历过线上检验, 比如说添加 tag 来关闭脚本 这样的经验积累. 不过考虑再三, 还是决定引入 swoole-jobs 来构建新的任务队列, 理由如下:

- 原来任务系统耦合在原有项目中, 需要做一定程度拆分重构才能单独成为任务队列
- 服务实现依赖 crontab, 就会受到 crontab 的局限, 比如 1 分钟运行一次, 需要多进程处理需要自己实现

而 swoole-jobs 在设计上的优势很明显:

- 基于 swoole process 实现服务化和进程管理, 修改配置项就可以改变进程数量, 同时开几百个进程没问题 (感谢 rango 大大的及时响应, 当时 @ 的时候还挺虚的)
- 抽象 Queue 实现, 可以基于 redis / rabbitmq 等不同驱动, 灵活拓展更多队列功能, 比如 topic 订阅, 任务优先级
- Job 类抽离, 可以有更多 Job 运行的方式: 类似 yii2 的 console 应用, 类静态方法调用, 类动态方法调用, 闭包

而所有这些, 不到 10 个文件 500 行代码, 推荐大家阅读源码.

> 欢迎 star / fork / 提交 pr, 我提交的 pr 已被作者采纳了哦.

## 支撑系统: api doc + phpunit

我以前也是那种啪啪啪疯狂输出, 然后提交完代码就 **万事大吉** 类型的, 现在再想来, 好听点叫 **too young**, 其实就是 **low**, 好好遇到了各种大大, 而不至于过早夭折.

> 用电竞圈的一个段子: 躲在大哥胯下疯狂输出.

谈到 **大哥**, 就不得不感谢一下一路走来遇到的各位腾讯技术出身的大大 -- 两任 CTO, 吃着中药熬着夜的刀哥, 传授了N多经验的红旗; 开发出 swoole 的 rango, 让我坚定 php 可以走的更远; 2大php开源项目的核心开发者朱新宇, `勇敢的少年啊快去创造奇迹`.

> 回到正题, 软甲开始是一个很难杜绝 bug 但是尽量要求 **正确** 并且 **持续正确** 的工作.

所以, 下面几点你当做最佳实践也好, 技能提升也好, 或者技术团队管理也好, 但是请你记住, 知道有这些方式方法经验教训的存在:

- 上面 git 部分的内容, 这里再重审一遍, 自己的写的代码, 一定要自测通过, 如果被发现没有自测并且存在明显(或者说低级)错误(你还好意思称之为 bug 么), 会受到集体鄙视;
- 提交代码前, 务必自己 review 一下, 那些调试语句没删掉, 数据库修改忘添加等, 都可以通过这道工序解决掉
- 接口开发一定要提供接口文档(api doc)

这里推荐一个 api 文档工具, [swagger ui](https://swagger.io/swagger-ui/), 使用非常简单, 只用编写 yaml 文件(这个多次碰到了, 赶紧 get 这个技能吧)就行. swagger ui 是静态页面, 放到 web server 下运行即可.

> 支付系统 api 文档: http://hp-api.dayday.tech (有 ip 限制, 公司内才可以访问)
> 也可以直接看官方的 demo: http://petstore.swagger.io/ (实际使用并不会用到这么多特性, 会非常简单)

- 测试先行, 或者说测试驱动开发, 是有效发现并减少 bug 的方法, 而且 phpunit 真的很简单易用, 上面的 gitlab ci, 就是集成了 phpunit 测试

phpunit 快速上手:

1. 全局可执行文件: `composer global require phpunit/phpunit`
2. 项目中添加依赖: `composer require phpunit/phpunit --dev`
3. 添加 `phpunit.xml` 配置文件: 好一个现有文件作为模板该改就好
4. 编写测试用例: 一般添加 unit test 作为功能测试(类, 函数 等), feature test 作为 api 接口测试

## 写在最后

行文至此, 颇有种 **酣畅淋漓** 之感.

> 也许这就是对自己所做的事产生的自豪感吧, 谦虚点是自鸣得意, 吹起牛皮来就是独孤求败了.
