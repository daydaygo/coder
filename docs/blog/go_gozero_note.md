# go-zero

- <https://github.com/tal-tech/go-zero>
- <https://www.yuque.com/tal-tech/go-zero>
- 微服务组件/工具: <https://github.com/zeromicro>

- goctl
  - [api](https://github.com/tal-tech/zero-doc/blob/main/doc/goctl.md)
  - [rpc](https://github.com/tal-tech/zero-doc/blob/main/doc/goctl-rpc.md)

- api gateway
  - **业务逻辑**
  - 鉴权 加解密
  - 日志 异常 监控报警 链路跟踪 数据统计 并发控制 超时 熔断 降级
- rpc
  - **业务逻辑**
  - 鉴权
  - 缓存
  - 日志 异常 监控报警 链路跟踪 数据统计 并发控制 超时 熔断 降级
- model
  - crud
  - 缓存: 穿透 击穿 雪崩 索引
  - 日志 慢查 熔断 连接

- doc
  - bloom
  - executors 任务池-批处理任务
  - fx stream-api 流式处理 -> java stream api
    - from() 数据流
    - map() filter() merge() ...
    - reduce() parallel() ...
  - store
    - sqlx 通用; sqlc 带缓存
    - redis-lock
  - limit: 滑动窗口(频率/并发上限) tokenBucket(瞬时流量)
  - rest: jwt breaker fingerprint/sign trace(span/spancontext)
  - zrpc: interceptor lb(p2c)
  - timeWheel
