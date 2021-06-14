# yii| 10 分钟开发 blog

因为最近可能工作需要, 又要开始使用 yii 框架. 想想都快 2 年没碰了, 最近一次接触到, 还是整理自己的 [wiki](http://wiki.dayday.tech), 关于 yii 的内容实在是既乱且杂. 整理的过程中发现自己当时做的笔记实在是 **不得要领**.

> 所以写这篇 blog 出来, 结合自己开发 php 框架的经验, 分享一下自己对 php 框架的见解, 希望能帮助到大家.

比较流行也比较重的几个框架(之后的讨论也是基于这几个框架): laravel5 thinkphp5 yii2. 大家在上手这些的框架的时候, 普遍都会感觉有一定的难度. 产生难度的原因, 大致有下面几个:

- 框架涵盖和解决的问题很多, 需要花一段时间消化
- 只关注到框架解决问题的快捷, 没有花足够的时间来理解框架的核心
- 对 MVC 很熟, 却忽视了从框架(或者说请求)的生命周期来进行更全面的理解

## 我们为什么需要框架

> 天下武功, 唯快不破

没错, 框架存在的理由, 就是加快开发速度, 增加产出, 把常见的问题给封装掉, 留出更多时间花在 **业务** 上.

这里简单梳理一下常见问题:

- 配置管理
- 日志管理
- 错误与异常处理
- 路由 + MVC
- 数据库服务, 包括关系型数据库, 非关系型数据库
- 缓存
- web应用与控制台应用
- 安全
- 测试

如果你的技术栈里面连这些基础的问题都没有思考过, 建议你把上面三个框架中的任何一个的文档好好看一遍.

简单过一下这几个问题, 希望能起到抛砖引玉的作用:

1. 配置管理:

- 有些内容直接写死到代码里面并不好, 这部分内容多以配置的形式存在
- 推荐 **面向约定** 编程, 而不是 **面向配置** 编程, 比如我们约定类和命名空间的首字母大写, 而不是依赖配置来解决, 尽管不这样写语法上面也没有问题
- 配置的解析: 最简单的是直接使用 php 数组, 遇到使用 `.ini/.yaml` 文件的也不要惊讶, 最终还是转化为数组
- 配置的使用: 直接数组式的访问, 也有 dotenv 形式的访问, 比如 `$arr['a']['b']['c']` 的 dotenv 访问类似于 `$arr[a.b.c]`
- 不同环境的配置文件管理: 正常项目至少有 2 套环境, 开发环境 + 发布环境, 不同环境下使用的配置文件不同, 为了防止配置文件提交到版本库中对其他环境(其他开发)产生影响, 会设置一些提交策略, 比如 `env.php.dev` 作为开发环境的配置文件提交, 其他开发获取代码后复制为 `env.php` 来使用

2. 日志管理:

- 日志等级: 要了解标准的日志等级, 实际使用中也至少要使用到 2 个级别 `error + info`, 前者需要马上处理, 后者方便后期定位问题
- 日志格式: 方便后期处理与查看
- 日志存储: 目前大部分直接存到项目的 `log` 目录, 单文件? 是否需要清理? 日志归集?
- 怎么打日志, 该打哪些内容到日志里面, 需要经验的积累, 建议新手在可能出现不同结果的代码处, 尝试加一下日志

3. 错误与异常处理

- 错误处理的三种方式: C语言式的 `return code`; 设置错误变量, 比如 `json_last_error()`; 异常 `Exception`
- 现在 php7 已经统一内部的错误处理机制到 **抛异常**
- 错误与异常处理往往都需要配合日志一起使用
- 最外层嵌套一层 `try-catch` 来捕获异常的方式, 现在也很流行

4. 路由 + MVC:

- MVC, 这个还用多说么, 不知道的赶紧恶补
- 路由的好处也很明显, 抽离出了一部分 http 协议的内容处理掉了, 比如解析 url 信息进行资源定位, http method
- router 和 controller 之间, 现在流行增加一层, laravel 中称之为 middleware(中间件), yii 称之为 filter(过滤器), 这个后面还会说到

5. 数据库服务(mysql)

- mysql 数据库相关的扩展, mysqli 和 pdo, 如果不清楚二者的区别, 甚至不知道扩展(ext), 恶补去吧
- 框架层做的基础部分: 数据库单例(至少要能手写单例模式)与 mysql 连接(连接 mysql 需要 mysql 协议);
- 框架层的高级部分: 执行原生 sql; `db query build`, 查询构造器, 一般采用链式调用; `Active Record`, 用 model 绑定数据中的表

6. 数据库服务(nosql)

- nosql 现在 mongo 还是用的多, 现在的框架想要添加一个外部服务简直不要太简单(这个后面会讲到)

7. 缓存

- memcache 和 redis 之争就不赘述了, 不过如果有人用这个问题来考你, 建议你先看看对方并发量再考虑是否要回答这个 **深入点** 的问题
- redis 被称之为 **程序员的瑞士军刀**, 也很适合用来多了解一下 **数据结构**
- TTL(time to live) / expiry, 缓存过期更新的问题
- 缓存是显著提高系统性能的常用手段, 建议仔细阅读框架这部分的文档, 比如你可以在 yii 中看到 **http 缓存**

8. web应用与控制台应用

- 这两个概念不清楚的话, 恶补吧
- 控制台应用与 crontab, 不知道就百度 `crontab`, 有兴趣也可以了解一下命令行应用的 argument 和 option, 了解一下 linux shell 的基本原理
- 如果区分不出来 `php-cli 和 php-cgi/php-fpm`, 恶补吧

9. 安全

- 不用你刻意去学, 框架设计者在写框架之前就需要考虑到, 所以, 去阅读框架的文档吧
- 密码等信息不能明文存储, 顺便八一八 `password_hash()` 这个函数, 可以折腾一下下面的代码

```php
public static function hash($str)
{
    $hash = password_hash($str, PASSWORD_DEFAULT);
    return str_replace('$2y$', '$2a$', $hash); // compatible with bcrypt of nodejs
}
```

10. 测试

- TDD, 测试驱动开发, 写代码要有 **测试先行** 的理念, 推荐 [<修改软件的艺术>](http://www.ituring.com.cn/book/1749), 可以培养一下这方面的意识
- phpunit 可以用起来了, 推荐 `unit test` 进行函数/方法级别的测试, `feature test` 进行 api 级别的测试
- 提到 api, 顺便推荐一下 **swagger UI**, 简单够用的 api 文档工具

## 框架之心

> 框架的核心是什么: 服务容器

这个词如果大家看完 laravel 的文档, 就会有很深的印象了, 这些框架的部分难点, 也在于理解这个概念上. 我尝试解释一下这个, 希望对大家有帮助.

- 正式些的解释: 服务容器通过依赖注入(DI), 实现控制反转(IoC).
- 说人话: 不用使用 `new` 了

直接看代码:

```php
// 没有服务容器
class Superman
{
    protected $power;

    public function __construct()
    {
        $this->power = new Power(999, 100);
    }
}

// 有服务容器
class Superman
{
    protected $power;

    public function __construct(Power $power)
    {
        $this->power = $power;
    }
}
```

但是, 需要 `new` 一个对象这件事不可能凭空消失吧, 没错, 服务容器就是帮我们干了这件事. 如果大家对对象的理解更深一点的话, 服务也可以抽象成对象来实现, 所以, 上面那句:

> 框架需要添加一个外部服务简直不要太简单

好了, 划了重点了, 下去记吧.

还是来看看代码实现, 来自我之前 [blog - 聊一聊 php 代码提示](http://www.jianshu.com/p/b3daadb3c4c5), 这里使用的  [hyperframework](http://hyperframework.com) 框架, 框架核心也是服务容器:

```php
// lib/Services/Redis.php 文件
<?php
namespace App\Services;

use Hyperframework\Common\Config;
use Hyperframework\Common\Registry;

class Redis
{
    /**
     * 将 redis 注册到 Hyperframework 的容器中
     */
    public static function getEngine()
    {
        return Registry::get('services.redis', function () {
            $redis = new \Redis();
            $redis->connect(
                Config::getString('redis.host'),
                Config::getString('redis.port'),
                Config::getString('redis.expire')
            );
            $redisPwd = Config::getString('redis.pwd');
            if ($redisPwd !== null) {
                $redis->auth($redisPwd);
            }
            return $redis;
        });
    }
...
}
```

建议阅读 [laravel官方文档 - container](https://laravel.com/docs/master/container) 加深理解.

强烈推荐 [学院君 - 深入理解控制反转（IoC）和依赖注入（DI）](http://laravelacademy.org/post/769.html), 写得太好了, 为学院君疯狂打 call.

## 框架(请求)的生命周期

> 我们需要故事

如果看过 <人类简史> 这本书, 人类的第一次革命是 **认知革命**, 简单说就是 **讲故事**. [<用数据讲故事>](http://www.ituring.com.cn/book/1763) 这本书也提到, 故事是最容易也最快让人接受的方式.

关于框架生命周期的故事:

- 起点是入口脚本 `index.php`, 这里会 new 一个 app 对象, 我们的框架核心, 也就是服务容器
- 服务容器调用 router 模块, 用来解析 url 里的路由信息, 用来找到对应 controller 中的 action
- 在 router 和 controller 之间, 可以使用 middleware(或者 filter), 来分离一些业务, 比如维护模式(所有请求直接跳转到 503 页面), 比如开启 session
- 好了, 我们到了 controller 里啦, controller 会有 2 个属性, Request 和 Response, 分别用来对应 http 的 request 和 response
- 如果业务涉及到数据(比如数据库), controller 通常找 model 来处理
- 当然, 展示给用户该怎么处理, controller 也不用关心, 它只需要提供数据给 view 就行了
- 如果还需要其他服务, 比如 redis, 服务容器很轻松就能解决

这里忽略了很多细节, 大家在看文档或者阅读的源码的过程中, 可以慢慢填充. 这里留一个常见的面试题, 大家可以试试讲故事的方法:

> 说一下从浏览器输入 url 到看页面发生了什么事, 越详细越好.

我的面试要求是至少要答出来 tcp/ip 4 层网络模型, dns, nginx + fpm 工作模式.

## 10 分钟?

> stay hungry stay foolish

回到标题, 大家平时可能也看到这样的标题, 比如 <使用 laravel 10 分钟搭建博客系统>, <21 天精通 xxx>, 也可能听过或者自己实践过得出 **xxx都是骗人的** 这样的结论. 首先, 我明确一下, 我是明确支持 **用 yii 10分钟搭建博客系统**, 下面是具体操作步骤:

- 感谢魏曦大大的 [视频教程](http://weixistyle.com/yii2.php), 视频质量高到超乎我的想象, 同样, 为魏曦大大疯狂打 call, 这次博客系统使用教程提供的源码
- 使用 coding 的动态 page 服务, 按照页面提示配置即可, 期间最耗时的任务是用 phpadmin 使用 sql 文件导入数据

没错, 到这里就结束了, 整个过程 10 分钟绰绰有余.

我是这样来看待的, 或者说这也算我对技术的一大认知:

哪怕只是配置一下域名到服务器, 新手从申请域名, 到配置时理解 A 记录和 CNAME 记录, 都不止 10 分钟, 但这只是我们上一步配置项里面的一个可选项而已. 技术其实就是由这样以及那样, 许许多多的知识点组成的, 了解的更多(基础打扎实), 越会讲故事(构建出自己技术栈和技能图谱), 最终就是能实现这样一个目标, 准确说, **小目标**.

## 一点个人见解

为什么 view 要这么重?

看 yii 框架, 很容易会得出这样一个感受, 特别是你发现 view 界面居然可以只有 php 代码. 为此, yii 框架里面有很多组件库, 帮助达到这样一个效果.

不只是 yii 框架, 其他框架也在这方面做了很多努力, 比如 laravel 的 blade 模板引擎和全套的前端编译构建环境.

我不反对做一个全栈工程师, 但是我的选择是优先做一个更好的 php 工程师. view 层的加重, 更多的是在分散注意力.

> 前后端分离: 也许我们需要的是一个前端, 最好还是一个妹子.

另外, 我也有个倡议, 不要再叫 **老程序员** 了.

> 虽然我们普遍看起来偏老, 但是用**长者**或者**大大**是否更好一点?
