# hyperf| 微服务之旅: 开篇

这里是 `hyperf 微服务之旅` 系列 blog 的开篇, 也是系列 blog 的导航, 微服务知识的总览, 持续更新中~

- [PHP 网络编程基础 - swoole](https://wiki.swoole.com)
- PHP 微服务框架
  - [hyperf](http://doc.hyperf.io)
- 微服务基础设施
  - 服务治理
  - consul
  - zookeeper
  - 配置中心: 简单配置(账号密码) 集群配置(不同环境) 开关控制(修改不用发版)
  - [Apollo](https://github.com/ctripcorp/apollo)
  - [qconf](https://github.com/Qihoo360/QConf)
  - aliyun acm - 应用配置管理
- 云平台基础设施
  - [cncf](https://www.cncf.io/)
  - aliyun k8s  - 容器服务k8s: 云原生
  - aliyun sls - 日志服务
  - aliyun pts - 性能测试
- php 基础
  - [psr](php_psr.md)
  - 框架: [yii](learn_yii.md)

参考一览:

- [swoole| swoole wiki 笔记](https://www.jianshu.com/p/12d645ac02b2)
- [tech| PTS 性能测试一览](https://www.jianshu.com/p/d31155c404ff)

blog 导航:

- [php| php 微服务之旅: devops](https://www.jianshu.com/p/10796001bf39)
- [php| php 微服务之旅: 配置中心](https://www.jianshu.com/p/9cb3fca076c1)

## todo

腾讯云TSF: <https://cloud.tencent.com/document/product/649>
tars: <https://github.com/TarsCloud>
阿里云EDAS: <https://help.aliyun.com/product/29500.html>

- 中间件(分布式事务)<https://www.aliyun.com/aliware>
- 分布式事务 <https://help.aliyun.com/product/48444.html>
- 使用mq避免分布式事务 <https://yq.aliyun.com/articles/10>
ali SOFA mesh <https://mp.weixin.qq.com/s?__biz=MzA3MDg5MTEzMA==&mid=500751401&idx=1&sn=42e3773740d8ab01eccb92af342e1fd4>
spring干货汇总 <https://mp.weixin.qq.com/s?__biz=MzAxODcyNjEzNQ==&mid=2247484574&idx=1&sn=0984db0da3dc0efda956fa0aaeabe479>
翟永超 <http://blog.didispace.com/>
spring全家桶: <https://gitee.com/yidao620/springboot-bucket>
微服务2.0技术栈选型手册 <https://mp.weixin.qq.com/s/OloZhn2pwfIrOQit_8jefA>

- <https://github.com/micro/go-micro>

微服务的10个挑战和解决方案
1.数据同步  – 我们使用事件源代码架构来使用异步消息传递平台解决此问题。传奇设计模式可以应对这一挑战。
2.安全性  – API网关可以解决这些挑战。Kong非常受欢迎，并且是开源的，并且正在被许多公司用于生产。还可以使用JWT令牌，Spring Security和Netflix Zuul / Zuul2为API安全性开发自定义解决方案。还有企业解决方案，如Apigee和Okta（两步认证）。Openshift用于公共云安全的顶级功能，如基于Red Hat Linux Kernel的安全性和基于命名空间的app-to-app安全性。
3.版本控制  – 这将由API注册表和发现API使用动态Swagger API处理，动态Swagger API可以动态更新并与服务器上的使用者共享。
4. 发现  – 这将由Kubernetes和OpenShift等API发现工具解决。它也可以在代码级使用Netflix Eureka完成。但是，使用业务流程层执行此操作会更好，并且可以通过这些工具进行管理，而不是通过代码和配置进行维护。
5.数据过期 –  应始终更新数据库以提供最新数据。API将从最近更新的数据库中获取数据。还可以为数据库中的每个记录添加时间戳条目，以检查和验证最近的数据。可以根据业务需求使用可定义的驱逐策略来使用和自定义缓存。
6.调试和记录  – 有多种解决方案。可以通过将日志消息推送到异步消息传递平台（如Kafka，Google PubSub等）来使用外化日志记录。客户端可以在标头中为REST API提供关联ID，以跟踪所有pod / Docker容器中的相关日志。此外，可以使用IDE或检查日志在每个微服务上单独完成本地调试。
7.测试 –  可以通过模拟REST API或集成/依赖API来解决此问题，这些API不可用于使用WireMock，BDD，Cucumber，集成测试，使用JMeter进行性能测试以及任何良好的分析工具（如Jprofiler）进行测试， DynaTrace，YourToolKit，VisualVM等
8.监控  – 监控可以使用开源工具，如Prometheus与Grafana结合使用，创建仪表和矩阵，Kubernetes / OpensShift，Influx DB，Apigee，结合Grafana和Graphite。
9. DevOps支持 –  使用最先进的DevOps工具（如GCP，Kubernetes和OpenShift与Jenkins）可以解决微服务部署和支持相关的挑战。
10.容错  – 如果给定SLA / ETA的API没有响应，Netflix Hystrix可用于断开电路。

30个用于微服务的顶级工具 <https://mp.weixin.qq.com/s/XmlmN2h7cCguJc6spBoD7w>
