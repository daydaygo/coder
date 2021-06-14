# swoft| 源码解读系列一: feature

> 官网: https://www.swoft.org/

> 源码解读: http://naotu.baidu.com/file/814e81c9781b733e04218ac7a0494e2a?token=f009094c71a791c5

> 号外号外, [欢迎大家 star](https://github.com/swoft-cloud/swoft), 我们开发组定了一个 star 1000+ 就线下聚一次的小目标

继续源码解读系列. php 里面的 yii/laravel 框架算是非常「重」的了. 这里的 **重** 先不具体到 **性能** 层面, 主要是框架的设计思想和框架集成的服务, 让框架可以既可以快速解决很多问题, 又可以轻松扩展.

这次解读 swoft 的源码 -- 基于 swoole2.0 原生协程的框架. 同时, swoft 使用了大量 swoole 提供的功能, 也非常适合阅读它的代码, 来学习如何造轮子. 其实解读过 yii/laravel 这样的框架后, 一些 **通用** 的框架设计思想就不赘述了, 主要讲解和 **服务器开发** 相关的部分, 思路也会按照官网的 feature list 展开.

前半部分聚焦框架常用的功能:

- 全局容器注入 & MVC 分层设计
- 注解进制(**亮点, 强烈推荐了解一下**)
- 高性能路由
- 别名机制 `$aliases`
- RestFul风格
- 事件机制
- 强大的日志系统
- 国际化(i18n)
- 数据库 ORM

后半部分聚焦服务器相关的功能:

- 基础概念(**亮点, 第一个基于 swoole2.0 原生协程的框架**)
- 连接池
- 服务治理熔断、降级、负载、注册与发现
- 任务投递 & Crontab 定时任务
- 用户自定义进程
- Inotify 自动 Reload

> PHP 框架的设计, 可以参考 [PSR(PHP Standards Recommendations)](http://www.php-fig.org/psr/)

## 全局容器注入 & MVC 分层设计

之所以把这 2 个放一起讲, 是因为一个是 **里**, 一个是 **表**. 只是新人听得比较多的是 **MVC** 的分层设计思想, 全局容器注入了解相对较少.

- MVC 分层设计: 更偏向于业务

MVC 是一种简单通用并且实用的 **对业务进行拆分然后加以实现** 的设计, 本质还是 **分层设计**. 更重要的, 还是掌握 **分层设计** 的思想, 这个在工程实践中大量的使用到, 比如 OSI 7 层网络模型 和 TCP/IP 4 层网络模型. 我分层设计可以有效的确定 **系统边界和职责划分**.

想要培养分层设计的思想, 其实可以从 **拆** 入手, 在拆轮子然后拼轮子的过程中, 你会惊奇的发现, 艺术就在其中.

> 榫卯 app: https://www.douban.com/note/373132669/

- 全局容器注入

在进入这个概念之前, 先要认清另一个概念: **面向对象编程**. 更常用的可能是 **面向过程编程 vs 面向对象编程**. 这里不会长篇大论, 只就思维方式来进行比较:

1. 面向过程编程: 一条接一条指令的执行, 这是计算机喜欢的方式
2. 面向对象编程: 通过对象来 **抽象** 里面不同的事物, 通过事物之间的联系, 来解决与之相关的业务.

从这个角度来看, **面向对象** 可能是更符合人类的思维方式, 或者说更智能的思维方式:

> 上者劳人. 抽象好管理对象, 从而更好的完成任务.

但是使用面向对象编程的过程中, 就会出现一个问题: `new`, 需要管理好对象之间依赖关系, 全局容器注入就是做这样一件事. 使用 `new`, 表明一个对象需要依赖另一个对象, 但是使用容器, 则是一个对象告诉容器它需要什么对象.

> 怎么实现我不管 -- 这就是使用 `new` 和容器注入的区别, 学名叫 **控制反转**.

所以, 容器是 **里**, 在处理具体业务时, 由容器按需提供相应的 MVC 对象来处理.

## 注解机制

在容器的实现上, 或者说框架的底层上, 其实各个框架都 **大同小异**. 这里说一下 swoft 不同的地方 -- 引入注解进制.

简单解释一下注解进制: 通过添加注释 & 解析注释, 将注释转化为一些特定的有意义的代码.

> 更简单一点: 注释是给程序员用的, 注解是给程序用的

实现起来其实也很简单, 只是可能接触的比较少 -- **反射**:

```php
// Bean\Parser\InjectParser
class InjectParser extends AbstractParser
{
    /**
     * Inject注解解析
     *
     * @param string $className
     * @param object $objectAnnotation
     * @param string $propertyName
     * @param string $methodName
     *
     * @return array
     */
    public function parser(string $className, $objectAnnotation = null, string $propertyName = "", string $methodName = "", $propertyValue = null)
    {
        $injectValue = $objectAnnotation->getName();
        if (!empty($injectValue)) {
            return [$injectValue, true];
        }

        // phpdoc解析
        $phpReader = new PhpDocReader(); // 将注释转化为类
        $property = new \ReflectionProperty($className, $propertyName); // 使用反射
        $propertyClass = $phpReader->getPropertyClass($property);

        $isRef = true;
        $injectProperty = $propertyClass;
        return [$injectProperty, $isRef];
    }
}
```

如果熟悉 java, 会发现里面有很多地方在方法前用到了 `@override`, 在 symfony 中也使用到了这样的方式. 好处是一定程度的内聚, 使用起来更加简洁, 而且可以减少配置.

## 高性能路由

首先回答一个问题, 路由是什么? 从对象的角度出发, 其实路由就对应 **URL**. 那 URL 是什么呢?

> URL, Uniform Resource Locator, 统一资源定位符.

所以, 路由这一层抽象, 就是为了解决 -- 找到 URL 对应需要执行的逻辑.

现在再来解释一下 swoft 提到的高性能:

```php
// app/routes.php: 路由配置文件
$router = \Swoft\App::getBean('httpRouter'); // 通过容器拿 httpRouter

// config/beans/base.php: beans 配置文件
'httpRouter'      => [
    'class'          => \Swoft\Router\Http\HandlerMapping::class, // httpRouter 其实对应这个
    'ignoreLastSep'  => false,
    'tmpCacheNumber' => 1000,
    'matchAll'       => '',
],

// \Swoft\Router\Http\HandlerMapping
private $cacheCounter = 0;
private $staticRoutes = []; // 静态路由
private $regularRoutes = []; // 动态路由
protected function cacheMatchedParamRoute($path, array $conf){} // 会缓存匹配到的路由
// 路由匹配的方法也很简单: 校验 -> 处理静态路由 -> 处理动态路由
public function map($methods, $route, $handler, array $opts = [])
{
    ...
    $methods = static::validateArguments($methods, $handler);
    ...
    if (self::isNoDynamicParam($route)) {
        ...
    }
    ...
    list($first, $conf) = static::parseParamRoute($route, $params, $conf);
}
```

高性能 = 路由匹配逻辑简单 + 路由缓存

## 别名机制 `$aliases`

用过 yii 的对这个就比较熟悉了, 其实是这样一个 **进化过程**:

- 使用 `__DIR__` / `DIRECTORY_SEPARATOR` 等拼接出绝对路径
- 使用 `define() / defined()` 定义全局变量来使用路径
- 使用 `$aliases` 变量替代全局变量

这里只展示一下配置的地方, 实现只是在类中开一个变 `$aliases` 属性存储一下就行了:

```php
// config/define.php
// 基础根目录
!defined('BASE_PATH') && define('BASE_PATH', dirname(__DIR__, 1));
// 注册别名
$aliases = [
    '@root'       => BASE_PATH,
    '@app'        => '@root/app',
    '@res'        => '@root/resources',
    '@runtime'    => '@root/runtime',
    '@configs'    => '@root/config',
    '@resources'  => '@root/resources',
    '@beans'      => '@configs/beans',
    '@properties' => '@configs/properties',
    '@commands'   => '@app/Commands'
];
App::setAliases($aliases);
```

## RestFul风格

restful 的思想其实很简单: **以资源为核心, 业务其实是围绕资源的增删改查**. 具体到 http 中:

- url 只作为资源标识, 有 2 种形式, `item` 和 `item/id`, 后者表示操作具体某个资源
- http method(get/post/put等)用来对应资源的 CRUD
- 使用 json 格式进行数据的 **输入输出**

实现起来也很简单: 路由 + 返回

## 事件机制

先用 3W1H(who what why how) 分析法的思路来解释一下 **事件机制**, 更重要的是, 这个有什么用.

*正常*的程序执行, 或者说人的思维趋势, 都是按照 **时间线性串行** 的, 保持 **连续性**. 不过现实中会存在各种 **打断**, 程序也不是永远都是 **就绪状态**, 那么, 就需要有一种机制, 来处理可能出现的各种打断, 或者在程序不同状态之间切换.

事件机制发展到现在, 有时候也算是一种预留手段, 根据你的经验在需要的地方 **埋点**, 方便之后 **打补丁**.

swoft 的事件机制基于 [PSR-14](https://github.com/php-fig/fig-standards/blob/master/proposed/event-manager.md) 实现, 高度内聚简洁.

由三部分组成:

- EventManager: 事件管理器
- Event: 事件
- EventHandler / Listener: 事件处理器/监听器

执行流程:

- 先生成 EventManager
- 将 Event 和 EventHandler 注册到 EventManager
- 触发 Event, EventManager 就会调用相应的 EventHandler

使用起来就更加简单了:

```php
use Swoft\Event\EventManager;

$em = new EventManager;

// 注册事件监听
$em->attach('someEvent', 'callback_handler'); // 这里也可以使用注解机制, 实现事件监听注册

// 触发事件
$em->trigger('someEvent', 'target', ['more params']);

// 也可以
$event = new Event('someEvent', ['more params']);
$em->trigger($event);
```

来看一下 swoft 在事件机制这里用来提升性能的亮点:

```php
namespace Swoft\Event;

class ListenerQueue implements \IteratorAggregate, \Countable
{
    protected $store;

    /**
     * 优先级队列
     * @var \SplPriorityQueue
     */
    protected $queue;

    /**
     * 计数器
     * 设定最大值为 PHP_INT_MAX == 300
     * @var int
     */
    private $counter = PHP_INT_MAX;

    public function __construct()
    {
        $this->store = new \SplObjectStorage(); // Event 对象先添加都这里
        $this->queue = new \SplPriorityQueue(); // 然后加入优先级队列, 之后进行调度
    }
    ...
}
```

稍微玩过 ACM 的人对 **优先级队列** 就不会陌生了, 基本所有 OJ 都有相关的题库. 不过 PHPer 不用太操心底层实现, 直接借助 SPL 库即可.

> [SPL, Standard PHP Library](http://php.net/manual/en/book.spl.php), 类似 C++ 的 STL, PHPer 一定要了解一下.

## 强大的日志系统

使用 [monolog/monolog](https://packagist.org/packages/monolog/monolog) 来实现日志系统基本已成为标配了, 当然底层还是实现 [PSR-3](http://www.php-fig.org/psr/psr-3/) 标准. 不过这个标准出现比较早, 发展到现在, 隐藏得比较深了.

> 这也是建立技术标准/协议的理由, 划定好 **最佳实践**, 之后的努力都是朝着越来越易用发展.

swoft 的日志系统, 由 2 部分组成:

- `Swoft\Log\Logger`: 日志主体功能
- `Swoft\Log\FileHandler`: 输出日志

至于另一个文件, `Swoft\Log\Log`, 只是对 Logger 的一层封装, 调用起来更方便而已.

当然, swoft 的日志系统和 yii2 框架有明显相似的地方:

```php
// 都在 App 中快读暴露日志功能
public static function info($message, array $context = array())
{
    self::getLogger()->info($message, $context); // 其实还是使用 Logger 来处理
}

// 都添加了 profile 功能
public static function profileStart(string $name)
{
    self::getLogger()->profileStart($name);
}
public static function profileEnd($name)
{
    self::getLogger()->profileEnd($name);
}
```

值得一提的是, yii2 框架的日志系统由三部分组成:

- `Logger`: 日志主体功能
- `Dispatch`: 日志分发, 可以将同一个日志分发给不同的 Target 处理
- `Target`: 日志消费者

这样的设计, 其实是将 `FileHandler` 的功能进行拆解, 更灵活, 更方便扩展.

来看看 swoft 日志系统强大的一面:

```php
private function aysncWrite(string $logFile, string $messageText)
{
    while (true) {
        // 使用 swoole 异步文件 IO
        $result = \Swoole\Async::writeFile($logFile, $messageText, null, FILE_APPEND);
        if ($result == true) {
            break;
        }
    }
}
```

当然, 也可以选择同步的方式:

```php
private function syncWrite(string $logFile, string $messageText)
{
    $fp = fopen($logFile, 'a');
    if ($fp === false) {
        throw new \InvalidArgumentException("Unable to append to log file: {$this->logFile}");
    }
    flock($fp, LOCK_EX); // 注意要加锁
    fwrite($fp, $messageText);
    flock($fp, LOCK_UN);
    fclose($fp);
}
```

**PS**: 日志统计分析功能开发团队正在开发中, 欢迎大家推荐方案~

## 国际化(i18n)

这个功能的实现比较简单, 不过 i18n 这个词倒是可以多讲一句, 原词是 `internationalization`, 不过实在太长了, 所以简写为 `i18n`, 类似的还有 `kubernetes -> k8s`.

## 数据库 ORM

ORM 这个发展很也成熟了, 看清楚下面的进化史就好了:

- Statement: 直接执行 sql 语句
- QueryBuild: 使用链式调用, 来实现拼接 sql 语句
- ActiveRecord: Model, 用来映射数据库中的表, 实际还是封装的 QueryBuild

当然这一层层的封装好处也很明显, 减少 sql 的存在感.

```php
// insert
$post = new Post();
$post->title = 'daydaygo';
$post->save();

// query
$post = Post::find(1);

// update
$post->content = 'coder at work';
$post->save();

// delete
$post->del();
```

要实现这样的效果, 还是有一定的代码量的, 也会遇到一些问题, 比如 [代码提示](http://www.jianshu.com/p/b3daadb3c4c5), 还有一些更高级的功能, 比如 [关联查询](http://www.jianshu.com/p/fd85383783eb)

## 基本概念

- 并发 vs 并行

抓住 **并行** 这个范围更小的概念就容易理解了, 并行是要 **同时执行**, 那么只能多 cpu 核心同时运算才行; 并发则是因为 cpu运行和切换速度快, 时间段内执行多个程序, 宏观上 **看起来** 像在同时执行

- 协程 vs 进程

一种简单的说法 **协程是用户态的线程**. 线程由操作系统进行调度, 可以自动调度到多 cpu 上执行; 同一个时刻同一个 cpu 核心上只有一个协程运行, 当遇到用户代码中的阻塞 IO 时, 底层调度器会进入事件循环, 达到 **协程由用户调度** 的效果

- swoole2.0 原生

具体的实现原理大家到官网查看, 会有更详细的 wiki 说明, 我这里从 **工具** 使用的角度来说明一下

1. 限制条件一: 需要 swoole2.0 的协程 server + 协程 client 配合
2. 限制条件二: 在协程 server 的 onRequet, onReceive, onConnect 事件回调中才能使用

```php
$server = new Swoole\Http\Server('127.0.0.1', 9501, SWOOLE_BASE);

// 1: 创建一个协程
$server->on('Request', function($request, $response) {
    $mysql = new Swoole\Coroutine\MySQL();
    // 协程 client 有阻塞 IO 操作, 触发协程调度
    $res = $mysql->connect([
        'host' => '127.0.0.1',
        'user' => 'root',
        'password' => 'root',
        'database' => 'test',
    ]);
    // 阻塞 IO 事件就绪, 协程恢复执行
    if ($res == false) {
        $response->end("MySQL connect fail!");
        return;
    }
    // 出现阻塞 IO, 继续协程调度
    $ret = $mysql->query('show tables', 2);
    $response->end("swoole response is ok, result=".var_export($ret, true));
});

$server->start();
```

**注意**: 触发一次回调函数, 就会在开始的时候生成一个协程, 结束的时候销毁这个协程, 协程的生命周期, 伴随此处回调函数执行的生命周期

## 连接池

swoft 的连接池功能实现, 主要在 `src/Pool` 下, 主要由三部分组成:

- Connect: 连接, 值得一提的是, 为了后续使用方便, 这里同时配置了 同步连接 + 异步连接
- Balancer: 负载均衡器, 目前提供 2 种策略, 随机数 + 轮询
- Pool: 连接池, 核心部分, 负责连接的管理和调度

**PS**: 自由切换同步/异步客户端非常简单, 切换一下连接就好

直接上代码:

```php
// 使用 SqlQueue 来管理连接
public function getConnect()
{
    if ($this->queue == null) {
        $this->queue = new \SplQueue(); // 又见 Spl
    }

    $connect = null;
    if ($this->currentCounter > $this->maxActive) {
        return null;
    }
    if (!$this->queue->isEmpty()) {
        $connect = $this->queue->shift(); // 有可用连接, 直接取
        return $connect;
    }

    $connect = $this->createConnect();
    if ($connect !== null) {
        $this->currentCounter++;
    }
    return $connect;
}

// 如果接入了服务治理, 将使用调度器
public function getConnectAddress()
{
    $serviceList = $this->getServiceList(); // 从 serviceProvider 那里获取到服务列表
    return $this->balancer->select($serviceList);
}
```

## 服务治理熔断、降级、负载、注册与发现

swoft 的服务治理相关的功能, 主要在 `src/Service` 下:

- Packer: 封包器, 和协议进行对应, 看过 swoole 文档的同学, 就能知道协议的作用了
- ServiceProvider: 服务提供者, 用来对接第三方服务管理方案, 目前已实现 Consul
- Service: RPC服务调用, 包含同步调用和协程调用(`deferCall()`), 目前添加 callback 实现简单的 **降级**
- ServiceConnect: 连接池中 Connect 的 RPC Service 实现, **不过个人认为放到连接池中实现更好**
- Circuit: 熔断, 在 `src/Circuit` 中实现, 有三种状态, 关闭/开启/半开
- DispatcherService: 服务调度器, 在 Service 之前封装一层, 添加 Middleware/Event 等功能

这里看看熔断这部分的代码, 半开状态的逻辑复杂一些, 值得参考:

```php
// Swoft\Circuit\CircuitBreaker
public function init()
{
    // 状态初始化
    $this->circuitState = new CloseState($this);
    $this->halfOpenLock = new \swoole_lock(SWOOLE_MUTEX); // 使用 swoole lock
}

// Swoft\Circuit\HalfOpenState
public function doCall($callback, $params = [], $fallback = null)
{
    // 加锁
    $lock = $this->circuitBreaker->getHalfOpenLock();
    $lock->lock();
    ...
    // 释放锁
    $lock->unlock();
}
```

## 任务投递 & Crontab 定时任务

swoft 任务投递的实现机制当然离不开 `Swoole\Timer::tick()`(`\Swoole\Server->task()` 底层执行机制是一样的) , swoft 在实现的时候, 添加了 **喜闻乐见** 的 crontab 方式, 实现在 `src/Crontab` 下:

- ParseCrontab: 解析 crontab
- TableCrontab: 使用 `Swoole\Table` 实现, 用来存储 crontab 任务
- Crontab: 连接 Task 和 TableCrontab

这里主要看一下 TableCrontab:

```php
// 存储原始的任务
private $originStruct = [
    'rule'       => [\Swoole\Table::TYPE_STRING, 100],
    'taskClass'  => [\Swoole\Table::TYPE_STRING, 255],
    'taskMethod' => [\Swoole\Table::TYPE_STRING, 255],
    'add_time'   => [\Swoole\Table::TYPE_STRING, 11]
];
// 存储解析后的任务
private $runTimeStruct = [
    'taskClass'  => [\Swoole\Table::TYPE_STRING, 255],
    'taskMethod' => [\Swoole\Table::TYPE_STRING, 255],
    'minte'      => [\Swoole\Table::TYPE_STRING, 20],
    'sec'        => [\Swoole\Table::TYPE_STRING, 20],
    'runStatus'  => [\Swoole\TABLE::TYPE_INT, 4]
];
```

## 用户自定义进程

自定义进程对 `\Swoole\Process` 的封装, swoft 封装之后, 想要使用用户自定义进程更简单了:

继承 `AbstractProcess` 类, 并实现 `run()` 来执行业务逻辑.

swoft 中功能实现在 `src/Process` 下, 框架自带三个自定义进程:

- Reload: 配合 `ext-inotify` 扩展实现自动 reload, 下面会具体讲解
- CronTimer: crontab 里的 task 在这里触发 `\Swoole\Server->tick()`
- CronExec: 实现协程 task, 实现中.

代码就不贴了, 这里再扩展一个比较适合使用自定义进程的场景: **订阅服务**

## Inotify 自动 Reload

服务器程序大都是常驻进程, 有效减少对象的生成和销毁, 提供性能, 但是这样也给服务器程序的开发带来了问题, 需要 reload 来查看生效后的程序. 使用 `ext-inotify` 扩展可以解决这个问题.

直接上代码, 看看 swoft 中的实现:

```php
// Swoft\Process\ReloadProcess
public function run(Process $process)
{
    $pname = $this->server->getPname();
    $processName = "$pname reload process";
    $process->name($processName);

    /* @var Inotify $inotify */
    $inotify = App::getBean('inotify'); // 自定义进程来启动 inotify
    $inotify->setServer($this->server);
    $inotify->run();
}

// Swoft\Base\Inotify
public function run()
{

    $inotify = inotify_init(); // 使用 inotify 扩展

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
    swoole_event_add($inotify, function ($inotify) { // 使用 \Swoole\Event
        // 读取有事件变化的文件
        $events = inotify_read($inotify);
        if ($events) {
            $this->reloadFiles($inotify, $events); // 监听到文件变动进行更新
        }
    }, null, SWOOLE_EVENT_READ);
}
```

## 写在最后

再补充一点, 在实现服务管理(reload stop)时, 使用的 `posix_kill(pid, sig);`, 并不是用 `\Swoole\Server` 中自带的 `reload()` 方法, 因为我们当前环境的上下文并不一定在`\Swoole\Server` 中.

想要做好一个框架, 尤其是一个开源框架, 实际上要比我们平时写 **业务代码** 要难很多, 一方面是业务初期的 **多快好省**, 往往要上一些 **能跑** 的代码. 这里引入一些关于代码的观点:

- 代码质量: bug 率 + 性能
- 代码规范: 形成规范可以提高代码开发/使用的体验
- 代码复用: 这是软件工程的难题, 需要慢慢积累, 有些地方可以通过遵循规范走走捷径

总结起来就一句话:

> 想要显著提高编码水平或者快速积累相关技术知识, 参与开源可以算是一条捷径.
