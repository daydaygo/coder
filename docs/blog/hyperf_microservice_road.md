# hyperf| 微服务之路

如果对 **微服务** 还不太了解, 可以先看这篇文章: https://doc.hyperf.io/#/zh-cn/microservice

最后, 不要为了微服务而微服务, 团队里有一对技术支持/支撑的人马来帮助解决微服务基础设施的问题, 你会发现微服务其实超好用?

> 让业务开心的写业务, 让技术支持做好服务 -- 技术人员的服务意识很重要

## rpc

服务间通过 rpc 通信, 性能更高, 而且约定清晰, 后续可以工具化自动生成

- rpc: 协议 服务契约 接口类(Interface)
    - json-rpc: 协议数据格式
    - rpc-server: 服务提供者
    - rpc-client: 服务消费者
- grpc: grpc 协议解析, grpc 实际上是基于 http2
    - grpc-server
    - grpc-client

## 服务中心

服务提供者自动注册到服务中心后, 服务消费者从服务中心获取到服务提供者的信息

- service-governance: 服务治理
    - consul: consul 客户端, 用来和 consul server 交互

## 服务降级/限流

熔断发生后, 自动使用降级方法, 避免引发 **雪崩**

- circuit-breaker
    - 熔断策略
        - 默认 超时策略: 超时时间+超时次数
    - 降级: 熔断发生后使用的替代方法, 比如 实时查询->缓存->固定值
- rate-limit: 令牌桶限流器

## 配置中心

why: 服务/配置增多, 方便管理; 配置修改, 不重启服务(db连接等还需要重启)

- config
    - Apollo
    - etcd(推荐)
    - consul
    - zookeeper / qconf(基于zookeeper)
    - aliyun acm

## 调用链追踪

需要一个调用链追踪系统来帮助我们动态地展示服务调用的链路，以便我们可以快速地对问题点进行定位，亦可根据链路信息对服务进行调优。

- tracer
    - 默认支持: guzzle redis db
    - 驱动: zipkin jeager 阿里云opentracking
    - 配置 span tag
    - 配置采样率, 全量采集可能影响性能

## 服务监控

微服务治理的一个核心需求便是服务可观察性

- metric
    - 驱动: prometheus(推荐) StatsD influxDB
    - 业务中实践: 底层实现 db/api 等监控, metric 数据写入统一 redis, 由 tools 项目提供 api, 供 prometheus 拉取数据

## 服务重试

- 网络通讯天然是不稳定的，因此在分布式系统中，需要有良好的容错设计
- 无差别重试是非常危险的

- retry
    - 最大尝试次数策略
    - 错误分类策略
    - 回退策略
    - 睡眠策略
    - 超时策略
    - 熔断策略
    - 预算策略

## 连接池

为啥需要连接池就不赘述了, 不明白的看这篇:

- [是否可以共用同一个redis/mysql连接](https://wiki.swoole.com/wiki/page/325.html): 连接池是标配

hyperf 中使用连接池真的真的是简单到发指: 把原来创建连接的代码, 稍微包装一下就行了:

- 这是 hyperf 官网的例子

```php
<?php
namespace App\Pool;

use Hyperf\Contract\ConnectionInterface;
use Hyperf\Pool\Pool;

class MyConnectionPool extends Pool
{
    public function createConnection(): ConnectionInterface
    {
        // 把原来创建连接的代码, 稍微包装一下就行
        return new MyConnection();
    }
}
```

- 这是使用 mongo eloquent 的例子

```php
// 初始化 mongo 连接
<?php
namespace Mt\Listener;

use Hyperf\Event\Contract\ListenerInterface;
use Hyperf\Contract\ConfigInterface;
use Hyperf\Framework\Event\BootApplication;
use Illuminate\Database\Capsule\Manager as Capsule;
use Jenssegers\Mongodb\Connection;

class BootAppConfListener implements ListenerInterface
{

    public function listen(): array
    {
        return [
            BootApplication::class,
        ];
    }

    public function process(object $event)
    {
        $config = container(ConfigInterface::class);

        // mongo 链接信息注册
        $mongo = config('mongo', []);
        if ($mongo) {
            $capsule = new Capsule;
            // 注册链接信息
            foreach ($mongo as $name => $config) {
                // 统一设置driver
                $config['driver'] = 'mongodb';
                $capsule->addConnection($config, $name);
            }
            $capsule->getDatabaseManager()
                    ->extend('mongodb', function ($config, $name) {
                        $config['name'] = $name;

                        return new Connection($config);
                    });
            // 设置全局静态可访问
            $capsule->setAsGlobal();
            // 启动Eloquent
            $capsule->bootEloquent();
        }
    }
}

// 设置配置文件 mongo.php
<?php
declare(strict_types=1);

return [
    // key = connection name = db name
    'log' => [
        'dsn' => env('MONGO'),
        // db name
        'database' => 'log',
        // options: 有则添加
        'options' => [],
    ],
];

// 设置 model
<?php

namespace App\MongoModel;

use Jenssegers\Mongodb\Eloquent\Model;

class Log extends Model
{
    public $connection = 'log';
    public $collection = 'log';
}

// 使用 model
$result = Log::query()->select('*')->limit(1)->get()->toArray();
var_dump($result);
```

## 协程 client

- redis

```php
// 封装静态方法, 方便使用
<?php
namespace Mt\Util;

use Hyperf\Redis\RedisFactory;

class Redis
{
    /**
     * @param $name string redis pool name
     *
     * @return \Redis
     */
    public static function connection($name = 'default')
    {
        return container(RedisFactory::class)->get($name);
    }
}

// 使用
$redis = \Mt\Util\Redis::connection('order');
$redis->get($key);
```

- guzzle: 别再纠结各类 curl 封装了, 上 guzzle

```php
// 封装静态方法, 方便使用
<?php
namespace Mt\Util;

use GuzzleHttp\Client;

class Guzzle
{
    /**
     * @param array $config
     * @return Client
     */
    public static function create(array $config = [])
    {
        // 如果在协程环境下创建，则会自动使用协程版的 Handler，非协程环境下无改变
        return container(ClientFactory::class)->create($config);
    }
}

// 使用
$http = \Mt\Util\Guzzle::create();
$http->get($url);
```

有了协程版 Guzzle 后, 后续基于 http 的 client, 都可以无缝协程化, 比如 es / consul / etcd 等

- es

```php
$es = container(ClientBuilderFactory::class)->create()->setHosts(['http://127.0.0.1:9200'])->build();
$info = $es->info();
```