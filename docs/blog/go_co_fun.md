# go| 感受并发编程的乐趣

学习了 [ccmouse - googl工程师](https://www.imooc.com/t/4898742) 在 [慕课网 - 搭建并行处理管道，感受GO语言魅力](https://www.imooc.com/learn/927), 获益匪浅, 也想把这份编程的快乐传递给大家.

强烈推荐一下ccmouse大大的课程, 总能让我生出 **Google工程师果然就是不一样** 之感, 每次都能从简单的 **hello world** 开始, 一步步 coding 到教程的主题, 并在过程中给予充分的理由 -- 为什么要一步步变复杂. 同时也会亲身 **踩坑** 示范, 干货满满.

内容提要:

- let's go: 为什么要使用 go ? swoole2.1 带来的 `go + channel` 编程体验
- go's design: go语言的设计
- go's hello-world: go 写 hello world 的各种姿势, 为之后并发编程解决的 big problem 埋下伏笔
- go's sort: 内排排序 -> 外部排序
- go's io: 二进制文件读写 + 大文件读写加速
- go 实现完整外部排序
- go 实现集群版(web版)外部排序

另外, ccmouse大大关于语言学习的方法也值得借鉴:

- 首先, 学习一下语言语法的要点
- 立刻找一个不那么简单的项目来做, 边做边查文档/stackoverflow

## let's go

swoole2.1 发布了([Swoole 2.1 正式版发布，协程+通道带来全新的 PHP 编程模式](https://segmentfault.com/a/1190000013239349)), 看着示例代码颇有点 **陌生** 之感, 看来真想要 **写好** 协程, 熟悉 go 应当是必选项了.

示例代码, go & channel:

```php
// coroutine
go(function () {
    co::sleep(0.5);
    echo "hello";
});
go("test");
go([$object, "method"]);

// go + channel
$c3 = new chan(2);
$c4 = new chan(2);
$c3->push(3);
$c3->push(3.1415);
$c4->push(3);
$c4->push(3.1415);
go(function () use ($c3, $c4) {
    echo "producer\n";
    co::sleep(1);
    $data = $c3->pop();
    echo "pop[1]\n";
    var_dump($data);
});
```

`go` 看起来就是一个接收 **闭包函数** 作为入参的函数, `channel`(通道)看起来也不过是一个类似 **栈** 的数据结构(`push()` 和 `pop()` 2 种操作). 看到示例代码却感到十分 **陌生** -- 为什么要这样写呀?

既然是 **借鉴至 go 语言**, 看来有必要了解一下 go, 来加深一下理解了

## go's design

Google内部的「标准」编程语言:

- c++: 性能保障部分, 如搜索引擎
- java: 复杂业务逻辑, 如 adwords, Google docs
- Python: 大量内部工具
- go: 新的内部工具, 及其他业务模块

语言对比:

- c/c++: 性能, 可以做系统开发; 没有繁琐的类型系统, 简单统一化的模块依赖管理, 编译速度飞快
- java: 垃圾回收 -> 慢, 会影响业务
- Python: 简单易学, 灵活类型, 函数式编程, 异步IO; 没有编译器静态类型检查

go 设计:

- 类型检查: 编译时
- 运行环境: 编译成机器码直接运行(不依赖虚拟机)
- 编程范式: 面向接口, 函数式编程, **并发编程**

go 并发编程:

- CSP, Communication Sequential Process 模型
- 不需要锁, 不需要 callback
- 并行计算 + 分布式

go 线程实现模型 MPG:

- M: Machine, 一个内核线程
- P: Processor, M所需的上下文环境
- G: Goroutine, 代表一段需要被并发执行的 go 语言代码的封装
- KSE: 内核调度实体
- 一个 M 和一个 P 关联, 形成一个有效的 G 运行环境, 每个 P 都会包含一个可运行的 G 的队列(runq)

![go 线程实现模型 MPG](http://upload-images.jianshu.io/upload_images/567399-3ed6d22cb14935af.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

进程/线程开多了都会涉及到系统进行调度导致的消耗, go 通过 MPG 模型进行映射, 可以并行很多 go协程, 底层自动实现调度

## go's hello world

4 个版本的 hello world:

- 原始版: 直接 print 输出
- http版: 使用 http 库输出到 web
- go版: 开始接触 go协程
- go+channel版: 开始接触 go协程 + channel数据传递

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    // 原始版
    helloWorld1()

    // http版
    helloWorld2()

    // go协程版
    for i := 0; i < 5000; i++ { // 协程 <-> 线程 之间存在隐射, 具体参考 <go并发编程> 一书中的「线程实现模型」
        go helloWorld3(i)
    }
    time.Sleep(time.Microsecond) // 添加延时, 协程才有机会在 main() 退出前执行

    // go + channel
    ch := make(chan string)
    for i := 0; i < 5; i++ {
        go helloWorld4(i, ch)
    }
    for {
        msg := <-ch
        fmt.Println(msg)
    }
}

func helloWorld1() {
    fmt.Println("hello world")
}

func helloWorld2() {
    // 为什么要使用指针: 因为参数可以被改变
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        //fmt.Fprintln(writer, "<h1>hello world</h1>")
        fmt.Fprintf(writer, "<h1>hello world %v</h1>", request.FormValue("name"))
    })
    http.ListenAndServe(":8888", nil)
}

func helloWorld3(i int) {
    fmt.Printf("hello world from goroutine %v \n", i)
}

func helloWorld4(i int, ch chan string) {
    for {
        ch <- fmt.Sprintf("hello world %v", i)
    }
}
```

## go's sort

先科普一下排序的知识:

- 排序分 **内部排序** + **外部排序** 两种, 区分在于数据量, 内部排序可以将数据全部放到内存中, 然后进行排序
- 常见的内部排序算法: 冒泡, 快排, 归并排序等, 其中 **快排** 是速度最快的不稳定排序算法, **归并排序**可以应用于外部排序
- 归并排序时, 可以一次归并多节点达到加速的效果, 例子中使用 **二路归并** 来简单演示

再来看看 go 实现的归并排序:

```go
package main

import (
    "sort"
    "fmt"
)

func main() {
    // 内部排序 -> 快排
    a := []int{3, 6, 2, 1, 9, 10, 8}
    sort.Ints(a)
    fmt.Println(a)

    // 外部排序 -> 归并排序
    c := Merge(inMemSort(arraySource(3, 6, 2, 1, 9, 10, 8)),
        inMemSort(arraySource(7, 4, 0, 3, 2, 8, 13)))
    // channel 中获取数据: 原始
    //for {
    //  // 风格一
    //  //if v, ok := <-c; ok {
    //  //  fmt.Println(v)
    //  //} else {
    //  //  break
    //  //}
    //  // 风格二
    //  v, ok := <-c
    //  if !ok {
    //      break
    //  }
    //  fmt.Println(v)
    //}
    // channel 中获取数据: 简写
    for v := range c {
        fmt.Println(v)
    }

}

func arraySource(a ...int) <-chan int { // 指定从 channel 获取数据
    out := make(chan int)
    go func() {
        for _, v := range a { // for + range + _
            out <- v
        }
        close(out) // 数据传递完毕, 关闭channel
    }()
    return out
}

func inMemSort(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        // read into memory
        a := []int{}
        for v := range in {
            a = append(a, v)
        }
        // sort
        sort.Ints(a)
        // output
        for _, v := range a {
            out <- v
        }
        close(out)
    }()
    return out
}

