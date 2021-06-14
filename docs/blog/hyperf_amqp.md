# hyperf| amqp

hyperf 开源有一段时间了, 从开发者交流群就能感受到热度. 这段时间下来, 有一个明显的现象, 某A 提了一个技术问题, 某B 直接抛一个官方文档的对应链接. 这种现象实在太常见, 甚至衍生出了 **欢迎进入 vip 交流群** 这样的商机. 不得不说:

> 花 2 个小时认真看一遍文档, 比遇到问题就卡住然后到处问要高效得多.

重要的事情说三遍:
- 认真看一遍文档
- 认真看一遍文档
- 认真看一遍文档

好了, 回到正题, 今天我们来玩 amqp.

## 项目准备

这些都可以在文档中找到, 所以直接上操作.

- 使用开发组提供的 docker

hold 不住 docker, 不用也行, 但是要能基于开发组的 Dockerfile 配置好环境, 如果既不会用 docker, 也无法自己配置好开发环境, **请一定要努力哦**.

使用 docker-compose 配置的全部环境:

```yaml
version: '3'
services:
    ms:
        image: hyperf/hyperf
        volumes:
            - ../:/data
        ports:
            - "9501:9501"
        environment:
            APP_ENV: dev
        tty: true
    mysql:
        image: mysql:5.7.26
        volumes:
            - ./config/my.cnf:/etc/mysql/conf.d/my.cnf
            - ./config/sql:/docker-entrypoint-initdb.d
            - ./data/mysql:/var/lib/mysql
        ports:
            - "3306:3306"
        environment:
            TZ: Asia/Shanghai
            MYSQL_ROOT_PASSWORD: root
    redis:
        image: redis:alpine
        volumes:
            - ./config/redis.conf:/etc/redis/redis.conf
            - ./data/redis:/data
        ports:
            - "6379:6379"
    rabbitmq:
        image: rabbitmq:management-alpine
        hostname: myrabbitmq
        volumes:
            - ./data/rabbitmq:/var/lib/rabbitmq/mnesia
        ports:
            - "5672:5672" # mq
            - "15672:15672" # admin
```

这里多说一句, docker / Dockerfile / docker-compose 只是满足开发环境的使用的话, 真的很简单, 记住几个常用的 docker 命令, 清楚 Dockerfile 几个常用的指令(RUN CMD 等), docker-compose 只是 yaml 格式的配置文件而已.

> 推荐一个好习惯: 一个文档专门记 docker / Dockerfile / docker-compose 的常用内容, 使用过程中逐渐对这个文件进行增删查改(CRUD), 不用多久, 你就会发现自己用起 docker 来, 贼 6 !

另一个好习惯是 **最佳实践**, 你要从无到有用起来很难, 但是跟着最佳实践走, 就能又快又好 ! 当然, 再上一层楼, **你也能成为最佳实践**.

- 安装项目

```
composer create-project hyperf/hyperf-skeleton hyperf-demo
```

安装过程选择自己需要的组件, **不清楚就先不要选**, 反正之后可以通过 `composer require` 安装. 其实我更想说的是:

> 安装的组件自己 hold 不住, 然后到处叫, 这样多没意思呀.

- 配置 composer.json

```
"repositories": {
    "hyperf": {
        "type": "path",
        "url": "../hyperf/src/*"
    },
    "packagist": {
        "type": "composer",
        "url": "https://mirrors.aliyun.com/composer"
    }
}
```

添加了 path, 从我本地加载 hyperf 组件, 方便开发, **如果不参与 hyperf 组件开发**, 可以忽略这一步.

- 添加 hyperf/amqp

```
# 安装
composer require hyperf/amqp 

# 添加配置文件
php bin/hyperf.php vendor:publish hyperf/amqp
```

- 修改配置, 启动并验证

我启动了 mysql / redis / rabbitmq, 配置相关组件的配置, 并启动框架进行验证

```
# config
vim .env
vim config/autoload/redis.php
vim config/autoload/database.php
vim config/autoload/amqp.php

# test
php bin/hyperf.php start
```

好了, 项目准备好了, 正式开始撸代码.

## 官方文档 amqp demo

文档有的, 还是直接贴:

```
# producer
php bin/hyperf.php gen:amqp-producer DemoProducer

# consumer
php bin/hyperf.php gen:amqp-consumer DemoConsumer

# 使用 command 盗用 DemoProducer 进行验证
php bin/hyperf.php gen:command TestCommand
```

producer 发个消息:
- 设置 command 的名字: `parent::__construct('t');`
- 使用 `@Inject()` 注解注入
- 发消息, 一行搞定: `$this->producer->produce(new DemoProducer('test'. date('Y-m-d H:i:s')));`

```php
<?php

declare(strict_types=1);

namespace App\Command;

use App\Amqp\Producer\DemoProducer;
use Hyperf\Amqp\Producer;
use Hyperf\Command\Command as HyperfCommand;
use Hyperf\Command\Annotation\Command;
use Hyperf\Di\Annotation\Inject;
use Psr\Container\ContainerInterface;

/**
 * @Command
 */
class TestCommand extends HyperfCommand
{
    /**
     * @var ContainerInterface
     */
    protected $container;

    /**
     * @Inject()
     * @var Producer
     */
    protected $producer;

    public function __construct(ContainerInterface $container)
    {
        $this->container = $container;

        parent::__construct('t');
    }

    public function configure()
    {
        $this->setDescription('Hyperf Demo Command');
    }

    public function handle()
    {
        $this->producer->produce(new DemoProducer('test'. date('Y-m-d H:i:s')));
    }
}
```

