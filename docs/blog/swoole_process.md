# Swoole| process in swoole

> [Swoole| Swoole 中 Process](https://www.jianshu.com/p/4b6326cdaaa7)

这篇 blog 折腾了很久才写出来, 问题主要还是在 **理解** 上. 有时候就是这样,

> 理解了之后就很简单, 不理解就很难; 知道了就很简单, 不知道往往就很难. 所以 **stay hungry stay foolish stay young** 真的很重要

本来计划开发 [swoft 框架](https://swoft.org/) 中的 Process 模块, 所以需要对 swoole 的 Process 模块要有比较深入的了解才行. 不过根据 [swoole 官方 wiki](https://wiki.swoole.com/wiki/page/p-process.html) 的实践过程中, 一直有未理解的部分. 之前虽然也做过多次 **多进程编程**, 但是当真正需要进行框架开发的时候, 就会发现以前学到的知识不够全面, 无法指导整体的设计. 好在一直在坚持, 奉上现在理解的程度.

内容一览:

- 进程相关基础操作: fork/exit/kill/wait
- 进程相关高级操作: 主进程退出子进程干完活后也退出; 子进程异常退出主进程自动重启
- 进程间通信(IPC) - 管道(pipe)
- 进程间通信(IPC) - 消息队列(message queue)
- swoole process 模块提供的更多功能

## 进程相关基础操作
> 进程是什么: 进程是运行者的程序

先来看看一个最简单的例子:

```php
<?php
echo posix_getpid(); // 获取当前进程的 pid
swoole_set_process_name('swoole process master'); // 修改所在进程的进程名
sleep(100); // 模拟一个持续运行 100s 的程序, 这样就可以在进程中查看到它, 而不是运行完了就结束
```

通过 `ps aux` 查看进程:

![未设置进程名](http://upload-images.jianshu.io/upload_images/567399-b09ad8a48f1f2e17.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![设置进程名](http://upload-images.jianshu.io/upload_images/567399-3a9eaffc6a22613e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

再来看看 swoole 中使用子进程的基础操作:

```php
use Swoole\Process;

$process = new Process(function (Process $worker) {
    if (Process::kill($worker->pid, 0)) { // kill操作常用来杀死进程, 传入 0 可以用来检测进程是否存在
        $worker->exit(); // 退出子进程
    }
});
$process->start(); // 启动子进程
Process::wait(); // 回收退出的子进程
```

- `new Process()`: 通过回调函数来设置子进程将要执行的逻辑
- `$process->start()`: 调用 `fork()` 系统调用, 来生成子进程
- `Process::kill()`: kill操作给进程发送信号, 常用来杀死进程, 传入 0 可以用来检测进程是否存在
- `Process::wait()`: 调用 `wait()` 系统调用, 回收子进程, 如果不回收, 子进程会编程 **僵尸进程**, 浪费系统资源

- `$worker->exit()`: 子进程主动退出

我在这里有一个疑问:

> 主进程的生命周期是怎么样的? 子进程的生命周期是怎么样的?

有这样一个疑问也来自于我之前的思维惯性: **理解一个事物时从事物的生命周期进行理解**. 结合 **进程是运行着的程序** 来一起理解:

- `new Process()`: 只有回调函数的逻辑会在进程中执行
- 除此之外的代码都是在主进程中执行

## 进程相关高级操作

- 主进程退出子进程干完活后也退出
- 子进程异常退出主进程自动重启

```php
<?php

use Swoole\Process;

class MyProcess1
{
    public $mpid = 0; // master pid, 即当前程序的进程ID
    public $works = []; // 记录子进程的 pid
    public $maxProcessNum = 1;
    public $newIndex = 0;

    public function __construct()
    {
        try {
            swoole_set_process_name(__CLASS__. ' : master');
            $this->mpid = posix_getpid();
            $this->run();
            $this->processWait();
        } catch (\Exception $e) {
            die('Error: '. $e->getMessage());
        }
    }

    public function run()
    {
        for ($i=0; $i<$this->maxProcessNum; $i++) {
            $this->createProcess();
        }
    }

    public function createProcess($index = null)
    {
        if (is_null($index)) {
            $index = $this->newIndex;
            $this->newIndex++;
        }
        $process = new Process(function (Process $worker) use($index) { // 子进程创建后需要执行的函数
            swoole_set_process_name(__CLASS__. ": worker $index");
            for ($j=0; $j<3; $j++) { // 模拟子进程执行耗时任务
                $this->checkMpid($worker);
                echo "msg: {$j}\n";
                sleep(1);
            }
        }, false, false); // 不重定向输入输出; 不使用管道
        $pid = $process->start();
        $this->works[$index] = $pid;
        return $pid;
    }

    // 主进程异常退出, 子进程工作完后退出
    public function checkMpid(Process $worker) // demo中使用的引用, 引用表示传的参数可以被改变, 由于传入 $worker 是 \Swoole\Process 对象, 所以不用使用 &
    {
        if (!Process::kill($this->mpid, 0)) { // 0 可以用来检测进程是否存在
            $worker->exit();
            $msg = "master process exited, worker {$worker->pid} also quit\n"; // 需要写入到日志中
            file_put_contents('process.log', $msg, FILE_APPEND); // todo: 这句话没有执行
        }
    }

    // 重启子进程
    public function rebootProcess($pid)
    {
        $index = array_search($pid, $this->works);
        if ($index !== false) {
            $newPid = $this->createProcess($index);
            echo "rebootProcess: {$index}={$pid}->{$newPid} Done\n";
            return;
        }
        throw new \Exception("rebootProcess error: no pid {$pid}");
    }

    // 自动重启子进程
    public function processWait()
    {
        while (1) {
            if (count($this->works)) {
                $ret = Process::wait(); // 子进程退出
                if ($ret) {
                    $this->rebootProcess($ret['pid']);
                }
            } else {
                break;
            }
        }
    }
}

new MyProcess1();
```

说明以下几点:

- 子进程运行结束后就会退出, 通过 `Process::wait()` 检测到子进程退出信号执行自动重启, 子进程就会一直执行下去
- 关于函数参数传 **引用/指针**, 一个很好的理解方式是: **参数可以被修改**

运行并模拟主进程异常退出:
![模拟主进程异常退出](http://upload-images.jianshu.io/upload_images/567399-433c30ea6e689c36.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![输出](http://upload-images.jianshu.io/upload_images/567399-66e09155ff9b20f0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 进程间通信(IPC) - 管道(pipe)

管道的几个关键词:

- 半双工: 数据单向流动, 一端只读, 一端只写.
- 同步 vs 异步: 默认为同步阻塞模式, 可以使用 `swoole_event_add()` 添加管道到 swoole 的 event loop 中, 实现异步IO
- 管道类型(数据格式): `SOCK_STREAM`, 流式, 需要用户自己处理数据的封包/解包; `SOCK_DGRAM`, 数据报, 每次收发都是一次完整的数据包 (`DGRAM/STREAM`)

注意, [swoole wiki - process->write()](https://wiki.swoole.com/wiki/page/216.html) 中提到 `SOCK_DGRAM` 并不会乱序丢包

先来看一个简单的例子, php从shell管道中读取数据:

```php
// get pip data
$fp = fopen('php://stdin', 'r');
if ($fp) {
    while ($line = fgets($fp, 4096)) {
        echo "php get pip data: ". $line;
    }
    fclose($fp);
}
```

![从shell管道读取数据](http://upload-images.jianshu.io/upload_images/567399-6318862a6b8da01e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


swoole process中的管道很强大, 支持 **子进程写, 主进程读** 以及 **主进程写, 子进程读**:

```php
use Swoole\Process;

// 子进程写, 父进程读
$process = new Process(function (Process $worker) {
    $worker->write("worker");
});
$process->start();
$msg = $process->read();
echo "from process: $msg", "\n";

// 父进程写, 子进程读
$process = new Process(function (Process $worker) {
    $msg = $worker->read();
    echo "from master: $msg", "\n";
});
$process->start();
$process->write('master');
```

![使用管道多次读写](http://upload-images.jianshu.io/upload_images/567399-9951c81418fdb977.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


注意区分 `$worker->write()` 和 `$process->write()`, 之前一直错误的以为这 2 个是相同的, 其实就是把 `$process` 误以为是子进程, 从而相当于 `$process->write()` 就是子进程写管道 -- 其实这里是主进程内执行的逻辑, 是主进程写数据到管道, 供子进程读取

swoole中其他管道相关操作:

- 异步IO

```php
use Swoole\Process;
use Swoole\Event;

// 异步IO
$process = new Process(function (Process $worker) {
    $GLOBALS['worker'] = $worker;
    Event::add($worker->pipe, function (int $pipe) { // 使用 swoole_event_add 添加管道到异步IO
        /** @var Process $worker */
        $worker = $GLOBALS['worker'];
        $msg = $worker->read();
        echo "from master: $msg \n";
        $worker->write("hello master");
        sleep(2);
        $worker->exit(0);
    });
});
$process->start();
$process->write("master msg 1");
$msg = $process->read();
echo "from process: $msg \n";
```

![异步IO](http://upload-images.jianshu.io/upload_images/567399-7680c9622252a29c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


- 设置超时

```php
use Swoole\Process;

// 设置管道超时
$process = new Process(function (Process $worker) {
    sleep(5);
});
$process->start();
$process->setTimeout(0.5);
$ret = $process->read();
var_dump($ret);
var_dump(swoole_errno());
```

![管道超时](http://upload-images.jianshu.io/upload_images/567399-e01316fa96063542.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


插播一个趣事, **@thinkpc 看完 [2017北京PHP开发者年会](http://www.itdks.com/eventlist/detail/1832), 就知道为啥会点赞了**

- 关闭管道

```php
// 关闭管道: 默认值0->关闭读写 1->关闭写 2->关闭读
$process->close();
```


## 进程间通信(IPC) - 消息队列(message queue)

消息队列:

- 一系列保存在内核中的消息链表
- 有一个 **msgKey**, 可以通过此访问不同的消息队列
- 有数据大小限制, 默认 8192, 可以通过内核修改
- 阻塞 vs 非阻塞: 阻塞模式下 `pop()`空消息队列/`push()`满消息队列会阻塞, 非阻塞模式可以直接返回

swoole 中使用消息队列:

- 通信模式: 默认为争抢模式, 无法将消息投递给指定子进程
- 新建消息队列后, 主进程就可以使用
- 消息队列不可和管道一起使用, 也无法使用 swoole event loop
- 主进程中要调用 `wait()`, 否则子进程中调用 `pop()/push()` 会报错

```php
use Swoole\Process;
$process = new Process(function (Process $worker) {
    // $worker->push('worker');
    echo "from master: ". $worker->pop(). "\n";
    sleep(2);
    // $worker->exit();
}, false, false); // 关闭管道
// 参数一为 msgKey, 这里是默认值
// 参数二为 通信模式, 默认值 2 表示争抢模式, 这里还加上了 非阻塞
$process->useQueue(ftok(__FILE__, 1), 2| Process::IPC_NOWAIT);
$process->push('hello1'); // 使用 useQueue 后, 主进程就可以读写消息队列了
$process->push('hello2');
echo "from woker: ". $process->pop(). "\n";
// echo "from woker: ". $process->pop(). "\n";
$process->start(); // 启动子进程
// 消息队列状态
var_dump($process->statQueue());
// 删除队列, 如果不调用则不会在程序结束时清楚数据, 下次使用相同 msgKey 时还可以访问数据
$process->freeQueue();
var_dump(Process::wait()); // 要调用 wait(), 否则子进程中 push()/pop() 会报错
```

todo: img

## swoole process 模块提供的更多功能

- `swoole_set_process_name()`: 修改进程名, 不兼容 mac
- `swoole_process->exec(string $execfile, array $args)` 执行外部程序

参数 `$execfile` 需要使用可执行文件的绝对路径, 参数 `args` 为参数数组

```php
// 比如 python test.py 123
swoole_process->exec('/usr/bin/python', ['test.py', 123]);

// 更复杂的例子
swoole_process->exec(('/usr/local/bin/php', ['/var/www/project/yii-best-practice/cli/yii', 't/index', '-m=123', 'abc', 'xyz']);

// 父进程 exec 进程进行管道通信
use Swoole\Process;
$process = new Process(function (Process $worker) {
    $worker->exec('/bin/echo', ['hello']);
    $worker->write('hello');
}, true); // 需要启用标准输入输出重定向
$process->start();
echo "from exec: ". $process->read(). "\n";
```

![父进程与exec进程通过管道通信](http://upload-images.jianshu.io/upload_images/567399-f02115d7eaf65816.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- `\Swoole\Process::kill($pid, $signo = SIGTERM)`: 向指定进程发送信号, 默认是终止进程, 传 0 可检测进程是否存在
- `\Swoole\Process::wait()`: 回收子进程, 如果主进程不调用此方法, 子进程会变成 **僵尸进程**, 浪费系统资源
- `\Swoole\Process::signal()`: 异步信号监听

```php
use Swoole\Process;

// 异步信号监听 + wait
Process::signal(SIGCHLD, function ($signal) { // 监听子进程退出信号
    // 可能同时有多个子进程退出, 所以要while循环
    while ($ret = Process::wait(false)) { // false 表示不阻塞
        var_dump($ret);
    }
});
```

`\Swoole\Process::daemon()`: 将当前进程变为一个守护进程

```php
use Swoole\Process;

// daemon
Process::daemon();
swoole_set_process_name('test daemon process');
sleep(100);
```

- `\Swoole\Process::alarm()`: 高精度定时器(微秒级), 对 `setitimer` 系统调用的封装, 可以配合 `\Swoole\Process::signal()` / `pcntl_signal` 使用

注意不可和 `\Swoole\Timer` 同时使用

```php
// signal + alarm
// 第一个参数表示时间, 单位 us, -1 表示清除定时器
// 第二个参数表示类型 0->真实时间->SIGALAM 1->cpu时间->SIGVTALAM 2->用户态+内核态时间->SIGPROF
Process::alarm(100*1000); // 100ms
Process::signal(SIGALRM, function ($signal) {
    static $i = 0;
    echo "#$i \t alarm \n";
    $i++;
    if ($i>20) {
        Process::alarm(-1); // -1 表示清除
    }
});
```

- `\Swoole\Process::setaffinity()`: 设置CPU亲和, 即将进程绑定到指定CPU核上

传值范围: `[0, swoole_cpu_num())`
CPU亲和: CPU的速度远远高于IO的速度, 所以CPU有多级缓存来解决IO等待的问题, 绑定指定CPU, 更容易命中CPU缓存

## 写在最后

资源推荐:

- [图灵社区 - 理解UNIX进程](http://www.ituring.com.cn/book/1081) + [「理解Unix进程」读书笔记](http://www.jianshu.com/p/9f6bf7d2a445)
- [blog - 「进程」编程](https://www.jianshu.com/p/97c77d257945)

todo:

- 使用输入输出重定向
- 管道类型为 `SOCK_STREAM` 时的情况, 是否需要 **封包/解包** 处理, 即 [swoole wiki - process->write()](https://wiki.swoole.com/wiki/page/216.html) 中提到的 **管道通信默认的方式是流式，write写入的数据在read可能会被底层合并**
- 多进程 + 异步IO 的注意事项

能理解 **因为子进程会继承父进程的内存和IO句柄** 这个会产生的影响, 但是给的示例并没有说明这个问题

```php
use Swoole\Process;
use Swoole\Event;

// 多个子进程 + 异步IO
$workers = [];
$workerNum = 3;
for ($i=0; $i<$workerNum; $i++) {
    $process = new Process(function (Process $worker) {
        $worker->write($worker->pid);
        echo "worker: {$worker->pid} \n";
    });
    $pid = $process->start();
    $workers[$pid] = $process;
    // Event::add($process->pipe, function (int $pipe) use ($process) {
    //  $data = $process->read();
    //  echo "recv: $data \n";
    // });
}
foreach ($workers as $worker) {
    Event::add($worker->pipe, function (int $pipe) use ($worker) {
        $data = $worker->read();
        echo "recv: $data \n";
    });
}
```

![多进程异步IO](http://upload-images.jianshu.io/upload_images/567399-721d74f6e812ea7b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

