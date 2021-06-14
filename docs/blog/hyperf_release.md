# hyperf| 发版

- 发版: pr review/merge 等
- 打版本 @李
- 获取发版信息(github - release): https://github.com/hyperf/hyperf/releases
- 获取 CI 信息(github - action - phpunit - 最新一条): https://github.com/hyperf/hyperf/actions/workflows/test.yml
- 发公众号 @李

## 模板

```markdown
# 复工大吉

大家享受完超长假期, 是不是满血复活准备大干一场, 祝大家复工大吉~

# 更新内容

本周主要进行小功能迭代优化和 🐛Bug 修复, 继续提升 hyperf 的稳定性, 发布于 [1.1.18](https://github.com/hyperf/hyperf/releases/tag/v1.1.18) 版, 建议使用到 Metric / async-queue 组件的用户更新此版本

直接访问 官网 [hyperf.io](https://hyperf.io) 或 文档 [hyperf.wiki](https://hyperf.wiki) 查看更新内容

## 新增

- [#1305](https://github.com/hyperf/hyperf/pull/1305) 为 `hyperf\metric` 组件添加预制的 `Grafana` 面板
- [#1328](https://github.com/hyperf/hyperf/pull/1328) 添加 `ModelRewriteInheritanceVisitor` 来重写 model 类继承的 `gen:model` 命令
- [#1331](https://github.com/hyperf/hyperf/pull/1331) 添加 `Hyperf\LoadBalancer\LoadBalancerInterface::getNodes()`
- [#1335](https://github.com/hyperf/hyperf/pull/1335) 为 `command` 添加 `AfterExecute` 事件
- [#1361](https://github.com/hyperf/hyperf/pull/1361) logger 组件添加 `processors` 配置

## 修复

- [#1330](https://github.com/hyperf/hyperf/pull/1330) 修复当使用 `(new Parallel())->add($callback, $key)` 并且参数 `$key` 并非 string 类型, 返回结果将会从 0 开始排序 `$key`
- [#1338](https://github.com/hyperf/hyperf/pull/1338) 修复当从 server 设置自己的设置时, 主 server 的配置不生效的 bug
- [#1344](https://github.com/hyperf/hyperf/pull/1344) 修复队列在没有设置最大消息数时每次都需要校验长度的 bug

## 变更

- [#1324](https://github.com/hyperf/hyperf/pull/1324) [hyperf/async-queue](https://github.com/hyperf/async-queue) 组件不再提供默认启用 `Hyperf\AsyncQueue\Listener\QueueLengthListener`

## 优化

- [#1305](https://github.com/hyperf/hyperf/pull/1305) 优化 `hyperf\metric` 中的边界条件
- [#1322](https://github.com/hyperf/hyperf/pull/1322) HTTP Server 自动处理 HEAD 请求并且不会在 HEAD 请求时返回 response baody

## 删除

- [#1303](https://github.com/hyperf/hyperf/pull/1303) 删除 `Hyperf\RpcServer\Router\Router` 中无用的 `$httpMethod`

# 关于 Hyperf

Hyperf 是基于 `Swoole 4.4+` 实现的高性能、高灵活性的 PHP 协程框架，内置协程服务器及大量常用的组件，性能较传统基于 `PHP-FPM` 的框架有质的提升，提供超高性能的同时，也保持着极其灵活的可扩展性，标准组件均基于 [PSR 标准](https://www.php-fig.org/psr) 实现，基于强大的依赖注入设计，保证了绝大部分组件或类都是 `可替换` 与 `可复用` 的。

框架组件库除了常见的协程版的 `MySQL 客户端`、`Redis 客户端`，还为您准备了协程版的 `Eloquent ORM`、`WebSocket 服务端及客户端`、`JSON RPC 服务端及客户端`、`GRPC 服务端及客户端`、`OpenTracing(Zipkin, Jaeger) 客户端`、`Guzzle HTTP 客户端`、`Elasticsearch 客户端`、`Consul 客户端`、`ETCD 客户端`、`AMQP 组件`、`Nats 组件`、`Apollo、ETCD、Zookeeper 和阿里云 ACM 的配置中心`、`基于令牌桶算法的限流器`、`通用连接池`、`熔断器`、`Swagger 文档生成`、`Swoole Tracker`、`Blade、Smarty、Twig、Plates 和 ThinkTemplate 视图引擎`、`Snowflake 全局ID生成器`、`Prometheus 监控` 等组件，省去了自己实现对应协程版本的麻烦。

Hyperf 还提供了 `基于 PSR-11 的依赖注入容器`、`注解`、`AOP 面向切面编程`、`基于 PSR-15 的中间件`、`自定义进程`、`基于 PSR-14 的事件管理器`、`Redis/RabbitMQ 消息队列`、`自动模型缓存`、`基于 PSR-16 的缓存`、`Crontab 秒级定时任务`、`Session`、`i18n 国际化`、`Validation 表单验证` 等非常便捷的功能，满足丰富的技术场景和业务场景，开箱即用。

# 框架初衷

尽管现在基于 PHP 语言开发的框架处于一个百花争鸣的时代，但仍旧未能看到一个优雅的设计与超高性能的共存的完美框架，亦没有看到一个真正为 PHP 微服务铺路的框架，此为 Hyperf 及其团队成员的初衷，我们将持续投入并为此付出努力，也欢迎你加入我们参与开源建设。

# 设计理念

`Hyperspeed + Flexibility = Hyperf`，从名字上我们就将 `超高速` 和 `灵活性` 作为 Hyperf 的基因。

- 对于超高速，我们基于 Swoole 协程并在框架设计上进行大量的优化以确保超高性能的输出。
- 对于灵活性，我们基于 Hyperf 强大的依赖注入组件，组件均基于 [PSR 标准](https://www.php-fig.org/psr) 的契约和由 Hyperf 定义的契约实现，达到框架内的绝大部分的组件或类都是可替换的。

基于以上的特点，Hyperf 将存在丰富的可能性，如实现 单体 Web 服务，API 服务，网关服务，分布式中间件，微服务架构，游戏服务器，物联网（IOT）等。

# 文档齐全

我们投入了大量的时间用于文档的建设，以解决各种因为文档缺失所带来的问题，文档上也提供了大量的示例，对新手同样友好。
[Hyperf 官方开发文档](https://hyperf.wiki)

# 生产可用

我们为组件进行了大量的单元测试以保证逻辑的正确，目前存在 `1312` 个单测共 `3848` 个断言条件，同时维护了高质量的文档，在 Hyperf 正式对外开放(2019 年 6 月 20 日)之前，便已经过了严酷的生产环境的考验，我们才正式的对外开放该项目，现在已有很多的大型互联网企业都已将 Hyperf 部署到了自己的生产环境上并稳定运行。

# 官网及交流

[Github](https://github.com/hyperf/hyperf) 👈👈👈👈👈 点 Star 支持我们
[Gitee 码云](https://gitee.com/hyperf/hyperf) 👈👈👈👈👈 点 Star 支持我们
[Hyperf 官网](https://hyperf.io)
[Hyperf 文档](https://hyperf.wiki)
QQ 群: 862099724
```
