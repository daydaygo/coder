# coder| 「进程」编程

> 图灵社区 - 理解UNIX进程: http://www.ituring.com.cn/book/1081
> 「理解Unix进程」读书笔记: http://www.jianshu.com/p/9f6bf7d2a445

博客来源:

- <理解Unix进程> 这本书
- rango和鸟哥的博客, 关于服务器编程的部分
- 自己工作学习中的积累

摘录几个进程相关的要点, 这玩意真的很重要:

- 所有的代码都是在进程中执行的(知道有多重要了吧)
- 用户空间 - 内核 - 系统调用(理解这三者的关系)
- Unix哲学：万物皆为文件(普遍一点的理解是 文件 + 资源, 但是这样抽象后, 可以统一进行处理了)
- 资源: FD(文件描述符)来跟踪资源; 三个公共资源 stdin stdout stderr
- 进程皆有环境(知道环境变量干啥用了吧)
- pid - ppid - 僵尸进程 - `copy-on-write`(进程之间的关系)
- `master/worker` 进程模型
- IPC(进程间通信): 信号 管道 socket
- 与Unix进程打交道事关两件事：抽象和通信

当然, 强烈推荐去读一下这本书, 1-2 小时就能读完.

## fork

先来对比一下:

```php
// 司空见惯的 if-else
if ($a == -1) {
    echo '-1';
} else if($a == 0) {
    echo '0';
} else {
    echo $a;
}

// 还是 if-else
$pid = pcntl_fork(); // 基于 ext-pcntl 扩展
if ($pid == -1) {
    echo 'fork fail';
} else if ($pid) {
    echo 'parent';
} else {
    echo 'child';
}
```

上面的例子, 只会有一个 `echo` 输出(`if-else` 语句就该这样呀), 但是下面的 `echo 'parent';` 和 `echo 'child';` 却都会输出.

是不是很难理解? 我第一次遇到这样的情况, 是来源于下面的代码:

```perl
sub main{
    my $zip_dir = $conf->{config}->getone('conf,zip_dir');
    my @zip_dir = split ',', $zip_dir;
    my @files = ();

    my $project = 'brand_item';
    for my $dir(@zip_dir){
        my @tmp_files = <$dir/$project/*.zip>;
        for my $file(@tmp_files){
            #if(-M $file > 0.01){ push @files, $file; }
            push @files, $file;
        }
    }

    while(@files){
        while(@files and (keys %$child < $process_num)){
            my @task_files;
            if(@files > $do_number_per_process){
                @task_files = @files[0..$do_number_per_process-1];
                @files = @files[$do_number_per_process..$#files];
            }
            else{
                @task_files = @files;
                @files = ();
            }

            if(my $pid = fork()){ $child->{$pid} = 1;}
            else{
                $conf->{log}->sayshort("fork child $$, " . join ' ', keys %$child);
                my $conf = get_config();
                eval{ &process_content($conf, \@task_files);}; #处理文件
                if($@){ $conf->{log}->err("child failed $@");}
                $conf->{log}->sayshort("child $$ die");
                exit;
            }
        }
        my $pid = wait();
        delete $child->{$pid};
    }

    while((my $pid = wait()) != -1){}
}
```

上面是一段 perl 爬虫脚本, 使用多进程处理爬取到的页面内容, 达到加速脚本执行的效果.

这里有 2 点需要提一下:

- 任务分解: 这里只处理整个爬虫任务的一部分 - 页面内容处理; 任务分解细化是解决是解决 **大规模 / 复杂问题** 常用思路
- 如果单进程跑, 任务大概要 2 小时, 使用多进程后, 15 分钟就能结束

不知道看完之后是否可以理解:

- fork 之后, 生成了子进程, 此时的进程和子进程都会执行之后的代码, 其实是有 2 个进程在跑 `if-else` 代码块, 这也是为什么 `if-else` 都有输出
- 使用 fork 后, 通常是 **子进程** 执行具体的任务, **原进程** 待子进程处理完后使用 `wait` 回收子进程

## server

看完进程的基本实践后(`fork + wait`), 我们再来看看服务器领域的例子.

其实上面的例子, 就包含了一个基本的进程模型 `master/worker`:

- 主进程(master)负责进程(work进程)管理
- worker进程负责执行具体的任务

```php
$serv = stream_socket_server("tcp://0.0.0.0:8000", $errno, $errstr) or die('create server failed');

// 直接 fork()
while (1) {
    $conn = stream_socket_accept($serv); // 阻塞在 accept 上
    if (pcntl_fork() == 0) {
        $request = fread($conn);
        // do something
        fwrite($conn, 'hello world');
        fclose($conn);
        exit(0);
    }
}

// 多进程/多线程 leader-follower模型 fpm/apache
for ($i=0; $i < 32; $i++) {
    if (pcntl_fork() == 0) {
        while (1) {
            $conn = stream_socket_accept($serv); // 阻塞在 accept 上
            if ($conn === false) continue;
            $request = fread($conn);
            // do something
            fwrite($conn, 'hello world');
            fclose($conn);
        }
        exit(0);
    }
}
```

上面 2 个实现其实都是 `master/worker` 模型, 仔细对比一下, 就会发现细微的差别:

- 下面的代码实现实现了进程池
- 为什么要用进程池? 因为进程的新建和销毁也需要消耗系统资源

`master/worker` 模型非常常见, `nginx / php-fpm` 都是基于此模型. swoole 也包含此模型, 但是因为需要处理更复杂的业务, 会有不同功能的进程, 进程模型也会更复杂一点.

当然, 无论进程还是服务器开发, 包含的内容都不止这么一点点. 这里写得很简单和简略, 是想要传达一个简单的认知:

> 吃的草多了, 你也会成为大牛
