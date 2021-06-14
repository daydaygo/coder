# go| go并发实战: 搭配 influxdb + grafana 实现实时日志监控系统

- [go| go并发实战: 搭配 influxdb + grafana 高性能实时日志监控系统](https://www.jianshu.com/p/f4d2b2ebafea)

> [Go并发编程案例解析](https://www.imooc.com/learn/982)

继续好玩的并发编程实战, 上一篇 [go| 感受并发编程的乐趣 前篇](https://www.jianshu.com/p/bd890e0f81f1).

实战内容: 实时处理读取/解析日志文件, 搭配 influxdb(时序数据库) 存储, grafana 展示, 并提供系统的简单监控.

# 0x00: 初始化, 面向过程编程

用面向过程的方式, 对问题进行简单的梳理, 代码如下:

```go
package main

func main() {
    // read log file

    // process log

    // write data
}
```

这里并没有写具体的实现, 因为到这里, 我们就可以开始考虑 **封装** 了

# 0x01: 过程封装, 使用 LogPorcess 结构体

引入 LogProcess 结构体, 将整个任务 **面向对象** 化, 伪代码如下:

```go
package main

import (
    "fmt"
    "strings"
)

type LogProcess struct {
    path string // 日志文件路径
    dsn string // influxdb dsn
}

func (lp *LogProcess) Read() {
    path := lp.path
    fmt.Println(path)
}

func (lp *LogProcess) Process() {
    log := "hello world"
    fmt.Println(strings.ToUpper(log))
}

func (lp *LogProcess) Write()  {
    dsn := lp.dsn
    fmt.Println(dsn)
}

func main() {
    lp := &LogProcess{
        path: "test path",
        dsn: "test dsn",
    }

    // read log file
    lp.Read()

    // process log
    lp.Process()

    // write data
    lp.Write()
}
```

# 0x02: 加上 go 和 chan, 并发就是如此简单

加上 go 关键字, 轻松实现协程:

```go
func main() {
    lp := &LogProcess{
        path: "test path",
        dsn: "test dsn",
    }

    // read log file
    go lp.Read()

    // process log
    go lp.Process()

    // write data
    go lp.Write()

    time.Sleep(time.Second) // 新手必知: 保证程序退出前, 协程可以执行完
}
```

加上 chan, 轻松实现协程间通信:

```go
type LogProcess struct {
    path string // 日志文件路径
    dsn string // influxdb dsn
    rc chan string // read chan
    wc chan string // write chan
}

func (lp *LogProcess) Read() {
    path := lp.path
    fmt.Println(path)

    lp.rc <- "test data"
}

func (lp *LogProcess) Process() {
    log := <- lp.rc
    lp.wc <- strings.ToUpper(log)
}

func (lp *LogProcess) Write()  {
    dsn := lp.dsn
    fmt.Println(dsn)

    data := <- lp.wc
    fmt.Println(data)
}
```

# 0x03: 引入 interface, 方便以后扩展

现在是从 **文件** 读取, 如果以后要从 **其他数据源** 读取呢? 这个时候就可以用上接口:

```go
type Reader interface {
    Read(rc chan string)
}

type ReadFromFile struct {
    path string
}

func (r *ReadFromFile) Read(rc chan string) {
    // read from file
}
```

同理, 数据写入到 **influxdb** 也可以加入接口, 方便以后扩展.

# 0x04: 读取文件的细节

实时读取日志文件要怎么实现呢? 直接上代码, 细节有很多, 注意 **注释**:

- **实时** 读取怎么实现: 从文件末尾开始读取
- 怎么一行一行的读取日志: `buf.ReadBytes('\n')`
- 输出怎么多了换行呢: 截取掉最后的换行符 `line[:len(line)-1]`

```go
func (r *ReadFromFile) Read(rc chan []byte) {
    f, err := os.Open(r.path)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    f.Seek(0, 2) // 文件末尾
    buf := bufio.NewReader(f) // []byte 数据类型, rc chan 的类型也相应进行了修改

    for {
        line, err := buf.ReadBytes('\n')
        // todo: 处理日志切割, inode 变化的情况
        if err == io.EOF {
            time.Sleep(500 * time.Millisecond)
        } else if err != nil {
            panic(err)
        } else { // 需要写到这里
            rc <- line[:len(line)-1]
        }
    }
}
```

还有一个需要优化的地方, 一般日志文件都会采取 **轮转** 策略(详见上篇blog [devops| 日志服务实践](https://www.jianshu.com/p/9dae3ba679e6)), 文件可能更新了, 所以读取文件时, 还需要加一个判断.

# 0x05: 日志解析, 又见正则

日志的解析比较简单, 按照日志的格式正则匹配即可:

```go
// 使用结构体来记录匹配到的日志数据
type Log struct {
    TimeLocal                    time.Time
    BytesSent                    int
    Path, Method, Scheme, Status string
    UpstreamTime, RequestTime    float64
}

func (l *LogProcess) Process() {
    // 正则
    re := regexp.MustCompile(`([\d\.]+)\s+([^ \[]+)\s+([^ \[]+)\s+\[([^\]]+)\]\s+([a-z]+)\s+\"([^"]+)\"\s+(\d{3})\s+(
\d+)\s+\"([^"]+)\"\s+\"(.*?)\"\s+\"([\d\.-]+)\"\s+([\d\.-]+)\s+([d\.-]+)`)

    loc, _ := time.LoadLocation("PRC")
    for v := range l.rc {
        str := string(v)
        ret := re.FindStringSubmatch(str)
        if len(ret) != 14 {
            log.Println(str)
            continue
        }

        msg := &Log{}
        t, err := time.ParseInLocation("02/Jan/2006:15:04:05 +0000", ret[4], loc)
        if err != nil {
            log.Println(ret[4])
        }
        msg.TimeLocal = t

        byteSent, _ := strconv.Atoi(ret[8])
        msg.BytesSent = byteSent

        // Get /for?query=t HTTP/1.0
        reqSli := strings.Split(ret[6], " ")
        if len(reqSli) != 3 {
            log.Println(ret[6])
            continue
        }
        msg.Method = reqSli[0]
        msg.Scheme = reqSli[2]
        // url parse
        u, err := url.Parse(reqSli[1])
        if err != nil {
            log.Println(reqSli[1])
            continue
        }
        msg.Path = u.Path
        msg.Status = ret[7]
        upTime, _ := strconv.ParseFloat(ret[12], 64)
        reqTime, _ := strconv.ParseFloat(ret[13], 64)
        msg.UpstreamTime = upTime
        msg.RequestTime = reqTime

        l.wc <- msg
    }
}
```

# 0x06: 上手 influxdb

influxdb 是时序数据库的一种, 包含如下基础概念:

- database: 数据库
- measurement: 数据库中的表
- points: 表里的一行数据

其中 points 包含以下内容:

- tags: 有索引的属性
- fields: 值
- time: 时间戳, 也是自动生成的主索引

使用 docker 快速开启 InfluxDb Server:

```yaml
    influxdb:
        image: influxdb:1.4.3-alpine
        ports:
            - "8086:8086"
        #     - "8083:8083" # admin
        #     - "2003:2003" # graphite
        environment:
            INFLUXDB_DB: log
            INFLUXDB_USER: log
            INFLUXDB_USER_PASSWORD: logpass
        #     INFLUXDB_GRAPHITE_ENABLED: 1
        #     INFLUXDB_ADMIN_ENABLED: 1
        # volumes:
        #     - ./data/influxdb:/var/lib/influxdb
