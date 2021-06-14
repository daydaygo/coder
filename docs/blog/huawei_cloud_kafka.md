# cloud| 华为云: kafka 实战训练营

- 课程地址: <https://education.huaweicloud.com:8443/courses/course-v1:HuaweiX+CBUCNXP017+Self-paced/about?isAuth=0&cfrom=hwc>
- 打卡地址: <https://w.url.cn/s/AqQf9K3>
- kafka 加速下载: <https://repo.huaweicloud.com/apache/kafka/1.1.0/kafka_2.12-1.1.0.tgz>

## day 1: basic

- 消息服务: 系统解耦 削峰填谷 数据交换 异步通知 日志通道
- 业界分布式 MQ 横向对比
- 基本概念
  - Broker: 集群中的服务实例
  - Topic: 消息类别
  - Partition: 分区, 物理上的概念, topic:partition = 1:n
    - offset: 只能通过追加增加消息; consumer 通过 offset 定位消息/记录消费位置
    - topic 根据分区个数分配给 consumer group 下的 consumer, 最多只能 partition:consumer = 1:1
    - 分区副本: 高可用; 分配到不同节点上; leader(一个)+ISR(其他副本通过 pull 模式同步 leader 消息)
  - producer
    - 批量生产: batch.size linger.ms
  - consumer
  - consumer group

```sh
# producer
bin/kafka-console-producer.sh --broker-list 192.168.0.180:9092 --topic test
# consumer
bin/kafka-console-consumer.sh --bootstrap-server 192.168.0.180:9092 --topic test --group testgroup --consumer-property enable.auto.commit=true --from-beginning
```

## day 2: producer demo

- 生产模型: BatchSize 打包大小; Linger.ms 发送等待时延; buffer.memory 内存缓存
- 参数调优:
  - tcp: receive.buffer.byte send.buffer.byte
  - acks: 0-不等待 1-等待leader all
- 建议规范:
  - 同步复制客户端配合使用 acks=all
  - 发送失败重试: retries = 3
  - 发送优化: linger.ms = 0
- 配置场景:
  - FIFO 消息保序: 生产消息指定 partiton + retries=0/max.flight.requests.per.connection=1
  - 高吞吐: topic 3分区2副本, acks=0/1
  - 相对可靠: topic 3分区3副本, min.insync.replicas=2, acks=-1
  - 高可靠: topic 3分区3副本, min.insync.replicas=2, flush.messages=1, acks=-1

```bash
# day 2
# producer
java -cp .:./libs/* dms.kafka.demo.KafkaProducerDemo 192.168.0.180:9092 test01
# consumer 同 day 1
```

```java
package dms.kafka.demo;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.clients.producer.RecordMetadata;

import java.util.Properties;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;

public class KafkaProducerDemo {
    public static void main(String[] args) throws InterruptedException, ExecutionException {
        if (args.length != 2) {
            throw new IllegalArgumentException("usage: dms.kafka.demo.KafkaProducerDemo bootstrap-servers topic-name.");
        }
        Properties props = new Properties();
        props.put("bootstrap.servers", args[0]);
        props.put("acks", "all");
        props.put("retries", 0);
        props.put("batch.size", 16384);
        props.put("linger.ms", 1);
        props.put("buffer.memory", 33554432);
        props.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        props.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        Producer<String, String> producer = new KafkaProducer<>(props);
        for (int i = 0; i < 100; i++) {
            Future<RecordMetadata> result = producer.send(new ProducerRecord<String, String>(args[1], Integer.toString(i), Integer.toString(i)));
            RecordMetadata rm = result.get();
            System.out.println("topic: " + rm.topic() + ", partition: " + rm.partition() + ", offset: " + rm.offset());
        }
        producer.close();
    }
};
```

## day3: consumer demo

- consumer: 拉取消息(pull) 确认消息(ack)
  - 消费模型: pull 模式, offset 记录在客户端, 服务端无状态
- consumer group: 实现 Topic 广播+单播
- Rebalance: group 内 consumer 以 topic 的 分区个数进行均衡分配
  - 触发条件: consumer 变化; topic 分区数变化
- assign 模式: 手动分配分区
- subscribe 模式: 自动分配分区

```bash
# day3 consumer
java -cp .:./libs/* dms.kafka.demo.KafkaConsumerDemo 192.168.0.18:9092,192.168.0.121:9092,192.168.0.206:9092 topic-1642673577 test-group
```

```java
package dms.kafka.demo;

import java.util.Arrays;
import java.util.Properties;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;

public class KafkaConsumerDemo {
    public static void main(String[] args) {
        if (args.length != 3) {
            throw new IllegalArgumentException("usage: dms.kafka.demo.KafkaProducerDemo bootstrap-servers topic-name group-name.");
        }
        Properties props = new Properties();
        props.put("bootstrap.servers", args[0]);
        props.put("group.id", args[2]);
        props.put("enable.auto.commit", "true");
        props.put("auto.offset.reset", "earliest");
        props.put("auto.commit.interval.ms", "1000");
        props.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        KafkaConsumer<String, String> consumer = new KafkaConsumer<>(props);
        consumer.subscribe(Arrays.asList(args[1]));
        while (true) {
            ConsumerRecords<String, String> records = consumer.poll(200);
            for (ConsumerRecord<String, String> record : records)
                System.out.printf("offset = %d, key = %s, value = %s%n", record.offset(), record.key(), record.value());
        }
    }
};
```

## day 4: kafka 架构与机制

- 总体架构:
  - zookeeper: 存储 kafka 元数据
  - broker 互为主备
  - topic 按分区存储
  - 副本分布在不同节点
- 节点角色
  - controller: partition 管理和副本管理; broker 节点状态管理; topic 分区状态管理
  - leader + follower
  - coordinator: consumer group 管理
- 核心流程
  - topic 新建/删除
  - leader 选举 / 副本迁移
  - 分区扩容
  - 生产请求流程
- data flow
  - producer
  - kafka -> broker -> topic -> partition
  - consumer group -> consumer

## day 5: kafka 工具

- topic 管理 `kafka-topics.sh`
  - bootstrap-server
  - zookeeper
  - create / delete / list / describe / topic(名称)
  - partitions: 指定分区数
  - replication-factor: 副本数
  - config
- producer 测试 `kafka-console-producer.sh`
  - broker-list
  - message-send-max-retries
  - producer.config: 客户端配置文件
  - propety: 客户端自定义配置
  - topic
- consumer 测试 `kafka-console-consumer.sh`
  - bootstrap-server
  - consumer-property
  - consumer.config
  - from-beginning: 若不存在消费进度, 从头开始消费
  - group
  - partition
  - topic
  - whitelist: 消费 topic 的正则表达式
- 消费组管理 `kafka-consumer-groups.sh`
  - bootstrap-server
  - comand-config: 自定义客户端配置
  - describe
  - group
  - list

```bash
root@ecs-s3-small-1-linux-20190812201034 /d/k/bin# pwd
/data/kafka_2.12-2.3.0/bin

# topic 管理
./kafka-topics.sh --help

# producer 测试
./kafka-console-producer.sh --help

# consumer 测试
./kafka-console-consumer.sh --help

# day 5
bin/kafka-console-producer.sh --broker-list 192.168.0.180:9092 --topic day5
bin/kafka-console-consumer.sh --bootstrap-server 192.168.0.180:9092 --topic day5 --group testgroup --consumer-property enable.auto.commit=true --from-beginning
```
