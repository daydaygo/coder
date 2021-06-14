# coder| 唯一ID引发的血案

写这篇 blog 原因是业务中遇到 生成唯一ID的场景 却没有按照需求生成唯一ID，由此引发了一番 *乱炖*，业务中目前使用的方案：

```php
// mysql 自增ID + 事务 + 时间 + 随机数
public function generateTradeNumber()
{
    $tradeTime = date('YmdHi', time());

    $lastTrade     = TradeNumber::findBySql('SELECT * FROM `Trade` ORDER BY id DESC LIMIT 1 FOR UPDATE');
    $lastTradeTime = '';
    if (!empty($lastTrade)) {
        $lastTradeNumber = $lastTrade->getTradeNumber();
        $lastTradeTime   = substr($lastTradeNumber, 0, 12);
        $lastTradeSerial = substr($lastTradeNumber, 12);
        if ($tradeTime == $lastTradeTime) {
            return $lastTradeTime . ($lastTradeSerial >= 99999 ? $lastTradeSerial + 1 : '0' . ($lastTradeSerial + 1));
        }
    }

    $initSerialNumber = rand(10000, 99999);
    return $tradeTime . '0' . $initSerialNumber;
}
```

简单解释一下：

- 唯一 TradeNumber 由 2 部分组成：当前时间 12 位 + 6 位数字
- 时间粒度到 **分**：`date('YmdHi', time())`
- 每次生成时，先锁死最后一条 Trade，当前分钟第一个进来的 Trade 会分配 `rand(10000, 99999)`（这样做是为了保持长度一致，以及**防止被人找出规律来**），之后的 Trade 在此基础上 +1

然而实际上并没有达到 唯一ID 的目的，查询数据库发现还是有重复 TradeNumber 存在

## php uniqid() 分析
> php manual uniqid()

详细的分析当然是直接看源码最清楚了，不过 php manual 已经说了：`based on the current time in microseconds`。并且 php manual 中已经详细说明了这个问题，详细可以查看 php manual note 中的内容，note 中提供了另外一种方案：

```php
function uniqidReal($lenght = 13) {
     // uniqid gives 13 chars, but you could adjust it to your needs.
     if (function_exists("random_bytes")) {
         $bytes = random_bytes(ceil($lenght / 2));
     } elseif (function_exists("openssl_random_pseudo_bytes")) {
         $bytes = openssl_random_pseudo_bytes(ceil($lenght / 2));
     } else {
         throw new Exception("no cryptographically secure random function available");
     }
     return substr(bin2hex($bytes), 0, $lenght);
 }
```

这种方案的原理：`随机数`。当然这个随机数概率就比 `rand(10000, 99999)` 小多了，毕竟加入了英文字母。那么，回到我们的主题，`随机数 === 唯一ID` ？显然，无论概率怎么小，还是存在重复的可能（-_-）。

## 唯一ID原理
> Linux多线程与同步：http://www.cnblogs.com/vamei/archive/2012/10/09/2715393.html
> 《高性能linux服务器编程》- 高性能 io - 进程相关部分
> 唯一ID生成原理与PHP实现：http://mp.weixin.qq.com/s/bagOgzdwLyZv_ITNVnYfoQ / https://github.com/liexusong/atom

其实非常的简单，只要保证**从同一个地方产生**，并按照**不重复的规律生成**。联想一下 mysql 的自增ID，就是 **同一张表，依次递增**。所以，自增ID真正的难点：**同一个地方生成**。但是，在多进程、多线程的场景下，这个 *地方* 就会成为 `竞态资源`。要了解race condition, mutex和condition variable的概念，请参考上面的博文 -- Linux多线程与同步。

多线程同步方式：