```

influxdb 使用 go 语言实现, 稍微修改一下官方文档中示例, 就可以使用 client:

> InfluxDB Client: https://github.com/influxdata/influxdb/tree/master/client

```go
// 写入也使用接口
type Writer interface {
    Write(wc chan *Log)
}

type WriteToInfluxdb struct {
    dsn string
}

// 只在官方示例代码上做了一点修改
func (w *WriteToInfluxdb) Write(wc chan *Log) {
    // dsn 示例: http://localhost:8086@log@logpass@log@s
    dsnSli := strings.Split(w.dsn, "@")

    // Create a new HTTPClient
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr:     dsnSli[0],
        Username: dsnSli[1],
        Password: dsnSli[2],
    })
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    // Create a new point batch
    bp, err := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  dsnSli[3],
        Precision: dsnSli[4],
    })
    if err != nil {
        log.Fatal(err)
    }

    for v := range wc {
        // Create a point and add to batch
        tags := map[string]string{
            "Path": v.Path,
            "Method": v.Method,
            "Scheme": v.Scheme,
            "Status": v.Status,
        }
        fields := map[string]interface{}{
            "bytesSent":   v.BytesSent,
            "upstreamTime": v.UpstreamTime,
            "RequestTime":   v.RequestTime,
        }

        pt, err := client.NewPoint("log", tags, fields, v.TimeLocal)
        if err != nil {
            log.Fatal(err)
        }
        bp.AddPoint(pt)

        // Write the batch
        if err := c.Write(bp); err != nil {
            log.Fatal(err)
        }

        // Close client resources
        if err := c.Close(); err != nil {
            log.Fatal(err)
        }
    }
}
```

# 0x07: 使用 Grafana 接入 InfluxDB 数据源

Grafana 使用 docker 也可以轻松部署:

```yaml
    grafana:
        image: grafana/grafana:5.1.0-beta1
        ports:
            - "3000:3000"
        environment:
            GF_SERVER_ROOT_URL: http://grafana.server.name
            GF_SECURITY_ADMIN_PASSWORD: secret
