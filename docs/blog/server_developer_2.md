# coder| 服务器开发系列 2

> blog: https://www.jianshu.com/p/19b4d1c527eb

经过 3 周的疯狂加班后, 服务器开发节奏终于可以放一放了, 也有空可以用「外在」的角度来好好看一下这次项目.

## 使用 swoole 裸写 tcp server

在没有「瑞士军刀」(熟悉的框架)的情况下, 裸写 swoole 就变成了下面这样:

- swoole tcp server 基本骨架

```php

require_once __DIR__ . '/../vendor/autoload.php'; // composer autoload
require_once __DIR__ . '/config.php'; // 配置文件

//---------------------server
$serv = new swoole_server("0.0.0.0", 9999);
$serv->set([
    'worker_num'            => 4,
    'task_worker_num'       => 8,
//    'daemonize'             => true,
    'pid'                   => __DIR__ . '/server.pid',
    'log_file'              => __DIR__ . '/../log/swoole.log',

    // 固定包头协议
    'open_length_check'     => 1,       // 开启协议解析
    'package_length_type'   => 'N',     // 长度字段的类型
    'package_length_offset' => 4,       // 第N个字节是包长度的值
    'package_body_offset'   => 8,       // 第N个字节开始计算长度
    'package_max_length'    => 2000000, // 协议最大长度
]);

// swoole table 是最开始的方案, 可以作为全局共享内存使用
//$swooleTable = new swoole_table('100000'); // 最多 10w 同时连接
//$swooleTable->column('auth', swoole_table::TYPE_INT, '1');
//$swooleTable->create();
//$serv->table = $swooleTable;

$serv->zoneId = $config['zone_id']; // 使用配置
$serv->userinfo = []; // 保存信息

// 必须在 onWorkerStart 回调中创建 redis/mysql 连接
$serv->on('workerStart', 'onWorkerStart');

// mysql 连接池
$serv->on('task', 'onTask');
$serv->on('finish', 'onFinish');

$serv->on('connect', 'onConnect');
$serv->on('receive', 'onReceive'); // 消息处理
$serv->on('close', 'onClose');

$serv->start();
```

- 消息处理

这里使用了很多了 `function` 来, 来分离开逻辑

```php
function onReceive(swoole_server $serv, $fd, $from_id, $data)
{
    // 使用 autoload psr-4 来加载协议解析的函数
    $data = decode($data);
    // 每个不同的消息都对应不同的函数来处理
    if ($data['msg_type'] == 0) {
        function_msg0($serv, $fd, $data);
    } else if ($data['msg_type'] == 1) {
        function_msg1($serv, $fd, $data);
    }
}
```

- 必须在 onWorkerStart 回调中创建 redis/mysql 连接

