# php-msf 源码解读

> php-msf: https://github.com/pinguo/php-msf

> 百度脑图 - php-msf 源码解读: http://naotu.baidu.com/file/cc7b5a49dfed46001d22222b1afa99ba?token=c9628331e99143c2

源码解读也做了一段时间了, 总结一下自己的心得:

- 抓住 **生命周期**, 让代码在你脑海中 **跑起来**
- 分析架构, 关键字 **分层** **边界** **隔离**

一个好的框架, 弄清楚 **生命周期** 和 **架构**, 基本就已经到了 **熟悉** 的状态了, 之后是填充细节和编码熟练了

这里再介绍几个次重要的心得:

- 弄明白这个工具擅长干什么, 适合干什么. 这个信息也非常容易获取到, 工具的文档通常都会显眼标注出来, 可以通过这些 功能/特性, 尝试以点见面
- 从工程化的角度去看这个项目, 主要和上面的 **架构** 区分, 在处理核心业务, 也就是上面的 功能/特性 外, 工程化还涉及到 安全/测试/编码规范/语言特性 等方面, 这些也是平时在写业务代码时思考较少并且实践较少的部分
- 工具的使用, 推荐我现在使用的组合: phpstorm + 百度脑图 + Markdown笔记 + blog

和 php-msf 的渊源等写技术生活相关的 blog 再来和大家八, 直接上菜.

## 生命周期 & 架构