```

官网效果图:

todo

# 0x08: 简单监控系统实现

作为一个 **实时** 系统, 需要后台常驻运行, 怎么查看系统的运行状态的呢?

加入一个简单的监控系统, 通过 http 请求查看系统实时运行状态:

```go
// 需要监控的系统状态
type SystemInfo struct {
    LogLine int `json:"logline"` // 总日志处理数
    Tps float64 `json:"tps"`
    ReadChanLen int `json:"readchanlen"` // read chan 长度
    WriteChanLen int `json:"writechanlen"` // write chan 长度
    RunTime string `json:"runtime"` // 运行总时间
    ErrNum int `json:"errnum"` // 错误数
}

// 监控类
type Monitor struct {
    startTime time.Time
    data SystemInfo
}

// 启动监控, 其实就是一个简单的 http server
func (m *Monitor) start(lp *LogProcess) {
    http.HandleFunc("/monitor", func(writer http.ResponseWriter, request *http.Request) {
        m.data.RunTime = time.Now().Sub(m.startTime).String()
        m.data.ReadChanLen = len(lp.rc)
        m.data.WriteChanLen = len(lp.wc)

        ret, _ := json.MarshalIndent(m.data, "", "\t")

        io.WriteString(writer, string(ret))
    })

    http.ListenAndServe(":9091", nil)
}

func main() {
    ...

    // 运行监控
    m := &Monitor{
        startTime: time.Now(),
        data: SystemInfo{},
    }
    m.start(l)
}
```

监控数据中的 TPS 稍微有点难处理:

- 启动一个定时器, 比如 5s
- 记录下时间间隔内的 `LogLine`(日志处理行数)

这样我们就可以用 `LogLine` 来估算系统的 TPS 了

# 0x09: 写在最后

并发编程实战, 总会给人带来 **又完成了了不起的任务** 的感觉, 特别是会了解更多的细节.

> 能够相遇, 也是一种快乐吧.

完整代码: https://gitee.com/daydaygo/codes/sc6tyr2odf58k39npub0v72
