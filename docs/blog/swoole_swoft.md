# swoole| swoole in swoft

上一篇 [blog - swoft 源码解读](http://www.jianshu.com/p/c3c472ff1414) 反响还不错, 不少同学推荐再加一篇, 讲解一下 swoft 中使用到的 swoole 功能, 帮助大家开启 swoole 的 **实战之旅**.

服务器开发涉及到的相关技术领域的知识非常多, 不日积月累打好基础, 是很难真正做好的. 所以我建议:

> [swoole wiki](https://wiki.swoole.com/) 最好看 3 遍, 包括评论. 第一遍快速过一遍, 形成大致印象; 第二遍边看边敲代码; 第三遍可以选择衍生的开源框架进行实战. [swoft](https://www.swoft.org/) 就是不错的选择.

swoole wiki 发展到现在已经 1400+ 页, 确实会有点难啃, **勇敢的少年呀**, 加油.

swoole 在 swoft 中的应用:

- `Swoole\Server`: swoole2.0 协程 Server
- `Swoole\HttpServer`: swoole2.0 协程 http Server, 继承自 `Swoole\Server`
- `Swoole\Coroutine\Client`: 协程客户端, swoole 封装了 tcp / http / redis / mysql
- `Swoole\Coroutine`: 协程工具集, 获取当前协程id，反射调用等能力

- `Swoole\Process`: 进程管理模块, 可以在 `Swoole\Server` 之外扩展更多功能

- `Swoole\Async`: 异步文件 IO
- `Swoole\Timer`: 基于 `timerfd + epoll` 实现的异步毫秒定时器，可完美的运行在 EventLoop 中
- `Swoole\Event`: 直接操作底层 `epoll/kqueue` 事件循环(EventLoop)的接口

- `Swoole\Lock`: 在 PHP 代码中可以很方便地创建一个锁, 用来实现数据同步
- `Swoole\Table`: 基于共享内存实现的超高性能数据结构

## Swoole\Http\Server

使用 swoole 的 http server 相较 tcp server 还是要简单一些, 只需要关心:

- `Swoole\Http\Server`
- `Swoole\Http\Request`
- `Swoole\Http\Response`

先看 http server:

```php
// \Swoft\Server\HttpServer
public function start()
{
    // http server
    $this->server = new \Swoole\Http\Server($this->httpSetting['host'], $this->httpSetting['port'], $this->httpSetting['model'], $this->httpSetting['type']);

    // 设置事件监听
    $this->server->set($this->setting);
    $this->server->on('start', [$this, 'onStart']);
    $this->server->on('workerStart', [$this, 'onWorkerStart']);
    $this->server->on('managerStart', [$this, 'onManagerStart']);
    $this->server->on('request', [$this, 'onRequest']);
    $this->server->on('task', [$this, 'onTask']);
    $this->server->on('pipeMessage', [$this, 'onPipeMessage']);
    $this->server->on('finish', [$this, 'onFinish']);

    // 启动RPC服务
    if ((int)$this->serverSetting['tcpable'] === 1) {
        $this->listen = $this->server->listen($this->tcpSetting['host'], $this->tcpSetting['port'], $this->tcpSetting['type']);
        $tcpSetting = $this->getListenTcpSetting();
        $this->listen->set($tcpSetting);
        $this->listen->on('connect', [$this, 'onConnect']);
        $this->listen->on('receive', [$this, 'onReceive']);
        $this->listen->on('close', [$this, 'onClose']);
    }

    $this->beforeStart();
    $this->server->start();
}
```

使用 swoole server 十分简单:

- 传入配置 server 配置信息, new 一个 swoole server
- 设置事件监听, 这一步需要大家对 swoole 的进程模型非常熟悉, 一定要看懂下面 2 张图
- 启动服务器

![进程流程图](http://upload-images.jianshu.io/upload_images/567399-a4212fa57c10af3e.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![进程/线程结构图](http://upload-images.jianshu.io/upload_images/567399-6b879104809d84d4.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

swoft 在使用 http server 时, 还会根据配置信息, 来判断是否同时新建一个 RPC server, 使用 swoole 的 [多端口监听](https://wiki.swoole.com/wiki/page/161.html) 来实现.

再来看 Request 和 Response, 提醒一下, 框架设计的时候, 要记住 **规范先行**:

> [PSR-7: HTTP message interfaces](http://www.php-fig.org/psr/psr-7/)

### Swoole\Http\Request

phper 比较熟悉的应该是 `$_GET $_POST $_COOKIE $_FILES $_SERVER` 这些全局变量, 这些在 swoole 中都得到了支持, 并且提供了更多方便的功能:

```php
// \Swoole\Http\Request $request
$request->get(); // -> $_GET
$request->post(); // -> $_POST
$request->cookie(); // -> $_COOKIE
$request->files(); // -> $_FILES
$request->server(); // -> $_SERVER

// 更方便的方法
$request->header(); // 原生 php 需要从 $_SERVER 中取
$request->rawContent(); // 获取原始的POST包体
```

这里强调一下 `$request->rawContent()`, phper 可能用 `$_POST` 比较 6, 导致一些知识不知道: post 的数据的格式. 因为这个知识, 所以 `$_POST` 不是所有时候都能取到数据的, 大家可以网上查找资料, 或者自己使用 postman 这样的工具自己测试验证一下. 在 `$_POST` 取不到数据的情况下, 会这样处理:

```php
$post = file_get_content('php://input');
```

`$request->rawContent()` 和这个等价的.

swoft 封装 Request 对象的方法, 和主流框架差不多, 以 laravel 为例(实际使用 symfony 的方法):

```php
// SymfonyRequest::createFromGlobals()
public static function createFromGlobals()
{
    // With the php's bug #66606, the php's built-in web server
    // stores the Content-Type and Content-Length header values in
    // HTTP_CONTENT_TYPE and HTTP_CONTENT_LENGTH fields.
    $server = $_SERVER;
    if ('cli-server' === PHP_SAPI) {
        if (array_key_exists('HTTP_CONTENT_LENGTH', $_SERVER)) {
            $server['CONTENT_LENGTH'] = $_SERVER['HTTP_CONTENT_LENGTH'];
        }
        if (array_key_exists('HTTP_CONTENT_TYPE', $_SERVER)) {
            $server['CONTENT_TYPE'] = $_SERVER['HTTP_CONTENT_TYPE'];
        }
    }

    $request = self::createRequestFromFactory($_GET, $_POST, array(), $_COOKIE, $_FILES, $server); // xglobal,

    if (0 === strpos($request->headers->get('CONTENT_TYPE'), 'application/x-www-form-urlencoded')
        && in_array(strtoupper($request->server->get('REQUEST_METHOD', 'GET')), array('PUT', 'DELETE', 'PATCH'))
    ) {
        parse_str($request->getContent(), $data);
        $request->request = new ParameterBag($data);
    }

    return $request;
}
```

### Swoole\Http\Response

`Swoole\Http\Response` 也是支持常见功能:

```php
// Swoole\Http\Response $response
$response->header($key, $value); // -> header("$key: $valu", $httpCode)
$response->cookie(); // -> setcookie()
$response->status(); // http 状态码
```

当然, swoole 还提供了常用的功能:

```php
$response->sendfile(); // 给客户端发送文件
$response->gzip(); // nginx + fpm 的场景, nginx 处理掉了这个
$response->end(); // 返回数据给客户端
$response->write(); // 分段传输数据, 最后调用 end() 表示数据传输结束
```

phper 注意下这里的 `write()` 和 `end()`, 这里有一个 http chunk 的知识点. 需要返回大量数据给客户端(>=2M)时, 需要分段(chunk)进行发送. 所以先用 `write()` 发送数据, 最后用 `end()` 表示结束. 数据量不大时, 直接调用 `end($html)` 返回就可以了.

在框架具体实现上, 和上面一样, laravel 依旧用的 `SymfonyResponse`, swoft 也是实现 PSR-7 定义的接口, 对 `Swoole\Http\Response` 进行封装.

## Swoole\Server

swoft 使用 `Swoole\Server` 来实现 RPC 服务, 其实在上面的多端口监听, 也是为了开启 RPC 服务. 注意一下单独启用中回调函数的区别:

```php
// \Swoft\Server\RpcServer
public function start()
{
    // rpc server
    $this->server = new Server($this->tcpSetting['host'], $this->tcpSetting['port'], $this->tcpSetting['model'], $this->tcpSetting['type']);

    // 设置回调函数
    $listenSetting = $this->getListenTcpSetting();
    $setting = array_merge($this->setting, $listenSetting);
    $this->server->set($setting);
    $this->server->on('start', [$this, 'onStart']);
    $this->server->on('workerStart', [$this, 'onWorkerStart']);
    $this->server->on('managerStart', [$this, 'onManagerStart']);
    $this->server->on('task', [$this, 'onTask']);
    $this->server->on('finish', [$this, 'onFinish']);
    $this->server->on('connect', [$this, 'onConnect']);
    $this->server->on('receive', [$this, 'onReceive']);
    $this->server->on('pipeMessage', [$this, 'onPipeMessage']); // 接收管道信息时触发的回调函数
    $this->server->on('close', [$this, 'onClose']);

    // before start
    $this->beforeStart();
    $this->server->start();
}
```

## Swoole\Coroutine\Client

swoole 自带的协程的客户端, swoft 都封装进了连接池, 用来提高性能. 同时, 为了业务使用方便, 既有协程连接, 也有同步连接, 方便业务使用时无缝切换.

同步/协程连接的实现代码:

```php
// RedisConnect -> 使用 swoole 协程客户端
public function createConnect()
{
    // 连接信息
    $timeout = $this->connectPool->getTimeout();
    $address = $this->connectPool->getConnectAddress();
    list($host, $port) = explode(":", $address);

    // 创建连接
    $redis = new \Swoole\Coroutine\Redis();
    $result = $redis->connect($host, $port, $timeout);
    if ($result == false) {
        App::error("redis连接失败，host=" . $host . " port=" . $port . " timeout=" . $timeout);
        return;
    }

    $this->connect = $redis;
}

// SyncRedisConnect -> 使用 \Redis 同步客户端
public function createConnect()
{
    // 连接信息
    $timeout = $this->connectPool->getTimeout();
    $address = $this->connectPool->getConnectAddress();
    list($host, $port) = explode(":", $address);

    // 初始化连接
    $redis = new \Redis();
    $redis->connect($host, $port, $timeout);
    $this->connect = $redis;
}
```

swoft 中实现连接池的代码在 `src/Pool` 下实现, 由三部分组成:

- Connect: 即上面代码中的连接
- Balancer: 负载均衡器, 目前实现了 随机/轮询 2 种方式
- Pool: 连接池, 调用 Balancer, 返回 Connect

详细内容可以参考之前的 [blog - swoft 源码解读](http://www.jianshu.com/p/c3c472ff1414)

## Swoole\Coroutine

作为首个使用 Swoole2.0 原生协程的框架, swoft 希望将协程的能力扩展到框架的核心设计中. 使用 `Swoft\Base\Coroutine` 进行封装, 方便整个应用中使用:

```php
public static function id()
{
    $cid = SwCoroutine::getuid(); // swoole 协程
    $context = ApplicationContext::getContext();

    if ($context == ApplicationContext::WORKER || $cid !== -1) {
        return $cid;
    }
    if ($context == ApplicationContext::TASK) {
        return Task::getId();
    }
    if($context == ApplicationContext::CONSOLE){
        return Console::id();
    }

    return Process::getId();
}
```

如同这段代码所示, Swoft 希望将方便易用的协程的能力, 扩展到 Console/Worker/Task/Process 等等不同的应用场景中

原生的 `call_user_func() / call_user_func_array()` 中无法使用协程 client, 所以 swoole 在协程组件中也封装的了相应的实现, swoft 中也有使用到, 请自行阅读源码.

## Swoole\Process

进程管理模块, 适合处理和 Server 比较独立的常驻进程任务, 在 swoft 中, 在以下场景中使用到:

- 协程定时器 CronTimerProcess
- 协程执行命令 CronExecProcess
- 热更新进程 ReloadProcess

swoft 使用 `\Swoft\Process` 对 `Swoole\Process` 进行了封装:

```php
// \Swoft\Process
public static function create(
    AbstractServer $server,
    string $processName,
    string $processClassName
) {
    ...

    // 创建进程
    $process = new SwooleProcess(function (SwooleProcess $process) use ($processClass, $processName) {
        // reload
        BeanFactory::reload();
        $initApplicationContext = new InitApplicationContext();
        $initApplicationContext->init();

        App::trigger(AppEvent::BEFORE_PROCESS, null, $processName, $process, null);
        PhpHelper::call([$processClass, 'run'], [$process]);
        App::trigger(AppEvent::AFTER_PROCESS);
    }, $iout, $pipe); // 启动 \Swoole\Process 并绑定回调函数即可

    return $process;
}
```

## Swoole\Async

swoft 在日志场景下使用 `Swoole\Async` 来提高性能, 同时保留了原有的同步方式, 方便进行切换

```php
// \Swoft\Log\FileHandler
private function aysncWrite(string $logFile, string $messageText)
{
    while (true) {
        $result = \Swoole\Async::writeFile($logFile, $messageText, null, FILE_APPEND); // 使用起来很简单
        if ($result == true) {
            break;
        }
    }
}
```

## Swoole\Event

服务器出于性能考虑, 通常都是 **常驻内存** 的, 传统的 `php-fpm` 也是, 修改了配置需要 reload 服务器才能生效. 也因为此, 服务器领域出现了新的需求 -- **热更新**. swoole 在进程管理上已经做了很多优化, 这里摘抄部分 wiki 内容:

```
Swoole提供了柔性终止/重启的机制
SIGTERM: 向主进程/管理进程发送此信号服务器将安全终止
SIGUSR1: 向主进程/管理进程发送SIGUSR1信号，将平稳地restart所有worker进程
```

目前大家采用的, 比较常见的方案, 是基于 Linux Inotify 特性, 通过监测文件变更来触发 swoole server reload. PHP 中有 Inotify 扩展, 方便使用, 具体实现在 `Swoft\Base\Inotify` 中:

```php
public function run()
{
    $inotify = inotify_init();

    // 设置为非阻塞
    stream_set_blocking($inotify, 0);

    $tempFiles = [];
    $iterator = new \RecursiveDirectoryIterator($this->watchDir);
    $files = new \RecursiveIteratorIterator($iterator);
    foreach ($files as $file) {
        $path = dirname($file);

        // 只监听目录
        if (!isset($tempFiles[$path])) {
            $wd = inotify_add_watch($inotify, $path, IN_MODIFY | IN_CREATE | IN_IGNORED | IN_DELETE);
            $tempFiles[$path] = $wd;
            $this->watchFiles[$wd] = $path;
        }
    }

    // swoole Event add
    $this->addSwooleEvent($inotify);
}

private function addSwooleEvent($inotify)
{
    // swoole Event add
    Event::add($inotify, function ($inotify) { // 使用 \Swoole\Event
        // 读取有事件变化的文件
        $events = inotify_read($inotify);
        if ($events) {
            $this->reloadFiles($inotify, $events);
        }
    }, null, SWOOLE_EVENT_READ);
}
```

## Swoole\Lock

swoft 在 CircuitBreaker(熔断器) 中的 HalfOpenState(半开状态) 使用到了, 并且这块的实现比较复杂, 推荐阅读源码:

```php
// CircuitBreaker
public function init()
{
    // 状态初始化
    $this->circuitState = new CloseState($this);
    $this->halfOpenLock = new \Swoole\Lock(SWOOLE_MUTEX); // 初始化互斥锁
}

// HalfOpenState
public function doCall($callback, $params = [], $fallback = null)
{
    // 加锁
    $lock = $this->circuitBreaker->getHalfOpenLock();
    $lock->lock();
    list($class ,$method) = $callback;

    ....

    // 释放锁
    $lock->unlock();

    ...
}
```

锁的使用, 难点主要在了解各种不同锁使用的场景, 目前 swoole 支持:

- 文件锁 SWOOLE_FILELOCK
- 读写锁 SWOOLE_RWLOCK
- 信号量 SWOOLE_SEM
- 互斥锁 SWOOLE_MUTEX
- 自旋锁 SWOOLE_SPINLOCK

## Swoole\Timer & Swoole\Table

定时器基本都会使用到, phper 用的比较多的应该是 crontab 了. 基于这个考虑, swoft 对 Timer 进行了封装, 方便 phper 用 **熟悉的姿势** 继续使用.

swoft 对 `Swoole\Timer` 进行了简单的封装, 代码在 `\Base\Timer` 中:

```
// 设置定时器
public function addTickTimer(string $name, int $time, $callback, $params = [])
{
    array_unshift($params, $name, $callback);

    $tid = \Swoole\Timer::tick($time, [$this, 'timerCallback'], $params);

    $this->timers[$name][$tid] = $tid;

    return $tid;
}

// 清除定时器
public function clearTimerByName(string $name)
{
    if (!isset($this->timers[$name])) {
        return true;
    }
    foreach ($this->timers[$name] as $tid => $tidVal) {
        \Swoole\Timer::clear($tid);
    }
    unset($this->timers[$name]);

    return true;
}
```

`Swoole\Table` 是在内存中开辟一块区域, 实现类似关系型数据库表(Table)这样的数据结构, 关于 `Swoole\Table` 的实现原理, rango 写过专门的文章 [swoole_table 实现原理剖析](https://segmentfault.com/a/1190000010853095), 推荐阅读.

`Swoole\Table` 在使用上需要注意以下几点:

- 类似关系型数据库, 需要提前定义好 **表结构**
- 需要预先判断数据的大小(行数)
- 注意内存, swoole 会更根据上面 2 个定义, 在调用 `\Swoole\Table->create()` 时分配掉这些内存

swoft 中则是使用这一功能, 来实现 crontab 方式的任务调度:

```php
private $originTable;
private $runTimeTable;

private $originStruct = [
    'rule'       => [\Swoole\Table::TYPE_STRING, 100],
    'taskClass'  => [\Swoole\Table::TYPE_STRING, 255],
    'taskMethod' => [\Swoole\Table::TYPE_STRING, 255],
    'add_time'   => [\Swoole\Table::TYPE_STRING, 11],
];

private $runTimeStruct = [
    'taskClass'  => [\Swoole\Table::TYPE_STRING, 255],
    'taskMethod' => [\Swoole\Table::TYPE_STRING, 255],
    'minte'      => [\Swoole\Table::TYPE_STRING, 20],
    'sec'        => [\Swoole\Table::TYPE_STRING, 20],
    'runStatus'  => [\Swoole\TABLE::TYPE_INT, 4],
];

// 使用 \Swoole\Table
private function createOriginTable(): bool
{
    $this->setOriginTable(new \Swoole\Table('origin', self::TABLE_SIZE, $this->originStruct));

    return $this->getOriginTable()->create();
}
```

## 写在最后

老生常谈了, 很多人吐槽 swoole 坑, 文档不好. 说句实话, **要敢于直面自己服务器开发能力不足的现实**. 我经常提的一句话:

> 要把 swoole 的 wiki 看 3 遍.

写这篇 blog 的初衷是给大家介绍一下 swoole 在 swoft 中的应用场景, 帮助大家尝试进行 swoole 落地. 希望这篇 blog 能对你有所帮助, 也希望你能多多关注 swoole 社区, 关注 swoft 框架, 能感受到服务器开发带来的乐趣.
