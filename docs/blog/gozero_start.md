# go-zero 快速上手

[官网文档](https://www.yuque.com/tal-tech/go-zero/rm435c) 其实说得很详细, 作为一个写几年代码, 但是接手 go 项目不久的 go 新手, 来聊聊我跑 demo 遇到的问题

- 解决环境问题一: etcd, mysql, redis
- 解决环境问题二: go mod ding
- 编码中遇到的配置问题
- 加速开发的一些工具
  - goland 配置 run/debug
  - http client 插件: 编码化实现 api test
  - goreman 轻松管理多服务

## 环境准备一: 安装etcd, mysql, redis

官网就一句话, 至于背后要怎么折腾, 就看开发者自己了, 当然目前最简单最推荐的方式 -- docker, 直接使用 docker-composer 启动

```yaml
version: '3.1'
services:
  mysql:
    image: mysql
    ports:
      - "3306:3306"
    environment:
      TZ: Asia/Shanghai
      MYSQL_ALLOW_EMPTY_PASSWORD: 'yes'
  redis:
    image: redis:alpine
    volumes:
      - ./config/redis.conf:/etc/redis/redis.conf
      - ./data/redis:/data
    ports:
      - "6379:6379"
  etcd:
    image: quay.io/coreos/etcd
    environment:
      ETCD_ADVERTISE_CLIENT_URLS: "http://0.0.0.0:2379"
      ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379"
      ETCDCTL_API: "3"
    ports:
      - 2379:2379 # http
```

启动的话也超级简单, 就一行:

```sh
docker-compose up -d
```

启动好了后一定要 `check`, 避免环境的问题一直遗留到代码才发现

```sh
# redis-client
brew install redis
redis-cli -h localhost

# mysql-client
brew install mycli # mycli 和 mysql 兼容, 会有代码提示, 推荐
mysql -hlocalhost -uroot -proot

# etcd-client
brew install etcd
etcdctl --endpoints <url> # 后面接相应的命令
member list # 查看 etcd server
put foo bar
get foo
del foo
```

PS: 建议做笔记, 这种入门级的内容往往最常用, 但是一段时间不接触, 又最容易遗忘, 比如 `etcdctl`, 不翻笔记真不记得

当然, 直接使用 docker 也行, docker-compose 只是 `docker run` 的 trick:

```sh
# 最近切换到了 podman, 感兴趣可以看文后的连接, 有 podman 快速上手
alias docker=podman
podman run -d -p6379:6379 redis:alpine
podman run -d -p3306:3306 -e TZ='Asia/Shanghai' -e MYSQL_ROOT_PASSWORD=root mysql
podman run -d -p2379:2379 -e ETCD_ADVERTISE_CLIENT_URLS='http://0.0.0.0:2379' -e ETCD_LISTEN_CLIENT_URLS='http://0.0.0.0:2379' -e ETCDCTL_API='3' quay.io/coreos/etcd
```

## 环境准备二: go mod

接触 go 比较早但是一直不深, 就会对 go 的环境停留在似懂非懂的状态, 然后经常遇到一些奇奇怪怪的问题, 这里简单总结了下 go 项目使用 go mod:

```sh
go mod init xxx # 初始化, 后续在项目执行 go run/build 就都会使用到 go mod
go mod vendor # copy 依赖到项目下的 vendor 目录, 方便 IDE 代码提示
go env -w GOPROXY=https://goproxy.cn/,direct # 设置代理, 这样 go mod 下载包才快
export GOPROXY=https://goproxy.cn/,direct # 也可以设置到环境变量里
go env # 检查 go env

# 下面的 env 可设可不设
GO111MODULE="on" # 现在默认是 auto, 后续默认是 on, 有 go.mod 的项目下执行 go 命令, 都会是 go mod 模式下
GOFLAGS="-mod=vendor" # 默认值, 配合 go mod vendor 使用
```

有了这些后, 基本就能无障碍跑官方 demo 了, enjoy

## 编码中遇到的配置问题

以 `add.yaml` 为例:

```yaml
# go 中 mysql dsn 的示例: user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
DataSource: root:root@tcp(localhost:3306)/gozero # 这里需要加上 pass 部分
Table: book
Cache:
  - Host: localhost:6379
    Pass: "123" # 需要注意, yaml 中会自动识别 int/string 类型, 这里 Pass 需要使用 string 类型, 需要加 "" 指定
```

这里遇到的几个问题:

- 需要熟悉 go 中 dsn 的配置, go-zero 中的配置 `DataSource` 其实是复用
- 配置的问题还不够完善, 需要有一定的源码阅读能力, yaml 文件会初始化到 `config.go` 中, 配置项和里面的 `struct` 可以对应起来
- yaml 中数据类型自动识别的小细节, 详情见上面的注释

## 加速开发的一些工具

### goland 配置 run/debug

以 api gateway 为例, 要运行, 在 api 目录下执行 go run 即可:

```sh
go run bookstore.go -f etc/bookstore-api.yaml # -f 可以不加, bookstore.go 有设置默认值
```

可以在 goland 中轻松配置 run/debug:

- menu > run > edit configurations...
- 使用 template 中的 `go build` 模板, 改一下项目执行的目录即可

todo: add img

- 配置好后, 可以直接使用 `⌃R` 运行 `⌃D` debug

详细文档: [goland > run/debug](https://www.jetbrains.com/help/go/creating-and-editing-run-debug-configurations.html#create-go-build-configuration)

### http client 插件: 编码化实现 api test

通常进行 api test 有 2 种方式:

- 简单情况下, 直接使用 `curl` 测试, 比如 `curl -i baidu.com`
- 复杂情况下, 使用图形化工具, 比如 postman 等

但是这些方案, 都在 `编码性` 方面很一般, 如果不能编码, **复用和修改** 能力都大打折扣

所以我推荐:

- idea 下使用 [http client](https://www.jetbrains.com/help/idea/http-client-in-product-code-editor.html) 插件
- vscode 下使用 [rest client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) 插件

2 者基本大同小异, 以 demo 中的 api 为例:

```http
###
http://localhost:8888/check?book=go-zero

###
http://localhost:8888/add?book=go-zero&price=10
```

- `###` 表示一个 api case
- coding 过程中会有代码提示, 自动处理 http 相关的细节
- 执行(run) 非常简单, 图形界面/快捷键 都有支持

运行的结果如下, 轻松查看请求细节:

```sh
GET http://localhost:8888/check?book=go-zero

HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 30 Dec 2020 05:53:47 GMT
Content-Length: 25

{
  "found": true,
  "price": 10
}

Response code: 200 (OK); Time: 49ms; Content length: 25 bytes
```

### goreman 轻松管理多服务

拿 demo 举例, 有 `api add check` 3 个服务, 要同时运行 3 个服务才能体验完整流程, 在 goreman 下可以这样做:

- 安装 goreman

```sh
brew install goreman
```

- 编写 `Procfile` 文件

```sh
# 定义好每个服务
api: cd api; go run bookstore.go
add: cd rpc/add; go run add.go
check: cd rpc/check; go run check.go
```

- run

```sh
goreman start add check api
```

## 写在最后

当然要秀一下 QPS 啦: `Requests/sec:  44110.50`

```sh
➜  bookstore wrk -t10 -c1000 -d40s --latency 'http://localhost:8888/check?book=go-zero'
Running 40s test @ http://localhost:8888/check?book=go-zero
  10 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    19.91ms    9.06ms 103.27ms   63.78%
    Req/Sec     4.44k   575.52     5.57k    80.67%
  Latency Distribution
     50%   20.29ms
     75%   26.22ms
     90%   31.21ms
     99%   42.30ms
  1766254 requests in 40.04s, 224.03MB read
  Socket errors: connect 0, read 1062, write 0, timeout 0
Requests/sec:  44110.50
Transfer/sec:      5.59MB
```

附本机配置:

todo: add img

- 完整代码: <https://github.com/daydaygo/gozero-bookstore>
- 对黑苹果感兴趣, 可以看这篇: <https://www.jianshu.com/p/e2eae28601b2>
- 对 podman 感兴趣, 可以看这篇: <https://www.jianshu.com/p/75a2fc08a118>
