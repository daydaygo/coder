# hyperf| hyperf 源码解读 2: start

- [hyperf| hyperf 源码解读 2: start](https://www.jianshu.com/p/0d1d89686ff9)

上篇我们跟着 `php bin/hyperf.php` 命令, 看到了框架的核心 `container`, 这篇我们跟着 `php bin/hyperf.php start`, 来会一会强大到爆炸的 `swoole`

开始之前, 请确保自己具备一定的 swoole 基础知识, 这篇适合你:

- [swoole| swoole wiki 笔记](https://www.jianshu.com/p/12d645ac02b2)

> 千万不要小看了 `基础知识`, 与其一直卡壳浪费时间, 不如沉下心来好好读一下 swoole wiki, 看似 `艰辛`, 绝对比没有这些基础知识导致的时间浪费要划算得多得多得多 ! 

## Command 基础

从上一篇的 blog 可知, [hyperf 源码解读 1: 启动](https://www.jianshu.com/p/c167f6e130df), `php bin/hyperf.php` 执行的命令, 是基于 `Symfony\Component\Console\Application` 提供的命令行应用提供的功能, 初始内置的 command, 是有 ConfigProvider 配置项中的 command 字段设置的. 那么, 我们将要执行的 `start` 命令, 是由哪个 command 提供的呢?

可以根据 `gen:command` 命令生成的 demo 代码, 了解到:

```php
// \App\Command\TestCommand
public function __construct(ContainerInterface $container)
{
    $this->container = $container;

    // 命令的名字通常在这里配置
    parent::__construct('t');
}
```

对应的 `start` 命令, 就可以通过全局搜索 `parent::__construct('start')` 查询到

> 搜索是很重要的能力, 甚至可以上升到 `搜商`. 你并不需要特别熟悉一个事物, 但是你依旧有 `能力` 进行处理, 搜商就是这样的一个基础能力. 阅读源码也是提高搜商的一个有力方法.

OK, 就定位到了我们 `start` 命令对应的代码: `\Hyperf\Server\Command\StartServer`

## StartServer 一览

```php
/**
 * @Command
 */
class StartServer extends SymfonyCommand
{
    /**
     * @var ContainerInterface
     */
    private $container;

    public function __construct(ContainerInterface $container)
    {
        parent::__construct('start');
        $this->container = $container;

        $this->setDescription('Start swoole server.');
    }

    protected function execute(InputInterface $input, OutputInterface $output)
    {
        \Swoole\Runtime::enableCoroutine(true);

        $this->checkEnvironment($output);

        $serverFactory = $this->container->get(ServerFactory::class)
            ->setEventDispatcher($this->container->get(EventDispatcherInterface::class))
            ->setLogger($this->container->get(StdoutLoggerInterface::class));

        $serverConfig = $this->container->get(ConfigInterface::class)->get('server', []);
        if (! $serverConfig) {
            throw new \InvalidArgumentException('At least one server should be defined.');
        }

        $serverFactory->configure($serverConfig);
        $serverFactory->start();
    }

    private function checkEnvironment(OutputInterface $output)
    {
        if (ini_get_all('swoole')['swoole.use_shortname']['local_value'] !== 'Off') {
            $output->writeln('<error>ERROR</error> Swoole short name have to disable before start server, please set swoole.use_shortname = \'Off\' into your php.ini.');
            exit(0);
        }
    }
}
```

- `@Command` 注解, container 初始化的时候会有 scan 注解, 就能把这个类解析为 command 来使用
- `execute()` 是 command 实际执行的方法, `\Symfony\Component\Console\Command\Command` 作为基类采用这个方法, 而 `\Hyperf\Command\Command` 作为基类, 则是使用 `handle()` 方法
- `$serverFactory = $this->container->get(ServerFactory::class)`: container 的经典使用场之一, 由 container 来解决类的初始化, 使用放直接使用即可

## ServerFactory 细节

还是上面的代码:

```php
$serverFactory = $this->container->get(ServerFactory::class)
  ->setEventDispatcher($this->container->get(EventDispatcherInterface::class))
  ->setLogger($this->container->get(StdoutLoggerInterface::class));

$serverConfig = $this->container->get(ConfigInterface::class)->get('server', []);
if (! $serverConfig) {
    throw new \InvalidArgumentException('At least one server should be defined.');
}

$serverFactory->configure($serverConfig);

$serverFactory->start();
```

- 实例化 ServerFactory 时, 配置好框架的基础组件 Event(事件模块) / Logger(日志模块)
- 从框架通的 Config(配置) 组件获取 server 配置, 即 `config/autoload/server.php` 文件的配置内容
- server start, 配置好了 swoole server 后, 使用 `Swoole\Server->start()` 启动服务即可

### server config 的实现细节

server config 的实现细节, 直接跟读代码定位到:

```php
// \Hyperf\Server\Server::initServers
    protected function initServers(ServerConfig $config)
    {
        $servers = $this->sortServers($config->getServers());

        foreach ($servers as $server) {
            $name = $server->getName();
            $type = $server->getType();
            $host = $server->getHost();
            $port = $server->getPort();
            $sockType = $server->getSockType();
            $callbacks = $server->getCallbacks();
            ...
```

- 先根据 swoole 的限制, 如果配置了多个 server, 需要先启动 http/ws server, 再通过 `addPort()` 的方式添加其余的 server, 这一步 hyperf 框架已经处理掉了
- 绑定 swoole server 需要 callback(回调函数)
- 触发 `BeforeMainServerStart / BeforeServerStart` 等事件

## 深入 server start

等等, server start 这么简单 ?! 再来回顾一下 `php bin/hyperf.php start`:

```bash
root@820d21e61cd8 /d/hyperf-demo# php bin/hyperf.php start
Scanning ...
Scan completed.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\Di\Listener\BootApplicationListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\Config\Listener\RegisterPropertyHandlerListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\RpcClient\Listener\AddConsumerDefinitionListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\Paginator\Listener\PageResolverListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\JsonRpc\Listener\RegisterProtocolListener listener.
[DEBUG] Event Hyperf\Framework\Event\BeforeMainServerStart handled by Hyperf\Amqp\Listener\BeforeMainServerStartListener listener.
[DEBUG] Event Hyperf\Framework\Event\BeforeMainServerStart handled by Hyperf\Process\Listener\BootProcessListener listener.
[DEBUG] Event Hyperf\Framework\Event\OnStart handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[DEBUG] Event Hyperf\Framework\Event\OnManagerStart handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[INFO] Worker#1 started.
[INFO] TaskWorker#2 started.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\AfterWorkerStartListener listener.
[DEBUG] Event Hyperf\Process\Event\BeforeProcessHandle handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[DEBUG] Event Hyperf\Process\Event\BeforeProcessHandle handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[DEBUG] Event Hyperf\Framework\Event\MainWorkerStart handled by Hyperf\Amqp\Listener\MainWorkerStartListener listener.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\AfterWorkerStartListener listener.
[INFO] Process[queue.default.0] start.
[DEBUG] Event Hyperf\Process\Event\BeforeProcessHandle handled by Hyperf\Process\Listener\LogBeforeProcessStartListener listener.
[INFO] Worker#0 started.
[INFO] Process[TestConsumer-hyperf.1] start.
[DEBUG] Event Hyperf\Process\Event\BeforeProcessHandle handled by Hyperf\Process\Listener\LogBeforeProcessStartListener listener.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[INFO] HTTP Server listening at 0.0.0.0:9501
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\AfterWorkerStartListener listener.
[INFO] TaskWorker#3 started.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[DEBUG] Event Hyperf\Framework\Event\AfterWorkerStart handled by Hyperf\Server\Listener\AfterWorkerStartListener listener.
[DEBUG] Event Hyperf\Process\Event\BeforeProcessHandle handled by Hyperf\Server\Listener\InitProcessTitleListener listener.
[INFO] Process[TestConsumer-hyperf.0] start.
[DEBUG] Event Hyperf\Process\Event\BeforeProcessHandle handled by Hyperf\Process\Listener\LogBeforeProcessStartListener listener.
```

是的, 使用了 Event 机制带来的灵活性, swoole Process 就是这个过程中启动的, 而 hyperf 中很多组件底层都是使用的 swoole Process 实现的, 这里只给出调用链路, 感兴趣的小伙伴自己去瞧哦:

```php
// 触发 BeforeMainServerStart 事件
// vendor/hyperf/server/src/Server.php:110
$this->eventDispatcher->dispatch(new BeforeMainServerStart($this->server, $config->toArray()));

// 事件监听器处理
\Hyperf\Process\Listener\BootProcessListener::process

// 启动 swoole process
\Hyperf\Process\AbstractProcess::bind

// 对应的 swoole 方法
\Swoole\Server::addProcess
```

## 写在最后

`start` 命令对源码阅读确实是一个相当有挑战的部分:

- 对 swoole 方法的封装, 要发挥出 swoole 超强能力, 所以说 swoole 的基础知识很重要
- hyperf 多个基础组件的联动, 包括 container / event / log / config

希望看到这里, 你可以感受到 hyperf 和 swoole 的强大之处, 头脑中大概能对 hyperf 运行的 `生命周期` 有个大致的了解

下篇我们开始协程的话题, 以及转到协程下编程必备的组件, 协程编程须知等内容.