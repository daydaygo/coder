# perl

- perl 的思维方式: 符号代替文字
- 换工作了, 也许以后就不会再用到 perl了, 估计也不会有新笔记了, 所以就整理发出来算了
- [cpan](http://www.cpan.org)

## 工具

- window

使用 activePerl(使用了其他的, 就是自带模块的多少不同而已, 但是 activeperl 安装模块要方便多了)
模块安装: cmd 下 使用 `ppm install xxx`, 或者使用图形化界面
如果运行图形化界面出现 `database disk image is malformed` 的错误, 是因为 SQLite 数据在操作的时候被强制中断了, 删除 `C:\Users\Administrator\AppData\Local` 目录下的文件即可

- linux

直接使用 `sudo apt-get install perl` 安装就行了
模块安装: 使用 cpan, 第一次使用的时候会进行配置, 默认配置就行了
```bash
apt-get install cpan # 如果服务器上面没有
cpan install xxx # 安装模块
cpan i /xxx/ # 查找
cpan m /xxx/ # 查找模块
```
linux crontab 中配置的perl脚本无法运行, 添加 `PATH="$PATH"` 到 crontab 中即可

模块安装需要注意的一点是: 名字一定要对应(大小写, ::)

## 解决方法不止一种

perl里面提供了很多灵活的 **识别** , 可以通过 更简单 / 更多样 的代码, 来完成需要做的事
定界符: m// s/// qw// , 还可以使用其他的定界符
函数的括号: 类似 shell 中的写法, 比如我就特别喜欢这样用: `say FILE Dumper $arr;`
默认变量: `$_` , 就是这么强大

## 奇怪的东东

`$_` : 默认变量
`$1-9`: 正则匹配中匹配 () 中的内容
`$! $@` : 错误信息, 可以配 die 或者 `eval{};` 使用
`$| = 1;` : 设置为直接输出, 不放到输出缓冲区
`$ARGV[0]` : 运行perl程序时传入的参数, 比如 `perl xxx.pl 201506`, 那么 `$ARGV[0]` 的值就是 201506
`$$` : 当前进程的pid

## 常量

- 数字

perl中的数都是存储为 double 类型, 所以 `10/3 = 3.33..`, 需要使用 `int 10/3`
浮点数: `-6.5e7`
其他进制: 0/0x/0b
比较: `< > == !=`
进制转换=> `hex();` 处理16进制 `oct();` 可以处理2,8,16进制,默认为8进制

- 字符串

应该都知道: 用 **引号** 的就是字符串
字符串连接: `'czl' . 'czl'`
字符串重复: `'czl' x 3`
`\l \u \L \U \Q \E` : 大小写转换
比较: `eq ne lt gt le ge`
heredoc结构:
```perl
$tmp =<<TMP; #注意分号在这里
xxxx
TMP
```

数字转字符串: 所见即所得
字符串转数字: 只能转成10进制

## 变量

`$var @arr %hash {$var}` , perl中的数据结构很简单, 就是 变量 / 数字=>变量(数组) / 变量=>变量(hash), 要想使用复杂的数据结构, 就需要使用指针, 比如 数字=>指针(指针也是变量, 如果这个指针指向一个数组, 那么就可以达到多维数组的效果了)
代码点: 使用16进制数来表示特殊字符 `ord('?') / chr(0x05D0)` , 这个思想在之后的 Unicode 字符转化的时候会用到的
bool值: 数字=> 0/其他; 字符串=> ''/其他
undef: 相当于C语言中的null

## 数组(列表)

`@arr`: 声明数组
`$arr[0] @arr[0] @arr->[0]`: 访问第一个元素
`$#arr`: 数组最后一个元素的索引
`$arr[-1] / $arr[$#arr]` : 数组最后一个元素(注意2者还是有区别的, 比如使用 `splice()` )
声明数组: `@arr = ('a', 'b', 'c'); @arr = ('a'..'z'); @arr = qw/a b c/;` 使用 qw 不需要使用引号
交换2个变量的值: `($tmp1,$tmp2) = ($tmp2,$tmp1);` perl 支持直接这样写
列表的赋值 `($tmp1,$tmp2,$tmp3) = qw(1 2 3);`
数组尾部操作 `pop(@arr); / push(@arr,0);`
数组前端操作 `shift(@arr); / unshift(@arr,0);`
从数组指定位置开始删除/替换 `@rmoved = splice(@arr,$start,[$len],[qw()]);`
不改变原列表,返回列表的逆序 `@arr2 = reverse(@arr1);`
不改变原列表,返回排序后的列表 `@arr2 = sort(@arr1);`
列表 => 标量 : `$a = @arr; / scalar(@arr);` 计算数组的长度
切片(只能用在数组上, hash 就考虑转化为数组): `my($a,$b,$c) = @arr[2,3,4]`

## hash

声明hash  `%hash = (key,val,key,val); %hash = (key=>val , key=>val );`
`%new_hash = reverse %old_hash;`  键值互换
获取键 `@k = keys %hash;`
获取值 `@v = values %hash;`
使用hash  `while(($k,$v) = each %hash){$k...$v};` , 数组也可以使用这样的方法
`foreach $k(sort keys %hash){...}` perl存储hash是有默认顺序的(对我们而言是无序的),需要自己进行排序
判断键是否存在 `if(exists $hash{'dino'}){...}`
删除指定键值 `delete $hash{'dino'};`
数组/hash的使用 , 注意 $arr / $hash 是指针 , 并不是真正的 数组或者哈希, 所以需要使用真正的值的时候 , 都需要使用相应的写法
数组: 声明=> my $arr = []; | 使用=> $arr->[0];/@$arr[0];(需要当作数组时,使用@$arr,这个看操作符与函数)
hash: 声明=> my $hash = {}; | 使用=> $hash->{a};

## 指针

```perl
$a = \@arr; # 将数组转化为指针
$a->[0]; # 访问数组第一个元素
push @$a,(xxx); # 给数组中添加一个元素
scalar @$a; # 求数组的长度
```

## 字符串与排序

```
index($big,$small); # 返回子字符串开始的位置
substr($str,$start,[$len]);
chomp($str) # 去掉字符串末尾的 换行符
chop($str) # 删除字符串末尾一个字符
join('-',@arr) # 将数组连接为字符串
split('-' , $str) # 将字符串分割为数组

@arr2 = sort {"\L$a" cmp "\L$b"} @arr; # 字符串从大到小排序
sort {$hash{$a} <=> $hash{$b}} keys %hash; # hash按键排序
{$hash{$a} <=> $hash{$b} or $a cmp $b}; # 多次排序
```

## 正则

. 匹配任意字符 | 或
量词: `* 0个或多个 + 一个或多个 ? 0个或者1个`
模式分组: `() $1-9`
字符集: `\d \w \s \W'
`[]`: 使用 `[]` 只匹配单个字符, 不能用来匹配中文
修饰符: `/s /i /g`
锚位: `^ $ \b`, 开头结尾, 单词边界

```perl
 $a = $str =~ /.../ ; # 这个返回的0/1
 ($a) = $str =~ /.../ ; # 这个将 $1 返回给 $a
 ($a = $str) =~ s///; #同上
m#/[^/]*?$#  # 匹配url中最后一个/../中的内容

# big_money
sub big_money{
	my $num = sprintf "%.2f",shift @_;
	1 while $num =~ s/^(-?\d+)(\d\d\d)/$1,$2/;
	$num =~ s/^(-?)/$1\$/;
	$num;
}
```

## 控制结构

必须要加{}

```perl
if(){}elsif(){}else{} # if结构, 注意是 elsif, unless:和if相反
while(){}
foreach $a(@arr){$a;} / foreach(@arr){$_;} # 可以使用 for
```

next:用于跳过一次循环( 相当于 continue)
last: 退出当前循环 (相当于 break) ( 注意 last 不能用于block中 , 例如 `eval{}` , 子函数(last不在一个循环里) , 否则会导致 last 终止这个block上一级的循环)

## 子程序(函数)

```
&do($a,$b);
sub do{
	my($a, $b) = @_; # 建议用这样的方式
	# 返回值直接放到最后一行, 或者使用 return
}
```

## 文件

目录操作,大部分操作和Linux下面的命令一致

```perl
chidr '/etc'; # 改变目录
@all_file = </etc/*> # 获取目录下的文件, 相当于执行 ls, 但是带有完整的路径
mkdir '...',0755; # 创建目录,0755表示8进制
rmdir '...'; # 和Linux下一样,只能删除空目录
unlink qw(...); # 删除文件,返回成功输出文件的个数
@all_file = stat 'C:/Users/czl/Desktop/ubuntu/www/test/daily.md'; # 8->atime 9->mtime 10->ctime, http://zhidao.baidu.com/question/93521537.html(详细介绍)
$mtime = (stat $file)[9]; #文件修改时间, 为时间戳

#修改文件名
my $dir = "C:/迅雷下载/权利的游戏1";
opendir DIR ,$dir;
while(my $file = readdir DIR){ if($file =~ /S01E(\d+)/){ rename "$dir/$file","$dir/$newfile"; }}

<stdin> #标准输入
chomp($text = <stdin>) #读取一行输入赋值给$text,并删除末尾的空格
chomp(@text = <stdin>) #读取多行输入赋值给@text,并删除末尾的空格,接收EOF信号时结束(linux:ctrl+d/window:ctrl+z)
while(<>){} / while(<stdin>){} / foreach(<stdin>){} #读取输入并操作, 直到接收EOF为止

use 5.01; say "xxx"; #自动加上换行
print #直接输出
printf #类似c语言中的printf
open IN,$file; while(<IN>){} #读文件操作
open OUT,'>',$file; say OUT 'xxx'; #写文件操作

open IN,'<',$file; #一次读取文件所有内容
local $/=undef;
my $in = <IN>;

while(<CSV>){ chomp; my @tmp = split ',';} #处理csv, 不能处理文本中含有 , 的情况
```

## 编码

```perl
# 当出现乱码时 , 可以考虑使用 use encoding 'utf-8'; 如 json->encode
use utf8; => 要求文件使用utf8格式保存, 但是加入这句后报wide character错误(不要使用)
use encoding 'utf8'; => 代码/输入/输出都适用utf8
open FILE,'<:utf8',$file; 使用utf8编码来打开文件

$content = decode("Detect",$content);
$content = encode("utf8",$content);
#检测编码
use Encode::Detect::Detector;
$charset = detect($content); # 服务器上面测试检测utf8 可以 , 但是非utf8问题多多

# xml, 速度慢, 还耗cpu, 所以我直接使用正则进行匹配
use XML::Simple;
my $xml = XMLin($content, ForceArray => 1);
my $items = $xml->{channel}->[0]->{item}; #注意 ->[0]

# 上面的基本不实用, 留着是防止以后我又碰到了

#将文本转化为utf8编码
use Encode;
Encode::from_to($content,'gbk','utf8');

use URI::Escape;
$url = uri_escape($url) / $url = uri_unescape($url); # 对url进行编码

#将代码点转成相应的特殊字符
$content =~ s/\\u(.{4})/Encode::encode('utf8',pack("U", hex($1)))/eg; #将字符串中的unicode转化为汉字 注意需要对替换的字符进行encode操作
$test =~ s/\\x(..)/chr hex $1/eg; #使用 /e 才能执行 chr/hex 函数的功能 , 否则当成普通字符进行处理

# json
use JSON;
my $data = encode_json{tmp => $tmp};

use JSON::Parse qw/json_to_perl/; # 这个的速度要比 JSON->decode() 快很多
$content = json_to_perl($content);
# say $content->[0]->{name}; 一定要注意这个 ->[0] , 被坑过
```

## 日期时间

```perl
time; #返回当前的时间戳 -> 1422618140
Time::HiRes::time; # 返回 当前的时间戳 -> 1422618140.30283

localtime; #格式化后的本地时间
($sec,$min,$hour,$mday,$mon,$year_off,$wday,$yday,$isdat) = localtime; #返回不同的本地时间数值
my($day , $mon , $year) = (localtime)[3,4,5]; $mon +=1; $year +=1900; #返回当前日期
my($day , $mon , $year) = (localtime(time()-86400))[3,4,5]; $mon +=1; $year +=1900; #求前一天的日期
my $date = sprintf("%04d-%02d-%02d",$year , $mon , $day);

chomp(my $date = `date +%Y-%m-%d`); #使用系统自带的时间日期命令 , 此为Linux系统

use Date::Calc qw/Today Today_and_Now Add_Delta_Days Add_Delta_YM Day_of_Week Days_in_Month/;
my $curdate = sprintf("%04d-%02d-%02d", Today());
sprintf("%04d-%02d-%02d %02d:%02d:%02d", &Today_and_Now());
sprintf("%04d%02d%02d", Add_Delta_Days(Today(), -$delta)); //$delta为整数
sprintf("%04d%02d", Add_Delta_YM(Today(), $Y, $M)); # 月份相加

use Date::Parse; #需要安装 datetime-format-dateparse
my $time = str2time('Mon, 04 May 2015 17:12:15 +0900'); #日期转化为时间戳

use POSIX qw(strftime);
my $date = strftime "%Y-%m-%d %H:%M:%S", localtime($time); #时间戳转化为日期
```

## 调试

```perl
perl -c xx.pl  # 看看脚本是否可以编译通过
perl -d xx.pl  # 进入调试模式, n: 下一个函数; s: 进入函数执行; q: 退出

use warnings; => 使用内置警告, 也可以在执行perl时使用 perl -w xxx.pl, 或者在脚本中使用 #! /usr/bin/perl -w
     use diagnostics; => 报告更为详细的警告信息(会使程序变慢)
     use autodie; => 自动输出错误信息
use strict; => 使用严格模式, 会强制要求使用 my($tmp) 定义变量来限制变量的作用范围
xxx or die 'xxx'; => 测试此处是否执行成功, 否则返回自己写的字符串
eval{xxx}; if($@){ $@; } => 将代码段放在eval中运行 , 即使代码段有严重错误, perl也会继续执行 , 注意大括号外的分号
use Data::Dumper; say Dumper($tmp1,$tmp2); => 查看数据,$arr输出数组, @arr输出数组中的值
```

## 进程管理

```perl
system 'xxx'; # 创建子进程,执行外部命令xxx, perl进程需要等待子进程执行完再处理剩余的部分
exec 'xxx'; # 结束perl进程,执行外部命令xxx, 配合fork使用,使进程在完全独立的子程序中运行
`xxx`; # 捕获外部命令的输出,等效于qx(xxx);
open DATE ,'date|';/open MAIL ,'|mail xxx'; # 采用句柄的方式

# 多进程
# 折腾了几个小时的perl多进程, 最后自己把自己坑住了:
# http://blog.chinaunix.net/uid-17196076-id-2817715.html 初步了解多进程
# http://www.linuxidc.com/Linux/2014-10/107468.htm perl捕捉系统信号
sub duo2{
    my $child = {};
    my $i = 0;
    while($i < 10){
        while ((keys %$child < 5) && $i<10) {
            my $pid = fork();
            while(!defined($pid)){
                $pid = fork();
            }
            if($pid){
                $child->{$pid} = 1;
            }else{
                # do something
                exit;
            }
        }
        my $pid = wait();
        delete $child->{$pid};
    }
    while((my $pid = wait()) != -1){}
}
```

## dbi

```perl
use DBI; # 使用mysql需要安装 DBD::mysql 模块
my $dbh = DBI->connect("DBI:mysql:haodf:114.80.77.92:3306","guahao","QBT094bt",{'RaiseError'=>1}); # 连接数据库
$dbh->do("SET NAMES 'utf8'"); # 设置编码
$row = $dbh->do($sql); # 执行一个sql , 如果是非select操作, 返回操作影响的行数

#查询数据
my $all = $dbh->selectall_arrayref($sql);  for my $row(@$all){ $row->[0]; $row->[1]; }#多列数据
my $all = $dbh->selectcol_arrayref($sql); for $row(@$all){ $row; }#单列数据
my $all = $dbh->selectrow_arrayref($sql);  { @$all[0]; }#单行数据

#插入数据
$arr = [];
push @$arr,(xxx, xxx, xxx, xxx) # $arr中插入数据的方式
sub do_insert{
	my ($arr) = @_;
	return unless ($arr and @$arr);

	my $sql = "insert ignore shop(xxx,xxx,xxx, xxx) values";
    $sql .= join ',' ,("(?,?,?,?)") x (@$arr / 4);

    eval{ $dbh->do($sql,undef,@$arr); };
    if($@){ say $@; }
}
```

一个防止多进程造成表死锁的思路: 子进程只进行 insert 操作, 插入到同一张表, 最后在主进程中执行一次 update 操作
遇到的一个问题: 把 dbi 封装到类中并使用 `eval{}`, 导致找了半天才发现是 dbi 连接的问题

## LWP

```perl
use LWP::UserAgent;
my $ua = LWP::UserAgent->new;
#无法获取到信息时 , 修改这2项
#$ua->cookie_jar({});
#$ua->default_header(''=>'',''=>'');
my $url = '';
my $content = $ua->get($url)->content; #通过get访问
Encode::from_to($content,'gbk','utf8');
my $ccontent = $ua->post($url,{ 'cityId'=>57, 'provinceId'=>1,})->content;#通过post访问
```

## 其他

perl中 use / require 的区别 : http://www.cnblogs.com/itech/archive/2010/11/22/1884345.html

```perl
# perl 入门例子
#! /usr/bin/perl # perl程序的安装目录
@lines = `perldoc -u -fatan2`; # 执行外部命令, 将命令执行的结果一行一行赋值给数组@lines
foreach(@lines){ # 按行处理@lines
	s/\w<([^>]+)>/\U$1/g; # 正则替换
	print; # 直接输出
}

# perl 中的重试机制
sub try3{
    my $try = 0;
    while(1){
        eval{ do something };
        if($@){ }else{ last; }
        last if $try++ > 3;
    }
}
```
