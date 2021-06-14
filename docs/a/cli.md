# 命令行 CLI

- 文件权限: ll rwx user/group/other 777
- process 进程
  - S status: RSX DZT(长时间=不正常)
  - signal 信号
- IO
  - 接口
    - fd = open(pathname, flags, mode)
    - rlen = read(fd, buf, count)
    - wlen = write(fd, buf, count)
    - status = close(fd)
  - io模型
- net
  - socket
    - status: TIME_WAIT一般通过优化内核参数能够解决；CLOSE_WAIT一般是由于程序编写不合理造成的

- dir/fold 目录
  - basic 基本: mv rm cp mkdir ln
  - wander 漫游: find/locate ls cd pwd
  - 管道 `|`;
  - 重定向: `>`; `&>` stdout+stderr
- text/string
  - view 查看: less tail cat
  - filter 过滤: grep awk sed diff/patch
  - count 统计: sort uniq
  - vim
- compress 压缩
  - tar bzip2-`.bz2` gzip-`.gz` unzip unrar
- ops 日常运维
  - shutdown mount chmod chown su/sudo yum/apt/apk password service/systemctl
- sys status 系统状态
  - 问题排查: 设备+cache+bottleneck; 信息收集->改动集合->问题抽象->问题排查
  - basic 概览: ps top free df ifconfig uname
  - cpu:
    - top: load wa sy,si,hi,st(5%) 各核均衡
    - vmstat: b 等待队列; cs 上下文切换; si/so 交换分区
    - sar(sysstat): -u 默认; -P ALL 各核; -q 队列; -w 上下文切换
    - mpstat/pidstat/dstat
    - `/proc`: `/proc/loadavg` top/uptime; `/proc/stat` 各核; `/proc/pid` process
  - mem 合理参数/优雅代码/禁用swap
    - swap: 很多性能场景的万恶之源，建议禁用
    - slab: 内核的缓存文件句柄等信息等的特殊区域
    - 程序内容: 低位-> text 程序代码 -> data 初始化静态变量 -> bss 未初始化静态变量 -> heap -> 未使用 -> stack 执行栈 -> system 命令行参数
    - volitile: 工作内存->主存
    - cpuCache/cacheLine/FalseSharing hugePage/TLB numa/swap
  - io 随机/顺序 读/写 cache-缓冲区/断电-丢数据
    - 设计
      - db BTree: 减少磁盘访问和随机读取
      - pgsql wal 日志 / es translog 日志 -> 预写-防丢
      - kafka: 顺序写(topic多->随机写) DMA-零拷贝(zero copy 内核态/用户态)->绕过内存直接发数据
      - redis: 内存模拟存储
  - net
    - ss ifstat nload iptraf
    - 抓包/流量复制: tcpdump+wireshark Fiddle2/charles gor
    - test: ping tracepath dig nslookup whois w/whoami nmap iperf nmon telnet nc
    - bench: wrk ab
    - iptables
- work 工作常用
  - basic: export xargs date whereis crontab
  - net: ssh/scp/rsync wget/curl
