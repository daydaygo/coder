# php| laravel 源码解读

yii/laravel 的源码质量超高, phper 一定要好好读一读

> 源码解读: http://naotu.baidu.com/file/9d4b5ab081c174ebfaf2af01db81bc5b?token=9fb122830b33d1c8

最近写的几篇 blog 都很长, 事无巨细, 基础差一点的同学看起来会 *舒服* 一些, 但是基础好一些的同学就会感觉不够深入, **干货** 不够. 而且写长文确实挺累的, 所以这篇尽量抓框架结构, 整体架构和思路, 细节地方的点, 以后再和大家慢慢分享啦.

## 生命周期 lifecycle

熟悉我的同学, 就知道我特别强调这个概念. 生命周期是连接具体业务与底层实现很好的桥梁, 熟悉生命周期, 既可以培养软件工程能力, 也可以培养对业务的理解能力.

> 重要的事情无限循环: 生命周期.

laravel 的生命周期:

### 初始化

- 使用 `$basePath` 初始化(`new`) `Illuminate\Foundation\Application`, Application 类其实一个容器, 用来解决类的依赖关系, 遵循 [PSR](http://www.php-fig.org/psr/) 标准
- Application 绑定(`bind`) 3 个单例: `App\Http\Kernel`, `App\Console\Kernel`, `App\Exceptions\Handler`, 请求由具体的 Kernel 来执行. 容器用来解决类的依赖关系是 **按需** 的, 这里都是申明的 **单例** 方式. 这里要注意 laravel 中亮眼的点: 绑定是按照 `Illuminate\Contracts\Xxx -> App\xxx` 的约束来实现, 全局统一. Contract 我们后面会继续讲到.

### 具体处理

- Application 根据场景(web / console), 实例化(`make`)化相应的 Kernel
- Kernel 实例化的过程, 先进入 bootstrap 阶段: 注册(`register`) ServiceProvider, 执行由 Application 进行, 注意 **注册** 是必须的 **关系申明**, 实际还是由容器来完成 **按需** 加载.
- 使用 php Global 变量里信息来初始化 Request 对象, 比如 `$_SERVER`
- Kernel 处理 Request 对象, 返回 Response 对象
- 最后 Kernel 执行 terminate() 收尾

### Kernel 处理 Request 流程

- 先执行定义的 Middleware
- 交给 Router 对象进行处理: 进行路由匹配, 匹配到 闭包 / Controller 进行逻辑处理, 返回数据
- 组装上面返回的数据, 并根据一些其他配置, 比如格式(format), 拼装成 Response 对象

大家会发现没有讲到自己熟悉的 MVC, 因为 MVC 也是有 Application 来进行管理一个个对象而已, 其他的所有服务也是.

> 有了 Container 这个基础, 再加上对生命周期的理解, 就可以根据业务按需填充代码进来了.

## 合约 Contract

生命周期和容器对于现代化框架来说, 基本是大同小异, 我们说说 laravel 不同的地方 -- Contract.

先介绍 2 个编程思想:

> 面向接口编程

> 约定大于配置

这 2 个思想其实是有重叠, 强调了在系统设计规范中, **定义** 与 **实现** 应该分离.

laravel 始终在践行这一思想, 永远 **定义** 先行, 框架提供的所有功能, 都能在这里找到 **源头**. 有人也喜欢说是 **影子**.

> 更通俗一点: 怎么实现我不管, 反正要按照我说的来.

没错, 现实其实就是这样的, 需要大家 **带着镣铐跳舞**.

在 laravel 中, 定义一个 功能/服务, 需要经历下面这一系列步骤:

- Contract 中定义接口
- Foundation 中定义基础实现, 大量的 Trait 和与这个服务相关的基本功能
- 具体的 功能/服务 实现, 需要按照 Contract 中的定义进行实现, 并且需要添加 ServiceProvider, 将服务注册到 Container 上
- Support 中提供一些通用, 与具体 业务/服务 无关的辅助方法, 供具体 功能/服务 实现时使用

这一套做完, 我们就可以按照 `XxxServiceProvider` 中 register() 方法中定义, 使用 `$app['xxx']` 这种形式来访问我们需要的服务.

比如:

```php
// AuthServiceProvider
protected function registerAuthenticator()
{
    $this->app->singleton('auth', function ($app) {
        // Once the authentication service has actually been requested by the developer
        // we will set a variable in the application indicating such. This helps us
        // know that we need to set any queued cookies in the after event later.
        $app['auth.loaded'] = true;

        return new AuthManager($app);
    });

    $this->app->singleton('auth.driver', function ($app) {
        return $app['auth']->guard();
    });
}

// 需要使用 Auth 服务
$auth = $app['auth']
```

**这里提示一点**, 通常都会设置一个类属性 `$app`, 赋值为框架初始化时的 Application, 这样所有的依赖关系, 都可以使用 `$app` 变量来统一处理, 既简洁又简单.

## 为什么 laravel 写起来这么舒服

这个标题其实是为了解释 laravel 的 slogan:

> Love beautiful code? We do too. The PHP Framework For Web Artisans

有了上面介绍的部分, 我们需要的核心功能都已经用了, 剩下的就看我们打算怎么用了. 其实已经介绍了一种方式: `$app['auth']`, 调用注册的 auth 服务.

### DI

DI, dependency injection, 依赖注入. 这种方式大家比较常见:

```php
// 框架提供的注册功能
/**
 * Handle a registration request for the application.
 *
 * @param  \Illuminate\Http\Request  $request
 * @return \Illuminate\Http\Response
 */
public function register(Request $request)
{
    $this->validator($request->all())->validate();

    event(new Registered($user = $this->create($request->all())));

    $this->guard()->login($user);

    return $this->registered($request, $user)
                    ?: redirect($this->redirectPath());
}
```

直接在函数参数中传递我们需要的类 `\Illuminate\Http\Request  $request`, 然后容器就会自动提供这个类给我们使用

### helper function

先说说广义上的 helper function, 通过 composer 的自动加载来实现:

```
// laravel/framework/src/composer.json
"autoload": {
    "files": [
        "src/Illuminate/Foundation/helpers.php",
        "src/Illuminate/Support/helpers.php"
    ],
    "psr-4": {
        "Illuminate\\": "src/Illuminate/"
    }
},
```

可以看到这里, laravel 注册了 2 类 helper function:

- `Foundation/helpers.php`: 框架功能有关的帮助函数
- `Support/helpers.php`: 框架功能无关的通用帮助函数, 比如数组的处理, 字符串的处理

这里就来讲讲框架功能有关的帮助函数:

```php
// 核心函数
function app($abstract = null, array $parameters = [])
{
    if (is_null($abstract)) {
        return Container::getInstance(); // 相当于上面的 $app
    }

    return Container::getInstance()->make($abstract, $parameters); // 相当于上面的 $app['auth']
}

// 使用效果
app() -> $app
app('auth') -> $app['auth']
```

看看这个有多么方便, 以 cache service 为例:

```php
// 实现
function cache()
{
    $arguments = func_get_args();

    if (empty($arguments)) {
        return app('cache');
    }

    if (is_string($arguments[0])) {
        return app('cache')->get($arguments[0], $arguments[1] ?? null);
    }

    if (! is_array($arguments[0])) {
        throw new Exception(
            'When setting a value in the cache, you must pass an array of key / value pairs.'
        );
    }

    if (! isset($arguments[1])) {
        throw new Exception(
            'You must specify an expiration time when setting a value in the cache.'
        );
    }

    return app('cache')->put(key($arguments[0]), reset($arguments[0]), $arguments[1]);
}

// 使用效果
cache($key); // 获取缓存
cache($key, $value); // 设置缓存
```

是不是 hin 方便?

### Facade

> Facade 设计模式: 通过外观包装, 隐藏具体细节对象, 降低应用程序的复杂度

先看效果, 同样用上面的 Cache 作为例子:

```php
\Cache::get($key);
\Cache::set($key, $value);
```

好像比 helper function 还是要复杂一点, 对比我们最后再做. 先看看 Facade 是怎么实现的.

先根据官方文档设置一个新的 Facede:

- 先到 `config/app.php` 中注册, ServiceProvider 的注册也在这里
- 新建一个 Facede 类, 继承自 `Illuminate\Support\Facades\Facade`

比如这里的 Cache Facade:

```php
namespace Illuminate\Support\Facades;

class Cache extends Facade
{
    protected static function getFacadeAccessor() // 只需要实现这个函数就好
    {
        return 'cache';
    }
}
```

那原理是啥呢? 继续看 `Illuminate\Support\Facades\Facade` 的代码:

```php
// 魔法函数
public static function __callStatic($method, $args)
{
    $instance = static::getFacadeRoot(); // 上面实现的 getFacadeAccessor() 就在这里使用到, 最后其实返回的 $app['cache']

    if (! $instance) {
        throw new RuntimeException('A facade root has not been set.');
    }

    return $instance->$method(...$args);
}
```

所以, 执行 `\Cache::get($key);` 最终执行了 `$app['cache']->get($key);`

分析到这里, 大家就会发现, 其实都是围绕着 `$app` 打转, **包装一下**, 实现了各种不同的书写方式.

### 3 种方式的对比

首先来看 DI 和 Facade, 不知道大家会不会和我一样有这样的疑问:

```php
// 书写控制器代码时, 有 2 种方式:

public function index(Request $request)
{
    $input = $request->all();
}

public function index()
{
    $input = Request::all();
}
```

当我们写 Request 时, 编辑器会提示 2 个出来 `\Illuminate\Http\Request` 和 `\Request`. 刚开始接触的时候, 这个地方疑惑了很久, 大家现在可以区分出来了吧, 一个是 DI, 一个是 Facade, 名字虽然 *一样*.

这个例子也可以看出 DI 和 Facade 的区别.

再来看看 Facade 和 helper function, 其实上面列举过一部分了, 至少上面的例子看起来 helper function 要简洁一点, 但是, 这个是有限的:

```php
Cache::get($key, $timeout); // 缓存可以设置过期时间, Cache Facade 直接加一下参数就行了

cache()->get($key, $timeout); // Cache helper function 没有封装, 所有要先用 cache() 来返回 $app['cache'], 然后再调用
```

这么多实现方式, 你是感觉 **有点晕**, 还是像官方说的 **艺术** ?

> 写法上有了更多的自由(简写), 提高生产力还是蛮明显的.

当然, 简写带来一个代码提示的问题, 会影响生产力的, 这个可以看我之前的 [blog - 聊一聊 php 代码提示](http://www.jianshu.com/p/b3daadb3c4c5).

## 最后再看一个套路

比如说 Cache 这个服务, 可以使用多种方式, 但是我们在使用的时候, 都是一样的方式, 这个是怎么实现的呢?

- 首先, 还是 Contract, **怎么实现我不管, 反正要按我说的来**, 这个是贯穿始终的
- 然后, 在有多种实现的方式时, laravel 惯用 `manager-driver` 这种方式

还是以 Cache service 为例:

```php
// CacheServiceProvider 中这样注册的服务
public function register()
{
    $this->app->singleton('cache', function ($app) { // 这就是我们常用的 $app['cache'], 实际使用的 CacheManager
        return new CacheManager($app);
    });

    $this->app->singleton('cache.store', function ($app) { // 继续看
        return $app['cache']->driver();
    });

    ...
}

// CacheManager
// 魔法函数, 比如执行 $app['cahe']->get($key), 其实最终执行的 $this->store()->get($key)
public function __call($method, $parameters)
{
    return $this->store()->$method(...$parameters);
}
public function store($name = null)
{
    $name = $name ?: $this->getDefaultDriver();

    return $this->stores[$name] = $this->get($name);
}
```

继续读代码, 会发现最终通过 `$name` 去寻找类中 `create{$ame}()` 方法来实例化 `$this->store()`

这种 `manager-driver` 的方式, 在 laravel 中大量使用, 比如 filesystem, Notification, Queue

## 写在最后

写着写着发现又是一个长文, 简直 **心累**, 明明是很简单的东西呀, 看着代码一点一点读就好了, 写起文字来却要 巴拉巴拉 这么多.

> 明明很简单的东西呀, 弄清楚生命周期和底层实现, 剩下的一点一点读代码就好了
