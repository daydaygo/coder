# swoole| swoole 协程用法笔记

> [swoole| swoole 协程用法笔记](https://www.jianshu.com/p/28e882352da5)

- [swoole源码解读-Coroutine实现](https://segmentfault.com/a/1190000019089997)
- [swoole协程之旅](https://wiki.swoole.com/wiki/page/1044.html)
- [swoole协程实现原理](https://wiki.swoole.com/wiki/page/p-coroutine_realization.html)

## 协程方法一览

协程方法简明笔记:
- [`Coroutine::set`](https://wiki.swoole.com/wiki/page/1036.html) 协程设置
    - `max_coroutine`
- `Coroutine::stats` 协程状态
- `Coroutine::list` 遍历当前进程中的所有协程
- `Coroutine::getBackTrace` 查看协程调用栈
    - 别名 `Coroutine::listCoroutines()`
- [`Coroutine::create`](https://wiki.swoole.com/wiki/page/774.html) 创建协程
    - 短名 `go()`
    - 协程开销: 推荐使用 php>=7.2, 创建初始分配的栈内存更小, 并且会自动扩容
- `Coroutine::defer` 类似 go 的 defer, PHP 有析构函数(`__destruct()`)和自动回收机制, defer 的使用范围没那么大
    - 短名 `defer()`
- `Coroutine\Channel` 类似 go 的 chan, 所有操作均为内存操作, 进程间内存隔离
    - 短名 `chan()`
    - 初始化是需要设置容量(capacity), 通道 `空/满` 会影响到后续的 push / pop
    - 必须在 swoole server 的 onWorkerStart 之后创建
- `Coroutine::getCid` 当前协程id, 即 cid
- `Coroutine::exist`
- `Coroutine::getPcid` 获取父协程cid
    - 协程的嵌套会带来初始的先后顺序(父子关系), 最终执行还是要看协程的调度(没有稳定的父子关系)
- `Coroutine::getContext`
- `Coroutine::yield` 让出当前协程的执行权, 需要配合 resume 使用, 由其他协程唤醒
    - 别名 `Coroutine::suspend`
- `Coroutine::resume` 唤醒其他协程

- `Coroutine::exec` 协程版 `shell_exec()`

- `Coroutine::gethostbyname` DNS查询, todo -> http client 是否需要
- `Coroutine::getaddrinfo`

- `Coroutine::statvfs` 获取文件系统信息(目前还不知道使用场景)

- `Coroutine\Client`
- `Coroutine\Http\Client`
- `Coroutine\Http2\Client`
- `Coroutine\Socket`
- `Coroutine\PostgreSQL`
    - 安装 swoole 时, 需要加编译参数 `--enable-coroutine-postgresql`
    - 系统需要安装 libpg 库

开启协程 runtime 后(`\Swoole\Runtime::enableCoroutine()`), 可以不再使用:
- `Coroutine::fread`
- `Coroutine::fgets`
- `Coroutine::fwrite`
- `Coroutine::sleep`
- `Coroutine::readFile`
- `Coroutine::writeFile`
- `Coroutine\Redis` 使用 ext-redis(phpredis) / predis
- `Coroutine\MySQL` 使用 mysqlnd 模式的 pdo、mysqli 扩展

## 阻塞代码检验

swoole 中使用协程的 2 个要点:
- 开协程: 这个容易, `go()` 一下就行了
- 协程中执行 **非阻塞代码**: 除了看官方文档, 下面提供一个简单的检测 demo

```php
go(function () {
    sleep(1); // 未开启协程 runtime, 此处会阻塞, 输出为 go -> main
    echo "go \n";
});
echo "main \n";
```

输出为: `go -> main`

```php
\Swoole\Runtime::enableCoroutine();

go(function () {
    sleep(1); // 开启协程 runtime, 此处为阻塞, 输出为 main -> go
    echo "go \n";
});
echo "main \n";
```

输出为: `main -> go`, 发生了协程调度.

使用时将 `sleep(1)` 替换为需要检测的代码即可.

## 对短名称的个人看法

- 建议关闭, 全部使用 `\Swoole\Coroutine` 命名空间保持一致性, 按需封装常用的几个, 比如 `go() chan() defer()`
- ini 配置: swoole.use_shortname = 'Off'

```php
if (! function_exists('go')) {
    function go(callable $callable)
    {
        \Hyperf\Utils\Coroutine::create($callable);
    }
}

if (! function_exists('defer')) {
    function defer(callable $callable): void
    {
        \Hyperf\Utils\Coroutine::defer($callable);
    }
}
```

## 其他

- swoole 协程中比较重要的参数设置 `max_coroutine`, 更科学方式(**看是否有需要**): 压测后查看内存占用, 进而进行调整
- swoole 协程参数设置: [`enable_preemptive_scheduler` 抢占式调度](https://segmentfault.com/a/1190000019253487)
- [wiki - 协程编程须知](https://wiki.swoole.com/wiki/page/851.html): 使用协程的注意事项
- [协程 go+chan+defer](https://wiki.swoole.com/wiki/page/p-csp.html): chan可以用在协程间交互, defer使用场景有待收集
- [实现 defer](https://wiki.swoole.com/wiki/page/p-go_defer.html): `__destruct()`
- [实现 waitgroup](https://wiki.swoole.com/wiki/page/p-waitgroup.html): chan+count计数, 可以进一步封装成更方便的写法
- [版本更新记录](https://wiki.swoole.com/wiki/page/p-project/change_log.html): 仔细看看, 就能体会到开发组在php 协程上的努力

## 并发调用

官方提供了 2 种方式:
- setDefer 机制: 延迟收包, 多个请求并发收取相应结果; 确保协程组件支持 setDefer 机制(绝大部分协程组件都支持)
- 子协程 + chan: 每个请求在子协程中执行, 通过 chan 实现请求收包时的协程调度

个人不建议使用, 并发调用可以达到更好的性能, 但是不符合常用的编程习惯, 多个请求需要同时完成, 推荐使用 `waitGroup`, 在此基础上, 还可以封装成更简单的写法

```php
class WaitGroup
{
    /**
     * @var int
     */
    private $counter = 0;

    /**
     * @var SwooleChannel
     */
    private $channel;

    public function __construct()
    {
        $this->channel = new chan();
    }

    public function add(int $incr = 1): void
    {
        $this->counter += $incr;
    }

    public function done(): void
    {
        $this->channel->push(true);
    }

    public function wait(): void
    {
        while ($this->counter--) {
            $this->channel->pop();
        }
    }
}
```