func Merge(in1, in2 <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        // 归并的过程要处理某个通道可能没有数据的情况, 代码非常值得一读
        v1, ok1 := <-in1
        v2, ok2 := <-in2
        for ok1 || ok2 {
            if !ok2 || (ok1 && v1 <= v2) {
                out <- v1
                v1, ok1 = <-in1
            } else {
                out <- v2
                v2, ok2 = <-in2
            }
        }
        close(out)
    }()
    return out
}
```

## go's io

文件读写算是一个 **常见且简单** 的任务, 但是:

- 如何读写二进制文件呢? int 大小是多少? 大端序小端序?
- 如何读写大文件, 比如 800m ? 什么是 buffer ? 什么是 flush ?

这一节的代码就用来处理这些问题:

```go
package main

import (
    "encoding/binary"
    "fmt"
    "io"
    "os"
    "math/rand"
    "bufio"
)

func main() {
    TestIo()
    largeIo()
}

// 文件读写测试
func TestIo() {
    // 写
    file, err := os.Create("small.in") // 二进制数据
    if err != nil {
        panic(err) // 遇到错误, 暂时不处理
    }
    defer file.Close()    // 运行结束后, 自动关闭文件
    p := randomSource(50) // small.in 的大小 = 8*50
    writerSink(file, p)
    // 读
    file, err = os.Open("small.in")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    p = readerSource(file)
    for v := range p {
        fmt.Println(v)
    }
}

