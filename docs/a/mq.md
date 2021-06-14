# mq

- noun名词: producer 生产者; consumer 消费者 broker=instance实例 pub=publish发布 sub=subscribe订阅 channel渠道 bind绑定
- 用处: 解耦 异步化 削峰填谷 数据交换 日志通道
  - 消息路由 消息确认与重发 定时/延迟消息 消息轨迹 消息回溯 死信队列 幂等
- 概念
  - producer=sender topic -> queue -> consumerGroup consumer=receiver
  - 消息如何路由 单播/广播 message=header+payload
  - 模式: consumer.push.pull(rate+batch 轮询+阻塞)
- 模式: p2p pub/sub
- ActiveMQ
  - JMSapi queue=队列 topic=发布订阅 broker protocol 存储与持久化 cluster&node
- RocketMQ: 延迟/自定义投递 广播模式 async刷盘/复制
- Scibe Flume: push模式

## rabbitmq amqp

- producer -> exchange -> exchangeType+bindingKey+routingKey -> queue -> consumer
  - 死信消息 -> 死信Exchange+BindingKey+死信RoutingKey+Header属性 -> 死信Queue
- exchangeType=fanout+Direct+topic`.#*`+alter备份AE+internal内建 attr=autoDelete+durable
  - direct->Header=attr`x-match`=all+any JMS
  - blackholed问题=Warrens: msg无法从exchange到queue=exchange没绑定queue+bingkey/routingkey错误
- queue arg=高可用+死信+消息过期+queue过期 attr=durable Mirrored
- msg: arg=requeue ack=reject+nack.一次拒绝多条消息+undeliver attr=persistent retry aliveTime delay延迟
- 逻辑: vhost=虚拟broker=exchange+binding+queue
- 物理: node->broker->cluster
- 1Conn.进程.物理=nChan.线程.逻辑
- metadata
- 广域网脑裂问题: HA系统中2个节点心跳突然跳开->分裂为2个独立个体->争抢对方资源

## kafka

- v2.8 Quorm控制器-kraft
- 副本机制与选举原理: zk健康检查; follower同步leader的写操作, 延时不能太久
- offset偏移量 -> partition分区(0分区随机选择broker.其他分区依次后移 nextReplicaShift副本位置偏移量.随机) -> broker=instance实例=服务 -> cluster
- producer -> topic -> consumerGroup(LB.1分区1消费者 组内有序组外无序) -> consumer.pull
- request.required.acks: 0 producer不等broker的ack; 1 等待leader的ack; -1 等待所有follow的ack
- 数据传输的事务定义: 1more最多一次 1least最少一次 1exactly
- fs
  - partition大文件=多个segment小文件(`xxx.index xxx.log` 大小默认1g->滚动生成+offset命名)->方便clean `log.dir`参数=分区目录.name=topic+分区id.目录最小文件夹新建
  - index->msg/resp大小; indexMetadata->memory->segment文件io↓; index稀疏存储->indexMetadata↓
  - 消息格式: len.4byte=1+4+n ver.1byte CRC.4byte msg.nbyte(kv格式 k可指定分区)
- diff: 持久化日志; 分布式; stream

## nsq

- nsqlookupd nsqd=node nsqadmin
- 消息只在单节点上
