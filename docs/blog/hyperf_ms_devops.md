# hyperf| 微服务之旅: devops

- [hyperf| 微服务之旅: devops](https://www.jianshu.com/p/10796001bf39)

本篇属于 `PHP 微服务之旅` 系列, 此系列会持续更新, 敬请期待.

如果说 微服务应用 对比 `传统巨石应用` 最大的难点在于 `基础设施` 的复杂度上, 那么本篇就是要抛砖引玉, 提供一个完整的 devops 实例, 证明一个简单的事实:

> 使用 devops 后, 微服务的开发也能有 传统巨石应用 的简单高效.

先提前预告一下效果, 对于一个参与到项目的开发而言:

- 本地只需要安装 git/docker, 容器启动后, 项目所需的开发环境全都有
- 开发者除了正常编写代码, 项目的 CI/CD 会在代码提交后自动触发, 开发者完全不用分心

## devops 流程一览

- github 私有库 / gitlab 私有部署 
- 绑定 Travis CI 持续集成
- 绑定 aliyun 容器镜像服务(cr), 设置自动构建规则
- 绑定 aliyun 容器服务swarm(cs), 设置镜像构建后自动部署
- 绑定 aliyun 日志服务(sls), 设置容器日志自动投递到日志服务

下面详细讲解下各部分

## git

类似对比 git 与 svn 这种事, 不会再出现在我的 blog 中, 有种去翻历史尘埃的感觉, 大概比较适合 `技术史` 一类的文章.

记一些 git 常用的功能吧:

- `git add/commit/stash/push` 这些基础命令不熟悉的话, 估计对应的 `工作去/暂存区/提交区/远程分支` 这些基础概念也不熟, 这是基础, 自行百度补
- `git branch/checkout/merge` 这些分支相关命令, 也要很熟悉
- `git tag` 只是在 commit 上做了一个标记, 建议使用统一的格式
- git clone 有 2 种方式, ssh 需要配置公钥, https 需要账号密码认证, 可以使用 `git config --global credential.helper "store"` 记住账号密码
- git 是分布式的, 本地就是一个完整的仓库(repository), 通常会添加一个远程仓库来完成大家的协作
    - 通常这个远程地址命名为 origin, 可以通过 `git remote add origin url-xxx` 来添加
    - 当然远程分支是可以修改的 `git remote set-url origin url-xxx`
    - 当然远程分支可以有多个
        - 比如 github 提 PR, 需要 fork 原项目, fork 的出来的项目除了有远程分支 `origin` 外, 还会多一个名为 `upstream` 的远程分支对应原项目
        - 当然也可以添加更多, 比如 `git remote set-url gitee url-xxx`, 添加一个国内的 git 库, 速度更快
- 推荐使用的 git 客户端: [github desktop](https://desktop.github.com/)
    - `命令行才是最好用的`, 但是有一个场景除外, `git diff`, 强调一点, `自己写的代码, 自己一定要 review`
    - github 自家的, 所以提 PR 使用快捷键 `cmd-r` 就可以了, 贼爽, 推荐给坚持开源的小伙伴
- 开发过程中通常使用 MR(merge request), 尽量做好 自测/code-review/CI, 从而尽量不发生 `revert`, 避免不了也不要慌, `gitlab revert MR` 体验贼好
- 添加 alias, 为 git 命令行操作加速

git alias 推荐:

```fish
alias gc='git checkout'
alias gb='git branch'
alias gt='git tag'
alias gs='git status'
alias gd='git diff'
alias gl="git log --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
alias ga='git add -A'
alias gm='git merge'
alias gp='git pull'
alias gcl='git clone --progress'
# alias gp='git pull --rebase'
alias gpp='git pull;git push'
alias gco='git commit -am'
alias gca='git commit -am "update";git pull;git push'

# 打 tag 也可以很轻松
function gtt
    set tag_name release-v(date +"%Y%m%d%H%M")
    echo $tag_name | pbcopy
    pwd
    git pull
    git push
    git tag $tag_name
    git push --tags
end

# 发版到测试环境
function gdt
    set b_ori (git symbolic-ref --short -q HEAD)
    git checkout dev
    git merge $b_ori -m 'merge'
    git pull
    git push
    git push gh dev
    git checkout -
end
```

我使用的 [fishshell](http://www.fishshell.com/), 再次疯狂安利一波, 几乎零配置, 比 zsh 更高效.

可供参考的教程:

- [GitHub 入门与实践](http://www.ituring.com.cn/book/1581)
- [git 常用操作脑图](http://blog.csdn.net/kehyuanyu/article/details/41278797)
- [git 简明指南](http://www.bootcss.com/p/git-guide/)
- [git tutorial](https://try.github.io)
- [廖雪峰的git教程](http://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000)
- [git 协作流程](http://kb.cnblogs.com/page/535581/)
- [版本控制入门 – 搬进 Github](http://www.imooc.com/learn/390)
- [Git实战_201703014](https://segmentfault.com/a/1190000008579494)
- [一张图看明白Git的四个区五种状态](http://imtuzi.com/post/git-four-areas-five-states.html)
- [回滚错误合并](https://segmentfault.com/a/1190000015792394)
- [git hook 代码检测](https://www.longlong.asia/2015/08/08/git-pre-receive-hook.html)

git 的部分稍显啰嗦, 后面的 docker 部分也会有点, 毕竟 git/docker 都是 `基础功`, 这是没法回避的部分, 后面的实施部分要简单不少, 几张截图就能搞定

## docker

一样的风格, 不再赘述 docker 好在哪, 以及 docker 的一些基础概念, 默认你已经了解以下基础知识:

- docker 的基本概念: 镜像 容器 网络(网络基础/端口映射) 文件挂载 ...
- `docker` `Dockerfile` `docker-compose` 的基础知识和基本用法

当然, 如果团队有 `老司机`, 你会发现使用 docker 是一件如此轻松的事, 几乎只是几个命令的事儿.

假设 git 仓库中提供了 `docker-compose.yml` 文件:

```yaml
version: '3.1'
services:
    s1: # 定义服务名称
        image: registry.cn-shanghai.aliyuncs.com/daydaygo/dev:hyperf
        volumes: # 挂载代码路径到容器中
            - ../:/data
        ports: # 端口绑定, 如果需要本地个可以访问容器中的服务
            - "80:9501"
        environment: # 设置环境变量
            APP_ENV: dev
        links: # 此服务依赖的其他服务, 按需配置
            - redis
            - mysql
            - rabbitmq       
        tty: true # 设置允许进入容器 tty
```

那么, 其他开发者日常只需要:

```bash
# 定义 alias, 后面省力不少
alias doc='docker-compose

doc up -d s1 # 启动服务(就是 docker 容器), 即由上面的 docker-compose.yml 文件定义的服务
dos stop/restart/rm s1 # 服务管理
doc exec s1 fish # 进入到容器tty, 使用 fishshell(是的, 再次安利 fishshell)
```

看到这里, 希望你注意到一个更加吃惊的事实:

> 本地除了 git/docker, 其他什么环境都不需要配置, 容器启动起来后, 项目所需的环境都有了!!!

## 绑定 aliyun 容器镜像服务(cr)

从这一部分开始, 后面的内容都会比较轻松, 几乎只是几张截图.

- 使用 aliyun 容器镜像服务, 划重点: `免费`.

![](http://qiniu.dayday.tech/20190602214017.png)

- 绑定代码源, 推荐 github/gitlab:

![](http://qiniu.dayday.tech/20190602214337.png)

- 配置自动构建规则

![](http://qiniu.dayday.tech/20190602214538.png)

除了默认的 tag 触发的自动构建(这也是上面配置 `git tag` alias 的原因之一), 我通常会配置 master 分支对应镜像 tag  latest, dev 分支对应镜像 tag test, 分别部署到 测试环境/线上环境. 当然, 需要增加其他环境, 再增加一个构建规则就好了.

而我们所需要的代码, 即 Dockerfile, 简单得不好意思贴进来:

- 项目镜像构建所需的 Dockerfile

```Dockerfile
# 获取基础镜像
FROM registry.cn-shanghai.aliyuncs.com/daydaygo/dev:hyperf
LABEL maintainer="daydaygo <daydaygo@vip.qq.com>" version="1.0"

# 添加项目源码
COPY . /data
# 设置环境相关配置
COPY .env.prod /data/.env
# 启动服务
ENTRYPOINT ["php", "/data/bin/hyperf.php", "start"]
```

## 绑定 aliyun 容器服务swarm(cs)

容器服务其实也是免费的, 收费的部分其实是集群对应的 ecs. 这里啰嗦一些 `集群` 的概念

- 集群的逻辑概念: 一个集群中可以有多个`应用`, 每个`应用`可能由多个`服务`协同, 比如上面的 docker-compose.yml 中的 s1 服务, 使用 `--links` 依赖了 `mysql/redis/rabbitmq` 3 个服务
- 集群的物理概念: 一个集群中可以有多台 ecs, 提供相应的 cpu/内存 等计算资源, 这些 ecs 被称之为 `节点`
- 集群其他资源: 网络 文件挂载 配置项 ...

使用也很简单, 新建一个集群, 配置好 ecs, 开发测试的话推荐使用 `按量付费`, 最好至少开 2 台来体验.

- 好了, 配置我们的第一个应用

![](http://qiniu.dayday.tech/20190602220924.png)

- 使用配置文件创建

```yml
hyperf-ms: # 定义服务名称
  # 使用容器镜像服务中的镜像
  image: 'registry-vpc.cn-shanghai.aliyuncs.com/daydaygo/hyperf-ms:latest'
  mem_limit: 0
  kernel_memory: 0
  memswap_reservation: 0
  # 容器自动重启
  restart: always
  shm_size: 0
  expose:
    - 9501
  links:
    - rabbitmq
  memswap_limit: 0
  # aliyun 容器服务提供的特殊功能标签
  labels:
    # routing: 简单 http/https 路由, 配置服务的 端口号->域名
    aliyun.routing.port_9501: ms.dayday.tech
    # scale: 应用数量
    aliyun.scale: '1'
    # log: 关联 aliyun 日志服务, 这里投递容器的 stdout 到日志服务中
    aliyun.log_store_stdout: stdout
    aliyun.log_ttl_stdout: 30
    aliyun.log.timestamp: false
    aliyun.depends: rabbitmq
rabbitmq:
    image: rabbitmq:3.7.8-management-alpine
    hostname: myrabbitmq
    ports:
        - "5672:5672" # mq
        - "15672:15672" # admi
```

- 配置触发器, 镜像更新后自动部署

![](http://qiniu.dayday.tech/20190602221639.png)

- 也可以在容器镜像服务中配置

![](http://qiniu.dayday.tech/20190602221822.png)

## 关联 aliyun 日志服务

关于日志服务, 我之前也啰嗦了不少, 有兴趣可以去看看之前的 blog:

- [devops| 日志服务实践](https://www.jianshu.com/p/9dae3ba679e6)
- [go| go并发实战: 搭配 influxdb + grafana 高性能实时日志监控系统](https://www.jianshu.com/p/f4d2b2ebafea)

上面配置好 容器服务投递日志到 日志服务 后, 接下来我们还有几步工作要做:

- 自动创建好的 日志 project, 先修改一下备注

![](http://qiniu.dayday.tech/20190602222702.png)

- 对应的 logstore

![](http://qiniu.dayday.tech/20190602222835.png)

- 修改 logtail 配置, 推荐使用 json 解析, 几乎零配置

![](http://qiniu.dayday.tech/20190602223051.png)

- json 解析, 需要日志输出为 json 格式, 这个也简单, 以 PHPer 使用的 `monolog` 为例, 只用配置一下 Formatter 就好了:

```php
$app_env = env('APP_ENV', 'dev');
if ($app_env == 'dev') {
    $formatter = [
        'class' => \Monolog\Formatter\LineFormatter::class,
        'constructor' => [
            'format' => "||%datetime%||%channel%||%level_name%||%message%||%context%||%extra%\n",
            'allowInlineLineBreaks' => true,
            'includeStacktraces' => true,
        ],
    ];
} else {
    $formatter = [
        // 看这里
        'class' => \Monolog\Formatter\JsonFormatter::class,
        'constructor' => [],
    ];
}

return [
    'default' => [
        'handler' => [
            'class' => \Monolog\Handler\StreamHandler::class,
            'constructor' => [
                'stream' => 'php://stdout',
                'level' => \Monolog\Logger::INFO,
            ],
        ],
        'formatter' => $formatter,
    ],
];
```

- 配置好索引
- 日志服务更多配置
  - 保存快查
  - 配置仪表盘 订阅仪表盘的日报 
  - 配置告警 推荐告警到钉钉

> 一言以蔽之: 日志服务真的是个`宝藏男孩`呀

## 绑定 Travis CI

> CI 的根本在于: 你得有自动化测试才行

PHPer 推荐使用 phpunit, 尤其是微服务化了以后, `服务通常以接口的形式提供`, 单测简直不要太好用

说回正题, 来看 Travis CI, 项目中添加一个 `.travis.yml` 的事儿:

```yml
branches:
  only:
  - master
  - dev

language: php

services:
  - docker 

before_install:
  - docker login registry.cn-shanghai.aliyuncs.com -u xxx -p xxx
  - docker build -t ci .
  - docker run -d --name ci ci
  - sleep 5

script:
  - docker exec ci /bin/sh -c "cd /data; phpunit"

notifications:
  email:
    - daydaygo@vip.qq.com
  webhooks:
    - https://oapi.dingtalk.com/robot/send?access_token=xxx
```

看配置项估计也能看明白: 
- 配置分支触发
- 根据项目下的 Dockerfile 构建镜像
- 使用镜像运行容器
- 执行容器中的 phpunit
- 通过 email/钉钉 通知 CI 构建信息

## 写在最后

内容看起来很多, 对参与开发的人而言, 其实很省心:

> 本地除了安装 git/docker 外, 并不需要其他环境; 除了写代码外, 不需要分心 CI/CD

最后再安利一下使用到的一些工具:

- git / github / [github desktop](https://desktop.github.com/) / gitlab
- docker / docker-compose
- [fishshell](http://www.fishshell.com/)
- aliyun: 容器镜像服务(cr) 容器服务swarm(cs) 日志服务(sls)
- 钉钉: 使用钉钉机器人来搜集消息, 当做消息box使用, 上面的所有流程, 都可以配置消息到钉钉