// 大文件读写
func largeIo() {
    const filename = "large.in"
    const n = 100000000 // 8*100*1000*100 = 800M
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    p := randomSource(n)
    writer := bufio.NewWriter(file) // io buffer 加快文件读写
    writerSink(writer, p)
    writer.Flush() // flush 掉 io buffer 中的数据
    // 读
    file, err = os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    // 写, 只测试前 100 个
    p = readerSource(bufio.NewReader(file))
    count := 0
    for v := range p {
        fmt.Println(v)
        count++
        if count >= 100 {
            break
        }
    }
}

func readerSource(reader io.Reader) <-chan int {
    out := make(chan int)
    go func() {
        buffer := make([]byte, 8) // int: 64bit -> 8byte
        for {
            n, err := reader.Read(buffer)
            if n > 0 { // 可能数据不足 8byte
                v := int(binary.BigEndian.Uint64(buffer)) // Uint64 -> int
                out <- v
            }
            if err != nil {
                break
            }
        }
        close(out)
    }()
    return out
}

func writerSink(writer io.Writer, in <-chan int) {
    for v := range in {
        buffer := make([]byte, 8)
        binary.BigEndian.PutUint64(buffer, uint64(v))
        writer.Write(buffer)
    }
}

func randomSource(count int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 0; i < count; i++ {
            out <- rand.Int()
        }
        close(out)
    }()
    return out
}
```

## go 实现完整外部排序

先来看看完整外部排序的设计图:

![图解: 外部排序](https://img4.mukewang.com/5a4213620001374419201080.jpg)

涉及到的功能大部分在上一章都有讲到, 整体流程:

- 从文件中读取数据: 注意这里使用了 **chunk** 的设计, 将文件进行分块读取, 而且 chunkSize 的设计很巧妙, 同时支持 **全文读取** 和 **chunk读取**
- 对读取到的 chunk 的数据进行内存排序(快排)
- 通过递归, 对排序后的 chunk 进行二路归并
- 将归并后的数据写入到文件中

从协程的角度来看待整个流程:

- goroutine1 进行 chunk读取, 写入 channel
- goroutine2 进行 内存排序, 排序后数据写入 channel
- goroutine3 进行 二路归并, 归并的过程中, 数据不断写入到 channel
- goroutine4 进行 文件写入, 将 channel 中的数据写入到文件

注意这里:

- goroutine1-4 可能是多个协程, 可能某一时刻是同一个协程, go 底层会有任务队列(runq)进行协程调度
- 可以通过数据流的角度来思考这个问题: 数据是怎么在 文件/channel/协程 之间进行流转的.
- 测试很重要, 示例中就先使用了small 数据进行测试, 检查程序的正确性, 再调整到 large 数据
- 日志很重要, 可以帮助我们获取到程序的更多信息, 比如 debug/性能调优

关于性能:

- **并行** 最终受限于 cpu 核数, 即 N 核cpu最多同时运行 N 个线程
- 协程间的抢占会带来性能损耗, 同理还有 进程/线程 的调度
- 协程+channel的机制方便并发编程扩展, 相对于单机内存操作自然性能要低一些

```go
package main

