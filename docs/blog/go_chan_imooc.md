# go| 慕课网分享邀请: go channel讲解

应慕课网内容分享邀请, 录制 go channel 讲解主题的视频

- 课程大纲
- 试讲指南: 时长>=30分钟 试讲大纲 PPT 准备材料 编写代码 构思讲解过程 发送视频给BD
- 录课教程: 微软雅黑20/变敲代码变讲解/光标选中讲解/英语发音准确

讲课思路:

- 总: 目标 场景 效果
- 分: 知识重难点 原理分析 详细讲解
- 总: 总结开发思路 梳理知识点

讲课技巧:
- 设问/反问来引导思考
- 轻松幽默 平等开放: 有意思的 欢乐的 抢眼球的 怪怪的 意料之外
- 抑扬顿挫
- 重点内容重点提示
- 利用互动: 举例 比喻 对比 设问 设置陷阱

很荣幸收到 [慕课网](https://www.imooc.com) 内容分享邀请, 录制 go channel 讲解主题的视频, 定位在 初级+基础, 希望帮助大家快速入门 go, 熟悉 go channel 的使用.

分享大纲:

- go 开发环境配置
- go channel 讲解: 为什么会有 go 和 channel, channel 编程上的细节
- 提升篇: 如何开始开发一个大型 go 项目; 深入理解协程之「快」

- 耳麦:
https://detail.tmall.com/item.htm?spm=a1z10.5-b.w4011-10825454581.67.I7Gqa1&id=10540935532&rn=722e1950fa735cf24af0a3f5126db532&abbucket=14&sku_properties=165354720:6536025
https://detail.tmall.com/item.htm?spm=a1z10.5-b.w4011-10825454581.59.I7Gqa1&id=531495180303&rn=722e1950fa735cf24af0a3f5126db532&abbucket=14&skuId=3169101492001

- 时间取决于自己
- 免费课1-2h 实战课>10h
- 大纲 -> 审核会讨论 -> ppt/代码 -> 视频(正式录制前: 改进和注意事项)

课程推荐:
- https://www.imooc.com/learn/980    这个老师的讲法是现在全平台认为，最适合入门级的讲法， 在讲法细节上可学习的地方很多
- https://www.imooc.com/learn/927 讲师非常厉害
- 很少口误， 不在讲课中组织语言（想好再说或者写稿子）， 有引入，有思路，有总结
- 一句话要说顺，如果出现口误、嗑巴的情况，就停几秒钟，重新说一遍  这样后期处理后，听起来也是顺畅的

go channel 容量/缓存 http://www.hi-roy.com/2018/06/04/GO语言学习笔记-缓冲区Channels和线程池/
☐ 慕课网 go channel 讲解
☐ go并发编程实战 - coding 练习 tmp/goc2p.v2 http://www.ituring.com.cn/book/1950

## 基础知识储备

- 怎么运行 go 代码: 安装 -> 运行 暂时不考虑完整项目
- 为什么要使用协程 -> 天下武功, 唯快不破 -> 要快, 要有很多很多的协程
- 为什么要使用 channel -> 承接上面, 因为有多个协程

## go channel 讲解

- go并发任务 blog
- go监控系统 blog

chan功能

- 数据通信: 协程间传递数据
- 协程调度: 往 chan 中发数据, 如果通道已满 / 从 chan 中取数据, 如果通道是空
- 通道编程细节: for写法 申明通道的方向 通道缓存大小

## 提升篇

- 写一个项目级的代码
- go 协程快的根源: CSP理论 + MPG模型

A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished

- 命令源码文件(package main) vs 库源码文件

分享大纲:

- go 开发环境配置
  - 搜索 -> 官网/golang中文社区 -> 下载安装
  - 环境变量: 进程皆有环境 / GOROOT GOPATH GOBIN
  - goland -> 命令行 -> go常用目录配置 src / pkg / bin -> goland 的几个常用快捷键
  - 编写 hello world -> package main -> main函数主协程
- go channel 讲解: 为什么会有 go 和 channel, channel 编程上的细节
- 提升篇: 如何开始开发一个大型 go 项目; 深入理解协程之「快」

## go 开发环境配置

- 安装
- 配置环境变量

**环境变量** 是非常常见的一个术语, 也是很多基础不太好的同学, 不太理解的一个概念. 要理解它其实很简单:

> 进程皆有环境  -- [理解UNIX进程](http://www.ituring.com.cn/book/1081)

程序运行之后, 在操作系统里面就是进程了, 而进程都有自己的环境, 环境变量属于环境的一种, 是进程可以访问到的资源.
其中非常常用的, 是 `PATH` 变量, 进程通过这个变量来寻找可执行文件. 如果没有这个变量的话, 就需要使用可执行文件的完整路径了.
比如 `go.exe` 可执行文件, 完整路径是 `c:/tools/go/go.exe`, 把 `c:/tools/go` 加入到 `PATH` 环境变量里, 才可以直接执行 `go run` 等命令, 否则就需要使用 `c:/tools/go/go run`.

起步阶段, 先只需要配置 `GOROOT` 和 `GOPATH` 2 个环境变量.

在 window 下配置环境变量有更快捷的操作: 使用 `win+r` 快捷键打开 `运行窗口` 输入 `sysdm.cpl`, 就可以快速打开环境变量的设置界面.

- hello go

环境配置好后, 就可以还是 `hello world` 之旅啦. 推荐 2 款工具: `vscode goland`.

### vscode: hello go

vscode 会自动监测代码文件使用的语言, 提示安装相应的插件, 按照提示安装即可. 推荐安装 `code runner` 插件, 安装后使用 `ctrl+alt+n` 就可以快速运行当前文件

### goland: hello go

goland 非常推荐, 强大的代码提示和很多带来效率提升的功能, 推荐使用 `view | distraction free mode`, 代码 **清爽** 的编码体验

## go channel 讲解

- hello.go: hello world

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello czl")
}
```

使用 `go run hello.go` 运行

### 为什么使用协程
> 天下武功, 唯快不破

为了体验协程之快, 需要在 `hello.go` 做一点改造, 模拟一件耗时的任务, 这样方便比较

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("hello czl")

    // 普通版
    doSomeThing()
    // 协程版
    go doSomeThing()
}

// 模拟一件耗时的任务
func doSomeThing() {
    time.Sleep(time.Second*1)
}
```

只用添加 `go` 关键词, 就可以将任务抛给协程来执行, 是不是 hin 简单

使用 `time go run hello.go` 分别来比较, 会发现, **协程版快很多** 呀

不过如果细心的话, 就会发现这个 **快得有问题** -- 耗时任务明明需要 1s, 怎么协程版居然 1s 内完成了?
改造一下模拟任务的函数, 就可以发现了:

```go
// 模拟一件耗时的任务
func doSomeThing() {
    time.Sleep(time.Second*1)
    // 增加输出, 方便查看程序的执行效果
    fmt.Println("do some thing")
}
```

这就是 go 程序运行机制导致的 **陷阱** -- go程序运行时, 会依次寻找 `package main` -> `func main()`, 依次执行 main() 中代码,
当使用 go 关键字把任务交给协程执行后, 程序继续向下执行, 然而下面并没有其他需要执行的代码了, 所以 main() 就退出了, 并没有等待协程执行完.

问题来了: **怎么让我们程序里的协程执行完呢?**

这里有三种方式:

- 方式一, 使用 sleep, 等协程执行完, 但是就不好比较协程版和非协程版的耗时了

```go
// 方式一: sleep
go doSomeThing()
time.Sleep(time.Second*4)
```

- 方式二: 使用 channel, 通过读写 channel 触发协程等待

```go
    // 方式二: channel
    ch := make(chan bool)
    go func() {
        doSomeThing()
        // 任务执行完了才给 channel 写数据
        ch <- true
    }()
    // 要等待 channel 里有数据
    <- ch
```

也正是 channel 的这一特性, 在后面 **channel功能二: 协程调度** 中会更详细的讲解

- 方式三: 调用 sync 系统库里的 WaitGroup API 实现

官方的定义如下:

```
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished.
At the same time, Wait can be used to block until all goroutines have finished
```

用代码翻译过来, 是四步:

```go
    // 方式三: WaitGroup
    var wg sync.WaitGroup // 第一步: 声明 WaitGroup
    wg.Add(1) // 第二步: 添加需要等待的协程, 需要等多少个协程, 就添加多少
    go func() {
        // 执行完之后需要做的操作
        defer wg.Done() // 第三步: 协程调用 done() 表示协程执行完了

        // 正常需要执行的逻辑
        doSomeThing()
    }()
    wg.Wait() // 第四步: 什么地方需要等待协程执行完, 就在上面地方加上 wait()
```

这里使用 go 中的 `defer` 关键字, defer 关键字可以在函数执行完后执行.
**为什么要在函数一开始调用却在函数最后执行呢?** 比如函数一开始打开了一个文件, 最后写了很多逻辑, 忘了关闭文件了, 使用 defer 就可以很好的避免这种情况.

准备工作, 或者说 **热身** 时间有点长, 来看看协程有多快, 多执行一点任务:

```go
    // 快在哪? 多来点任务不就知道了
    for i := 0; i < 4; i++ {
        doSomeThing()
    }

    // 协程版
    var wg sync.WaitGroup
    wg.Add(4)
    for i := 0; i < 4; i++ {
        // wg.Add(1)
        go func() {
            defer wg.Done()

            doSomeThing()
        }()
    }
    wg.Wait()
```

使用 `time go run hello.go`, 会发现协程版要快很多. 可以把 for 循环再调多点, 感受就更明显了

```
> 23:47 src $ time go run hello.go
hello world
do some thing
do some thing
do some thing
do some thing

real    0m4.978s
user    0m0.000s
sys     0m0.015s

> 00:23 src $ time go run hello.go
hello world
do some thing
do some thing
do some thing
do some thing

real    0m2.200s
user    0m0.000s
sys     0m0.031s
```

## 小热身: go 开发利器 goland

- 演示: 10s 一个 hello
- 演示: 初始界面的几个功能
- 演示: 多种方式新建文件 多种方式查找
- 实现方式这么多, 是不是学习成本太高了 -- 不要在更新了, 学不动了. NONONO, 选择的自由, 也是编程的乐趣之一, happy coding
- 界面设置: distraction free mode

## go channel 讲解: 为什么要使用 channel

- 协程调度

方式二: 使用 channel, 通过读写 channel 触发协程等待

```go
    // 方式二: channel
    ch := make(chan bool)
    go func() {
        doSomeThing()
        // 任务执行完了才给 channel 写数据
        ch <- true
    }()
    // 要等待 channel 里有数据
    <- ch
```

- 协程间数据通信

## 写在后面

感受到 **协程之快** 以及 **需要多个协程来执行任务** 后, 就引入了下一个话题:

> 为什么要使用 channel? -- 因为有多个协程.