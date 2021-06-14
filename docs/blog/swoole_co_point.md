# swoole| swoole 协程知识点小结

本文要点:
- swoole 协程现状一览: 学不动? 其实是更简单了
- 使用 swoole 协程很简单: 开个协程, 协程里写非阻塞代码
- 展望 swoole 协程未来

## swoole 协程现状一览

swoole 一直保持着 **颇为快速** 的迭代速度, 快到什么程度呢 -- 「快别更新了, 学不动了」
- 半年前还是 v4.0 支持完整的协程编程(CSP, go+chan), 现在已经迭代到 v4.3.3
- v4.3 版本做了一次大更新, 项目拆分成了 [`swoole`](https://github.com/swoole/swoole-src) 和 [`swoole_async`](https://github.com/swoole/ext-async)
- 官方 wiki 修改了很多, 协程的部分的文档增加了不少, 而且提前到更加显眼的地方

> 来一句灵魂叩问: 改动这么大, 那是不是真的 「学不动了」?

并不是! swoole 是一直在为 **世界上最好的语言** 添砖加瓦:
- 更为完整的协程编程支持, 直观的效果是更加 **无缝无感** 的编程切换体验(后面细说), 意味着需要了解和注意的语法细节更少, 编程更轻松
- v4.3 版本做的一次大更新, 实际是优化 swoole 项目的架构, 主项目 focus 协程模式的网络编程, 更多网络编程相关的功能, 使用 **扩展(ext)** 的方式提供(具体可以参考 swoole 的 github 主页: https://github.com/swoole, 扫一眼下面有的项目, 就能有所启发)
- 官方 wiki 一直以信息量著称(同时也意味着 **可以学到很多东西**), 但是如果具备了 **网络编程 + 协程** 的基础知识, 然后 focus 到 [swoole 协程部分文档](https://wiki.swoole.com/wiki/page/p-coroutine.html) 上, 就会发现其实都是一些 **编程语法**, so easy~

## 使用 swole 协程

如何使用协程: 
- 使用 `go()(\Swoole\Coroutine::create() 的简写)` 创建一个协程
- 在 go() 的回调函数中, 加入协程需要执行的代码, 注意是 **非阻塞代码**

```php
use Swoole\Coroutine as Co; // 常用的缩写方式

go(function () { // 创建协程, 回调函数中写需要在协程中执行的代码
    echo "daydaygo";
    Co::sleep(1); // 不能是阻塞代码
});
```

> swoole 中协程就是这么简单: 开个协程, 协程里写非阻塞代码. 官方协程部分文档看起来很多, 牢记这两点, 其实很简单!

## 开协程

- 上文提到的, 使用 `go()` 创建一个协程
- swoole server 中, 底层自动在 onRequet, onReceive, onConnect 等事件回调之前自动创建一个协程
    - [开启 `enable_coroutine` 参数后的影响范围](https://wiki.swoole.com/wiki/page/949.html): 主要还包括 Timer 定时器
- 使用 `task_enable_coroutine` 开启的协程版 Task 进程, 会在 onTask 回调之前自动创建一个协程
- 进程和进程池支持开启协程, 开启后创建的子进程会自动创建协程

```php
// tcp/udp server, 可以在此基础可封装 rpc
$s = new \Swoole\Server();

// http server, 替代传统的 fpm
$s = new \Swoole\Http\Server();

// 开启 http2 支持: https://wiki.swoole.com/wiki/page/326.html
$s = new \Swoole\Http\Server();
$s->set([
    'open_http2_protocol' => true,
]);
// 进而可以实现基于 http2 的服务, 比如 grpc

// websocket server
$s = new \Swoole\WebSocket\Server();
```

## 非阻塞代码

协程中必须编写 **非阻塞代码**, 看到上面 `Co::sleep(1)` 的小伙伴会有疑问了: 

> 连个 sleep 都要 swoole 提供一个协程版, 我得掌握多少 swoole 协程版 API 才够呀?

所以问题的关键点来了:

- 协程中一定要使用 **非阻塞代码**(一定要牢记, 多次重复了, 可以心里再默念三遍)
- 怎么区分哪些是阻塞的, 哪些是非阻塞的: 可以参考 [官方wiki - runtime](https://wiki.swoole.com/wiki/page/p-runtime.html)
- 随着 swoole 的迭代, 对协程的支持越来越完整, `区分哪些阻塞, 哪些非阻塞`, 越来越无感

swoole 更新后, 添加了开启协程 runtime 功能:

```php
// 没有开启协程 runtime, 需要协程版 API
use Swoole\Coroutine as Co;

go(function () {
    echo "daydaygo";
    Co::sleep(1); // 需要使用 swoole 提供的协程版 API
});

// 开启协程 runtime
\Swoole\Runtime::enableCoroutine();
go(function () {
    echo "daydaygo";
    sleep(1); // 和原来编程一样了
});
```

协程 runtime 开启后, 支持的列表:
- redis扩展
- 使用mysqlnd模式的pdo、mysqli扩展，如果未启用mysqlnd将不支持协程化
- soap扩展
- `file_get_contents、fopen`
- `stream_socket_client (predis)`
- `stream_socket_server`
- `stream_select`(需要4.3.2以上版本)
- fsockopen
- [文件操作](https://wiki.swoole.com/wiki/page/991.html) 底层使用 AIO 线程池模拟实现
    - fopen / fclose
    - fread / fwrite 
    - fgets / fputs
    - `file_get_contents / file_put_contents`
    - unlink / mkdir / rmdir
- [sleep系列函数](https://wiki.swoole.com/wiki/page/992.html)
    - sleep / usleep
    - `time_nanosleep / time_sleep_until`

不支持的列表:
- mysql：底层使用libmysqlclient
- curl：底层使用libcurl （即不能使用CURL驱动的Guzzle）
- mongo：底层使用mongo-c-client
- `pdo_pgsql / pdo_ori / pdo_odbc / pdo_firebird`

## 协程 runtime 还不支持怎么办

需要的功能协程 runtime 下还没支持怎么办? 你有三种方法:

- 协程 runtime 之前, 官方和社区已经贡献了很多协程版 API 可供使用, 比如上面 `Co::sleep(1)`, PostgreSQL Client `Swoole\Coroutine\PostgreSQL`
- 官方和社区没有, 可以使用 swoole 提供的协程版 client 进行封装, 可以参考 [官方 amqp client 封装](https://github.com/swoole/php-amqplib), 将 socket() 函数实现的 tcp client, 使用 swoole 协程版 tcp client 实现即可

```php
public function connect()
{
    // 使用 Swoole\Coroutine\Client 
    $sock = new Swoole\Coroutine\Client(SWOOLE_SOCK_TCP);
    if (!$sock->connect($this->host, $this->port, $this->connection_timeout))
    {
        throw new AMQPRuntimeException(
            sprintf(
                'Error Connecting to server(%s): %s ',
                $sock->errCode,
                swoole_strerror($sock->errCode)
            ),
            $sock->errCode
        );
    }
    $this->sock = $sock;
}
```

PS: 这里只是抛砖引玉, 原库中使用 `stream_socket_client()`, 现在 swoole 协程 runtime 已经支持了

- 传统的阻塞解决方案(当然是在现有的协程方式都不行, 才会继续使用传统的方式): 抛给 swoole 的 task 进程, 使用 MQ 异步掉, 等等

## 展望 swoole 协程的未来

到目前为止, 希望小伙伴们已经 get 到了 swoole 中协程编程的要点(我喜欢用 **姿势**, 人最紧要的是姿势好看~), 让我们展望一下未来:

- 解锁更多协程使用: chan, defer, select, waitgroup, 这些官方都提供了 demo( [韩天峰 - PHP 协程：Go + Chan + Defer](https://segmentfault.com/a/1190000017243966)), 看完后自己也能封装一份
- 并不是 **完整(100%, one hundred percent)** 的支持, 要是一不小心踩到不支持的怎么办? swoole 的后续版本将支持检测协程环境下是否有阻塞调用
- 随着 swoole 官方在协程编程上的持续发力, 基于 swoole 实现的全协程式 PHP 开发框架也将更为简单, 从基础/底层的网络编程到整个微服务架构的道路也将更为平坦, 比如马上将要迎来大版本升级的 [swoft2](https://github.com/swoft-cloud/swoft)

## 写在最后

官方协程部分的文档看起来多, 其实多是对协程 API 的介绍, 并没有在知识结构理解的复杂度上有所增加. swoole 中协程的编程语法, 都在 `\Swoole\Coroutine` 命名空间下可以找到.

最后回到一个经典问题, 学习 swoole 的协程好, 还是学习 go 的协程好? 我谈谈我个人的观点:
- 所需要的基础知识: 网络编程 + 协程, 不会因为你是用 swoole 还是 go 而有所减少, 基础不大好, 表现出来了就是学着学着就容易卡住, 效率上不来
- 以为你写的是 swoole, 不不不, 写的是一个又一个功能的 API, go 也同样(要用到 redis/mysql/mq, 相应的 API 你还是得学得会), 区别在于, swoole 趋势是在底层实现支持(比如 协程runtime), 这样 PHPer 可以无缝切换过来, 而 Gopher 则需要学习一个又一个基于 go 协程封装好的 API. 当初在 PHP 中学习的这些 API, 到 go 里面, 一样需要再熟悉一遍
- 最后来谈谈性能, 请允许我用一个傲娇一点的说, **你用 swoole 达不到的性能, 换个语言, 呵呵呵.** 难易程度排行: `加机器 < 加程序员 < 加语言`.