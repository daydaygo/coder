# shell

- why: 脚本, 简化管理操作
- env: sh/bash/zsh
  - [shortcut](http://blog.chinaunix.net/uid-361890-id-342066.html): `⌃ a/e u/k w d` `⎋ b/f`
  - `#! /bin/bash`
- syntax
  - type: string`'' ""` array(`arr=(a0 a1 a2)` `${arr[0]} ${arr[*]} ${#arr[*]}`)
  - var
    - userDefine 局部变量: define `k=v` use `$k ${k}`
    - preDefine shell变量: `$n $0 $# $* $@` `$?退出状态 $$当前pid $!最后pid $-`
    - ENV 环境变量
  - op: 算术.数字`+-*/% == !=` 关系.数字`-eq -ne -gt -lt -ge -le` bool`!非 -o或 -a与` 逻辑`&& ||`
    - string`== != -z长度0 -n长度不为0 $是否为空`
    - 文件测试`-b块设备 -c字符设备 -d目录 -f普通文件 -g是否设置SGID -k是否设置StickyBits -p管道 -u是否设置SUID -r可读`
    - 重定向 `> < >> >& <& <<tag将tag之间的内容作为输入 /dev/null`
- [shell-菜鸟教程](https://www.runoob.com/linux/linux-shell.html)

```sh
# env
export source env # ENV; declare -x; .bashrc .bash_history /etc/issue /etc/motd
export PATH=$PATH:~/bin # array
echo $SHELL # cat /etc/shells

name='foo'${a}"$b" # echo $name; 字符串拼接
echo ${arr[*]} # 输出数组所有值
echo ${arr} # 输出数组第一个值

set -u # 查看变量, 当变量不存在时提示; unset
echo -e "\e[1;31m xxx \e[0m" # \033[31m[error]\033[0m; \x
echo -e "a\tb\n" # 开启转义
printf "%s\t%s\n" "a" "b"
read -p '提示信息' -t 30 -n 30 -s passwd # prompt time nchar secure

# 文件包含
. file
source file

# 位置参数变量
$1 -$9 ${10} # 给脚本传递变量
$# # 参数个数
$* # 所有参数, 当做一个
$@ # 所有参数, 当做多个
# 预定义变量
$? # 上一条命令的返回值
$$ # 当前进程的 pid
$! # 后台运行的最后一个进程的 pid

$(expr $a + $b) # expr/let, 进行数组运算, 必须要有空格
$((运算式)) / $[运算式]

# op
n1 -eq -ne -gt -lt -ge -le n2 # 数组判断
-z -n == != # # 字符串判断
-a -o ! # and/or/not # 多重条件判断
# 文件判断
test -e file # 文件类型
[ -e file] # 注意要有空格
-w file # 判断是否有 w 权限, 注意不区分用户
file1 -nt -ot -ef file2 # new old equeal: 使用 iNode 判断

# ctl
env|grep USER|cut -d '=' -f 2 # 获取登录用户
df -h|grep vda1|awk '{print $5}'|cut -d '%' -f 1 # 获取系统盘使用率
[ -n $(ps aux|grep nginx|grep -v grep)] # 判断进程是否运行
[ -z $(echo $num|sed 's/[0-9]//g')] # 判断是否为全数字

# if
[condition] && echo yes || echo no # 类似3元操作符的栗子
if [condition]; then # then 可以换行
    todo
elif [condition]; then
    todo
else
    todo
fi

# for
for i in $(seq 1 10); do # {1..10} `ls` $* $str /proc/* $(ls *.sh)
    echo $i
done
for((i=1;i<=10;i++)); do # 不能使用 i++
  echo $i
done

# while
i=1
while [ $i -lt 10 ]; do
  echo $i
  let i++
done

# case
case $var in
"yes")
    echo
    ;;
"no")
    echo
    ;;
*)
    echo
    ;;
esas

# func
$(date) # `date`
# 换行+注释
ls \
-l `# long`
```

- Bash Shell Function Library https://github.com/SkypLabs/bsfl
- bash lib https://github.com/aks/bash-lib