- 互斥锁(mutex)：一个线程申请了互斥锁，然后进行 **原子操作**，其他进程必须要等待此线程释放互斥锁，才能成功申请。
- 条件变量(condition variable)：配合互斥锁一起使用，在 **多个线程等待某个条件发生** 时的场景下使用
- 读写锁(reader-writer lock)：和互斥锁类似，锁有 3 种状态 -- R、W、unlock。如果资源上 R 锁，其他线程可以继续申请 R 锁；如果需要申请 W 锁，必须等待其他进程都释放 R 锁；如果资源上 W 锁，其他线程必须等待其释放

多进程的同步方式：

- 管道
- 共享内存

是不是有点峰回路转，怎么突然跑到操作系统层面了？whatever，你曾经学的操作系统原理，就是这么有用。

再来简单的剖析一下 *唯一ID生成原理与PHP实现* 这篇博文中提到的 snowflake算法：

- 唯一ID组成：64bit， 1bit（不用）+ 41bit（时间戳）+ 10bit（机器id）+ 12bit（序列号）
- 时间毫秒级（可以用2082年），每毫秒最多 `2^12=4096` 个请求，如果超过了，就分配到下一毫秒
- 为什么分机器ID：达到分布式计算，避免不同机器间同步带来的性能损耗

代码就不贴了，大家自己看，挺简单的。

## 唯一ID的php实现
> 韩天峰(Rango) - 从零开始编写第一个PHP扩展：http://wiki.swoole.com/wiki/page/238.html
> 淘宝信海龙 php 扩展开发：http://www.bo56.com/category/programming-language/php-programming-language
> 唯一ID生成原理与PHP实现：http://mp.weixin.qq.com/s/bagOgzdwLyZv_ITNVnYfoQ / https://github.com/liexusong/atom
> 韩天峰(Rango) Yii/Yaf/Swoole3个框架的压测性能对比：http://rango.swoole.com/archives/254

由于 php-fpm 是基于进程管理，每个pfm子进程都是相互独立的，想要实现 **同一个地方生成**，就需要依靠扩展。而php扩展开发，不少业内大牛都贴了教程，无耻引用几个。

推荐学习次序：

- 韩天峰(Rango) - 从零开始编写第一个PHP扩展
- 淘宝信海龙 php 扩展开发系列教程
- 查看 `https://github.com/liexusong/atom` 源码：github 上的文档只提到了 2 个函数，还是直接读源码吧，主要是 `atom.c`

测试：

- cli 下的测试

```
// php test
for ($i=0; $i < 20; $i++) {
    $id = atom_next_id(); // different
    $info = atom_explain($id); // the same
    echo $id . "\t" . date('YmdHis', $info['timestamp']) . "\t" . $info['datacenter'] . "\t" . $info['worker'] . "\n";
}

// ext atom source code
retval = ((current - context->twepoch) << context->timestamp_left_shift)
    | (context->datacenter_id << context->datacenter_id_shift)
    | (context->worker_id << context->worker_id_shift)
    | context->sequence;
```

## 写在最后
> 《高性能mysql》
> MOOC - 《linux操作系统原理与应用》
> swoole官方教程和 rango 的博客
> 《Just for Fun》linus 自传

由于本身的技术局限，对 mysql 中 这套 `SELECT FOR UPDATE` 不太熟悉，这个点有待提高，推荐《高性能mysql》。

操作系统原理到底有没有用？如果看到 进程、线程、共享内存、管道 等等名词感觉不太熟悉（熟悉名词和知道怎么用还隔着好远呢），最好还是多了解一点。

- 推荐《linux操作系统原理与应用》：这个是大学教材，找个MOOC网看看在线视频，网易公开课、中国大学MOOC、学堂在线 等等
- 《高性能linux服务器编程》：还有鼎鼎大名的《Unix Network Programing》，就是实在是太厚了
- swoole官方教程和 rango 的博客：c10k问题、时间循环、swoole/nginx/go/nodejs 实现原理对比、线程/进程/协程 对比 等等

最后的最后，写在这个没有出去浪的元旦假期：

- 且将新火试新茶，诗酒趁年华
- Just for Fun