愉快的玩耍起来:

```
# produce
php bin/hyperf.php t

# consume
php bin/hyperf.php start # 会使用 swoole process 启动 DemoConsumer

# 也可以访问 rabbitmq admin 控制台
http://localhost:15672
```

## 撸一撸 rabbitmq 官网 tutorial

跟着 rabbitmq 官网 tutorial, 见识一下 hyperf 中的 amqp 有多简单

- [hello world](https://www.rabbitmq.com/tutorials/tutorial-one-php.html)

```php
// consumer
/**
 * @Consumer()
 */
class DemoConsumer extends ConsumerMessage
{
    protected $exchange = 'hello';
    protected $type = Type::FANOUT;
    protected $queue = 'hello';

    public function consume($data): string
    {
        var_dump($data);
        return Result::ACK;
    }
}

// producer
/**
 * @Producer()
 */
class DemoProducer extends ProducerMessage
{
    protected $exchange = 'hello';
    protected $type = Type::FANOUT;
    protected $routingKey = 'hello';
    public function __construct($data)
    {
        $this->payload = $data;
    }
}
```

- [work queue](https://www.rabbitmq.com/tutorials/tutorial-two-php.html)

设置一下 `nums` 参数, 就可以多进程.

```php
// Consumer
/**
 * @Consumer(nums=2)
 */
class DemoConsumer extends ConsumerMessage
{
    protected $exchange = 'task';
    protected $type = Type::FANOUT;
    protected $queue = 'task';

    public function consume($data): string
    {
        var_dump($data);
        return Result::ACK;
    }
}

// producer
/**
 * @Producer()
 */
class DemoProducer extends ProducerMessage
{
    protected $exchange = 'task';
    protected $type = Type::FANOUT;
    protected $routingKey = 'task';
    public function __construct($data)
    {
        $this->payload = $data;
    }
}
```

- [pub/sub](https://www.rabbitmq.com/tutorials/tutorial-three-php.html)

和上面的 `hello world` 一致

- [routing](https://www.rabbitmq.com/tutorials/tutorial-four-php.html)

终于看到 `routing_key` 的作用了


```
// consumer
/**
 * @Consumer()
 */
class DemoConsumer extends ConsumerMessage
{
    protected $exchange = 'routing';
    protected $type = Type::DIRECT;
    // 这个 consumer 只消费 error 级别的日志
    protected $queue = 'routing.error';
    protected $routingKey = 'error';

    public function consume($data): string
    {
        var_dump($data);
        return Result::ACK;
    }
}

/**
 * @Consumer()
 */
class Demo2Consumer extends ConsumerMessage
{
    protected $exchange = 'routing';
    protected $type = Type::DIRECT;
    // 这个 consumer 消费所有级别的日志
    protected $queue = 'routing.all';
    protected $routingKey = [
        'info',
        'warning',
        'error',
    ];

    public function consume($data): string
    {
        var_dump($data);
        return Result::ACK;
    }
}

// producer
/**
 * @Producer()
 */
class DemoProducer extends ProducerMessage
{
    protected $exchange = 'routing';
    protected $type = Type::DIRECT;
    public function __construct($data, $routingKey)
    {
        $this->routingKey = $routingKey;
        $this->payload = $data;
    }
}

// produce
$this->producer->produce(new DemoProducer('info'. date('Y-m-d H:i:s'), 'info'));
$this->producer->produce(new DemoProducer('warning'. date('Y-m-d H:i:s'), 'warning'));
$this->producer->produce(new DemoProducer('error'. date('Y-m-d H:i:s'), 'error'));
var_dump('done');
```

- [topics](https://www.rabbitmq.com/tutorials/tutorial-five-php.html)

和的, 和上面的 routing 差不多

```php
// consume
/**
 * @Consumer()
 */
class DemoConsumer extends ConsumerMessage
{
    protected $exchange = 'topics';
    protected $type = Type::TOPIC;
    protected $queue = 'topics.t1';
    // protected $routingKey = '#'; // all
    // protected $routingKey = 'kern.*';
    // protected $routingKey = '*.critical';
    // protected $routingKey = 'kern.critical';
    protected $routingKey = [
        'kern.*',
        '*.critical',
    ];

    public function consume($data): string
    {
        var_dump($data);
        return Result::ACK;
    }
}

// produce
/**
 * @Producer()
 */
class DemoProducer extends ProducerMessage
{
    protected $exchange = 'topics';
    protected $type = Type::TOPIC;
    public function __construct($data, $routingKey)
    {
        $this->routingKey = $routingKey;
        $this->payload = $data;
    }
}
```

可以看到, 想要用 amqp, 自动生成好代码后, 改一改属性就成, so easy~

## 再聊 amqp

amqp 难不难用? 至少基础的使用还是很好掌握的, 下面有一张图可供参考

把 `producer consumer connection vhost channel exchange queue routing_key bind publish consume msg` 几个概念了解了, 基础使用就能很顺手. 而在 hyperf 中, 得益于对一些常用使用的方式的封装, **自动生成代码 + 改改类属性** 就能把 amqp 用起来 !

## 写在最后

无他, 唯手熟尔.

- [php| 初探 rabbitmq](https://www.jianshu.com/p/6bbdcce31663)

- [hyperf: 喜欢就请给个 star 哦~](https://github.com/hyperf-cloud/hyperf)