# linux

- 内核锁.同步机制=spinlock+信号量.睡眠锁+大内核锁+rwlock+rcu.readCopyUpdate+顺序锁
- 内核态.无限制.模块+用户态.应用程序.链接库函数 软件中断实现系统调用 软中断=可延迟函数vs工作队列
- 内存: 页分配器 slab分配器 申请大内存? zoneBasedBuddySystem buddySystem TLB页表缓存.线性地址

- fs
  - basic: format(fat32 ntfs; 4k block; iNode link dir; path) name(ext) time(a m c) size
    - type: 执行 普通 目录 链接 设备.网卡例外=字符+块 管道
    - `ls -l` umask 777 `struct super_block/inode/file/dentry file_operations`
  - [`/proc`](https://c.isme.pub/2019/02/18/linux-proc): meminfo
  - `/dev`设备: shm stdin/stdout/stderr null `2>&1` `&>`
  - `/etc`: os-release
    - `/etc/ysyconfig/i18n` locale LANG
  - /usr
- disk
  - partition分区: 主(<=4) ext扩展(<=1) logic逻辑 mount挂载
- io: zeroCopy零拷贝
  - pageCache算法=writeBack=writeBehind: lazyWrite dirty
- ENV [shell](shell.md) [fish](fish.md) [cli](cli.md)

```sh
# tmpfs
df -hT
free -m
mount -o remount,size=666M tmpfs /dev/shm
cat /etc/fstab
```

- 编辑不同用户的 crontab: `crontab -e -u xxx`
- 杀进程: `ps aux|grep logistic|grep -v 'grep'|awk '{print $2}'|xargs kill -9`
- 查看端口占用: lsof -i:3110
- 定位高cpu: top -> ps -ef pid(查看进程) -> ps -Lp pid(查看线程)

## ubuntu

- [ppa](https://launchpad.net/ubuntu/+ppas)
- [中文方块](http://wiki.ubuntu.org.cn/Qref/More)

```sh
# fish
sudo apt-add-repository ppa:fish-shell/release-3
sudo apt-get update
sudo apt-get install fish

do-release-upgrade

sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
apt-get update && \
apt-get install -y --no-install-recommends tzdata
# rm -rf /var/lib/apt/lists/*
```

## ubuntu wiki 几点忠告

- 不要当传教士
- 不要强迫自己
- 不要 ‘玩’ linux
- 不要挑剔发行版
- 不要盲目升级
- 不要配置你不需要的东西
- 不要习惯使用 root；只在需要的时候使用
- 不要用商业眼光对待 linux
- 干的正事去

## Unix哲学（Unix编程艺术）

- Doug Mcilroy:

1.让每个程序就做好一件事。如果有新任务，就重新开始，不要往原程序中加入新功能而搞得复杂。
2.假定每个程序的输出都会成国另一个程序的输入，哪怕那个程序还是未知的。输出中不要有无关的信息干扰。避免使用严格的分栏格式和二进制格式输入。不要坚持使用交互式输入。
3.尽可能早地将设计和编译的软件投入试用，哪怕是操作系统也不例外，理想情况下，应该是在几星期之内。对拙劣的代码别犹豫，扔掉重写。
4.优先使用工具而不是拙劣的帮助来减轻编程的负担。工欲善其事，必先利其器。
5.一个程序只做一件事，并做好。程序要能协作。程序要能处理文本流，因为这是最通用的接口。

- Rob Pike:

1.你无法断定程序会在什么地方耗费运行时间。瓶颈经常出现在想不到的地方，所以别急于胡找个地方改代码，除非你已证实那儿就是瓶颈所在。
2.估量。在你没对代码进行估量，特别是没找到最耗时的那部分之前，别去优化速度。
3.花哨的算法在n很小时通常很慢，而n通常很小。花哨算法的常数复杂度很大。除非你确定n总是很大，否则不要用花哨算法（即使n很大，也优先考虑第2条）
4.花哨的算法比简单算法更容易出bug、更难实现。尽量使用简单的算法配合简单的数据结构。
5.数据压倒一切。如果已经选择了正确的数据结构并具把一切都组织得井井有条，正确的算法也就不言自明。编程的核心是数据结构，而不是算法。
给我看流程图而不让我看（数据）表，我仍会茫然不解；如果给我看（数据）表，通常就不需要流程图了；数据表足够说明问题了。
6.原则6：没有原则6

- Ken Thompson:

模块原则：使用间洁的接口拼合简单的部件。
清晰原则：清晰胜于机巧。
拿不准就穷举。
组合原则：设计时考虑拼接组合。
分离原则：策略同机制分离，接口同引擎分离。
简洁原则：设计要简洁，复杂度能低则低。
吝啬原则：除非确无它法，不要编写宠大的程序。
透明性原则：设计要可见，以便审查和调试。
健壮原则：健壮源于透明与简洁。
表示原则：把知识叠入数据以求逻辑质朴与健壮。
通俗原则：接口设计避免标新立异。
缄默原则：如果一个程序没什么好说的，就沉默。
补救原则：也现异常时，马上退也并给出足够错误信息。
经济原则：宁花机器一分钟，不花程序员一秒钟。
生成原则：避免手工hack，尽量编写程序去生成程序。
优化原则：雕琢前先要有原型，跑之前先学会走。
多样原则：决不相信所谓“不二法门”的断言。
扩展原则：设计着眼未来，未来部比预想来得快。

## mark

- [Linux快速工具教程](http://linuxtools-rst.readthedocs.org/zh_CN/latest/)
- [linux命令查询](http://wangchujiang.com/linux-command/)
- [linux kernel](https://www.kernel.org/)
- [图灵社区 - 理解UNIX进程](http://www.ituring.com.cn/book/1081)
- [用iptables搭建一套强大的安全防护盾](http://www.imooc.com/learn/389)
- [Linux系统扫描技术及安全防范](http://www.imooc.com/learn/344)
- graphical user interfaces make easy tasks easy, while command line interfaces make difficult tasks possible
- 没有消息就是最好的消息
- 一切皆文件.fd
- 进程皆有环境
- 上千个 linux 版本: <https://distrowatch.com/images/other/distro-family-tree.png>
  - LFS: 从0构建一个linux
  - gentoo: 追求极限的配置/性能
  - archlinux: 滚动升级
  - Ubuntu: 个人/桌面
  - kali: 安全 渗透测试
