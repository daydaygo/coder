# php| 初探 rabbitmq

- [php| 初探 rabbitmq](https://www.jianshu.com/p/6bbdcce31663)

经常看到消息队列( MQ ), 实战中比较少, 说说我的一些粗线的理解:

- 引入消息队列, 使系统之间解耦 -> 当然还有很多 **小型项目** 使用 **重项目** 的方式(系统拆分, 不存在的!); 解耦这部分的内容, 后来还会讲到
- 通过将流程 **异步化**, 增加每部分的吞吐能力, 从而实现最终增加系统性能, 这一点有点类似服务器领域的 **同步 -> 异步/协程**, 可以参考 [swoole| swoole 协程初体验](https://www.jianshu.com/p/745b0b3ffae7)

当然, **实践出真知**, 还是希望能在业务中多实战, 或者参与相关的开源项目~

## 相关教程

老生常谈, **生命周期思维方式**, 想一想消息是怎么在整个消息队列系统中流动的.

- [RabbitMQ与AMQP协议详解](https://www.cnblogs.com/frankyou/p/5283539.html): **强烈推荐**, 理解基础概念非常好的一篇文章

- [PHP消息队列实现及应用](www.imooc.com/learn/852): 简单入门级, 不用纠结代码, 关注应用场景和解决方式

应用场景: 冗余 解耦 流量削峰 异步通信 扩展新 排序保证 -> 队列结构的中间件
队列介质: mysql redis 消息队列服务(rabbitmq kafka)
触发机制: 死循环while 定时脚本cron 守护进程daemon(fpm)
场景一 订单系统/配送系统解耦: 订单系统 -> 队列表 -> 配送系统
场景二 流量削峰: redis list类型实现定长队列, 请求先入队列, 超出队列长度后的请求丢弃
rabbitmq架构和原理: 完整实现AMQP 集群简化 持久化 跨平台

- [RabbitMQ消息中间件极速入门与实战](https://www.imooc.com/learn/1042)

基于 AMQP(advanced message queue protocol, 高级消息队列协议)
集群模式: 表达式配置 HA模式 镜像队列模式
保证数据不丢失/高可靠/高可用

## 安装 rabbitmq

老生常谈, 直接上 docker:

- [docker环境快速安装](https://cr.console.aliyun.com/cn-hangzhou/mirrors): 阿里云提供, 安装/加速 全搞定

docker-compose, 使用 rabbitmq 官方镜像:

```yaml
version: '3'
services:
    rabbitmq: # https://hub.docker.com/_/rabbitmq/
        image: rabbitmq:3.7.7-management-alpine
        hostname: myrabbitmq
        ports:
            - "5672:5672" # mq
            - "15672:15672" # admin
```

就这么几行, rabbitmq 就配置好了, 使用 `docker-compose up -d rabbit` 启动, 大功告成

这里使用的 `management` 版本的 rabbitmq, 管理控制台地址 `http://localhost:15672`, 初始密码 `guest/guest`

进入容器内部看看:

```bash
# 进入容器
docker-compose exec rabbitmq bash

# 查看启动的服务
bash-4.4# ps aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
rabbitmq     1  0.0  0.0   1612  1012 ?        Ss   06:25   0:00 /bin/sh /opt/rabbitmq/sbin/rabbitmq-server
rabbitmq    90  0.0  0.0   1068   660 ?        S    06:25   0:00 /usr/lib/erlang/erts-9.3/bin/epmd -daemon
rabbitmq   158  0.5  3.9 1707624 81152 ?       Sl   06:25   0:37 /usr/lib/erlang/erts-9.3/bin/beam.smp -W w -A 64 -MBas ageffcbf -MHas ageffcbf -MBlmbcs 512 -MHlmbcs 512 -MMmcs 30 -P 1048576 -t 5000000 -s
rabbitmq   265  0.0  0.0    752   536 ?        Ss   06:25   0:00 erl_child_setup 1048576
rabbitmq   306  0.0  0.0    772     4 ?        Ss   06:25   0:00 inet_gethost 4
rabbitmq   307  0.0  0.0    772    32 ?        S    06:25   0:00 inet_gethost 4
root      3258  0.0  0.0   6364  2004 pts/0    Ss   07:25   0:00 bash
root      6813  0.0  0.0   5696   628 pts/0    R+   08:23   0:00 ps aux

# rabbitmq 相关命令行命令
bash-4.4# rabbitmq
rabbitmq-defaults     rabbitmq-diagnostics  rabbitmq-env          rabbitmq-plugins      rabbitmq-server       rabbitmqadmin         rabbitmqctl
```

- rabbitmqctl: rabbitmq control
- rabbitmq-plugins: 也可以通过 rabbitmq-plugins 来开启管理控制台

更多配置, 可以访问 [rabbitmq镜像官网](https://hub.docker.com/_/rabbitmq/) 进行查看

## 快速开始

[rabbitmq 官方文档](https://www.rabbitmq.com/documentation.html) 非常完善与清晰, 值得花时间看看

[新手快速入门文档](https://www.rabbitmq.com/getstarted.html), [代码 - rabbitmq/rabbitmq-tutorials](https://github.com/rabbitmq/rabbitmq-tutorials)

rabbitmq 支持多种语言的 client, 这里说明支持的 php client:

- php-amqplib: compoer package, 入门 rabbitmq 最简单的方式, `composer require` 即可完成安装
- ext-amqp: PHP扩展, 需要安装 `ext-amqp` 扩展, 性能更优

`queue-interop` 按下不表, 还没进入 PSR.

官方的快速入门手册:

- hello world

消息队列最基础的三个概念, **生产者producer + 消息队列MQ + 消费者consumer**

![image](http://upload-images.jianshu.io/upload_images/567399-9461a4685ae84c25.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- work queues

如果 **消费者的消费能力不足** 怎么办? 多开几个 consumer 呗. 多个 consumer 怎么分担 MQ 中的消息呢? 所以在 消息队列MQ 和 消费者consumer 之间进行 **负载均衡LB**. LB并不是准去的说法, 通常的说法是消费者 **订阅** 消息队列的内容. 在我看来, LB 更有表现力 -- 消费能力不足导致需要多个消费者, 怎么和 web server 并发不够, **加机器加 LB** 有些像呢?

> 这里我使用的 LB 这样的概念, 目的在于 LB 在服务器领域太常见了, 包含很多内容, 需要细细体会

![image](http://upload-images.jianshu.io/upload_images/567399-c45b820d2d267ebb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- pub/sub

发布订阅, 看起来像多个 `hello world`, 注意图里多了一个 rabbitmq 中的新概念 `交换器exchange`(图中简写为 X), 在 生产者producer 和 消息队列MQ 之间, 由 exchange 来决定消息 **分发** 到哪个 MQ 中

![image](http://upload-images.jianshu.io/upload_images/567399-7f0da55111f8d3bf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- routing/topic

两个新的概念, **路由功能** 和 **主题订阅** 功能, 而这些都和 交换器exchange 有关, 涉及到的配置: `交换器类型(exchange type)` + `路由key(routing key)` + `绑定key(binding key)`. 为啥会这么复杂呢? 干嘛要这么多配置?

> 一言以蔽之, 决定 生产者producer 产生的消息, 投递到哪个 消息队列MQ 中, 最终又由哪个 consumer 消费

![routing](http://upload-images.jianshu.io/upload_images/567399-d434c609962a6c0c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![topic](http://upload-images.jianshu.io/upload_images/567399-f4831235c2fe776d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- RPC

新玩法, 通过消息队列实现 **远程过程调用RPC** 的效果

![image](http://upload-images.jianshu.io/upload_images/567399-3bfdfab1e7a0779c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我在这里并没有贴具体代码, 一则因为官方给的代码确实适合上手, 更重要是因为, **理解 rabbitmq 究竟是个什么玩意** 更重要

## rabbitmq 究竟是啥

- 首先是 MQ 最基础的概念: 生产者producer + 消息队列MQ(狭义指存储消息的队列类型的数据结构) + 消费者consumer
- consumer 怎么从 MQ 中获取消息 + consumer获取消息后确认(ack): 可以近似理解 MQ 和 consumer 之间需要一层 LB
- producer 怎么投递消息到 MQ 中: exchange + exchange type + routing key + binding key -> 消息路由/主题订阅
- 怎么扩展 MQ 的性能呢: **集群**, 涉及到集群, 又会新增很多概念了

放几张图辅助理解:

![amqp模型](https://upload-images.jianshu.io/upload_images/567399-6022a812e01c3615.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 实际使用过程中, 还需要 tcp 连接相关的概念: connection + channel

![rabbitmq模型](https://upload-images.jianshu.io/upload_images/567399-9587f0707603bb2e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 这张图可以看到消息msg 的 **生命周期**

![rabbitmq消息流转](https://upload-images.jianshu.io/upload_images/567399-3afd224ad9eefff9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 写在最后

正儿八经(准确说, 这叫 **专业**) 的消息队列, 架构设计上确实有一定的复杂性, 而且为了满足高性能, 还会有很多性能优化的点 -- 学习消息队列, 长路漫漫, 可以好好折腾一番了~
