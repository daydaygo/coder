# architect 架构

- 架构=问题+技术组合?=软件设计=tradeOff 高可用=多节点+LB+服务化 成长路线图=技能图谱=技术栈
- 微服务/云原生/devops
- LB [LVS负载均衡 LVS简介、三种工作模式、十种调度算法](https://blog.csdn.net/weixin_40470303/article/details/80541639) eagerLoad
- zookeeper: provider(container) consumer registry monitor leader选举机制 watch机制 数据模型 应用场景
- 计算机思维
  - 复用; 黑箱思维; 编程思维; 算法思维; 互联网思维; 生命周期=运行时序图 问题导向(应用场景和解决方式) 基础概念和架构设计; 分布式思维
  - 分层(tcp/ip 4层) 分类(分治的算法思想) 类比 面向过程-生命周期(浏览器地址栏输入url)
  - 面向对象-概念/整体/分类 函数式-幂等
  - 组件化(依赖管理/代码复用)
  - 御术底层OS: static/yushuos.webp
  - 框架思维=套路: 更高更共性的抽象, 比如大部分算法技巧本质上都是树的遍历问题

## 分布式

- version=发展=演进: 单机扩展->大型分布式 单体=巨石->soa->微服务 高可用/可伸缩/高性能
- 理论: CAP定理(强一致性) base定理(最终一致性) -> 共识算法: raft paxos zap
- cdn静态文件 sso
- cache
- se搜索引擎: elasticSearch(高级查询)
- transaction: 2PC两阶段提交 3pc=柔性事务 TCC补偿事务
- dispatch调度 schedule计划 task: quartz elasticJob crond/crontab xxl-job
- lock共享资源的互斥访问: redis zk consulKV consul信号量; 全局锁; 超时清理
- id: snowflake 美团leaf
- trackingf链路监控: GoogleDapper tree/span/annotation reqid sampler抽样 collector收集器
- logMonitor: ELK(elastic logstash kibana)
- 动态扩容 容灾

## microService 微服务

- 分封制 中台
- 设计: 小即是美 单一职责 尽可能早创建原型 可移植性比效率更重要
- 原则: 单一职责 自治 轻量通信 粒度
- component
  - framework: hyperf spring gokit
  - service istio
    - registerDiscovery 服务注册发现: consul etcd zk eureka nacos
    - flow/traffic流量: rateLimit限流 重试 circuitBreaker熔断(hystrix(dashboard turbine聚合) sentinel)=fallback降级=服务雪崩
    - client: lb(ribbon) rpc(jet) http(feign guzzle)
  - rpc: grpc thift dubbo json
  - net: tcp/udp websocket(socket.io)
  - gateway=对外+路由+鉴权+并发控制+LB+apidoc: kong zuul(信号量) nginx
  - config: apollo nacos acm; 敏感信息
  - mq消息队列: bus消息总线 stream消息驱动 EDA事件驱动 rocketmq ampq=rabbitmq nats nsq kafka activeMQ
  - Distributed分布式: transaction(TCC seata tx-lcn) job(schedulerX) id(snowflake) lock(恶汉/懒汉/队列 zk)
  - tracing: opentracing jaeger zipkin skyWalking pinpoint sleuth
  - log: elk sls solr es fluentd
  - metric/monitor: prometheus
  - tool: grafana oss sms admin/dashboard 中间件
- spring: feign config sleuth admin管控 actuator监控 turbine消息聚合
  - netfix: eureka hystrix zuul ribbon
- alibaba: sentinel nacos rocketMQ dubbo seata acm oss schedulerX sms
- mark: [go微服务38讲](../blog/go_ms_38.md) [Microservices](https://martinfowler.com/articles/microservices.html)

## cloudnative 云原生

## devops

- ci
- cd: 灰度=金丝雀 蓝绿=滚动

## 技能图谱

- 工程
  - ds&algo dp
  - pl: web server cli; tool debug test
  - ops: monitor(sys log flow流量) alert log
  - app
    - accessLayer
    - framework: spring dubbo thrift
    - middleware中间件 msgBus消息总线 lib第三方库
    - cache nosql db
    - hardware: cpu mem disk net
    - os-linux
- sys系统
  - basic: ext扩展/伸缩 available可用 Reliable可靠 Consistent一致 LB/过载保护/灾备
  - protocol accessLayer
  - logicLayer: 连接池 串行化 masterSlave batchWrite configCenter 去中心化
    - 通讯: sync/async mq cron rpc
    - db: cache(高可用 允许cacheMiss) DAO&ORM 双主 masterSlave 读写分离
- performance性能优化
  - code: 关联代码优化 cache对齐 分支预测 copyOnWrite 内联优化
  - tool
  - sys: cache 延迟计算 数据预读 async/并发 轮询&通知 内存池 模块化
- soft软技能
  - communicate沟通 ProblemSolving解决问题 learn Innovative创新 PM项目管理 paperReading论文阅读 summingUp总结归纳
  - doc文档=standard规范=bestPractice最佳实践?: 需求 技术升级 java ui db sys api
  - 图: classDiagram类图 sequenceDiagram时序图

## zk zookeeper

- Google.chubby开源实现=fs+notify 读.任意机器.监听器 写.达成一致后.全局有序.zxid.zkTransactionId
  - 原子广播=server间同步=zab协议=恢复模式.选主+广播模式.同步 恢复模式->leader选举(basicPaxos fastPaxos)->同步状态 leader=1台执行其他共享
  - serverStatus=look+lead+follow 宕机.半数节点
  - zxid.64bit=epoch.leader.32bit+cnt.32bit proposal.zxid
- fs 树状目录结构 1m数据上限 znode=(persistent+ephemeral)*sequential
- notify通知机制=watch事件=注册.getData.exists.getChildren+触发.create.delete.setData=本地jvm.Callback
- 命名服务(全局唯一路径->资源/服务地址) 配置管理 集群管理(机器+- master选举.最小) 分布式锁(独占 时序.最小) 队列管理(全员到达 PERSISTENT_SEQUENTIAL.最小)
- 数据复制 容错.ext.client就近访问 writeMaster写主=读写分离 writeAny写任意=延迟复制最终一致性
