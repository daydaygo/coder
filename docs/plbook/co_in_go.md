# book| Concurrency in Go: tools and techniques for developers-O'Reilly Media (2017)

- go语言并发之道 https://github.com/kat-co/concurrency-in-go-src

- sytax: go sync(waitGroup mutex rwmutex cond once pool) chan select GOMAXPROCS
- co pattern
- scale
- go runtime

- raceCondition: `must exec in correct order` dataRace(同一个数据上同时进行读写) raceDetect
- atom原子性: op + context/scope
- mem.sync -> dataRace? `sync.Mutex`
- deadlock: Coffman Conditions
- livelock: 死循环
- Starvation: cannot get all the resources it needs to perform work
- CSP CommunicatingSequentialProcesses https://dl.acm.org/doi/10.1145/359576.359585: Do not communicate by sharing memory. Instead, share memory by communicating
- Concurrency Parallelism: Concurrency is a property of the code; parallelism is a property of the running program.
- Amdahl law 阿姆达尔/安达尔: $S=\dfrac{1}{1-a+a/n}$ 并行存储的性能分析
- https://github.com/golang/go/wiki/MutexOrChannel: mutex=cache/state
- chan: op-read/write/close state-nil/full-open/close
- pattern
  - for-select
  - goLeak: If a goroutine is responsible for creating a goroutine, it is also responsible for ensur‐ ing it can stop the goroutine.
  - or-chan
  - err-handle
  - pipeline generator
  - fan-out fan-in
  - or-done-chan
  - bridge
  - queue
  - context
