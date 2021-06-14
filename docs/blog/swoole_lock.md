# swoole| lock in swoole

时间不早了, 尽量言简意赅一些.

怎么快速学习锁:

- 锁是什么 -> 锁的概念
- 为什么要用锁 -> 锁的使用场景
- 怎么使用锁 -> **show me the code**

需要了解的相关概念:

- 多进程 / 多线程
- 临界区 / 竞态资源

要培养 **锁** 的意识:

- 并发场景下的数据安全, 比如 session 覆盖([PHP的Session数据同步问题](http://rango.swoole.com/archives/74))

## 与锁相关问题的几个维度

- 锁的种类 : 锁的规则 : 锁的场景 = 1 : 1 : 1

swoole 中支持的锁(lock), 相关代码在 [`swoole/src/lock`](https://github.com/swoole/swoole-src/tree/master/src/lock) 下, 代码很清晰, 适合阅读:

1. 互斥锁 SWOOLE_MUTEX
2. 自旋锁 SWOOLE_SPINLOCK
3. 读写锁 SWOOLE_RWLOCK
4. 文件锁 SWOOLE_FILELOCK
5. 信号量 SWOOLE_SEM

- 锁的生命周期: 创建锁 -> 加锁 -> 解锁

```php
$lock = new \Swoole\Lock(SWOOLE_MUTEX); // 参数是锁的种类, 下面细说
$lock->lock();
$lock->unlock();
```

**注意**: 这里并没有看到锁的销毁, **受限于 PHP 的多进程模型，只能这样**, 所以官网特地提到

> 请勿在onReceive等回调函数中创建锁，否则底层的GlobalMemory内存会持续增长，造成内存泄漏

- 是否阻塞: 申请加锁是否会阻塞

```php
$lock->lock(); // 阻塞, 直到成功申请到锁

$lock->trylock(); // 不会阻塞, 立即返回抢锁的结果
```

**抢锁 = 申请锁**, 不要被中文误导, 这里是同一个意思

因为锁的规则不同, **信号量 SWOOLE_SEM** 不支持 `trylock()` 这样非阻塞调用方式

- 锁的实现方式

```php
$lock = new \Swoole\Lock(SWOOLE_FILELOCK, string $lockfile); // 文件锁

$lock = new \Swoole\Lock(SWOOLE_MUTEX); // 其他锁
```

**文件锁 SWOOLE_FILELOCK** 使用文件来作为锁, **其他锁** 使用内存, 所以 [**其他类型的锁必须在父进程内创建**](https://wiki.swoole.com/wiki/page/128.html).

## 互斥锁 SWOOLE_MUTEX

这个概念理解起来最容易: **互斥锁 = 独占锁 = 排它锁 = 写锁**.

- 锁的规则: 顾名思义, 就是这个锁当前只能有一个 进程/线程 使用, 其他 进程/线程 只能等待
- 实现: swoole 底层使用 `pthread_mutex_xxx` 系列函数实现.
- 场景: 任何需要 **排他/独占** 的场景都适用

## 自旋锁 SWOOLE_SPINLOCK

自旋锁和互斥锁类似, 只是在 **阻塞** 时的状态不一样:

- 互斥锁没有申请到锁时, 通常会进入睡眠状态, 这样当锁释放时, 就需要 **唤醒**; 自旋锁则会一直循环在那里等待锁
- 实现: swoole 底层使用 `pthread_spin_xxx` 系列函数实现.
- 场景: 锁保持时间非常短的情况, 减少 **睡眠 -> 唤醒** 的成本

## 读写锁 SWOOLE_RWLOCK

```php
$lock->lock(); // 写锁

$lock->lock_read(); // 读锁
```

写锁上面已经见过了, 读写锁由 **读锁 + 写锁** 组成

- 锁的规则: 如果当前有 **读锁**, 可以继续申请 **读锁**, 但是不能申请 **写锁**; **写锁** 的规则和上面一致
- 实现: swoole 底层使用 `pthread_rwlock_xxx` 系列函数实现.
- 场景: 读多写少的场景, 减少 **锁** 的阻塞带来的性能损耗

## 文件锁 SWOOLE_FILELOCK

除了使用文件作为锁, 规则和场景和 **读写锁 SWOOLE_RWLOCK** 一致

- 实现: swoole 底层使用 `fcntl()` 函数实现.

## 信号量 SWOOLE_SEM

信号量的规则稍微会复杂一些, 会涉及到 2 个操作: `P V`

- 锁的规则: 如果信号量 >= 0, 则可以执行 P 操作, 并将信号量减一; 如果信号量 < 最大值, 则可以执行 V 操作, 并将信号量加一
- 实现: swoole 底层调用 `<sys/sem.h>` 实现.
- 场景: 信号量操作可以想象成 **仓库进货出货** 的过程, 仓库是空的, 就不能用 P 操作来出货了, 仓库是满的, 就不能用 V 操作来继续进货了

注意: 信号量没有 `trylock()` 方法; swoole 没有提供 **容量** 的设置, 退化为容量为 1 的情况了, 其实和 **互斥锁 SWOOLE_MUTEX** 差不多了

## 写在最后

到目前为止, swoole 当中使用到的锁都聊到了, 最后说一个**彩蛋**. 在 [2017北京PHP开发者年会](http://www.itdks.com/eventlist/detail/1832) 上, thinkpc 讲到 **exec的梗**:

使用 `\Swoole\Process::exec('tail -f')` 导致服务一直阻塞在那里, 后来和韩老大说这个事, 韩老大就给加上了一个 timeout.

```php
$lock->lockwait($timeout);
```

所以我猜测 `lockwait()` 这个只有 **Mutex** 类型才支持的方法, 是不是也是这种情况.

[swoole wiki - `\Swoole\Lock`](https://wiki.swoole.com/wiki/page/p-lock.html) 还是优先推荐的第一手资料.

理解一个不太熟悉的概念时, 可以尝试划一划 **不同维度**, 在这些 **维度** 上面找 **相同 / 不同**.

Yii 框架中特定提供了 **mutex** 模块, 并提供了不同种类的实现的方式, 可以移步看我之前的 [blog - yii源码解读](https://www.jianshu.com/p/fd85383783eb)

使用 file 实现的 mutex:

```php
// 加锁
protected function acquireLock($name, $timeout = 0)
{
    $file = fopen($this->getLockFilePath($name), 'w+');
    if ($file === false) {
        return false;
    }
    if ($this->fileMode !== null) {
        @chmod($this->getLockFilePath($name), $this->fileMode);
    }
    $waitTime = 0;
    while (!flock($file, LOCK_EX | LOCK_NB)) { // 加锁
        $waitTime++;
        if ($waitTime > $timeout) { // 超时逻辑
            fclose($file);

            return false;
        }
        sleep(1);
    }
    $this->_files[$name] = $file;

    return true;
}

// 解锁
protected function releaseLock($name)
{
    if (!isset($this->_files[$name]) || !flock($this->_files[$name], LOCK_UN)) { // 解锁
        return false;
    } else {
        fclose($this->_files[$name]);
        unlink($this->getLockFilePath($name));
        unset($this->_files[$name]);

        return true;
    }
}
```

核心其实是 `flock()` 函数:

- LOCK_EX, 英文对应 exclusive lock, 排它锁
- LOCK_UN, 英文对应 unlock, 解锁
- LOCK_SH, 英文对应 share lock, 共享锁
- LOCK_NB, 英文对应 non block, 非阻塞

使用 mysql 实现的 mutex:

```php
// 加锁
protected function acquireLock($name, $timeout = 0)
{
    return (bool) $this->db
        ->createCommand('SELECT GET_LOCK(:name, :timeout)', [':name' => $name, ':timeout' => $timeout])
        ->queryScalar();
}

// 解锁
protected function releaseLock($name)
{
    return (bool) $this->db
        ->createCommand('SELECT RELEASE_LOCK(:name)', [':name' => $name])
        ->queryScalar();
}
```

调用 mysql 提供的函数: `GET_LOCK() / RELEASE_LOCK()` 即可