import (
    "io"
    "encoding/binary"
    "os"
    "bufio"
    "sort"
    "fmt"
    "time"
)

var startTime time.Time

func main() {
    fileIn := "small.in"
    fileOut := "small.out"
    p := createPipeline(fileIn, 512, 4) // 按照cpu核数设置节点数, 减少协程间抢占带来性能损耗
    writeToFile(p, fileOut)
    printFile(fileOut, -1)

    startTime = time.Now() // 添加日志
    fileIn = "large.in"
    fileOut = "large.out"
    p = createPipeline(fileIn, 800000000, 4)
    writeToFile(p, fileOut)
    printFile(fileOut, 100)
}

func createPipeline(filename string, fileSize, chunkCount int) <-chan int {
    chunkSize := fileSize / chunkCount // fileSize/8/chunkCount = int/chunk, 这里简单处理, 设置为可以整除的参数
    sortResults := []<-chan int{}      // 传递给 mergeN() 的已排序切片
    for i := 0; i < chunkCount; i++ {
        file, err := os.Open(filename) // 为什么没有用 defer file.close() ? 因为需要在函数外去关闭掉, 比较麻烦, 这里暂时省略
        if err != nil {
            panic(err)
        }
        file.Seek(int64(i*chunkSize), 0) // 定位到每个 chunk 的起始位置
        s := readerChunk(bufio.NewReader(file), chunkSize)
        sortResults = append(sortResults, memSort(s))
    }
    return mergeN(sortResults...)
}

func writeToFile(ch <-chan int, filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    writer := bufio.NewWriter(file)
    defer writer.Flush() // defer 是 LIFO

    for v := range ch {
        buffer := make([]byte, 8)
        binary.BigEndian.PutUint64(buffer, uint64(v))
        writer.Write(buffer)
    }
}

func printFile(filename string, count int) {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    p := readerChunk(file, -1) // -1 的作用体现出来了, 这里就可以读取全部文件
    if count == -1 {
        for v := range p {
            fmt.Println(v)
        }
    } else {
        n := 0
        for v := range p {
            fmt.Println(v)
            n++
            if n >= count {
                break
            }
        }
    }
}

// 递归解决两两归并
func mergeN(ins ...<-chan int) <-chan int {
    if len(ins) == 1 {
        return ins[0]
    }
    m := len(ins) / 2
    // ins[0..m) + ins[m..end)
    return merge(mergeN(ins[:m]...),
        mergeN(ins[m:]...))
}

func merge(in1, in2 <-chan int) <-chan int {
    out := make(chan int, 1024) // 性能优化, 给 channel 添加 buffer, 而不是收一个就发一个
    go func() {
        // 归并的过程要处理某个通道可能没有数据的情况, 代码非常值得一读
        v1, ok1 := <-in1
        v2, ok2 := <-in2
        for ok1 || ok2 {
            if !ok2 || (ok1 && v1 <= v2) {
                out <- v1
                v1, ok1 = <-in1
            } else {
                out <- v2
                v2, ok2 = <-in2
            }
        }
        close(out)
        fmt.Println("merge done: ", time.Now().Sub(startTime))
    }()
    return out
}

// 添加 chunk 来读取文件,
func readerChunk(reader io.Reader, chunkSize int) <-chan int {
    out := make(chan int, 1024) // 性能优化, 给 channel 添加 buffer, 而不是收一个就发一个
    bytesRead := 0
    go func() {
        buffer := make([]byte, 8) // int: 64bit -> 8byte
        for {
            n, err := reader.Read(buffer)
            bytesRead += n
            if n > 0 { // 可能数据不足 8byte
                v := int(binary.BigEndian.Uint64(buffer))
                out <- v
            }
            // 使用 -1 表示不添加 chunk 大小限制
            // 使用是 >=, 读取区间是 [0, chunkSize)
            if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
                break
            }
        }
        close(out)
    }()
    return out
}

