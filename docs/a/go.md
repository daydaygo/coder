# go

- env
  - IDE 选择 `goland`, 推荐上手 goland 时, 用 [`IDE feature trainer`](https://plugins.jetbrains.com/plugin/8554-ide-features-trainer) 练习感受下, **神器果然是神器**
    - goland 设置: `file watch|go fmt|current file` `code style|go import`
    - generate(`⌘+n`): test benchmark interface prof
  - [code example](https://github.com/daydaygo/coder/tree/main/src/go)
- syntax
  - Visible Encapsulation 大小写
  - comment: `Deprecated` `+build`(go doc go/build) packageComment(doc.go)
  - type
    - baisc: bool string number(int float string complex `byte = uint8` `rune = int32`)
    - aggregate: array(value-type) struct(`struct{}{}` 0-mem-use)
    - ref: slice(pointer+len+cap扩容) map ptr chan func
    - interface
    - declare alias trans`T(x)` asset`x.(T)` `switch x.(type)`
  - var const `_`
  - control流程控制: for range(值拷贝) if switch(fallthrough 注意和其他语言的区别)
  - func: `func value with state` `_ init()` `new()` `make() slice/map/chan` `append() cap扩容`
  - method: composition, not inheritance; associated with user-defined type
  - interface
  - co并发
    - CSP通信 [MPG线程模型](https://povilasv.me/go-scheduler/) go chan.buffered select.multiplex cancel.broadcast.close context(requestScoped value across api/process)
    - shm: raceCondition raceDetector ync(Cond Locker Map Mutex/RWMutex Once Pool WaitGroup) sync.atomic(lock-free, 减少 mutex)
- func
  - builtin: math(rand) container(list) sort fmt(v默认 +v结构体 #v语法表示 T类型) flag log(logrus zap) io(os->io->bufio) path text html image mime hash crypto archive compress datebase time net reflect unsafe runtime
  - github.com: golang(go net =golang.org/x?) pkg(errors)
    - google/protobuf
    - uber-go/goleak
    - mattn/goreman 多服务管理 `brew install goreman` `Procfile .env` `goreman start xxx`
    - isdamir/gotype
      - base: LNode链表 BNode二叉树 TrieNode AVLNode stack/queue/set(使用sync.RWMutex) `container/list` priority_queue
      - util: 创建/打印链表 中序创建/打印二叉树 层序打印二叉树
- project
  - package: dir internal
  - go tool
    - code: `go fmt/vet` [lint](https://golangci-lint.run/)
      - [style](https://github.com/golang/go/wiki/CodeReviewComments)
      - [layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)
    - test: ide生成 test/benchmark/example/race `-v -run -race` table-driven/rand `f(x)=y, want z`
      - benchmark: `-bench -benchmem -cpu=1 -benchtime=1s`
      - cover `go tool cover` `-cover -coverprofile`
      - prof: `go tool pprof` `-cpuprofile -blockprofile -memprofile -run=NONE`
        - `go test -run=NONE -bench=ClientServerParallelTLS64 -cpuprofile=cpu.log net/http` `go tool pprof -text -nodecount=10 ./http.test cpu.log`
        - pprof: allocs内存分配信息 block阻塞情况信息 cmdline显示程序启动命令 groutinue当前所有协程堆栈信息 heap 堆上内存使用情况(会下载文件) mutex锁竞争情况 profilecpu占用情况(会下载文件) threadcreate系统线程创建信息 trace 程序运行跟踪信息(会下载文件)
      - tarce: `go test -bench=. -trace trace.out` `go tool trace trace.out`
    - doc `go doc http.ListenAndServe`
    - list `go list ...` `...xml...` cmd/go
    - cgo `go tool cgo`
    - golang.org/x: benchmarks/perf crypto/image/net/sync/sys/text time(rate 令牌桶) debug/mobile/exp blog/build/pkg/tour review/tools(godoc goimports gorename)
- eco
  - core architect security
    - GC: 减少对象分配 reuse stack preallocation; 降低 scan 成本
  - http://go.dev http://pkg.go.dev [扩展库](http://golang.org/x)
  - [doc](http://golang.org/doc): EffectiveGo pkg cmd spec MemoryModel.sync releaseHistory
    - [slice](https://blog.golang.org/slices-intro)
    - [map](https://cloud.tencent.com/developer/article/1468799)
    - [mutex](https://colobu.com/2018/12/18/dive-into-sync-mutex/)
  - github.com
    - tal-tech/go-zero go-kratos go-chassis grab-kit tarsgo ttdb nightgale bitxhub goplus

## code

- basic: demo co(go&sync) net(tcp&http) grpc
- algo
- pref性能优化

## 资料

- qtt go实战培训: 第一课(slice array string map sync.map mem) 并发(GMP chan atom&lock context 并发模式) 实战(mem http调优 设计模式 qodis) 实战2(pprof 泛型) 微服务(rpc http框架 优雅重启 业务踩坑)