官方文档制作了一张非常好的图: [处理请求流程图](https://pinguo.gitbooks.io/php-msf-docs/chapter-4/4.0-%E6%A1%86%E6%9E%B6%E7%BB%93%E6%9E%84.html). 推荐各位同仁, 有闲暇时制作类似的图, 对思维很有的帮助.

根据这张图来思考 **生命周期 & 架构**, 这里就不赘述了, 这里分析一下 msf 中一些技术点:

- 协程相关知识
- msf 中技术点摘录

## 协程

我会用我的方式来讲解, 如果需要深入了解的, 可以看我后面推荐的资源.

**类 vs 对象** 是一组很重要的概念. 类代表我们对事物的抽象, 这个抽象的能力在我们以后会一直用到, 希望大家有意识的培养这方面的意识, 至少可以起到触类旁通的作用. 对象是 **实例化** 的类, 是 **真正干活的**, 我们要讨论的 **协程**, 就是这样一个 **真正干活的** 角色.

> 协程从哪里来, 到哪里去, 它是干什么的?

想一想这几个简单的问题, 也许你对协程的理解就更深刻了, 记住这几个关键词:

- 产生. 需要有地方来产生协程, 你可能不需要知道细节, 但是需要知道什么时候发生了
- 调度. 肯定是有很多协程一起工作的, 所以需要调度, 怎么调度的呢?
- 销毁. 是否会销毁? 什么时候销毁?

现在, 我们再来看看协程的使用方式对比, 这里注意一下, 我没有用 **协程的实现方式对比**, 因为很多时候, 需求实际是这样的:

> 怎么实现我不管, 我选最好用的.

```php
// msf - 单次协程调度
$response = yield $this->getRedisPool('tw')->get('apiCacheForABCoroutine');

// msf - 并发协程调用
$client1 = $this->getObject(Client::class, ['http://www.baidu.com/']);
yield $client1->goDnsLookup();
$client2 = $this->getObject(Client::class, ['http://www.qq.com/']);
yield $client2->goDnsLookup();
$result[] = yield $client1->goGet('/');
$result[] = yield $client2->goGet('/');
```

**大致** 是这样的一个等式: `使用协程 = 加上 yield`, 所以搞清楚哪些地方需要加上 yield 就好了 -- 有阻塞IO的地方, 比如 文件IO, 网络IO(redis/mysql/http) 等.

当然, **大致** 就是还有需要注意的地方

- [协程调度顺序](https://pinguo.gitbooks.io/php-msf-docs/chapter-5/5.1-%E5%8D%8F%E7%A8%8B.html#%E5%8D%8F%E7%A8%8B%E7%9A%84%E8%B0%83%E5%BA%A6%E9%A1%BA%E5%BA%8F), 如果不注意, 就可能会退化成同步调用.
- 调用链: 使用 yield 的调用链上, 都需要加上 yield. 比如下面这样:

```php
function a_test() {
    return yield $this->getRedisPool('tw')->get('apiCacheForABCoroutine');
}
$res = yield a_test(); // 如果不加 yield, 就变成了同步执行
```

对比一下 swoole2.0 的协程方案:

```php
$server = new Swoole\Http\Server("127.0.0.1", 9502, SWOOLE_BASE);

$server->set([
    'worker_num' => 1,
]);

// 需要在协程 server 的异步回调函数中
$server->on('Request', function ($request, $response) {

    $tcpclient = new Swoole\Coroutine\Client(SWOOLE_SOCK_TCP); // 需要配合使用协程客户端
    $tcpclient->connect('127.0.0.1', 9501，0.5)
    $tcpclient->send("hello world\n");

    $redis = new Swoole\Coroutine\Redis();
    $redis->connect('127.0.0.1', 6379);
    $redis->setDefer(); // 标注延迟收包, 实现并发调用
    $redis->get('key');

    $mysql = new Swoole\Coroutine\MySQL();
    $mysql->connect([
        'host' => '127.0.0.1',
        'user' => 'user',
        'password' => 'pass',
        'database' => 'test',
    ]);
    $mysql->setDefer();
    $mysql->query('select sleep(1)');

    $httpclient = new Swoole\Coroutine\Http\Client('0.0.0.0', 9599);
    $httpclient->setHeaders(['Host' => "api.mp.qq.com"]);
    $httpclient->set([ 'timeout' => 1]);
    $httpclient->setDefer();
    $httpclient->get('/');

    $tcp_res  = $tcpclient->recv();
    $redis_res = $redis->recv();
    $mysql_res = $mysql->recv();
    $http_res  = $httpclient->recv();

    $response->end('Test End');
});
$server->start();
```

使用 swoole2.0 的协程方案, 好处很明显:

- 不用加 `yield` 了
- 并发调用不用刻意注意 `yield` 的顺序了, 使用 `defer()` 延迟收包即可

但是, 没办法直接用 `使用协程 = 加上 yield` 这样一个简单的等式了, 上面的例子需要配合使用 swoole 协程 server + swoole 协程 client:

- server 在异步回调触发时 **生成协程**
- client 触发 **协程调度**
- 异步回调执行结束时 **销毁协程**

这就导致了 2 个问题:

- 不在 swoole 协程 server 的异步回调中怎么办: 使用 `Swoole\Coroutine::create()` 显式生成协程
- 需要使用其他的协程 Client 怎么办: 这是 Swoole3 的目标, Swoole2.0 可以考虑用协程 task 来伪装

这样看起来, 好像 `使用协程 = 加上 yield` 这样要简单一些? 我不这样认为, 补充一些观点, 大家自己斟酌:

- 使用 yield 的方式, 基于 php 生成器 + 自己实现 PHP 协程调度器, 想要用起来不出错, 比如上面 **协程调度顺序**, 你还是需要去弄清楚这块的实现
- Swoole2.0 的原生方式, 理解起来其实更容易, 只需要知道协程 **生成/调度/销毁** 的时机就可以用好
- Swoole2.0 这样异步回调中频繁创建和销毁协程, 是否十分损耗性能? -- 不会的, 实际是一些内存操作, 比进程/对象小很多

想要继续深入了解的同学, 可以继续阅读下面的文章:

> php7 下协程实现: https://segmentfault.com/a/1190000012457145
> 在PHP中使用协程实现多任务调度 | 鸟哥: http://www.laruence.com/2015/05/28/3038.html
> php-msf doc - 协程原理: https://pinguo.gitbooks.io/php-msf-docs/chapter-2/2.3-%E5%8D%8F%E7%A8%8B%E5%8E%9F%E7%90%86.html
> swoole 底层和 php yield 实现协程, 本质是一样的, 毕竟都是 c 语言实现: https://github.com/php/php-src/blob/master/Zend/zend_generators.c

有时候 **书读百遍其义自见** 还是很有道理的. 我希望我使用的 **协程生成/调度/销毁** 的角度, 能给大家带来帮助.

> 感谢韩老大, msf 的开发者, swoft 的小伙伴, 在这个过程中, 对我给予的耐心帮助.

## msf 中技术点摘录

msf 在设计上有很多出彩的地方, 很多代码都值得借鉴.

### 请求上下文 Context

这是从 fpm 到 swoole http server 非常重要的概念. fpm 是多进程模式, 虽然 `$_POST` 等变量, 被称之为超全局变量, 但是, 这些变量在不同 fpm 进程间是隔离的. 但是到了 swoole http server 中, 一个 worker 进程, 会异步处理多个请求, 简单理解就是下面的等式:

```
fpm worker : http request = 1 : 1
swoole worker : http request = 1 : n
```

所以, 我们就需要一种新的方式, 来进行 request 间的隔离.

> 在编程语言里, 有一个专业词汇 scope(作用域). 通常会使用 `scope/生命周期`, 所以我一直强调的生命周期的概念, 真的很重要.

swoole 本身是实现了隔离的:

```php
$http = new swoole_http_server("127.0.0.1", 9501);
$http->on('request', function ($request, $response) {
    $response->end("<h1>Hello Swoole. #".rand(1000, 9999)."</h1>");
});
$http->start();
```

msf 在 Context 上还做了一层封装, 让 Context 看起来 **为所欲为**:

```php
// 你几乎可以用这种方式, 完成任何需要的逻辑
$this->getContext()->xxxModule->xxxModuleFunction();
```

细节可以查看 `src/Helpers/Context.php` 文件

### 对象池

对象池这个概念, 大家可能比较陌生, 目的是减少对象的频繁创建与销毁, 以此来提升性能, msf 做了很好的封装, 使用很简单:

```php
// getObject() 就可以了
/** @var DemoModel $demoModel */
$demoModel = $this->getObject(DemoModel::class, [1, 2]);
```

注意一下这行注释, 加上这个才有代码提示的效果的, 原理可以看我之前的 [blog - 聊一聊 php 代码提示](https://www.jianshu.com/p/b3daadb3c4c5)

对象池的具体代码在 `src/Base/Pool.php` 下:

- 底层使用反射来实现对象的动态创建

```php
public function get($class, ...$args)
{
    $poolName = trim($class, '\\');

    if (!$poolName) {
        return null;
    }

    $pool     = $this->map[$poolName] ?? null;
    if ($pool == null) {
        $pool = $this->applyNewPool($poolName);
    }

    if ($pool->count()) {
        $obj = $pool->shift();
        $obj->__isConstruct = false;
        return $obj;
    } else {
        // 使用反射
        $reflector         = new \ReflectionClass($poolName);
        $obj               = $reflector->newInstanceWithoutConstructor();

        $obj->__useCount   = 0;
        $obj->__genTime    = time();
        $obj->__isConstruct = false;
        $obj->__DSLevel    = Macro::DS_PUBLIC;
        unset($reflector);
        return $obj;
    }
}
```

感兴趣的同学可以去了解一下 [**反射**](http://php.net/manual/en/book.reflection.php), 可以给语言增加很多灵活性

- 使用 SplStack 来管理对象

```php
private function applyNewPool($poolName)
{
    if (array_key_exists($poolName, $this->map)) {
        throw new Exception('the name is exists in pool map');
    }
    $this->map[$poolName] = new \SplStack();

    return $this->map[$poolName];
}

// 管理对象
$pool->push($classInstance);
$obj = $pool->shift();
```

> msf doc 这块的文章非常值得一读, 特别是 **php进程内存优化**, 对我触动很大: https://pinguo.gitbooks.io/php-msf-docs/chapter-5/5.6-%E5%AF%B9%E8%B1%A1%E6%B1%A0.html

### 连接池 & 代理

- 连接池 Pools

连接池的概念就不赘述了, 我们来直接看 msf 中的实现, 代码在 `src/Pools/AsynPool.php` 下:

```php
public function __construct($config)
{
    $this->callBacks = [];
    $this->commands  = new \SplQueue();
    $this->pool      = new \SplQueue();
    $this->config    = $config;
}
```

这里使用的 `SplQueue` 来管理连接和需要执行的命令. 可以和上面对比一下, 想一想为什么一个使用 `SplStack`, 一个使用 `SplQueue`.

- 代理 Proxy

代理是在连接池的基础上进一步的封装, msf 提供了 2 种封装方式:

- 主从 master slave
- 集群 cluster

查看示例 `App\Controllers\Redis` 中的代码:

```php
class Redis extends Controller
{
    // Redis连接池读写示例
    public function actionPoolSetGet()
    {
        yield $this->getRedisPool('p1')->set('key1', 'val1');
        $val = yield $this->getRedisPool('p1')->get('key1');

        $this->outputJson($val);
    }

    // Redis代理使用示例（分布式）
    public function actionProxySetGet()
    {
        for ($i = 0; $i <= 100; $i++) {
            yield $this->getRedisProxy('cluster')->set('proxy' . $i, $i);
        }

        $val = yield $this->getRedisProxy('cluster')->get('proxy22');
        $this->outputJson($val);
    }

    // Redis代理使用示例（主从）
    public function actionMaserSlaveSetGet()
    {
        for ($i = 0; $i <= 100; $i++) {
            yield $this->getRedisProxy('master_slave')->set('M' . $i, $i);
        }

        $val = yield $this->getRedisProxy('master_slave')->get('M66');
        $this->outputJson($val);
    }
}
```

代理就是在连接池的基础上进一步 **搞事情**. 以 **主从** 模式为例:

- 主从策略: 读主库, 写从库

代理做的事情:

- 判断是读操作还是写操作, 选择相应的库去执行

### 公共库

msf 推行 **公共库** 的做法, 希望不同功能组件可以做到 **可插拔**, 这一点可以看 laravel 框架和 symfony 框架, 都由框架核心加一个个的 package 组成. 这种思想我是非常推荐的, 但是仔细看 [百度脑图 - php-msf 源码解读](http://naotu.baidu.com/file/cc7b5a49dfed46001d22222b1afa99ba?token=c9628331e99143c2) 这张图的话, 就会发现类与类之间的依赖关系, **分层/边界** 做得并不好. 如果看过我之前的 [blog - laravel源码解读](https://www.jianshu.com/p/b7ea3f2a55f6) / [blog - yii源码解读](https://www.jianshu.com/p/fd85383783eb), 进行对比就会感受很明显.

但是, 这并不意味着 **代码不好**, 至少功能正常的代码, 几乎都能算是好代码. 从功能之外建立的 **优越感**, 更多的是对 **美好生活的向往** -- 还可以更好一点.

### AOP
> php AOP 扩展: http://pecl.php.net/package/aop
> PHP-AOP扩展介绍 | rango: http://rango.swoole.com/archives/83

AOP, 面向切面编程, 韩老大 的 [blog - PHP-AOP扩展介绍 | rango](http://rango.swoole.com/archives/83) 可以看看.

需不需要了解一个新事物, 先看看这个事物有什么作用:

> AOP, 将业务代码和业务无关的代码进行分离, 场景有 日志记录 / 性能统计 / 安全控制 / 事务处理 / 异常处理 / 缓存 等等.

这里引用一段 [程序员DD - 翟永超的公众号](http://blog.didispace.com/) 文章里的代码, 让大家感受下:

- 同样是 CRUD, 不使用 AOP

```java
＠PostMapping("/delete")
public Map<String, Object> delete(long id, String lang) {
  Map<String, Object> data = new HashMap<String, Object>();

  boolean result = false;
  try {
    // 语言（中英文提示不同）
    Locale local = "zh".equalsIgnoreCase(lang) ? Locale.CHINESE : Locale.ENGLISH;

    result = configService.delete(id, local);

    data.put("code", 0);

  } catch (CheckException e) {
    // 参数等校验出错，这类异常属于已知异常，不需要打印堆栈，返回码为-1
    data.put("code", -1);
    data.put("msg", e.getMessage());
  } catch (Exception e) {
    // 其他未知异常，需要打印堆栈分析用，返回码为99
    log.error(e);

    data.put("code", 99);
    data.put("msg", e.toString());
  }

  data.put("result", result);

  return data;
}
```

- 使用 AOP

```java
＠PostMapping("/delete")
public ResultBean<Boolean> delete(long id) {
  return new ResultBean<Boolean>(configService.delete(id));
}
```

代码只用一行, 需要的特性一个没少, 你是不是也想写这样的 CRUD 代码?

### 配置文件管理

先明确一下配置管理的痛点:

- 是否支撑热更新, 常驻内存需要考虑
- 考虑不同环境: dev test production
- 方便使用

热更其实可以算是常驻内存服务器的整体需求, 目前 php 常用的解决方案是 inotify, 可以参考我之前的 [blog - swoft 源码解读](https://www.jianshu.com/p/c3c472ff1414) .

msf 使用第三方库来解析处理配置文件, 这里着重提一个 `array_merge()` 的细节:

```php
$a = ['a' => [
    'a1' => 'a1',
]];

$b = ['a' => [
    'b1' => 'b1',
]];

$arr = array_merge($a, $b); // 注意, array_merge() 并不会循环合并
var_dump($arr);

// 结果
array(1) {
  ["a"]=>
  array(1) {
    ["b1"]=>
    string(2) "b1"
  }
}
```

msf 中使用配置:

```php
$ids = $this->getConfig()->get('params.mock_ids', []);

// 对比一下 laravel
$ids = cofnig('params.mock_ids', []);
```

看起来 laravel 中要简单一些, 其实是通过 composer autoload 来加载函数, 这个函数对实际的操作包装了一层. 至于要不要这样做, 就看自己需求了.

## 写在最后

msf 最复杂的部分在 **服务启动阶段**, 继承也很长:

`Child -> Server -> HttpServer -> MSFServer -> AppServer`, 有兴趣可以挑战一下.

另外一个比较难的点, 是 `MongoDbTask 实现原理`.

msf 还封装了很多有用的功能, RPC / 消息队列 / restful, 大家根据文档自己探索即可.