func memSort(in <-chan int) <-chan int {
    out := make(chan int, 1024) // 性能优化, 给 channel 添加 buffer, 而不是收一个就发一个
    go func() {
        // read into memory
        a := []int{}
        for v := range in {
            a = append(a, v)
        }
        fmt.Println("read into memory: ", time.Now().Sub(startTime))
        // sort
        sort.Ints(a)
        fmt.Println("sort done: ", time.Now().Sub(startTime))
        // output
        for _, v := range a {
            out <- v
        }
        close(out)
    }()
    return out
}
```

## go 实现集群版(web版)外部排序

网络版的设计:

![图解: 网络版的修改](https://img1.mukewang.com/5a45e1650001cc8219201080.jpg)

网络版只是在完整外排序的版本上, 新增了从网络读写数据, 并相应修改 pipeline 即可

```go
package main

import (
    "net"
    "bufio"
    "encoding/binary"
    "os"
    "strconv"
    "time"
    "fmt"
    "sort"
    "io"
)

var startTime time.Time

func main() {
    startTime = time.Now()

    // 测试 net server
    //netPipeline("small.in", 512, 4) // 按照cpu核数设置节点数, 减少协程间抢占带来性能损耗
    //time.Sleep(time.Hour)

    // 测试 small
    p := netPipeline("small.in", 512, 4)
    writeToFile(p, "small.out")
    printFile("small.out", -1)

    // 测试 large
    //p := netPipeline("small.in", 512, 4)
    //writeToFile(p, "small.out")
    //printFile("small.out", -1)
}

func netPipeline(filename string, fileSize, chunkCount int) <-chan int {
    chunkSize := fileSize / chunkCount // fileSize/8/chunkCount = int/chunk, 这里简单处理, 设置为可以整除的参数
    sortAddr := []string{}
    for i := 0; i < chunkCount; i++ {
        file, err := os.Open(filename) // 为什么没有用 defer file.close() ? 因为需要在函数外去关闭掉, 比较麻烦, 这里暂时省略
        if err != nil {
            panic(err)
        }
        file.Seek(int64(i*chunkSize), 0) // 定位到每个 chunk 的起始位置
        s := readerChunk(bufio.NewReader(file), chunkSize)

        addr := ":" + strconv.Itoa(7000 + i) // 设置不同端口号来设置不同的 server
        netSink(addr, memSort(s)) // 注意 pipeline 的设计思路是建立执行流程, 真正开始执行在 pipeline 创建之后
        sortAddr = append(sortAddr, addr)
    }

    //return nil // 测试 net server

    sortResults := []<-chan int{}
    for _,addr := range sortAddr {
        sortResults = append(sortResults, netSource(addr))
    }
    return mergeN(sortResults...)
}

func writeToFile(ch <-chan int, filename string) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    writer := bufio.NewWriter(file)
    defer writer.Flush() // defer 是 LIFO

    for v := range ch {
        buffer := make([]byte, 8)
        binary.BigEndian.PutUint64(buffer, uint64(v))
        writer.Write(buffer)
    }
}

func printFile(filename string, count int) {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    p := readerChunk(file, -1) // -1 的作用体现出来了, 这里就可以读取全部文件
    if count == -1 {
        for v := range p {
            fmt.Println(v)
        }
    } else {
        n := 0
        for v := range p {
            fmt.Println(v)
            n++
            if n >= count {
                break
            }
        }
    }
}

