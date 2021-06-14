# coder| 服务器开发系列 1

> [服务器开发系列 1](https://www.jianshu.com/p/1633fa196c43)

算是第一次在实际项目中写 tcp server, 确实有些吃力, 不过投入产出还不错, Mark 一下, 和大家一起学习.

用 php 来写服务器, swoole 当然是首选. 当然, swoole 是将网络底层都打包好了, 应用层的服务治理发现 / 分布式 / 框架 等等, 还是需要自己基于 swoole 来写了. 不过 swoole 现在的生态链很好, 开源项目也多. 至于之前一直被诟病的文档, rango 在 php 开发者大会 上说过今年 swoole 的开发工作, 其中一部分就是文档. 所以, 期待 swoole 越来越好.

## 入门例子

```php
$serv = new swoole_server("127.0.0.1", 9501); // 绑定的本地 ip, 所以只能本地访问
$serv->set(array(
    'worker_num' => 8,
    'daemonize' => true, // 后台服务, 测试时设置成 false, 方便查看打印的信息
));
$serv->on('connect', function ($serv, $fd){
    echo "Client:Connect.\n";
});
$serv->on('receive', function ($serv, $fd, $from_id, $data) {
    $serv->send($fd, 'Swoole: '.$data);
    $serv->close($fd);
});
$serv->on('close', function ($serv, $fd) {
    echo "Client: Close.\n";
});
$serv->start();
```

上面的代码就是一个 **异步** 的 tcp server, 这里简单解释一下 **同步 / 异步**:

- 同步代码, 会同步阻塞, 使用进程模型, 一个进程只能处理一个请求, 依赖进程多少来来处理并发, 但进程越多, 进程间切换开销会越来越大
- 异步代码, 使用事件驱动, 所有地方都要改为异步代码, 当某个请求进入等待后, 进程并不会等在这里, 而是切换去处理就绪请求, 所以只需要设置为 1-4 倍的 cpu 核数即可

补充一个知识点: php 中回调的 4 种写法 -- 闭包(也叫 匿名函数) / 函数 / 类方法 / 类静态方法

## 协议

为什么要使用协议(protocol)? 因为 tcp 协议是流式的, 可能多次信息合并到一个包, 也可能一个信息分多个包传输, 所以应用层就需要自定义协议, 进行「分包」「合包」, 来确定数据的边界(即确定一次消息). 常用的自定义协议有 2 种: EOF 协议 和 固定包头协议.

- EOF 协议: 固定结尾符

读取消息, 直到遇到自定义结尾符. 优点是简单, 但是需要保证发送的消息中不包含「结尾符」, 否则就被拆开了.

swoole 中要设置 EOF 协议非常轻松:

```
$serv->set([
    'open_eof_split' => true,   // 开启EOF检测
    'package_eof' => '/r/n' ,   // 设置EOF标记
]);
```

实际中可能并不常使用, 但是在**测试**时非常有用, 方便直接用 `telnet` 连接服务器进行调试.

- 固定包头协议: 先读取固定包头, 获取包体大小信息, 然后读取这样大小的数据, 就是包体了

swoole 中设置固定包头协议, 也非常轻松:

```php
$serv->set([
   'open_length_check'     => 1,       // 开启协议解析
   'package_length_type'   => 'N',     // 长度字段的类型
   'package_length_offset' => 0,       // 第N个字节是包长度的值
   'package_body_offset'   => 4,       // 第N个字节开始计算长度
   'package_max_length'    => 2000000, // 协议最大长度
]);
```

`package_length_type` 这里使用的 `N`, 代表 4 字节 uint 型网络序. 对应的类型, 可以参考 [php manual - pack()](http://php.net/manual/en/function.pack.php).

这里解释一下网络序:

- 最小的计算机单位是 bit, 即 **位**, 只能表示 0 和 1, 实际使用时, 最小的单位是 byte, 即 **字节**(8 bit)
- 对于多字节数据, 按照字节进行划分, 就出现了一个排序的问题, 即高位放在前面还是后面的问题, 于是就产生了大端序和小端序
- 究竟使用的大端序还是小端序, 不同的机器是不一样的, 这个就叫做机器序
- 不统一当然不行了, 网络传输的数据就不一致了, 所以出现了网络序, 统一使用 大端序 来传输数据

固定包头还有一个玩法, 前 4 字节用来表示消息编号, 然后再用 4 字节来表示数据包大小, 而这只需要修改一下 `package_length_offset` 和 `package_max_length` 参数即可

比较简单的做法 `固定包头 + json 包体`

当然, 还可以指定更加复杂的协议, 不过本质上都抛不开这 2 种方式, 比如 http 协议中会有 `content-lentth`, 还有我之前所在的游戏公司, 使用 `固定整形签名 + 校验和 + 包体大小 + 包` 作为协议.

## protobuf

有 Google 当爹, 这个我就不过多介绍了. 我这里说明一下 php 如何快速入门 protobuf.

- protobuf 现在已经支持 php 了(我大 php 在服务器领域还是后劲十足的)
- protobuf 由 2 部分组成, protoc(protobuf compile), 用来将 proto 文件, 编译输出为不同语言可以使用的文件; protobuf runtime, 用来执行这些文件

php 使用 protobuf 需要做的准备:

- 到官网下载 protoc 的可执行文件, 安装到系统中
- 下载 protobuf 扩展, 即 protobuf runtime for php

```
# 解压后, 进入 protoc 的可执行文件的文件夹
cp bin/* /usr/local/bin/ # 复制 protoc 文件
cp -r include/* /usr/local/include/

# 使用 pecl 安装 protobuf 扩展
pecl install protobuf
pecl install protobuf-xxx # 指定不同斑斑
pecl install localfile # 使用本地文件安装
```

好了, 接下来定义我们的 protobuf 消息. protobuf 的消息类型(数据结构)很少, 扫一下 [官方文档](https://developers.google.com/protocol-buffers/docs/reference/php-generated) 即可:

```
# game.proto 文件
syntax = "proto3";

package game.protobuf;

message Auth {
    uint32 msgType = 1;
    int64 uid = 2;
    string token = 3;
    uint32 roomId = 4; // 认证并尝试上机
};

# 生成 php 可以使用的文件
protoc --php_out=build/ game.proto

# 复制 build/ 文件夹中的内容, 加入到我们的项目中, 修改 composer.json, 实现自动加载
{
 "autoload": {
    "psr-4": {
      "Game\\Protobuf\\": "protobuf/Game/Protobuf",
      "GPBMetadata\\": "protobuf/GPBMetadata"
    }
  }
}
composer dumpautoload
```

好了, 来一发:

```php
$auth = new Auth();
$auth->setMsgType(1);
$auth->setUid(1);
$auth->setToken('daydaygo');
$auth->serializeToString();

$auth = new  \Game\Protobuf\Auth();
$auth->mergeFromString($data);
echo $auth->getToken();
```

> PS: 折腾了 protobuf 很久, 一个重要的原因的就是没有好好的阅读官方文档, 采取直接百度「php protobuf」这样的方式直接找「实战」, 但是, 理解相关的概念更重要
> PPS: 把文档里面下载 protoc 文件的文件名看错了, 然后一直跑不起来, 浪费了好长时间

## tcp auth

需要做一个简单的认证: tcp 连接后, 客户端必须先发送 `Auth` 消息来认证, 认证不过就会断开连接. 这个需求的难点在于: **怎么确定用户是第一次给你发消息?**

和负责另一个 tcp server 的同事(java)讨论, 他那边将每个 **连接** 都抽象成了对象, 有私有变量 `_auth` 来表示是否认证. 虽然是放在对象里, 其实本质还是使用内存保存了当前连接的状态的. 既然如此, 我也可以直接申请内存来保存这个状态.

于是, `swoole_table` 参上:

```
$swooleTable = new swoole_table('100000'); // 最多 10w 同时连接
$swooleTable->column('auth', swoole_table::TYPE_INT, '1'); // 判断是否 auth
$swooleTable->create();
$swooleTable->set(1, ['auth' => 1]);
var_dump($swooleTable->get(2, 'auth'));

$serv->on('receive', function (swoole_server $serv, $fd, $from_id, $data) use ($swooleTable) {
    // 协议解析
   $data = decode($data);
   $auth = new  \Game\Protobuf\Auth();
   $response = 'your token: '. $auth->getToken();
   $auth->mergeFromString($data);
   $serv->send($fd, encode($response));

    // auth 后
    if ($swooleTable->get($fd, 'auth')) {
        echo $data, "\n";
        return;
    }
    echo "$data\n";
    $arr = explode(':', $data);
    var_dump($arr);
    if ($arr[0] !== 'token') {
        echo "need auth $fd\n";
        $serv->close($fd);
    } else {
        echo "auth ok $fd $data";
        $swooleTable->set($fd, ['auth' => 1]);
    }
});
```

需要注意:

- `swoole_table` 中, row 用来代表数据个数, column 用来表示数据可以具有的属性;
- `swoole_table` 运行后不能再次动态分配

> PS: 欢迎大大们提供更好的方案, 总感觉这个一杯茶熬出来的方案不是最常用的

## mq

有 2 个 tcp server 需要进行数据交互, 考虑到解耦, 于是引入了 mq(message queue , 消息队列). 实际场景中, mq 中的数据量不大, 所以直接使用 redis 的 `pub/sub`, 当然, 最终还是要做压测的, 以压测数据为准. 目前我们看靠了这篇 blog [redis的pub/sub性能测试](http://blog.csdn.net/dreamvyps/article/details/71123974).

```
// pub.php
$redis = new \Redis();
$res = $redis->connect('127.0.0.1', 6379);
$res = $redis->publish('test','hello,world'); // channel + msg

// sub.php
$redis = new \Redis();
$res = $redis->pconnect('127.0.0.1', 6379,0);
$redis->subscribe(['test'], 'callback');

// 回调函数,这里写处理逻辑
function callback($instance, $channelName, $message) {
 echo $channelName, "==>", $message,PHP_EOL;
}
```

好了, 这周主要把这次 tcp server 需要使用的技术梳理了一遍, 下周就要开始写业务代码了, 期待最后的压测的结果.
