# 拉勾教育 - go 微服务实战 38 讲

> <https://kaiwu.lagou.com/course/courseInfo.htm?courseId=287>

- <https://github.com/longjoy/micro-go-course>

## 重点

- docker vs k8s
  - docker 开发自己使用, 一条docker命令(短暂测试)/docker-compose(长期稳定使用) 起各种服务; 环境隔离与迁移
  - k8s 必须搭配 devops, 做到业务方务方无需感知; 大规模部署与运行

## 笔记

- 云原生=devops+CD+container+microservice 微服务 DDD serviceMesh
  - devops: 不可变基础设施(immutable infrastructure)
  - Scrum 是敏捷的一种具体实践
  - 声明式api declarative vs 过程式设计 imperative
  - [云原生12因素 12-factors heroku](https://12factor.net/)
  - go-kit 标准库 vs go-micro 可插拔RPC框架
  - DDD: 领域分层 UI/application/domain/infrastructure
  - service mesh: sidecar 边车; istio(流量管理 策略控制 可观测性 安全认证)
- go: 基础=语法+流程控制 并发 web开发 DDD+货运平台 微服务部署+容器编排+CI+自动化测试
  - MPG CSP(使用通信来共享内存)
  - <https://github.com/longjoy/micro-go-course>
  - mod=module=模块: src/package=包=目录 pkg=编译后 bin=可执行 GOPROXY
    - init-初始化 download-手动修改go.mod后更新 tidy-移除没被使用的模块
  - rpc: rpcx(原生 rpc 增强) grpc
- go-kit: transport-http/grpc endpoint-日志/限流/熔断 service-业务
  - **纯手写太啰嗦了, 需要工具支持**
- 微服务基础组件(服务注册发现 RPC调用 网关 容错处理 LB auth=认证+授权 trace) 业务案例
  - 服务注册发现
    - 分布式CAP理论
    - Consul的交互图/Consul的架构图/consul etcd zookeeper 对比
  - 容错
    - 冗余-单点 熔断-雪崩
    - 限流: 拒绝/降级/优先级/延时/弹性 漏桶/令牌桶
    - 降级 无状态 幂等 超时 重试 缓存 实时监控度量
    - 熔断 Hystrix
  - LB: 随机 轮询/加权轮询 hash/一致性hash 最小连接数
  - auth: OAuth2 分布式session JWT RBAC/ACL
    - OAuth2: client -> resource-owner auth-server resource-server
    - jwt: header.payload(user/role/permiss).signature
  - tracing
    - span基本单元 trace-span集合 annotation注解-特定事件相关信息(CS/SR/SS/CR 带外数据/带内数据)
    - opentracing 规范: Jaeger(云原生)、Zipkin、SkyWalking(国产)和Pinpoint
- 经验: log err co metric/monitor

![Consul 的交互图](https://s0.lgstatic.com/i/image/M00/3E/CF/CgqCHl8tP0uAfPqfAAC1xfaVTwQ927.png)
![consul 架构图](https://s0.lgstatic.com/i/image/M00/3E/CE/CgqCHl8tP0OAVC4_AAIBrjsMQhU949.png)
![consul etcd zookeeper 对比](https://s0.lgstatic.com/i/image/M00/3E/C3/Ciqc1F8tPxCAT_4RAADBzRFlUA0352.png)
![RPC框架图](https://s0.lgstatic.com/i/image/M00/43/EE/Ciqc1F887bWAQUMOAACCOORZi64063.png)
![网关选型](https://s0.lgstatic.com/i/image/M00/4A/82/Ciqc1F9R4GCAPyX4AAFMzQx0Fy8655.png)
![hystrix原理](https://s0.lgstatic.com/i/image/M00/4F/EC/Ciqc1F9hodmASLFTAADzDRuBp1g798.png)
![tracing metric log](https://s0.lgstatic.com/i/image/M00/5E/73/Ciqc1F-GvOSAV_BRAAEKN28KEAQ070.png)
![tracing 组件对比](https://s0.lgstatic.com/i/image/M00/62/81/Ciqc1F-SjXKAKnGuAAGS0Nd0F1o697.png)

## 目录

- 开篇词 | 掌握 Go 和微服务，跟进未来服务端开发的主流趋势
- 模块一：云原生与微服务架构基础
  - 01 | 为什么说云原生重构了互联网产品开发模式？
  - 02 | 云原生基础架构的组成以及云原生应用的特征
  - 03 | 微服务架构是如何演进的？
  - 04 | DDD 领域场景分析的战略模式
  - 05 | 为什么说 Service Mesh 是下一代微服务架构？
- 模块二：Go 语言基础和示例应用
  - 06 | Go 语言开发快速回顾：语法、数据结构和流程控制
  - 07 | 如何使用 Go 更好地开发并发程序？
  - 08 | 如何基于 Go-kit 开发 Web 应用：从接口层到业务层再到数据层
  - 09 | 案例：货运平台应用的微服务划分
  - 10 | 案例：微服务 Docker 容器化部署和 Kubernetes 容器编排
  - 11 | 案例：如何结合 Jenkins 完成持续化集成和自动化测试？
- 模块三：服务注册与发现
  - 12 | 服务注册与发现如何满足服务治理？
  - 13 | 案例：如何基于 Consul 给微服务添加服务注册与发现？
  - 14 | 案例：如何在 Go-kit 和 Service Mesh 中进行服务注册与发现？
- 模块四：微服务之间的 RPC 调用
  - 15 | 微服务间如何进行远程方法调用？
  - 16 | Go RPC 如何实现服务间通信？
  - 17 | gRPC 和 Apache Thrift 之间如何进行选型？
  - 18 | 案例：Go-kit 如何集成 gRPC？
- 模块五：微服务网关
  - 19 | 微服务网关如何作为服务端统一入口点？
  - 20 | 如何进行网关选型？
  - 21 | 案例：如何使用 Kong 进行网关业务化定制？
- 模块六：微服务容错处理
  - 22 | 如何保障分布式系统的高可用性？（上）
  - 23 | 如何保障分布式系统的高可用性？（下）
  - 24 | 如何实现熔断机制？
  - 25 | 如何实现接口限流和降级？
  - 26 | 案例：如何通过 Service Mesh 实现熔断和限流？
- 模块七：负载均衡
  - 27 | 负载均衡如何提高系统可用性？
  - 28 | 案例：如何在 Go 微服务中实现负载均衡？
    - 模块八：统一认证与授权
  - 29 | 统一认证与授权如何保障服务安全？
  - 30 | 如何设计基于 OAuth2 和 JWT 的认证与授权服务体系?
  - 31 | 案例：如何自定义授权服务器？
  - 32 | 案例：如何保证微服务实例资源安全？
- 模块九：分布式链路追踪
  - 33 | 如何追踪分布式系统调用链路的问题？
  - 34 | OpenTracing 规范介绍与分布式链路追踪组件选型
  - 35 | 案例：如何在微服务中集成 Zipkin 组件？
- 模块十：Go 微服务开发的其他要点
  - 36 | 如何使用 ELK 进行日志采集以及统一处理？
  - 37 | 如何处理 Go 错误异常与并发陷阱？
  - 38 | 案例：如何使用 Prometheus 和 Grafana 监控预警服务集群？
- 结束语 | 云原生不是服务端架构的终点
