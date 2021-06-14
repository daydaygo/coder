# hyperf| 微服务之旅: 配置中心

- [hyperf| 微服务之旅: 配置中心](https://www.jianshu.com/p/9cb3fca076c1)

这篇我们来撸配置中心.

为啥要用配置中心呢? 我用个讨巧的方式来回答这个问题:

- 携程配置中心 [apollo](https://github.com/ctripcorp/apollo)
- 360配置中心 [QConf](https://github.com/Qihoo360/QConf)
- aliyun 应用配置管理 [acm](https://help.aliyun.com/product/59604.html)

这是调研并在框架层适配过的 3 个配置中心, 在他们的文档里, 你都可以找到使用配置中心的理由, 有些理由(或者说场景), 也许正好直击你的痛点.

> 也许仅仅只是为了 `更改配置不用发版`, 也应该尝试一下配置中心试试看.

使用配置中心, 远远没有我们想的那么复杂, 在做选型的时候, 往往会被这个工具或者那个工具新增的一些 `定义` 绕进去. 那么, 我们从自己使用的框架开始, 先看我们使用配置的需求, 然后再来看配置中心 `怎样满足我们的需求`.

## PHPer 项目中使用的配置

以 PHPer 为例, 项目中通常有 2 类配置:

- env 环境相关的配置

以下面的 db 配置为例, 不同环境需要使用不同的 db 配置, 从 `.env` 配置文件中获取

```php
return [
    'db' => [
        'driver' => 'mysql',
        'host' => env('DB_HOST'),
        'database' => 'test',
        'username' => env('DB_USER'),
        'password' => env('DB_PWD'),
    ],
];
```

- config 项目中所有配置

```php
$container = ApplicationContext::getContainer();
/** @var ConfigInterface $config */
$config = $container->get(ConfigInterface::class);
$config->get('a.b.c', 'default');
```

这里有记住 2 个原则(`约定大于配置`):

- env 只在配置文件中使用, 所有配置使用的地方, 都使用 config
- 所有的配置都可以通过 `$config` 对象来获取, 所有配置都由 `$config` 来管理

> 明白了这点以后, 无论什么配置中心, 都只需要增加一个 package 来适配, 最终将配置更新到 `$config` 中即可

## 以适配 aliyun 应用配置管理(acm) 为例

我们新增一个 `hyperf/config-aliyun-acm` 的 package, 用来适配 aliyun 应用配置管理(acm). 这是个免费的服务, 适合用来体验配置中心, 少了一部配置中心服务运维管理的步骤.

总共只有 2 步:
- 从配置中心获取最新配置

```php
// vendor/hyperf/config-aliyun-acm/src/Client.php
<?php

declare(strict_types=1);

namespace Hyperf\ConfigAliyunAcm;

use Closure;
use Hyperf\Contract\ConfigInterface;
use Hyperf\Guzzle\ClientFactory as GuzzleClientFactory;
use Psr\Container\ContainerInterface;
use RuntimeException;

class Client implements ClientInterface
{
    /**
     * @var array
     */
    public $fetchConfig;

    /**
     * @var Closure
     */
    private $client;

    /**
     * @var ConfigInterface
     */
    private $config;

    /**
     * @var array
     */
    private $servers;

    public function __construct(ContainerInterface $container)
    {
        $this->client = $container->get(GuzzleClientFactory::class)->create();
        $this->config = $container->get(ConfigInterface::class);
    }

    public function pull(): array
    {
        $client = $this->client;
        if (! $client instanceof \GuzzleHttp\Client) {
            throw new RuntimeException('aliyun acm: Invalid http client.');
        }

        // ACM config
        $endpoint = $this->config->get('aliyun_acm.endpoint', 'acm.aliyun.com');
        $namespace = $this->config->get('aliyun_acm.namespace', '');
        $dataId = $this->config->get('aliyun_acm.data_id', '');
        $group = $this->config->get('aliyun_acm.group', 'DEFAULT_GROUP');
        $accessKey = $this->config->get('aliyun_acm.access_key', '');
        $secretKey = $this->config->get('aliyun_acm.secret_key', '');

        // Sign
        $timestamp = round(microtime(true) * 1000);
        $sign = base64_encode(hash_hmac('sha1', "{$namespace}+{$group}+{$timestamp}", $secretKey, true));

        if (! $this->servers) {
            // server list
            $response = $client->get("http://{$endpoint}:8080/diamond-server/diamond");
            if ($response->getStatusCode() !== 200) {
                throw new RuntimeException('Get server list failed from Aliyun ACM.');
            }
            $this->servers = array_filter(explode("\n", $response->getBody()->getContents()));
        }
        $server = $this->servers[array_rand($this->servers)];

        // Get config
        $response = $client->get("http://{$server}:8080/diamond-server/config.co", [
            'headers' => [
                'Spas-AccessKey' => $accessKey,
                'timeStamp' => $timestamp,
                'Spas-Signature' => $sign,
            ],
            'query' => [
                'tenant' => $namespace,
                'dataId' => $dataId,
                'group' => $group,
            ],
        ]);
        if ($response->getStatusCode() !== 200) {
            throw new RuntimeException('Get config failed from Aliyun ACM.');
        }
        return json_decode($response->getBody()->getContents(), true);
    }
}
```
- 更新到 `$config` 对象中

```php
// vendor/hyperf/config-aliyun-acm/src/Process/ConfigFetcherProcess.php
<?php

declare(strict_types=1);

namespace Hyperf\ConfigAliyunAcm\Process;

use Hyperf\ConfigAliyunAcm\ClientInterface;
use Hyperf\Contract\ConfigInterface;
use Hyperf\Process\AbstractProcess;
use Hyperf\Process\Annotation\Process;
use Psr\Container\ContainerInterface;
use Swoole\Server;

/**
 * @Process(name="aliyun-acm-config-fetcher")
 */
class ConfigFetcherProcess extends AbstractProcess
{
    /**
     * @var Server
     */
    private $server;

    /**
     * @var ClientInterface
     */
    private $client;

    /**
     * @var ConfigInterface
     */
    private $config;

    /**
     * @var string
     */
    private $cacheConfig;

    public function __construct(ContainerInterface $container)
    {
        parent::__construct($container);
        $this->client = $container->get(ClientInterface::class);
        $this->config = $container->get(ConfigInterface::class);
    }

    public function bind(Server $server): void
    {
        $this->server = $server;
        parent::bind($server);
    }

    public function isEnable(): bool
    {
        return $this->config->get('aliyun_acm.enable', false);
    }

    public function handle(): void
    {
        while (true) {
            $config = $this->client->pull();
            if ($config !== $this->cacheConfig) {
                if ($this->cacheConfig !== null) {
                    $diff = array_diff($this->cacheConfig ?? [], $config);
                } else {
                    $diff = $config;
                }
                $this->cacheConfig = $config;
                $workerCount = $this->server->setting['worker_num'] + $this->server->setting['task_worker_num'] - 1;
                // 通过进程间通信, 投递配置信息到每一个启动的进程
                for ($workerId = 0; $workerId <= $workerCount; ++$workerId) {
                    $this->server->sendMessage($diff, $workerId);
                }
            }
            sleep($this->config->get('aliyun_acm.interval', 5));
        }
    }
}
```

## 以 apollo 为例

代码适配其实和上面类似, 也是通过调用 apollo 提供的 api 去获取最新的配置, 然后进行更新, 不同的是, 你得部署一套 apollo 配置中心的环境. 以本地开发测试为例:

```yml
version: '3.1'
services:
    apollo: # https://github.com/ctripcorp/apollo/tree/master/scripts/docker-quick-start
        image: nobodyiam/apollo-quick-start
        ports:
            - "8080:8080"
            - "8070:8070"
        links:
            - mysql:apollo-db
        tty: true
    mysql:
        image: mysql:5.7.26
        restart: always
        volumes:
            - ./config/my.cnf:/etc/mysql/conf.d/my.cnf
            - ./config/sql:/docker-entrypoint-initdb.d
            - ./data/mysql:/var/lib/mysql
        ports:
            - "3306:3306"
        environment:
            TZ: Asia/Shanghai
            MYSQL_ALLOW_EMPTY_PASSWORD: 'yes'
```

这只是本地测试, 所以只启动了一个配置中心, 如果是线上环境, 要做集群化处理, 配置中心这样重要的服务, 要确保高可用.

## 写在最后

QConf 的理念和部署又是另一个样子了, 使用了 zookeeper 来确保服务的可用.

但是无论是使用 apollo QConf 还是 acm, 在框架层其实只需要使用 `composer require` 增加相应的适配包即可, 项目中的代码完全不用修改.

配置中心可以改进的部分:
- 目前使用 http 轮询的方式, 部分配置中心提供了长连接, 可以进行适配
- 将配置中心注册到 `服务注册发现` 服务中, 统一从 服务注册发现服务中获取服务信息