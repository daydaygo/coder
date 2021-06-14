# TS| docker 入门

> slide: https://c.dayday.tech/landslide/TS20171222.html

团队内很多同学对 docker 感兴趣, 于是准备了这期分享, 希望可以帮助大家快速入门. 主要包含以下内容:

- why docker
- docker 基础概念
- 入门实战

## why docker

先讲讲我自己的经历. 作为一个典型的工具控, 这个可以通过 [我的 wiki](http://wiki.dayday.tech) 看出一二, 比如 tmux / fish 等工具. 对环境的折腾几乎是无底洞般消耗着, 而且在以前经常以重装系统告终. 一方面是自己重装系统实在确实 6, 大学里还找过 「计导」 的老师尝试自己讲一堂相关的课, 一方面是 **太年轻**, 基础不够好, 解决问题的方式方法也不在路上, 经常导致问题 **解决不了**. 分享 2 句话:

> 在手里只有一把锤子的人眼里,世界就是一个钉子. -- <穷查理宝典>

> 工匠应该专注于作品的创意，不应该浪费精力，没限制地在折腾自己的工具.

简单记录一下自己开发环境的变迁史:

- v4 2017年

上面的方式其实推荐一个机器只部署一个项目，每个项目还是分开的，但是我只有一台云主机，却有很多项目，所以重新架构了一次：

独立一个通用的 docker 环境，目前包括 nginx（lnmp）+ cli（php-cli、npm、git 等）
每个项目依旧在不同的 git 仓库中维护
一台腾讯云上面部署所有项目

- v3 2016年

尝试到了 docker-compose 的甜头，开始信奉 dockerfile + 项目源码 的方式

- v2 2015年

项目全部进 git
开发环境：vagrant 大量使用（当时 laravel 的 homestead），开始接触 docker
部署环境：申请到一年的腾讯云，裸搭、lnmp 一键环境

- v1 2014年

项目只包含源码，有的使用 git，有的没使用
本地开发正式接触 virtualbox 和 vagrant，大部分还是 window 下直接搭建或者集成环境
当时的只有 blog 项目，尝试了 各种 gitpage；开始接触腾讯云

期间还有比如选 win / Ubuntu / mac 这样的纠结, 也花了大量时间尝试, 这是我最后的答案:

- win: 在我看来, win 的 GUI 即是更熟悉的, 也是更易用的(没有挑起圣战的意思, 同时为我的 小米笔记本Pro 打 call)
- docker: run, run, run
- 重度命令行使用者

**PS**: 我本人是 win 环境, 下面的讲解都以此为基础.

再分享 2 句话:

> 关于 docker: 不要因为环境, 限制了我们的想象力

> 关于命令行: graphical user interfaces make easy tasks easy, while command line interfaces make difficult tasks possible

## docker 基础概念

在讲解 docker 之前, 先对比一下 docker 和传统虚拟机:

- 虚拟机

![虚拟机](http://docker_practice.gitee.io/introduction/_images/virtualization.png)

- docker

![docker](http://docker_practice.gitee.io/introduction/_images/docker.png)

![docker vs 虚拟机](http://www.aixchina.net/home/attachment/201504/3/59140_14280451259eP2.png)

docker 使用的关键技术:

- go 语言实现
- 操作系统层面的虚拟化技术
- 基于 Linux 内核的 cgroup，namespace
- AUFS

Docker 包括三个基本概念

- 仓库（Repository）
- 镜像（Image）
- 容器（Container）

> 从生命周期的角度来理解和学习新的概念和知识.

![docker daemon](http://s4.51cto.com/wyfs02/M01/7A/F5/wKiom1bC8hryMLQeAAB6H6AJL2c652.png)

仓库的概念我们可以类比 git 仓库类来理解, 只是 git 仓库我们放的代码, 而 docker 仓库, 我们放置的是镜像. 其他仓库相关的操作, 都是一样一样滴.

镜像和容器, 可以类比 **类 和 对象**, 同样的类比, 还有 **程序 和 进程**. 运行容器, 就像实例化一个对象, 类起到定义的作用, 真正干活的是对象.

> 是不是感觉到触类旁通?

## 入门实战

入门了解 仓库/镜像/容器 这些概念就可以开始实战了. 在实践的过程中, 要 **刻意** 去理解这些概念以及 **生命周期**.

### 工具链

docker: docker 环境安装好后, 这是我们第一个接触到的工具. 更确切的说法是 `docker client`. 用来和 docker daemon 进行通信(流程见上面的图). 了解几个常用的就好:

```
# 镜像相关
docker pull image:tag
docker images
docker rmi

# 容器相关
docker run xxx
docker ps
docker rm

# other
docker info
docker help

docker build # 马上 Dockerfile 就要讲到
```

到了之类大家有没有疑问:

- 镜像哪来的?
- 刚才用 **类** 来类比镜像, 类是可以用代码定义的, 那镜像可以么?

先回答第一个问题, 我们可以通过容器来构建镜像, 命令是 `docker commit`, 和 git 仓库确实很相似. 但是在上面的列表里, 我并没有列举这个, 因为 **非主流** -- 我们也可以用代码来定义镜像, 这就是 Dockerfile.

直接看个例子就好:

```
FROM mysql:8
LABEL maintainer="1252409767@qq.com"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ADD my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod -R 644 /etc/mysql/conf.d/*

CMD ["mysqld"]

EXPOSE 3306
```

Dockerfile 就是这样一条一条 **指令** 组成, 我们从 **基础镜像** 开始, 一步一步构建出我们需要的环境的. 每一条指令, 对应我们在容器中执行相关操作, 然后把这个容器 commit 成新的镜像.

到这一步, 怎么构建环境和运行环境折腾清楚了, 然后继续实践, 比如 phper 的经典的环境: linux + nginx + php + mysql. 传统的做法, 我们把这个都安装到一个镜像里面就好了.

此处可以有停顿.

既然镜像都能一步一步的慢慢构建起来, 为什么我们不通过一个一个服务来组合成我们需要的环境呢? 于是到了 docker-compose 登场.

直接看例子就能明白:

```
fpm:
    build:
        context: ./server
        dockerfile: fpm.Dockerfile
    volumes:
        - ../:/var/www
        - ./logs/nginx/:/var/log/nginx
        - ./logs/fpm/:/var/log/php7
    # 注意这里, 通过 links 表示 fpm 这个服务需要 redis 和 fpm 服务
    links:
        - mysql
        - redis
        # - rabbitmq
    dns:
        - 223.5.5.5
        - 223.6.6.6
    extra_hosts:
        - "fpm:127.0.0.1"
    ports:
       - "80:80"
       - "443:443"
mysql:
    build:
        context: ./server
        dockerfile: mysql.Dockerfile
    volumes:
        - ./data/mysql:/var/lib/mysql
    ports:
        - "3306:3306"
    environment:
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
    ports:
        - "6379:6379"
```

这个是 docker-compose 的配置文件, yaml 格式, yaml 的语法就不多讲了, 简单程度和 json 一个级别. 配置内容也很简单, 看 key 基本就能看懂.

然后, 跑起来吧:

```
docker-compose up -d fpm # 启动 fpm, 应为 fpm 需要 redis 和 mysql, 所以也会跟着启动
docker-compose stop fpm
docker-compose logs fpm
docker-compose exec fpm bash # 执行容器中的命令
```

### 环境安装 & 配置

win 下的安装现在非常简单, 按照 [阿里云容器 - 镜像加速器](https://cr.console.aliyun.com/#/accelerator) 中的文档操作就好了.

- 推荐 win10 + docker for window, 当然 docker toolbox 也能玩, 希望现在没有我以前踩过的坑了
- 配置阿里云镜像加速
- 配置文件挂载
- 配置 docker 硬件资源, 推荐 2核4G

ok, 安装及环境就好了

### docker 的 lnmp 之旅

其实这部分内容, 上面在讲工具链的时候, 基本都提到了, 这里就演示一下最终效果就好.

## 写在最后

> 程序员认为自己需要交付的是代码，只要代码逻辑正确就好了，但实际上项目需要交付的是 代码 + 代码运行所需的环境

纯工具的角度来使用 docker, 其实相当轻松. 这次分享之后, 希望对 docker 有兴趣:

- 一定要实践
- 实践中巩固这些概念
- 站在巨人的肩膀上 -- 读官方镜像的 Dockerfile, 看其他人的源码

资源推荐:

- 入门值得一看的源码 [laradock](https://github.com/laradock/laradock)
- 开源入门好书 [《Docker — 从入门到实践》](http://docker_practice.gitee.io/)
- 深入讲解的好书 [<Docker——容器与容器云（第2版）>](http://www.ituring.com.cn/book/1899)