func netSink(addr string, in <-chan int) {
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        panic(err)
    }
    go func() {
        defer listener.Close()
        conn, err := listener.Accept() // 通常 accept() 要放到 for{} 中不断的接收请求, 这里只处理一次就关闭了
        if err != nil {
            panic(err)
        }
        defer conn.Close()
        writer := bufio.NewWriter(conn)
        defer writer.Flush() // 别忘了 flush buffer
        for v := range in {
            buffer := make([]byte, 8)
            binary.BigEndian.PutUint64(buffer, uint64(v))
            writer.Write(buffer)
        }
    }()
}

func netSource(addr string) <-chan int {
    out := make(chan int)
    go func() {
        conn, err := net.Dial("tcp", addr)
        if err != nil {
            panic(err)
        }
        defer conn.Close()
        r := readerChunk(bufio.NewReader(conn), -1)
        for v := range r {
            out <- v
        }
        close(out)
    }()
    return out
}

// 递归解决两两归并
func mergeN(ins ...<-chan int) <-chan int {
    if len(ins) == 1 {
        return ins[0]
    }
    m := len(ins) / 2
    // ins[0..m) + ins[m..end)
    return merge(mergeN(ins[:m]...),
        mergeN(ins[m:]...))
}

func merge(in1, in2 <-chan int) <-chan int {
    out := make(chan int, 1024)
    go func() {
        // 归并的过程要处理某个通道可能没有数据的情况, 代码非常值得一读
        v1, ok1 := <-in1
        v2, ok2 := <-in2
        for ok1 || ok2 {
            if !ok2 || (ok1 && v1 <= v2) {
                out <- v1
                v1, ok1 = <-in1
            } else {
                out <- v2
                v2, ok2 = <-in2
            }
        }
        close(out)
        fmt.Println("merge done: ", time.Now().Sub(startTime))
    }()
    return out
}

// 添加 chunk 来读取文件,
func readerChunk(reader io.Reader, chunkSize int) <-chan int {
    out := make(chan int, 1024) // 性能优化, 给 channel 添加 buffer, 而不是收一个就发一个
    bytesRead := 0
    go func() {
        buffer := make([]byte, 8) // int: 64bit -> 8byte
        for {
            n, err := reader.Read(buffer)
            bytesRead += n
            if n > 0 { // 可能数据不足 8byte
                v := int(binary.BigEndian.Uint64(buffer))
                out <- v
            }
            // 使用 -1 表示不添加 chunk 大小限制
            // 使用是 >=, 读取区间是 [0, chunkSize)
            if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
                break
            }
        }
        close(out)
    }()
    return out
}

func memSort(in <-chan int) <-chan int {
    out := make(chan int, 1024)
    go func() {
        // read into memory
        a := []int{}
        for v := range in {
            a = append(a, v)
        }
        fmt.Println("read into memory: ", time.Now().Sub(startTime))
        // sort
        sort.Ints(a)
        fmt.Println("sort done: ", time.Now().Sub(startTime))
        // output
        for _, v := range a {
            out <- v
        }
        close(out)
    }()
    return out
}
```

## 写在最后

go 的「强制」在编程方面感觉优点大于缺点:

- 强制代码风格: 读/写代码都轻松了不少
- 强制类型检查: 出错时的错误提示非常友好

书写过程中, 基本根据编译器提示, 就可以把大部分 bug 清理掉.

go语言三大特色:

- 面向接口, 比如示例中的 Reader/Writer, 从而可以轻松添加 buffer 进行性能优化
- 函数式, go语言中函数式一等公民
- 并发编程: go + channel

再次推荐一下 go, 给想要写 **并发编程** 的程序汪, 就如 ccmouse大大的教程所说:

> 感受并发编程的乐趣

资源推荐:

- 首推 ccmouse大大的视频教程: [慕课网 - 搭建并行处理管道，感受GO语言魅力](https://www.imooc.com/learn/927)
- 同时推荐看完一本书 [go并发编程实战ed2](http://www.ituring.com.cn/book/1950): 讲解并发编程相关知识和 go并发编程原理 上面非常透彻, 几个实战的项目也适合 ccmouse大大推荐的学习方式 -- **不那么简单的项目来练手**
