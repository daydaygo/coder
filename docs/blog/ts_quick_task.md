# TS| 技术分享: 脚本慢如何优化?

> [tech| 技术分享: 脚本慢如何优化?](https://www.jianshu.com/p/438aacff6090)

业务上有很多功能通过后台脚本运行, 有时会遇到 **脚本还没跑完, 又要加班了** 这种情况, 这篇聊聊脚本优化提速的话题.

优化的 2 个大方向:

- 物力
- 人力

先上干货, 聊聊物力, 怎么想办法发挥出计算机的性能, 物力优化常见的有 2 方面:

- 多进程
- 多协程

在继续下面的内容之前, 请确保你熟悉这些基础知识:

- unix系统架构图
- 系统调用和库函数
- 用户态和内核态
- 程序是如何执行的
- 进程 线程 协程

这篇文章非常不错, 推荐阅读

> 编程基础知识: https://mp.weixin.qq.com/s/nxdFeLGGQLgBcy5zq5rsvw

## 进程 线程 协程

- 什么是进程

进程就是运行着的程序.

```php
// test.php
<?php
sleep(100);
```

查看进程:

```
/var/www/coding/php # php test.php &
/var/www/coding/php # ps aux
PID   USER     TIME   COMMAND
  156 root       0:00 php test.php
  157 root       0:00 ps aux
```

- 什么是线程

线程是操作系统的最小执行单位, 真正干活的不是进程, 而是进程中的线程

- 什么是协程

非常形象的说法: **用户态线程**. 为什么协程比线程快, 下面还会讲到.

思考一个问题: 使用协程的时候, 到底是谁在干活?

## 多进程

### 基础实现

[fork 系统调用](https://www.cnblogs.com/bastard/p/2664896.html)

```c
#include <unistd.h>
#include <stdio.h>
int main ()
{
    pid_t fpid; //fpid表示fork函数返回的值
    int count=0;
    fpid=fork();
    if (fpid < 0)
        printf("error in fork!");
    else if (fpid == 0) {
        printf("i am the child process, my process id is %d/n",getpid());
        printf("我是爹的儿子/n");//对某些人来说中文看着更直白。
        count++;
    }
    else {
        printf("i am the parent process, my process id is %d/n",getpid());
        printf("我是孩子他爹/n");
        count++;
    }
    printf("统计结果是: %d/n",count);
    return 0;
}
```

运行结果:

```
i am the child process, my process id is 5574
我是爹的儿子
统计结果是: 1
i am the parent process, my process id is 5573
我是孩子他爹
统计结果是: 1
```

是不是很神奇, **if-else 代码块都执行了**

PHP 中多进程相关扩展: **pcntl/posix**

更多 PHP 中多进程编程的知识: [rango blog](http://rango.swoole.com)

### 更简单的方式

swoole 的协程池, 轻松开启多进程:

```php
// 进程数
$pool = new \Swoole\Process\Pool($workNum);
$pool->on('workerStart', function ($pool, $workerId) {
    // 业务逻辑
});
$pool->start();
```

具体实例:

```php
public function actionOpinionMq($workNum = 1)
{
    $pool = new Pool($workNum);
    $pool->on('workerStart', function ($pool, $workerId) {
        $callback = function (AMQPMessage $msg) {
            $msgBody = $msg->body;
            $this->info($msgBody);
            $msgBody = json_decode($msgBody, true);
            if (isset($msgBody['id'])) {
                $row = Yii::$app->getDb()->createCommand("SELECT id,content FROM opinion WHERE source_id='{$msgBody['id']}'")->queryOne();
                if ($row) {
                    // 业务代码
                }
            }
            /** @var AMQPChannel $ch */
            $ch = $msg->delivery_info['channel'];
            $ch->basic_ack($msg->delivery_info['delivery_tag']);
        };
        $connection = new AMQPStreamConnection('rabbitmq', 5672, 'guest', 'guest');
        $channel = $connection->channel();
        $channel->queue_declare('crawl-opinion', false, false, false, false);
        $channel->basic_qos(null, 1, null);
        $channel->basic_consume('crawl-opinion', '', false, false, false, false, $callback);

        while (count($channel->callbacks)) {
            $channel->wait();
        }

        $channel->close();
        $connection->close();
    });
    $pool->start();
}
```

## 多协程

### 协程为什么快

- 用户态 内核态
- cpu密集型 IO密集型

```php
$n = 4;
// 普通版
for ($i = 0; $i < $n; $i++) {
    sleep(1);
    echo microtime(true) . ": hello $i \n";
};
// 单协程
go(function () use ($n) {
    for ($i = 0; $i < $n; $i++) {
        Co::sleep(1);
        echo microtime(true) . ": hello $i \n";
    };
});
for ($i = 0; $i < $n; $i++) {
    go(function () use ($i) {
        Co::sleep(1);
        // sleep(1);
        echo microtime(true) . ": hello $i \n";
    });
};
```

推荐这篇文章对协程的解读: [swoole| swoole 协程初体验](https://www.jianshu.com/p/745b0b3ffae7)

swoole 现在支持原生 mysql/redis 无缝切换到协程

业务中的一次实践, 更新现有手机号的运营商信息:

```php
for ($i=0; $i< 500; $i++) {
    go(function () use ($i, $sms_job_id){
        // 取模进行任务分片
        $sql = "SELECT id,phone FROM sms_job_phone WHERE sms_job_id=$sms_job_id AND ops_type=0 and id%500={$i} LIMIT 500";
        $rows = static::getDb()->createCommand($sql)->queryAll();
        foreach ($rows as $row) {
            echo $row['id'], "\n";
            $m = static::findOne($row['id']);
            // 判断用户手机号运营商
            $m->ops_type = Helper::getMobileOperator($row['phone']);
            $m->save();
        }
    });
}
```

## 怎么把任务拆成多个

上面出现过的案例:

- sql 中 **取模** / 分段(`id>:id limit 1000`)
- 消息队列

![rabbitmq_queue](http://qiniu.dayday.tech/rabbitmq_queue.png)

## 物力

为何会优先划分 **人力/物力** 这 2 个范畴呢? 因为绝大多数场景, **并不需要去榨干计算机性能**. 反而是编程时没有采取一些 **最佳实践** 或者一些 **失误** 导致程序运行的效果 **比较糟糕**.

## 开发规范

**历史总是惊人的相似**, 积累一些开发规范, 可以少踩一些坑

- [aliyun redis 开发规范](https://yq.aliyun.com/articles/531067)
- [DBA 谈互联网 MySQL 开发规范](https://yq.aliyun.com/articles/87755)

## 案例分享: 变通思路

- 《聊斋志异》手稿本卷三《驱怪》篇末，有“异史氏曰：黄狸黑狸，得鼠者雄”！
![喵星人看鬼书](http://qiniu.dayday.tech/deng_cat.jpeg)

比如上面手机运营商的例子, 开了 500 协程还是很慢, 3w 数据超过 10 分钟才执行完, 还能更快一些么? 可以, 业务中直接读取的本地文件.

### out of memory

有个定时脚本从 mongo 中取数据进行处理, 循环到数据处理完, 上线后执行一段时间报错: `out of memory`

- `ini_set('memory_limit', '512m');` 调整脚本运行内存限制, 执行一段时间, 依旧报错
- `memory_get_usage() / memory_get_peak_usage()` 添加日志, 记录脚本内存使用, 定位问题, 发现每条需要处理的数据超过 20m
- `gc_collect_cycles() / unset()` 添加强制 gc, 主动回收变量, 有效果, 但是内存依旧会持续增长
- `0 22 * * *` -> `* 22 * * *` + `limit 50` 原脚本每天 22 点执行一次, 改成每天 22 点每分钟执行一次, 每次只处理 50 条数据

### connection gone away

有个统计脚本统计比较复杂, 需要统计三层数据: 查询出第一层数据, 统计后, 拿获取的数据再去查询, 查询后再统计, 比如 **一定维度的用户 + 这些用户相关的订单 + 这些订单相关的账单**, 测试时发现脚本运行一段时间后报错 `connection gone away`

关于连接超时:

- Navicat 中的 keepalive
- mysql 中 timeout 相关配置 `SHOW VARIABLES LIKE '%timeout%';`

代码中:

- 最简单的例子, `mysqli_ping()`

```php
// 进程长时间运行时的长连接
if ($this->_linkr && mysqli_ping($this->_linkr)) {
   $this->_link = $this->_linkr;
   return true;
}
$this->_linkr = $this->_connect($host);
```

- yii 框架中

vendor/yiisoft/yii2/db/Connection.php

```php
/**
    * Returns a value indicating whether the DB connection is established.
    * @return bool whether the DB connection is established
    */
public function getIsActive()
{
    return $this->pdo !== null;
}
```

common/helpers/GlobalHelper.php

```php
/**
    * connect db 重新连接数据库
    * @param string $db
    * @return \yii\db\Connection 返回数据库操作句柄
    */
public static function connectDb($db = 'db') {
    $db_handle = ($db instanceof \yii\db\Connection) ? $db : \yii::$app->$db;
    if (empty($db_handle)) {
        return false;
    }

    try {
        return $db_handle->createCommand("select 1")->queryScalar();
    }
    catch (\Exception $e){
        $db_handle->close();
        $db_handle->open();
    }
}
```

- 给报错的代码打上补丁(伪代码)

```php
// 获取指定用户
// 可能在这里断开连接, 打上补丁
GlobalHelper::connectDb();
$data1 = $db->execute($sql_user);
foreach ($data1 as $v1) {
    $data2 = $db->execute($sql_order);
    foreach ($data2 as $v2) {
        $data3 = $db->execute($sql_bill);
    }
}
```

问题依旧存在, 只是变成了: **补丁要打在哪里**

- 现实中的代码比这个更复杂, 给每个 db 连接的地方都打补丁?
- 直接改动框架底层, 风险怎么评估?

那问题怎么解决呢? 同样的套路: limit 数据量 + 多次执行.

## 写在最后

提升技能的过程中, 不妨多掌握一些 **套路**:

- 积累/制定 相应技术的开发规范: **历史总是惊人的相似**
- 有些知识知道了(理解了)就很简单, 没那么玄乎
- 《聊斋志异》手稿本卷三《驱怪》篇末，有“异史氏曰：黄狸黑狸，得鼠者雄”！

推荐阅读:

- [编程基础知识](https://mp.weixin.qq.com/s/nxdFeLGGQLgBcy5zq5rsvw)
- [多进程编程 - fork 系统调用](https://www.cnblogs.com/bastard/p/2664896.html)
- [服务器编程 - rango blog](http://rango.swoole.com)
- 推荐这篇文章对协程的解读: [swoole| swoole 协程初体验](https://www.jianshu.com/p/745b0b3ffae7)
