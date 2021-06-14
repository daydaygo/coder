# go| jupiter 学习笔记

生命不息, 学习不止, 这次我们来折腾 [jupiter 框架](https://github.com/douyu/jupiter)

> [【斗鱼】没人比我更懂微服务-Go 微服务框架Jupiter](https://www.bilibili.com/video/BV1FT4y1J7tV)

## helloworld

把官网的 example 都实现了一遍, 才发现 helloworld 应该是这样的:

- 最简单版

```go
package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/xlog"
)

func main() {
	var app jupiter.Application
	app.Startup() // 启动框架, 可以使用框架的各种功能了
	xlog.Info("hello world")
}
```

- 稍微来点封装

```go
package main

import (
    "github.com/douyu/jupiter"
    "github.com/douyu/jupiter/pkg/xlog"
)

func main() {
    var app jupiter.Application
    app.Startup(testLog) // 支持在框架初始化后, 执行特定的方法
}

func testLog() error { // 封装成方法
    xlog.Info("hello world")
    return nil
}
```

- 更复杂点, **套个壳**

```go
package main

import (
	"fmt"
	"github.com/douyu/jupiter"
    "github.com/douyu/jupiter/pkg/xlog"
)

func main() {
	eng := NewEngine()
	fmt.Println(eng)
}

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	eng.Startup(testLog)
	return eng
}

func testLog() error {
	xlog.Info("hello world")
	return nil
}
```

> PS: 为了下面讲解代码方便, 均不使用套壳版

- 再深入点, 看看 `Startup` 干了些啥

```go
func (app *Application) Startup(fns ...func() error) error {
	app.initialize() // 初始化 app
	if err := app.startup(); err != nil { // 初始化 falg/log/config/trace/governor 等模块
		return err
	}
	return xgo.SerialUntilError(fns...)() // 这是为啥支持传入多个方法
}
```

## 生命周期

- 直接上完整的例子

```go
package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/conf" // conf 模块
	"github.com/douyu/jupiter/pkg/registry/compound"
	"github.com/douyu/jupiter/pkg/registry/etcdv3" // 除了 registry, 还是 client 的使用例子
	"github.com/douyu/jupiter/pkg/server"
	"github.com/douyu/jupiter/pkg/server/xecho"
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/worker"
	"github.com/douyu/jupiter/pkg/worker/xcron"
	"github.com/douyu/jupiter/pkg/xlog" // log 模块
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
    var app jupiter.Application
    
    // 初始化框架的功能, 这里额外传入了
    app.Startup(fileWatcher)

    // 修改 xlog.DefaultLogger, 从而改变 xlog 的行为
    // 后面会具体讲解 config/log 模块
	xlog.DefaultLogger = xlog.StdConfig("default").Build()

    // 可以启动多个 server
	app.Serve(startEcho())
    app.Serve(startGin())
    
    // 可以设置注册中心, server 启动是会自动注册进去, 这里使用 etcd 作为注册中心
	app.SetRegistry(compound.New(etcdv3.StdConfig("etcd").Build()))

    // 设置 worker
	app.Schedule(startWorker())

    // 启动应用
	app.Run()
}

func fileWatcher() error {
	go func() {
		peopleName := conf.GetString("people.name")
		xlog.Info(peopleName)
		time.Sleep(time.Second*10)
	}()
	return nil
}

func startEcho() server.Server {
	s := xecho.DefaultConfig().Build()
	s.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "echo")
	})
	return s
}

func startGin() server.Server {
	s := xgin.StdConfig("http").Build()
	s.GET("/gin", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
	return s
}

func startWorker() worker.Worker {
	cron := xcron.DefaultConfig().Build()
	cron.Schedule(xcron.Every(time.Second*10), xcron.FuncJob(func() error {
		xlog.Info("cron")
		return nil
	}))
	return cron
}
```

- 对应的配置

```toml
# jupiter 默认提供, governor 用于服务治理
[jupiter.server.governor]
enable = false
port = 2345

# server 配置
# http server: echo gin goframe
# grpc server
[jupiter.server.http]
#enable = false
port = 1234

# registry: registry + 具体实现(这里是 etcd)
[jupiter.registry.etcd]
configKey = "jupiter.etcdv3.default"
timeout = "1s"
[jupiter.etcdv3.default]
endpoints = ["127.0.0.1:2379"]
secure = false

[jupiter.cron.test]
withSeconds = false
concurrentDelay= -1
immediatelyRun = false

[jupiter.logger.default]
debug = true
enableConsole = true
async = false

# 自定义配置
[people]
name = "daydaygo"
```

### 框架的执行流程如下

- `app.Startup(fileWatcher)`: 上一步讲到, 初始化框架的功能, 这里传入了 `fileWatcher`, 可以使用动态更新配置, 后面会详细讲 `-watch` 功能
- `app.Serve()`: 设置 server
- `app.Schedule()`: 设置 worker
- `app.run()`: 启动 app, 执行 server/worker 等内容

### 看一下 `app.run()` 源码就明白了

```go
func (app *Application) Run(servers ...server.Server) error {
	app.smu.Lock()
	app.servers = append(app.servers, servers...) // app.Serve() 其实就是设置 app.servers 变量
	app.smu.Unlock()

	app.waitSignals() //start signal listen task in goroutine
	defer app.clean()

	// todo jobs not graceful
	app.startJobs()

	// start servers and govern server
	app.cycle.Run(app.startServers) // 这里完成 server + server 注册到注册中心
	// start workers
	app.cycle.Run(app.startWorkers) // 这里执行 worker

	//blocking and wait quit
	if err := <-app.cycle.Wait(); err != nil {
		app.logger.Error("jupiter shutdown with error", xlog.FieldMod(ecode.ModApp), xlog.FieldErr(err))
		return err
	}
	app.logger.Info("shutdown jupiter, bye!", xlog.FieldMod(ecode.ModApp))
	return nil
}
```

## jupiter 的几大模块

- config

默认配置文件使用 toml 格式, 使用 `--config` flag 来使用本地配置文件

```sh
go run main.go --conifg=config.toml
```

属于 jupiter 的模块, 使用 `[jupiter.模块名.名字]` 来使用, 比如 `[jupiter.server.http]`, 则是一个 jupiter server 的配置, 这个 server 名字为 http

jupiter 中通过 2 类配置来初始化模块:

```go
// 使用默认配置
xlog.DefaultConfig().Build()

// 使用配置文件: [jupiter.logger.default]
xlog.xlog.StdConfig("default").Build()
```

理解了上面这些, 就掌握了配置的核心用法, 使用 Apollo/etcd 等配置中心, 配置文件的 filewatch 都是在此基础之上

- log

上面其实已经看到 log 的模块的用法了, 需要修改 log 的行为, 只需要修改配置, 并且使用如下代码设置生效即可:

```go
// 设置 DefaultLogger 即可
xlog.DefaultLogger = xlog.StdConfig("default").Build()

// 看一下 xlog.info 的源码就能知道答案
func Info(msg string, fields ...Field) {
	DefaultLogger.Info(msg, fields...)
}
```

只要理解了这一点, 就已经理解了日志的核心用法, 日志 level, 日志输出到 stdout/file 都在此基础之上

- server registry governor

server 这部分内容是 jupiter 的重中之中, jupiter 增加了对 echo/gin/frame/grpc 等 server 的适配使用 xecho/xgin/xframe/xgrpc 等进行配置和使用, 非常的简洁方便

使用 registry 适配配置中心, 目前适配了 etcd

使用 governor 进行服务治理(在 `app.startuUp` 阶段就设置好了, 在 `app.run` 阶段启动)

**理解了这几个模块之间的关系**, 就很容易理解 server 模块的核心用法

- worker

worker 比较简单, 对应 `[jupiter.cron.xxx]` 下的配置, 按需设置即可

## jupiter 其他内容

- jupiter 默认支持一些 flag(命令行参数), 可以使用 `go run main.go -h` 查看
- `-watch` 的场景:
    - 修改 log level, info -> debug, 方便线上有问题时搜集更多日志进行分析
    - 修改自定义配置, 可以实时生效
- 自己遇到的一些问题
    - log 如果没有配置 `async`, 在 server 启动后, 每隔 30s 输出一次, 这导致我通过 log 来验证的场景, 以为是遇到 bug 了
    - 我测试代码的时候喜欢忽略 err, 虽然代码看起来 **简单很多**, 给 debug 增加了难度, 同时不利于养成好的工程习惯
    - debug 遇到 context timeout, 这个属于没有经验, context timout 不会因为 debug 的单步调试停止计时, 导致我绕进去了很久, 才发现是 `context timeout` 触发了

## 快速配置 etcd 开发环境

jupiter 很多功能都需要 etcd 支持, 可以使用 docker-compose, 本地快速起起来:

```yaml
version: '3'
services:
    etcd:
        image: quay.io/coreos/etcd
        environment: 
            ETCD_ADVERTISE_CLIENT_URLS: "http://0.0.0.0:2379"
            ETCD_LISTEN_CLIENT_URLS: "http://0.0.0.0:2379"
            ETCDCTL_API: "3"
        ports: 
            - 12379:2379 # http, 本地的端口自己设置
            # - 2380:2380 # 节点间
            # - 4001:4001
    etcda: # 简单的管理界面
        image: evildecay/etcdkeeper
        environment: 
            HOST: 0.0.0.0
        ports:
            - 10280:8080 # 本地的端口自己设置
        links: 
            - etcd
```

也可以直接使用 etcdctl 来测试:

```sh
# install
brew install etcd

# use
etcdctl --endpoints=127.0.0.1:12379 get '/hello'
```

## 开源逗逼唠

jupiter 的这次开源在我这个开源老兵(github star 4k+ 和 star 3k+ 框架的核心开发者)看来看来确实有些仓促, 主要集中在文档这块, 至于源码, 目前 **实力不允许**, 总得多看看多写写, 能拿出足够多的干货时再 BB

从目前文档看到的几个问题:

- 文档基于 [vuepress](https://www.vuepress.cn/), 简单实用上手快, 不过 jupiter 源码和文档是分 2 个不同项目的, 这就导致 `edit on github` 一直 404, [我已经给开发组提了 PR](https://github.com/douyu/jupiter-website/pull/5)
- 部分 url 404, 这种算是非常低级的错误了, 通常因 **年久失修** 会比较多, 但是 jupiter 才开源多久
- 部分贴的代码实例有错误, 所以关于代码, 一是要使用源码中提供的 example, 二是一定要自己动手跑起来, 文档贴代码因为 **上下文不全**, 人为失误等, 一向是重灾区, 受欢迎的开源项目文档有多人参与贡献, 这块要好很多
- 文档在 **组织** 上对新人并不是特别友好, 或者说文档没有遵循一定的 **套路**, 导致引起一些不必要的麻烦(我踩了几个, 后面一一列出来)

关于文档中错误的部分, [我也一并提交了一个 PR](https://github.com/douyu/jupiter-website/pull/7)

最后来几句开源老兵的叨逼叨:

- 希望不是一个 KPI 项目, 虽然多看源码总是有帮助的, 但是, 那感觉会像吃了苍蝇一样
- 时间是开源软件的朋友, 时间稍微拉长一点, 是否 **真的开源**, 一目了然, 这里并不是 **结果导向**, 开源确实需要付出很多, 才能做好

## 写在最后 -- 如何快速上手一个框架?

- 熟悉文档和 api, 勤做笔记和练手
- 生命周期思考法, 了解框架的执行流程
- 翻源码, 有奇效