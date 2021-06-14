# c

## makefile

- [Make 命令教程](http://www.ruanyifeng.com/blog/2015/02/make.html)

```makefile
target ... : prerequisites ... # taget: obj exe phony
 command # sh command
 ...
 ...

objects = main.o kbd.o command.o display.o \
  insert.o search.o files.o utils.o
cc = gcc # var

edit : $(objects)
 cc -o edit $(objects)

main.o : defs.h # auto add: main.c
kbd.o : defs.h command.h
command.o : defs.h command.h
display.o : defs.h buffer.h
insert.o : defs.h buffer.h
search.o : defs.h buffer.h
files.o : defs.h buffer.h command.h
utils.o : defs.h

-include foo.make *.mk # - 忽略文件读取不到

.PHONY : clean # 伪目标
clean :
 rm edit $(objects)
```

## c 基础

环境配置尝试:

- get-mingw: 需要下载安装, 没有好用的源
- cygwin: gcc 版本有点老
- codeblock: 有自带 mingw 的版本, gcc 版本有点老

## C语言入门

> 慕课网: <http://www.imooc.com/learn/249>

注释: `// /**/`
标识符: 给变量或者函数起名; 严格区分大小写; 不能使用关键字; 对应内存中的存储位置
int 类型的大小和编译环境有关
自动类型转换: char -> int -> double(小字节 -> 大字节)
强制类型转换: 不修改原值; 没有四舍五入
自增运算符: ++x 先运算, 再取值
算数 / 比较 / 逻辑 / 三目 运算符
break: 只跳出当前循环; switch
continue: 直接进入下次循环
goto: 反正我没用
形参: 只在函数内部有效, 对应调用函数时使用的 **实参**
递归函数: 递归都可以写成循环, 只是难易度而已
变量: 函数体(包括 if) -> main函数 -> 全局
变量类型: 感觉就 **static** 需要记一下
内部函数(static) / 外部函数(extern): 内部函数只在当前源文件有效

抽象数据类型的表现&实现:

1. 预定义常量&类型: `#define`, `int` 等
2. 数据结构(存储结构): `typedef`
3. 基本操作算法(函数)
4. 赋值: 包括 三元赋值
5. 选择语句: if, switch
6. 循环: for, while, do-while
7. 结束: return, break, continue, exit(异常)
8. 输入输出: scanf, printf
9. 注释
10. 基本函数
11. 逻辑运算

```c
#include<stdio.h> // 头文件
#define PI 3.14

int main(){ // 主函数
    int arr[] = {1,2,3};

    // 字符串
    strlen(s);
    strcmp(s1,s2);
    strcpy(s1,s2);
    strcat(s1,s2);
    atoi("100");

    if(){}else if(){}else{}
    while(){} // 一定要在 while 改变执行条件
    do{}while(); // 一定会运行一次
    for(;;){}

    // 冒泡, 从小到大
    for(i=0;i<n-1;i++){
        for(j=i+1;j<n;j++){
            swap(a[i], a[j]);
        }
    }

    printf("hello world"); // 输出
    puts(str);
    return 0; // 主函数返回值, 这是为什么 linux 下面 error code 为 0 就是正常的
}
```

## c语言指针&内存

> 慕课网: <http://www.imooc.com/learn/394>

## 明解 C 语言 2017-10-1 21:52:26

> 入门篇: <http://www.ituring.com.cn/book/1671>
> 中级篇: <http://www.ituring.com.cn/book/1810>

```c
scanf()
printf()
puts()

concat(string &t, sgring s1, string s2); // 串联接, 设置 uncut 判断是否发生截断(顺序存储)
substr(string &t, int pos, int len); // 子串, 参数非法时 return ERROR;

// 头插法: 表尾到表头
L->next = null; (p->next = L->next; L->next = p;)

// 尾插法: 表头到表尾
L->next = null; *e = L; (e->next = p; e = p;) e->next = null;

// 串的模式匹配
// 简单算法
if(s[i] == t[j]){ // 匹配到则继续
    ++i; ++j;
}else{
    i=i-j+2; j=1;        // 指针后退重新开始
}

// KMP
if(j==0 || s[i] == t[j]){
    ++i; ++j;            // 匹配到则继续
}else{
    j=next[j];            // 模式串向右移动
}
// KMP 求 next[] 数组
if(j==0 || t[i]==t[j]){
    ++i; ++j;

    // 原版
    next[i] = j;
    // 改进版
    if(t[i] != t[j])    next[i] = j;
    else                 next[i] = next[j];
}else    j = next[j];
```
