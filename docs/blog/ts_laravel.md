# TS| 技术分享 - 人生苦短, 我用 laravel

应聘 [慕课网](https://www.imooc.com/) 的讲师, 选题是 laravel 相关的, 仔细斟酌之下, 定了这个选题:

> 人生苦短, 我用 laravel

希望这次分享, 能帮助大家解答:

- 为什么要学习 laravel 框架
- 它有什么优势

## 为什么要学习 laravel 框架

要回答这个问题, 我们先要弄清楚, **框架** 是什么?

根据 [<编程风格：好代码的逻辑>](http://www.ituring.com.cn/book/1724) 这本书中的定义:

> 框架: 一类特殊库/可重用组件, 提供一个能被进一步开发的通用应用程序功能

我来给这句话划一下重点:

- 可重用 = 圈内的黑话 **轮子**
- 通用程序功能 = 前人栽树, 后人乘凉
- 能被进一步开发 = 这就是我们要做的, 通常我们叫它 **业务代码**

从这些定义可以看出, 评估一个框架的好坏有以下直观的标准:

- 提供丰富的可重用/通用功能, 避免让我们 **重复造轮子**
- 让我们写起 **业务代码** 来, 能更快更简单

首先来看功能丰富, 看过 [laravel 官方文档](https://laravel.com/docs/master) 的人, 都会感觉 **吃力** -- 这么多内容呀. 这里我要纠正大家一个观点:

> 多是因为 **功能多**, 核心其实就那么一点点, 按需使用自己需要的功能就好了.

laravel 就是这样一个 **大而全** 的框架, 基本不需要 **重复造轮子**.

再来谈谈 laravel 的怎么让我们写起 **业务代码** 来更快更简单.

> Love beautiful code? We do too. -- The PHP Framework For Web Artisans

laravel 的宣传标语还是蛮 **傲娇** 的 -- 为 Web 艺术家而生. 用几个同义的关键字:

- simple
- clean
- easy

这样的代码写起来, 当然快啦.

我们来实操感受下, 先来看 **路由**, 文件地址 `routes/web.php`:

```php
// 默认路由
Route::get('/', function() {
    return view('welcome');
});

// 开始玩起来
Route::get('/test', function() { // 自定义路由
    return "czl"; // 直接返回字符串
    return view('welcome'); // 返回 html 页面
    return ["a" => "czl", "b" => "daydaygo"]; // 直接写数组就可以返回 json 格式
});
```

这还只是 **路由** 功能的冰山一角, 可以参考 [官方文档 - routing](https://laravel.com/docs/master/routing) 解锁更多姿势.

我们接着来看一下 Web 开发中经常会使用到的缓存功能, 我们还是在路由这里进行测试:

```php
use Illuminate\Support\Facades\Cache;

Route::get('/', function() {
    // 使用 Facade
    Cache::set('czl', 'daydaygo'); // 设置缓存
    $cache = Cache::get('czl'); // 获取缓存

    // 常用操作, 还有 helper function 可以使用, 更快捷
    cache('czl', 'daydaygo'); // 设置缓存
    cache('czl'); // 获取缓存
})
```

是不是 hin 轻松?

再来体验一下 laravel 的快, 使用 `php artisan` 轻松生成代码, 开发效率倍增

![php artisan make - 代码生成](http://upload-images.jianshu.io/upload_images/567399-7de5d1f7aba73251.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 它有什么优势

其实, 应该说 **还有** 什么优势:

- *难*
- 社区活跃

### 难

是不是感觉很奇怪, *难* 怎么变成优势了?

先说一下 laravel 难在哪里 -- 框架的架构(architecture):

- 服务容器 service container
- 服务提供者 service provider
- 门面 Facade(设计模式的一种)
- 合约 contract

尤其是第一次接触 laravel 框架的同学, 看到这些 **概念**, 简直头大. 但是, 得益于 **世界上最好的语言**, 你只要稍微深入一点, 就会发现这个其实并不难, 并且当你跨过这座 **高山**,  就会有 **一览众山小** 的感觉:

> 因为 laravel 的架构, 算是 php 框架中最复杂的.

这也是 laravel NB 的地方, 我们在写 **业务代码** 的时候, 并不会感知到底层的复杂.

### 社区活跃

这么好的框架, 当然吸引了不少 **有识之士** 啦. laravel 的社区非常活跃, 相关的站点也非常多, 这样也带来了明显的好处: 遇到了问题, 基本百度都能找到 **前车之鉴**, 不太容易被一些技术细节及难点卡住, 能获得十分流畅的学习的体验.

资源很充足:

- [中文社区: laravel-China](https://laravel-china.org/)
- [中文文档](http://d.laravel-china.org/)
- [中文书籍/教程](https://fsdhub.com/books/laravel-essential-training-5.5)
- 还有 [Laravel Dineer](https://laravel-china.org/topics/3448/laravel-dineer-03-shanghai-railway-station-rush-together-looking-forward-to-the-next-party) 线下面基

## 写在最后

不知道介绍到这里, 有没有激起大家想要 **撸一把** laravel 的冲动呢. 还是那句话:

> laravel, 为 Web 艺术家而生
