# docker

> <https://l.cncf.io>
> [Docker精华学习资料集锦](https://yq.aliyun.com/articles/65145)

![docker architecture](http://qiniu.dayday.tech/docker_architecture.png)

![docker command](https://docker_practice.gitee.io/appendix/_images/cmd_logic.png)

* k8s: kubectl helm OAM
* docker
  * dockerd daemon
  * cli client docker
  * docker-compose 服务编排
  * image镜像
    * Dockerfile: 标准化定义 环境+执行程序
    * docker repository: 镜像仓库, 用来分享镜像
    * 优化: 更小的基础镜像 alpine/scratch/busybox; 分阶段 build, 如 go
  * container容器=运行态镜像

---

* [安装与镜像加速](https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors)

``` sh
# ~/.docker/daemon.json
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://c3ywro5t.mirror.aliyuncs.com"]
}
EOF
sudo systemctl daemon-reload # systemd 管理
sudo systemctl restart docker
sudo service docker restart # service 管理
```

* docker 常用命令

``` bash
docker system df
docker system prune # -a, 没有容器使用的docker镜像和容器

docker info
# registry, 格式(user)/(repo_name)
docker search/history/push (image-name)
# image
docker pull imgage:tag # 下载镜像
docker images -a # 查看 images
# -f --filter, -q, --digests
docker image ls --format "table {{.ID}}\t{{.Repository}}\t{{.Tag}}" # go 模板格式语法
docker tag <image> <tag>
docker push <tag name>
docker rmi $(docker images -f "dangling=true" -q) # 先 stop、rm 容器，再删除名称为 none 的镜像

# container
docker commit container-id image:tag # build image from container; 不推荐，推荐使用 Dockerfile
docker ps -a # 查看所有容器, 默认只显示正在运行的容器
docker rm `docker ps -a -q` # 删除所有容器
docker run -ti --rm --name <container-name> image:tag /bin/bash # -d daemon; 查看镜像内容, 方便写 Dockerfile
-e XXX=xx
docker run --name xxx-app -d -p 8080:80 xxx # 创建容器, 绑定端口
docker start/stop/restart/rm/attach/logs/kill <container>
docker top
docker exec
docker cp <contaner> <local>
docker inspect # 查看容器详情
docker ps | awk '{print $1}' | xargs docker stop # 批量操作

docker volume ls
docker volume inspect my-vol

docker network create dev
docker network ls

docker stack # 编排能力: route
docker secret # 优雅实现安全编排
```

## Dockerfile

 `docker build . -t tag -f Dockerfile`

* . 当前context, **注意**会将当前context下的所有内容都发送到 docker daemon, 建议建子文件夹
* -t image:tag
* -f 默认读取 context 下的 Dockerfile, 文件名不同需要 -f 指定

``` Dockerfile
# 格式 INSTRUCTION argument
# 明确指定 image:tag, 避免基础镜像更新导致重新构建
FROM php:7.2.5-cli-alpine3.7
LABEL maintainer="1252409767@qq.com"
# 设置中文源加速
RUN echo -e "http://mirrors.ustc.edu.cn/alpine/v3.7/main\nhttp://mirrors.ustc.edu.cn/alpine/v3.7/community" > /etc/apk/repositories && \
    apk update
# 设置时区
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone
# 支持正则 go-filepath.Match: https://golang.org/pkg/path/filepath/#Match
COPY src dst
# 更高级复制: url 压缩文件解压
ADD src dst
# string => sh -c $cmd; array => $cmd
CMD /bin/echo
CMD ["/bin/echo"]
# 1. 镜像作为命令使用 2. 镜像启动前的准备工作
ENTRYPOINT
# 暴露端口
EXPOSE 80:8080
EXPOSE 80
# 影响其他命令使用相对路径 RUN / CMD / ENTRYPOINT / COPY
WORKDIR /path/to/workdir
ENV <key> <value>
# 只在 build 时有效, --build-arg key=value
ARG <key> <value>
# 切换用户执行, 使用 gosu 替换 su/sudo
USER <uid>
# 挂载, 数据持久化
VOLUME ["/data"]
# 触发器, 基于当前镜像构建镜像时, 触发器的内容才会执行
ONBUILD COPY . /app
# 健康检查
HEALTHCHECK

# 分阶段build
FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app
RUN cd /app && go build -o goapp
FROM alpine
RUN apk update && \
 apk add ca-certificates && \
 update-ca-certificates && \
 rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/goapp /app/
EXPOSE 8080
ENTRYPOINT ./goapp
```

## docker-compose

* docker-compose 命令行

``` sh
# install
sudo curl -L "https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

abbr -a doc docker-compose # 简写, 方便使用
doc up -d <service> -f <yml file>
doc ps/rm/stop/restart
```

* docker-compose.yml, [常见示例](https://github.com/daydaygo/coder/tree/master/src/docker)

``` yaml
version: '3.1' # 语法版本
networks:
    default:
        external:
            name: mt
services: # 定义服务
    service-name:
        image: image:tag # 明确指定镜像版本
          labels: # 配置路由服务; 基于nginx负载均衡; 基于api服务发现
            aliyun.routing.port_8080: 'http://tomcat-sample'
            aliyun.scale: '3'
        build: dir
        build:
            context: dir # context, 上下文
            dockerfile: Dockerfile
            args: # 替换Dockerfile中的 ARG 参数
                arg1:val1
        command: ["cmd"] # 用来覆盖缺省命令/添加参数
        restart: always # 自动重启, 比较适合 mysql/redis 等稳定的基础服务, 业务镜像可能导致一直失败重启
        environment: # 字典/数组 格式
            env: val
        links: #
            - service-name
        ports:
            - "local:container"
        volumes:
            - local:container
            - ./xxx.conf:/etc/xxx.conf # 挂载配置文件
            # https://docs.docker.com/docker-for-mac/osxfs-caching/
            - local:container:default # delegated cached consistent
        volumes_from:
            - volume-name
        extra_hosts: # /etc/hosts
            - "host:ip"
        network_mode: "bridge" # 网络模式: bridge none container:name host
        networks: # 搭配上面的 networks 配置使用, 使用特定 network 配置
            mt:
        dns: 8.8.8.8
        dns:
            - 8.8.8.8
        tty: true # 打开此配置才可以 exec 进入容器
        # 其他配置参考 docker run 命令参数
    db: # 使用阿里云RDS
        external:
            host: rds******.mysql.rds.aliyuncs.com
            ports:
            - 3306
        environment:
            - MYSQL_DATABASE=blog
            - MYSQL_USER=ghost
            - MYSQL_PASSWORD=***********
```

## podman

* <http://docs.podman.io/>

```sh

```

## etcd

* etcdctl

``` sh
brew install etcd
etcd --endpoints xxx
member list
put foo bar
get foo
del foo
```

## consul

![Consul 的交互图](https://s0.lgstatic.com/i/image/M00/3E/CF/CgqCHl8tP0uAfPqfAAC1xfaVTwQ927.png)
![consul 架构图](https://s0.lgstatic.com/i/image/M00/3E/CE/CgqCHl8tP0OAVC4_AAIBrjsMQhU949.png)
![consul etcd zookeeper 对比](https://s0.lgstatic.com/i/image/M00/3E/C3/Ciqc1F8tPxCAT_4RAADBzRFlUA0352.png)

```sh
/v1/agent/service/register // 服务注册接口
/v1/agent/service/deregister/${instanceId} // 服务注销接口
/v1/health/service/${serviceName} // 服务发现接口
```

## kong

* <https://konghq.com/install/>
* [plugin](https://konghq.com/hub/): jwt prometheus zipkin

## httpbin

* ennethreitz/httpbin
* <http://httpbin.org/>
  * /status/502

## 问题 & 资源

* docker容器无法访问外网, 可能是 [dns 引起的](https://segmentfault.com/a/1190000010261378)
