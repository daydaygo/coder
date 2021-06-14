# 谈一谈编程语言 2


> 也许不会用到多种编程语言, 但并不妨碍你去了解它, 因为真的很有趣
> 推荐 Yawenina 小姐姐的 ES6 教程: https://www.codecasts.com/series/es6-from-scratch

## 变量作用域

这期先讲简单的, 指针和引用先放放.;

ES6 现在有 3 种变量定义了, `var / let / const`, 还是先来看代码

```js
for(let i=1; i<10; i++) { // var
    console.log(i);
    setTimeout(function() {
        console.log(`i:${i}`)
    });
}
console.log(`i:${i}`)
```

试着把上面的 let, 改成 var 或者 const, 就能看到区别了. 其实就是变量的作用域的问题:

- global scope: 全局作用域, 当前文件全局有效
- function scope: 函数作用域, 当前函数中有效, 这里对应 var
- block scope: 块级作用域, 当前代码块(比如 for 循环)中有效, 这里对应 let / const
- const 更进了一步, 不允许修改

这里补充说明一下 `global scope` 和 `function scope`, 每个函数都有自己的独立内存空间, 其实可以把 `global scope` 所在的地方, 看成是一个「大函数」, 2 个函数之间, 自然是隔离的.

看到这里了, 我们再看看 php 和 c 的一个例子:

```php
// php
for ($i=0; $i < 10; $i++) {
    echo $i;
}
echo $i;
```

```c
// c
#include <stdio.h>

int main()
{
    for (int i = 0; i < 10; ++i)
    {
        printf("%d\n", i);
    }
    printf("%d\n", i);
    return 0;
}
```

运行试试看, 可以理解吧.


## staic self

面向对象编程, 就少不了那些关键字: `parent self static const`, 讲讲比较难区分的 2 个 static 和 self

```php
class BaseModel
{
    // 获取 model 对应的数据库 table 名
    public static function getTableName()
    {
        $class = new \ReflectionClass(static::class); // 使用反射获取类名
        return strtolower(preg_replace('/((?<=[a-z])(?=[A-Z]))/', '_', $class->getShortName())); // snake_case
    }
}

class GameSession extends BaseModel
{

}
```

注意, 这里要使用 `static`, 如果你使用 `self` 得到的就是 `BaseModel` 了. 至于一个简单的理解 static & self 的方式: static 是指当前内存中运行的实例, 所以永远都是 **所见即所得**.

## 同步 异步

首先声明一下, 这里的同步和异步, 指的是「同步代码」和「异步代码」, 服务器领域经常会出现这 2 个词, 可以看我的「服务器编程系列」blog.

先来看我们最常见的同步写法(php):

```php
echo 'foo';
sleep(3);
echo 'bar';
```

然后再看看 js:

```js
console.log('foo');
ajax.post(url, data); // long time job, like ajax here
window.location.href('http://www.baidu.com');
```

然后, 稍微 js 经验多一点的程序员就会告诉你, 这样做是错了, 任务耗时太长, 可能会 **先执行页面跳转**

正确的写法应该是这样:

```js
ajax.post(url, data, function() {
    window.location.href('http://www.baidu.com');
});
```

继续, 我们再来对比一下 nodejs:

```js
// 同步
var fs = require("fs");

var data = fs.readFileSync('input.txt'); // 同步阻塞

console.log(data.toString());
console.log("程序执行结束!");

// 异步
var fs = require("fs");

fs.readFile('input.txt', function (err, data) { // 异步非阻塞
    if (err) return console.error(err);
    console.log(data.toString());
});

console.log("程序执行结束!");
```

一图胜千言 - Node.js 事件循环:

![](http://www.runoob.com/wp-content/uploads/2015/09/event_loop.jpg)

强烈建议一直写同步代码的小伙伴, 可以看一下 [nodejs 菜鸟教程](http://www.runoob.com/nodejs/nodejs-tutorial.html), 可能也会和我一样, 感受到明显的冲击