- mac
  - [iproute2mac](https://github.com/brona/iproute2mac): ip vs netstat/ifconfig/ndp/arp/route/networksetup

```sh
mkdir -p # -p 包含父目录
tail -f -n # -f 滚动查看; -n 行数
sort -t -k # -t 分隔符; -k 排序
grep -rn --color -C10 # -r recursion; -n num; -C 前后n行; -v invert
gzip file # gzip -d xx.gz
tar cvfz archive.tar.gz dir/ # tar xvfz archive.tar.gz
zip -r xx.zip file # unzip xx.zip
mount /dev/sdb1 /xiaodianying
chmod a+x a.sh
systemctl restart mysqld
kill -l # 查看 linux signal
uname -a # /etc/issue
ps -ef|grep java
top -H -p pid
pgrep -u hchen # name -> pid
pstree -a|grep server.php # 查看进程树
strace -pf # 查看线程
ifconfig en0 # ip addr show
netstat -ant # -a all socket; -n addr number; -t tcp
export PATH=$PATH:/home/xjj/jdk/bin # 设置环境变量
find ./ -name "*.o" -exec rm {} \; # -exec; -type f/i/d; -atime/mtime/ctime
find . | grep .class$ | xargs rm -rvf # 读取输入源 -> 逐行处理
locate updatedb # /var/lib/mlocate /etc/updatedb.conf
ls *.rmvb | xargs -n1 -i cp {} /xxx # cp 所有 rmvb 到新目录
scp -r /dir ip:/dir
wget -c url # -c continue
wget -r -p -np -k $url # 下载整站

# user
id www
groupadd www
useradd -g www -s /sbin/nologin www
visudo: sudo: joe ALL=(ALL) NOPASSWD: ALL
passwd xxx

# cpu
ps -eo %cpu,pid |sort -n -k1 -r | head -n 1 |  awk '{print $2}' |xargs  top -b -n1 -Hp | grep COMMAND -A1 | tail -n 1 | awk '{print $1}' | xargs printf 0x%x

# mem 物理内存
top -> M # 排序查看内存
ps -p 75 -o rss,vsz # 进程使用物理内存
free -h / sar -r # used free buffers(dir/iNode FIFO) cached(slab; file; LRU; 文件IO高)
cat /proc/meminfo # 内存每个区详情
slabtop
# mem 内存泄露
cat /proc/28806/status |grep RSS # RSS包含共享内存
smem -p|grep USS # USS才是PHP申请的内存

# io
sync # buffer 落盘
mkdir /memdisk # 内存盘
mount  -t tmpfs -o size=1024m  tmpfs /memdisk/
time dd if=/dev/zero of=test.file bs=4k count=200000
top # 查看 wa
vmstat 1 # wa 和 io(bi/bo)
sar -b 1 2 # 性能相关 io 情况
iostat -x 1 # 问题相关 io 情况
iotop # 使用 io 最多线程
lsof # 查看 fd
df -h # 磁盘
du -h file
cat /dev/null > file # 清空文件

# net
curl
nc -z -w2 ip port # tcp
hostname -I # curl cip.cc
nslookup # dns /etc/resolv.conf dig
netstat -antp | awk '{a[$6]++}END{ for(x in a)print x,a[x]}' # 统计当前系统的连接
watch cat /proc/net/dev
tcpdump -i eth0 -nn -s0 -v port 80 -w test.pcap # 抓包: -i interface; -nn 域名解析; -s 包长度; -v verbose; -w write
tcpdump -A -s0 port 80 # -A ascii; -X hex
tcpdump -i eth0 host 10.10.1.1 dst 10.10.1.20 # 特定ip
ipcs -q # linux ipc

# work
ssh root@ip "sudo runuser -l www -c 'cd /data/service.base; whoami; pwd;git pull'"
cut -f 2 -d ':' file # f: 列 d: 分隔符; 要注意 分隔符
grep -vi -C10
sort -f -n -r -t -k n,m # f 忽略大小写 n 以数值排序, 默认为字符串 r 反向排序 t 分隔符, 默认为 \t k 指定范围
wc -l -w -m # l line w word m char
```

- [iproute2](http://www.policyrouting.org/iproute2.doc.html)
- [tcpdump](https://hackertarget.com/tcpdump-examples)

## sed

```sh
sed [option -n -e] 'action' file
sed -n '1,4 p' file # 参数: -n; 范围: 1,4; 操作: pdw aic

5 # 选择第5行。
2,5 # 选择2到5行，共4行。
1~2 # 选择奇数行。
2~2 # 选择偶数行。
2,+3 # 和2,5的效果是一样的，共4行。
2,$ # 从第二行到文件结尾。
/sys/,+3 # 选择出现sys字样的行，以及后面的三行。
/\^sys/,/mem/ # 选择以sys开头的行，和出现mem字样行之间的数据

sed '/^sys/,/mem/s/a/b/g' file # 操作: s 替换模式; flag: gpwi
```

## awk

- 内置变量: FS 分隔符; OFS 输出分隔符; NF 列号; NR 行号; RS 记录分隔标志; ORS 记录输出分隔标志; FILENAME
- 编程语言特性: math string

```sh
awk 'pattern {action}' file
awk -F "," '/^a/ {print $1,$2}' file # 参数: -F ","; 范围 /^a/; 操作: {print $1,$2}

netstat  -ant |
awk ' \
  BEGIN{print  "State","Count" } \ # BEGIN: 参数 表头 变量; {FS=":"} 默认 \t
  /^tcp/ \ # Pattern: 可选, 正则匹配
  { rt[$6]++ } \ # Action: 按行处理 统计打印
  END{ for(i in rt){print i,rt[i]} }' # END: 汇总 输出
```

## iproute

> net-tools(netstat) -> iproute2

| 用途 | net-tools | iproute   |
| ---- | --------- | --------- |
| 统计 | ifconfig  | ss        |
| 地址 | netstat   | ip addr   |
| 路由 | route     | ip route  |
| 邻居 | arp       | ip neigh  |
| VPN  | iptunnel  | ip tunnel |
| VLAN | vconfig   | ip link   |
| 组播 | ipmaddr   | ip maddr  |

```sh
ss -s # 统计信息
ss -atr # 正在监听的 tcp 连接
ss -atn # 仅 ip
ss -alt # 系统中所有连接
ss -ltp | grep 444 # port <-> pid
ss -a -u # socket 类型: -u udp; -t tcp; -w raw; -x unix
ss dst 10.66.224.130 # 某个ip的所有连接: :http :443
ss dport = :http # 所有 http 连接
```

|        | Recv-Q                                                       | Send-Q                                                       |
| ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| LISTEN | 代表建立的连接还有多少没有被accept，比如Nginx接受新连接变的很慢 | 代表listen backlog值                                         |
| ESTAB  | 内核中的数据还有多少(bytes)没有被应用程序读取，发生了一定程度的阻塞 | 代表内核中发送队列里还有多少(bytes)数据没有收到ack，对端的接收处理能力不强 |

## 命令行工具推荐

- fish shell-一款易于上手的终端shell
  - franciscolourenco/done ——在长时间运行的脚本完成后发送通知
  - evanlucas/fish-kubectl-completions
  - fzf——将 fzf 工具与 Fish 集成在一起的插件
- starship-款强大的shell提示工具
  - git 状态
  - 编程语言环境信息
  - 错误提示
- Z-一个快速切换文件路径的命令工具 <https://github.com/rupa/z>
- fzf-一款好用的模糊查找工具 <https://github.com/jethrokuan/fzf>
  - 模糊查找 文件/历史命令/进程/git 等
- fd-升级版的find查找工具
- ripgrep-like grep but better升级版的grep工具
- htop and glances-一款给力的系统监控工具
- virtualenv and virtualfish-Python虚拟环境管理工具
- pyenv,nodenv,and rbenv-一款对Python,Node以及Ruby进行不同版本管理的工具
- pipx-一款Python依赖管理以及环境管理工具
- ctop and lazydocker-款给力的Docker监控工具
- homebrew-一款MacOS下的软件包管理工具
- asciinema-款终端会话记录工具，支持从动画中进行拷贝
- colordiff and diff-so-fancy-升级版的diff工具
- tree-一款用于展示目录树状结构的命令工具
- bat-一款升级版的cat工具
- httpie-一款升级版的cur1工具
- tldr-Too long,Don't read,简化版的man pages查看工具。
- exa-一款升级版的ls命令工具
- mas-App Store的命令行工具
- ncdu-一款磁盘使用分析的命令行工具

```sh
lsof -i:80 # 查看端口上的进程
curl -sOL -i # -s slient -O output-file -L location -i include -H header
pstree -p # 查看进程数
ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime # 修改时区

jq # 命令行 json
htop # 替代 top

# cheatsheet
curl https://cht.sh/:cht.sh > ~/bin/cht
cht go/md5 # go 语言 md5 怎么写

lrzsz/sz/rz
rysnc -avz /from user@host:/to

# tmux session http://www.ruanyifeng.com/blog/2019/10/tmux.html
tmux new -s <session-name>
tmux detach # c-b d
tmux ls
tmux attach -t 0 # <session-name>

# 压测
ab -n 100 -c 10 -p 'post.txt' -T 'application/x-www-form-urlencoded' 'http://test.api.com/ttk/auth/info/'
hey -z 30s -c 90 --host "coffee.default.example.com" "http://47.100.6.39/?sleep=100"
wrk -t10 -c1000 -d40s --latency URL
```

## man 查看帮助文档

```sh
man -k # apropos
man -f # whatis
man -w hub # 查看 man page 文件位置, 通常在命令 share 目录下
/usr/local/Cellar/hub/2.14.2/share/doc/hub-doc/ # 可以看看应用的 share 目录下是否有 html 版的 man page
groff -man -Thtml (man -W hub) > ~/tmp/t.html && open ~/tmp/t.html # mac 下实现查看 html 版 man page
# PAGER=cat # man 文件转成纯文本
# man less | col -b > less.txt

tree -L 2 --dirsfirst -d # -L level; -d onlyDir

ldd /usr/bin/java # 查看可执行文件所使用的 ldd
```

## ag the_silver_searcher 查找工具

```sh
ag xxx # 当前文件夹下递归搜索文件内容
ag --html xxx # 限制文件类型
ag -g xxx # 查找文件名

# 根据目录名查找
find <dir> -type d -iname <name> # -i 不区分大小写; -type, d 为目录
```

## mark

- 22款好用的 cli 工具: <https://mp.weixin.qq.com/s/i3qEnIF9XeKstKFebKlsDw>
- graphical user interfaces make easy tasks easy, while command line interfaces make difficult tasks possible
- GUI 的优点和缺点: what you see is what you get; what you see is all you get.
