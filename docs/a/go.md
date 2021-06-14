# go

- env
  - IDE 选择 `goland`, 推荐上手 goland 时, 用 [`IDE feature trainer`](https://plugins.jetbrains.com/plugin/8554-ide-features-trainer) 练习感受下, **神器果然是神器**
    - goland 设置: `file watch|go fmt|current file` `code style|go import`
    - generate(`⌘+n`): test benchmark interface prof
- syntax
  - Visible Encapsulation 大小写
  - comment: `Deprecated` `+build`(go doc go/build) packageComment(doc.go)
  - type
    - bool int float string complex `byte = uint8` `rune = int32`
      - json: 使用了反射; string key 直接使用, int 会被转为 string, 不可使用interface; mailru/easyjson [JSON-to-Go](https://mholt.github.io/json-to-go) msgpack=BJson
    - array(value-type) slice(pointer+len+cap扩容) map struct(`struct{}{}` 0-mem-use)
  - var const
    - 空标识符 `_` => 垃圾桶; 占位符/匿名变量; 包的初始化
  - control流程控制: for range(值拷贝) if switch(fallthrough 注意和其他语言的区别)
  - func
    - named return
    - err errors: panic recover
    - defer延迟调用: LIFO 资源释放(fd lock conn) 参数解析`defer func(){}()`; call defered but field/method immediately
    - main() -> programming; init() -> package
    - new vs make: make 只用于 slice/map/chan
    - append: cap扩容
  - method
    - composition, not inheritance; associated with user-defined type
    - receiver: 传入值如何使用 `map vs struct`
  - interface
    - 嵌入 & 反转依赖
    - 接口(interface)
      - 一种类型(type), 一种抽象类型
      - 一组方法method 的集合, 定义一个协议(规则)
      - 面向接口编程, 可以使用 ide 自动生成
      - 接口型函数
  - co并发
    - CSP通信 [MPG线程模型](https://povilasv.me/go-scheduler/)
    - go
    - chan: 无缓冲的通道只有在有人接收值的时候才能发送值 `fatal error: all goroutines are asleep - deadlock!`
    - context(requestScoped value across api/process)
    - select
    - memShared并发读写: sync(Cond Locker Map Mutex/RWMutex Once Pool WaitGroup) sync.atomic(lock-free, 减少 mutex)
- func 功能
  - fmt: v 默认格式; +v 结构体; #v go语法表示; T 类型
  - flag: 解析命令行参数
  - log -> 第三方库 logrus zap
  - IO: os(exec signal user) -> io(fs ioutil) -> bufio(Reader Writer ReaderWriter Scanner); path.filepath
  - html(escaping html text).template(against code injection)
  - text: scanner tabwriter template.parse
  - image(2d image): color.palette draw(image compose) gif jpeg png
  - mime: multipart quotedprintable
  - hash: adler32 crc32 crc64 fnv maphash
  - crypto: cipher(mode: cbc cfb) aes des dsa ecdsa ed25519 elliptic rc4 rsa hmac md5 rand sha1 sha256 sha512 subtle tls x509.pkix
  - archive: tar zip
  - compress: bzip2 flate(DEFLATE RFC1951) gzip lzw zlib
  - database.sql.driver
  - time(Time Duration Timer Since) time.tzdata
  - net(netIO tcp/ip udp dns unixDomainSocket)
    - http: cgi cookiejar fcgi httptest httptrace httputil internal pprof
    - url
    - rpc.jsonrpc
    - mail smtp textproto(text-based req/resp protocol -> http/nntp/smtp)
  - `reflect`反射: json gorm `cache[typ.Field(i).Name] = i`
  - low-level program
    - unsafe(Pointer)
    - runtime(interact with go's runtime system): cgo(runtime support for `go tool cgo`) debug metrics msan pprof race(race detection) trace
    - debug: dwarf(debug) elf(elf .o) macho(mach-o .o) plan9obj gosym(go symbol) pe
    - embed: file embedded in go program
    - expvar: /debug/vars cmdline/memstats `import _ "expvar"`
    - syscall(low-level os primitive).js(wasm)
    - plugin
  - github.com: golang(go net =golang.org/x?) pkg(errors)
    - google/protobuf
    - isdamir/gotype
      - base: LNode链表 BNode二叉树 TrieNode AVLNode stack/queue/set(使用sync.RWMutex) `container/list` priority_queue
      - util: 创建/打印链表 中序创建/打印二叉树 层序打印二叉树
- project
  - package: dir internal
  - go tool: compile fmt
    - test: ide生成 test/benchmark/example `-v -run` table-driven/rand `f(x)=y, want z`
      - benchmark: `-bench -benchmem`
      - cover `go tool cover` `-cover -coverprofile`
      - prof: `go tool pprof` `-cpuprofile -blockprofile -memprofile -run=NONE`
        - `go test -run=NONE -bench=ClientServerParallelTLS64 -cpuprofile=cpu.log net/http` `go tool pprof -text -nodecount=10 ./http.test cpu.log`
    - cmd(go tool): addr2line(GNU addr2line tool -> pprof) api asm(.go -> .o) buildid cgo(go call C code) compile(go package -> command line) dist(go distribution) fix(old -> new) go(manage go src) gofmt link(link main+dep -> exe) nm(list symbols) objdump(disassembles exe) pack(Unix ar tool) test2json trace vet(examines code)
      - doc `go doc http.ListenAndServe`
      - list `go list ...` `...xml...`
      - pprof: allocs内存分配信息 block阻塞情况信息 cmdline显示程序启动命令 groutinue当前所有协程堆栈信息 heap 堆上内存使用情况(会下载文件) mutex锁竞争情况 profilecpu占用情况(会下载文件) threadcreate系统线程创建信息 trace 程序运行跟踪信息(会下载文件)
    - go: ast(syntax tree) build(gather package info).constraint constant doc(go ast -> src doc) format(go src std fmt) importer parser printer(print ast node) scanner token(lexical token) type(type check)
    - golang.org/x: benchmarks/perf crypto/image/net/sync/sys/text/time debug/mobile/exp blog/build/pkg/tour review/tools(godoc goimports gorename)
    - code: `go fmt/vet` [lint](https://golangci-lint.run/)
      - [style](https://github.com/golang/go/wiki/CodeReviewComments)
      - [layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)
- architect
- security
- core
  - [slice](https://blog.golang.org/slices-intro)
  - [map](https://cloud.tencent.com/developer/article/1468799)
  - [mutex](https://colobu.com/2018/12/18/dive-into-sync-mutex/)
  - goroutine
    - 限制/泄露 context timeout channel select
      - <https://github.com/uber-go/goleak>
      - runtime.NumGoroutine()string
      - pprof/goroutine
  - GC
    - 减少对象分配: reuse stack preallocation
    - 降低 scan 成本
- eco
  - http://go.dev http://pkg.go.dev http://golang.org/doc [spec](https://golang.org/ref/spec) [cmd](https://golang.org/doc/cmd) [标准库](http://golang.org/pkg) [扩展库](http://golang.org/x) `go list ...`
  - go microservice framework
    - bilibili: go-kratos
    - huawei / shopee: go-chassis
    - tantan: ttdb
    - grab-kit
    - baidu: BFE 流量转发
    - didi: nightgale 夜莺运维平台
    - 趣链科技: bitxhub
    - Tencent: tarsgo
    - [tal: go-zero](https://github.com/tal-tech/go-zero)
    - qiniu: goplus
- algo
  - 双向链表: `container.list`
  - 二分查找 / 快排 qsort: `sort.Search` `sort.Sort`
  - 令牌桶: `golang.org/x/time/rate`
  - 随机数: `math/rand`
  - hash: `hash/crc32`
  - 密码: `crypto` `encoding`
    - hmac_sha256: `crypto/hmac` - `crypto/sha256` - `encoding/base64` +`encoding/hex`

``` sh
brew install go

set -gx GOPATH ~/go # GOROOT GOPATH 工作目录 src/pkg/bin
go env/get/build/run

set -gx GO111MODULE on # 必开, 包管理, 加速下载都有用到
set -gx GOPROXY 'https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct'
go mod init/download/tidy/vendor/graph/why/replace

# 单测
go test -v -run=Group # -v 详情, -run 正则匹配
# 覆盖度
go test -cover -coverprofile=c.out -covermode=atomic # 可以 >> 合并多个
go tool cover -html=c.out
# 基准测试
go test -bench=Split -benchmem # 显示使用 -bench 指定
# 性能比较函数
go test -bench=. -cpu=1 # 指定cpu数量
go test -bench=Fib40 -benchtime=20s # -benchtime 默认为 1s, 可以设置更长的时间, 从而运行更多次数, 获取更可靠的数据
# example
go test -run=Example # 配合 godoc 产生示例代码
# race
go test -race # race 检查, 出现 race 比性能问题更严重

# prof
go test -bench=. -cpuprofile cpu.prof
go tool pprof -svg cpu.prof > cpu.svg # svg 可以使用 chrome 打开, 查看热点函数, 定位性能瓶颈
go tool pprof -http :9900 http://127.0.0.1:9522/debug/pprof/heap # 内存分配信息 -> 下拉VIEW选择Flame Graph 即可看到火焰图了
go tool pprof -http :9900 http://127.0.0.1:9522/debug/pprof/profile # cpu 信息
# trace
go test -bench=. -trace trace.out
go tool trace trace.out # 自动打开浏览器, 可查看 cpu 使用情况

# goreman: 多服务管理 https://github.com/mattn/goreman
brew install goreman
# goreman 配置文件: Procfile
api: cd api; go run bookstore.go -f etc/bookstore-api.yaml
# .env: 设置环境变量
goreman start api xxx # 可启动多个
```

```go
type newInt int // 类型定义
type byte = uint8 // 类型别名
T() // 类型转换

// string byte
s:= "hello"
fmt.Println(s[0]) // 104
fmt.Println(string(s[0])) // h
s - "daydaygo" // 字符串拼接
s := "hello" // strings.Builder
var b strings.Builder
b.Grow(len(s)) // 多次拼接, 可以预先分配内存
b.WriteString(s)
b.String()

// byte 和 rune
s := "baidu百度"
for i:=0; i< len(s); i+- {
    fmt.Printf("%v(%c)", s[i], s[i])
}
fmt.Println()
for _, r := range s {
    fmt.Printf("%v(%c)", r, r)
}

// 修改字符串, 需要 2 次类型转换
s:= "hello 你好"
s2 := []rune(s)
s2[len(s2)-2] = '您'
fmt.Println(string(s2)) // hello 您好

// array
a := [...]int{1,2,3}
a := [3]int{}

// slice
s := []int{1,2,3}
s := make([]int, 0, 3) // make([]int, 3) 会导致前 3 个已赋值, append() 从第 4 个开始
s:= "hello 你好"
fmt.Println(string(s[0:2])) // [i,j)
type slice struct { // slice 底层
    array unsafe.Pointer //指向存放数据的数组指针
    len   int            //长度有多大
    cap   int            //容量有多大, cap 不够用会导致重新分配内存
}

// map
// init
m2 := map[string]int{ // var m map[int]int; m := make(map[string]int)
    "a": 1,
    "b": 2,
}
m3 := make(map[string]interface{}) // 获取 interface 的值需要注意类型转换
// check value
if v, ok := m2["c"]; ok { // 简写可以缩小变量的作用域
    fmt.Println(v)
}
// 遍历
for k, v := range m2 { // 可以只遍历 k, 也可以都忽略
    fmt.Println(k, v)
}
delete(m2, "a")

// err
fmt.Errorf() // errors.New()
defer func() { // defer 必须在 panic 前
    if err := recover(); err != nil { // recover() 必须在 defer func(){} 内
        fmt.Println(err)
    }
}()
// github.com/pkg/errors
// 同一个 struct err 简化
type Reader struct {
    r   io.Reader
    err error
}
func (r *Reader) read(data interface{}) {
    if r.err == nil {
        r.err = binary.Read(r.r, binary.BigEndian, data)
    }
}

// struct
var p Person
p2 := new(Person) // 指针: &Person
func (p *Person) PointerTest() {
  fmt.Printf("Pointer: %p\n", p)
}

// interface
v, ok := x.(string) // x.(T) 接口var的value的类型推断
v := x.(string) // 可以省略 ok
x.(type) // 配合 switch

// go
go expression // 并发执行代码
// chan
var name chan T // 双向 channel
var name chan <- T // 只能发送消息的 channel
var name T <- chan // 只能接收消息的 channel
channel <- val // 发送消息
val := <- channel // 接收消息
val, ok := <- channel // 非阻塞接收消息
ch := make(chan T, sizeOfChan)
i, ok := <-ch1 // 通道关闭后再取值ok=false
// 多路复用: chan - timeout
select {
case val := <- ch1: // 从 ch1 读取数据
  fmt.Printf("get value %d from ch1\n", val)
case ch2 <- 2 : // 使用 ch2 发送消息
  fmt.Println("send value by ch2")
case <-time.After(2 - time.Second): // 超时设置
  fmt.Println("Time out")
  return
}
// context 上下文, 一个请求的多个 go 协同工作 - 取消/超时
func testContext(ctx context.Context, t time.Duration) {
  select {
  case <-time.After(t): // 模拟耗时任务
  case <-ctx.Done(): // ctx cancel
    fmt.Println(ctx.Err()) // ctx err
    // clean
  }
}
func main() {
  ctx := context.Background()          // new empty context
  ctx = context.WithValue(ctx, "a", 1) // add context info
  ctx, cancel := context.WithTimeout(ctx, time.Second*2)
  defer cancel()
  go testContext(ctx, time.Second*4) // task1
  go testContext(ctx, time.Second*3) // task2

  time.Sleep(time.Second - 5) // let all go run
}
// sync
var lock sync.Mutex
var wg sync.WaitGroup
var x int
func atomAdd() {
    // lock.Lock()
    for i := 0; i < 5000; i+- {
        x = x+1 // 需要保证当前只有一个协程能访问变量 x
    }
    // lock.Unlock()

    wg.Done()
}
func main() {
    wg.Add(2)
    go atomAdd()
    go atomAdd()
    wg.Wait()

    fmt.Println(x)
}

// fmt

// json & tag
var l []int
json.Unmarshal([]byte(costTpl.DefaultExpress), &l) // [132,1]
type person struct {
    Name string `json: name`
    Age int
    gender uint8 // json 序列化时不可见
}
func main() {
  p := &person{
    Name: "dayday",
    Age: 18,
    gender: 1,
  }
  data, _ := json.Marshal(p)
  fmt.Printf("%#v \n", data)
}
// protobuf
proto.Marshal()

// reflect
reflect.DeepEqual() // 深度比较, 单测中会使用
reflect.Indirect(reflect.ValueOf(dest)).Type() // ORM 中大量使用

// signal
errc := make(chan error)
go func() {
    c := make(chan os.Signal)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    errc <- fmt.Errorf("%s", <-c)
}()
```