可以参考这篇 wiki: [是否可以共用1个redis或mysql连接](https://wiki.swoole.com/wiki/page/325.html)
下面初始化了 2 个 redis 连接, 一个用来做 `cache`, 一个用来做 `pub/sub`

```php
function onWorkerStart(swoole_server $serv, $id){
    // 只在 worker 进程中使用
    if ($id < $serv->setting['worker_num']) {
        // cache
        $cache = new \Redis();
        $cache->connect($config['cache']['host'], $config['cache']['port'], $config['cache']['timeout']);
        $cache->auth($config['cache']['auth']);
        $serv->cache = $cache;

        // pub
        $pub = new \Redis();
        $pub->connect($config['pub_sub']['host'], $config['pub_sub']['port'], $config['pub_sub']['timeout']);
        $pub->auth($config['pub_sub']['auth']);
        $serv->pub = $pub;
    }
}

// 之后就可以直接这样使用了
$serv->cache->set('key1', 'value1');
$serv->pub->publish('topic1', 'data1');
```

- mysql 连接池

虽说标题叫连接池, 这里其实只实例化了一个 mysql 连接对象, 原理其实是一样的

```php
function onTask($serv, $task_id, $from_id, $data) {
    static $link = null;
    if ($link == null) {
        $link = mysqli_connect($config['mysql']['host'], $config['mysql']['user'], $config['mysql']['password'], $config['mysql']['database']);
        if (!$link) {
            $link = null;
            return;
        }
    }
    // 这里做了一层封装, 需要在原有 sql 语句上加标记来判断是什么样的原句, 可以优化
    list($queryType, $sql) = explode('|', $data);
    $result = $link->query($sql);
    if ($result) {
        if ($queryType == 'select') {
            $result = $result->fetch_all(MYSQLI_ASSOC);
        } else if ($queryType == 'insert') {
            $result = mysqli_insert_id($link);
        }
        return $result;
    }
}

function onFinish($serv, $data)
{
    //
}

// 之后就可以直接这样使用了
$res = $serv->taskWait('select|select name from user where id=xxx');
```

## 使用订阅时踩到的坑

swoole 中订阅的实现, 可以先参考这篇 blog: [如何实现从 Redis 中订阅消息转发到 WebSocket 客户端](https://segmentfault.com/a/1190000010986855)

同样还是要在 onWorkerStart 回调函数中启动 `redis sub`

```php
function onWorkerStart(swoole_server $serv, $id){
//        if ($id == 0) { // 只启动一个 sub
            $sub = new swoole_redis(); // swoole_redis 支持异步
            $sub->on('message', function (swoole_redis $redis, $result) use ($serv, $config) {
                if ($result[0] == 'message') {
                    list($userId, $status) = explode(':', $result[2]);
                    // 解析出 fd, 用来给 client 发送消息
                    $userFd = $serv->userinfo[$userId]['fd'] ?? 0;
                    if ($userId && $userFd && in_array($status, [0, 1, 2])) {
                        foreach ($serv->connections as $fd) {
                            if ($userFd == $fd) { // 只发给对应客户端

                                // 省略业务逻辑

                                // 给 client 发送消息
                                $serv->send($fd, encode('foo'));
                                break;
                            }
                        }
                    }
                }
            });
            $sub->connect($config['pub_sub']['host'], $config['pub_sub']['port'], function (swoole_redis $redis, $result) use ($config) {
                $redis->auth($config['pub_sub']['auth'], function (swoole_redis $redis, $result) use ($config) {
                    $redis->subscribe('game_result_'. $config['zone_id']);
                });
            });
//        }
    }
}
```

仔细看代码, 会发现有这样一行注释 **「只启动一个 sub」**, 这是由业务决定的, 收到订阅消息的时候, 只需要转发给特定的用户.
但是, 在 onWorkerStart 回调函数中启动, 并不能实现. 原因我们需要先了解一下 swoole 的进程模型:

- server 启动时, 首先会开启一个 `master` 进程, master 进程会启动 `reactor` 线程, 用来管理 tcp 连接和 tcp 数据的收发
- `master` 进程启动 `manager` 进程和 `reactor` 线程
- `reactor` 线程, 用来管理 tcp 连接和 tcp 数据的收发
- `manager` 进程用来管理 `worker` 进程和 `task_worker` 进程, 根据上面的 `worker_num/task_worker_num` 配置
- `worker` 进程处理 `reactor` 线程转发过来的数据, 处理完业务逻辑后, 将数据发给 `reactor` 线程, 由 `reactor` 线程转发给用户
- `worker` 进程会将耗时任务投递给 `task_worker` 进程, `task_worker` 进程处理完后触发 `onFinish` 事件回调

所以, 我们可以根据使用 `$work_id < $serv->setting['worker_num']` 来判断我们当前是在 `worker` 进程还是 `task_worker` 进程

第一版在写的时候, 我是按照业务需求, 来限定只在 `$work_id = 0` 的进程上开启 redis sub. 但是, 问题马上就来了: 当前用户的 fd, 并不一定在 `$work_id = 0` 的进程上, 这就会导致下面这段代码失效:

```php
foreach ($serv->connections as $fd) { // $serv 其实对应的当前的 worker 进程
    if ($userFd == $fd) {
        // do something
    }
}
```

但是, 如果不加 `$work_id = 0` 的限制, 就会导致我们开了多少个 worker 进程, 就会有多少个 sub, 导致消息的重复订阅, 重复的业务逻辑处理.

这时候, 就必要了解一下, swoole 提供的 [Process](https://wiki.swoole.com/wiki/page/p-process.html) 进程管理模块, 我们只需要单独起一个进程, 用来维护 sub 就好了

## 使用 swoole 裸写 server 发现的问题

很明显, 上面的业务逻辑还不够复杂, 使用的服务也不多, 但是整个开发下来的不舒适感是非常明显:

- 开发环境和测试环境的搭建: 编译 swoole, 安装 redis/mysql
- 配置管理: 快速开发时写死到业务里, 到优化时抽到 `config.php` 配置文件中
- 服务部署: 开始尝试官方 wiki 里 systemd daemonize 的方案, 结果产生了大量僵尸进程
- 连接池: 如果需要更高的性能, 连接池会很有必要, 无论是 redis 还是 mysql
- 协议处理: 我们使用了 `固定包头 + protobuf` 的自定义协议, 将协议和业务分离开才是良好的设计
- 学习成本: 官方 wiki 第一次读产生大致映像, 第二次边读边实现 wiki 中的例子, 第三次根据业务需求去细读 wiki 相应章节. 但是 wiki 尽管接近 1400 页, 还是会有新问题.

相比而言, php 的 web 框架如此之多, MVC 大行其道, 是否也有 php 的 server 框架, 可以解决上面这些共性问题呢?

这里推荐一下 [swoole distribution](http://docs.youwoxing.net), 重构的时候选择了这个框架, 简单说一下优点:

- docker 配置开发环境, 不过 docker for window 通过目录挂载会导致无法热更新(当然也有方案解决)
- 不依赖其他服务(systemd supervisor)进行服务化部署
- Pack 模块解决协议解析
- 经典的 MVC 结构, 只需要稍微修改 route, 就可以在 controller 和 model 中书写业务逻辑
- 自带连接池, 修改配置文件即可
- 没错, 还有协程

```php
$value = yield $this->redis_pool->getCoroutine()->get('key1');
```

- 没错, 还有 Process

```php
namespace app\Process;

use Server\Components\Process\Process;

class MyProcess extends Process
{
    public function start($process)
    {
        parent::start($process);

        // 可以把 redis sub 的逻辑放这里了
    }

    // 可以在 controller 中使用 rpc 调用此方法
    public function getData()
    {
        return '123';
    }
}
```

## 写在最后

确实, 之前并没有写过服务器, 一直停留在「纸上谈兵」的阶段, 真正写起来的时候才发现「这活真累」.
不过, 那些年你读的书, 刷的技术 blog, 参加的技术大会, 总归是有用的, 拦在入门地方的, 并不是语言, 而是这个领域的「基础」, 这些你都可以通过这些方式获取, 缺的是需要自己将它系统化.
当然, 接着就是 coding, `practice makes perfect` 对一线程序员会一直有用.
