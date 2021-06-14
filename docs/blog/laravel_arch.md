# php| laravel 架构一览

第一次写貌似有点深度的 **blog** , 有点想到哪写到哪的感觉, 勿怪.

## 一次请求的生命周期

### 请求 -> 服务器

环境: `nginx + php-fpm + laravel`
简单讲, 我们在浏览器输入了 url，浏览器发起了一次 http 请求
首先要知道一点 **http** 的基础知识:

- 构成

1. 协议: http / https, 当然 https 是大势所趋, 未来的 http2( **应用层** )
2. ip, 或者 **域名**, 这样才能在网络上找到你的服务器 ( ip + 端口 就是 **网络层**)
3. 端口, http 默认 80, https 默认 443, 也可以使用不同端口，在 nginx 中配置即可
4. path(路径) 或者 route(路由), 在nginx这里其实是 path，通过 server 配置，根据域名定位到不同的项目路径，然后将请求转发给 fpm，框架层都统一到入口脚本`index.php` 中，之后的 path 被框架层解析为 route
5. http method（get、post ）、restful、get params、post body data

假如我是浏览器中输入url: 首先，浏览器会帮我补上http协议，然后根据 ip（dns会帮我把域名转成ip） 找到服务器, 再根据端口找到服务器上面的服务(这里是 nginx 的 server 配置), 然后根据 nginx 的配置到项目的根目录, 然后 nginx 把请求交给 fpm，fpm 其实是执行框架的入口脚本`index.php` ，然后根据 path 去找 route。

一个 nginx + php-fpm 的实例:

```conf
server {
        listen       80; # 端口
        server_name  laravel.dev laravel.dev; # 域名
        root   /data/web/laravel/public; # 项目根目录
        index  index.html index.htm index.php; # 默认执行文件
        autoindex  on; # 开启目录浏览功能
        location / { # 所有文件的执行规则
            try_files $uri $uri/ /index.php?$query_string; # 转换url, 不显示 index.php
        }
        location ~ \.php(.*)$ { # php 后缀文件的执行规则, 其实就是交给 php-fpm 来执行
            fastcgi_pass   127.0.0.1:9000; # 也可以使用 unix socket
            fastcgi_index  index.php;
            fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
            include        fastcgi_params;
        }
}
```

这里有一点要注意: nginx 是先 **执行文件** , 比如你有一个 url为 `http://xxx.com/test`, 根目录下刚好有一个名为 `test` 的文件, 那么nginx就直接去下载这个文件了, 而不会到你项目中 test 对应的 `controller/action`

### 请求 -> laravel 框架

通过上面可以知道 **请求** 被 `public/index.php` 文件来执行.
`index.php` 这个文件干了 3 件事[^1]:
1. 实现 composer 的 类的自动加载
2. 载入 laravel app, 核心是一个 service container, 这个后面慢慢分析
3. return response, 返回 http response

- 关于类的自动加载, 可以参考下面2个教程:

1. 站在巨人的肩膀上写代码—SPL: http://www.imooc.com/learn/150, 里面详细的讲了类的自动加载是如何 **发展&实现** 的
2.  PHP实现手机归属地查询: http://www.imooc.com/learn/604, 类的自动加载的 **简单** 实现

- laravel 处理请求的简单流程

1. 创建 laravel app / service container 实例
2. 请求交给 `app/Http/Kernel.php` 处理
3. `app/Http/Kernel.php` 扩展自 `Illuminate\Foundation\Http\Kernel`, 里面定义了 `bootstrappers`, 在请求执行前运行; 还定义了 ` middleware`, 类似 **管道** 的作用.
4. 请求经过 `app/Http/Route.php` 文件进行匹配, 找到相应的执行方法, 执行后获取 response

## service container
service container 用来关联 service provider, 用来为应用提供各种服务, 比如: db/队列/缓存 等等, 创建应用时, 会执行所有 service provider 的 `register` 方法, 用来绑定 service provider 到 service container 中, 当所有 service provider 都 **注册** 后, 执行 `boot` 方法.

- laravel5 应用结构

```
app/                    你写的主要代码都在这里
    console/            php artisan 脚本
        commands/       php artisan 脚本
        Kernel.php      cli 内核, 添加的 php artisan 指令需要在这里注册
    events/
    exceptions/
    Http/               http request
        Controllers/    控制器
        Middleware/     中间件
        Requests/       自定义请求, 可以将表单验证独立出来到这里
        Kernel.php      http 内核, middleware 需要在这里注册
    Jobs/               可以从 controller 中抽取一部分代码到这里
    Listeners/
    Policies/           达到颗粒化权限认证的效果
    Providers/
    User.php            自带 auth 认证使用的 ORM 模型
bootstrap/
    cache/              启动优化的缓存文件
config/                 配置文件, 注意这里的文件是被直接 require 的, 所以只使用配置数组, 不要使用语句
database/               数据库相关文件, 也可以将 sqlite 搭建在此目录下
    factories/          数据工厂, 用来 db:seed 使用
    migrations/         数据库迁移文件
    seeds/              数据生成, 比如生成测试数据, 创建 admin 账号
public/                 asset(js/css/font)
resources/              raw resource(原生资源)
    assets/             LESS, SASS, CoffeeScript 等
    lang/               多语音配置文件
    views/              视图文件
storage/                文件存储目录
    framework/          包括使用 file 作为驱动的 session / cookie, blade模板编译后的文件
    logs/               日志文件
tests/
vendor/                 composer 资源, 包括 laravel 框架源码
index.php               入口脚本
```
- service provider

为框架提供各种各样的服务, 比如 db / cache / 文件存储, service provider 简单而言就实现了2件事: 绑定服务到服务容器; 按需加载

- service container

laravel的核心, 本质上是一个高级的 Ioc 容器, 能轻松的实现 **依赖注入** 和 **控制反转**, 这样就可以轻松的解决 类&类 之间的依赖的关系, 写出更 **优雅** 的代码.

## why

其实整个请求进入laravel的过程, 框架实例化一个 `app` 类, 然后这个 app 类的调用各种服务 request / response / db / cache 来完成各种事务.

为什么要采取这样一种方式呢?

1. 扩展性: 需要什么功能, 添加服务即可.
2. 面向接口编程: 订立契约精神, `app` 不用管服务的具体实现, 只要自己需要的服务器能被提供就行了, 不同的实现方式可以通过实现接口来实现.

推荐一篇讲解 laravel 服务容器的干货: http://laravelacademy.org/post/769.html [^2]

[^1]: 这里只讨论 http 请求, 不讨论 cli (控制台)

[^2]: laravel 学院
