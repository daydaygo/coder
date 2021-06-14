# 聊一聊 php 代码提示

这次我们来聊一聊 php 的代码提示, 不使用 IDE 的同学也可以瞧瞧看, PHP IDE 推荐 phpstorm.

> phpstorm 使用代码提示非常简单, 只需要将代码提示文件放到项目中就好, 我目前放到 `vendor/` 目录下

## 起源

1. 最近开发的项目中, 有使用到 `PHP 魔术方法` 和 `单例模式`, 导致了需要代码提示的问题
2. 最近在尝试用 swoole 写 tcp server, 有需要用到 `swoole IDE helper`, [swoole wiki](https://wiki.swoole.com/)首页就有推荐

## 数据库模型

在 laravel 中, 如果有一张数据表 lessons 如下:

```
CREATE TABLE `lessons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `intro` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `published_at` timestamp NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

那么可以建立一个 Lesson 模型和他对应:

```
<?php
namespace App;

use Illuminate\Database\Eloquent\Model;

class Lesson extends Model
{
    //
}
```

之后, 我们就可以直接使用下面的方法了:

```php
$lesson = new Lesson();
$lesson->title = 'far'; // set

$lesson = Lession::find(1);
echo $lesson->title; // get
```

这样写是不是很舒服, 或者说很「优雅」?

而实现起来, 非常简单, `__get() / __set()` 就可以了:

```php
// laravel 文件: Illuminate\Database\Eloquent\Model

/**
 * Dynamically retrieve attributes on the model.
 *
 * @param  string  $key
 * @return mixed
 */
public function __get($key)
{
    return $this->getAttribute($key);
}

/**
 * Dynamically set attributes on the model.
 *
 * @param  string  $key
 * @param  mixed  $value
 * @return void
 */
public function __set($key, $value)
{
    $this->setAttribute($key, $value);
}
```

在 laravel 中, 这样的实现方式随处可见, 比如:

```php
// Illuminate\Http\Request $request

$request->abc; // 获取 http 请求中的 abc 参数, 无论是 get 请求还是 post 请求
$request->get('abc'); // 和上面等效
```

好了, 原理清楚了. 写起来确实「舒服」了, 但是, 代码提示呢? 难道要自己去数据库查看字段么?

在我们的另一个使用 [hyperframework](http://hyperframework.com)框架的项目中, 我们使用了 **代码自动生成的方法**:

```
// lib/Table/Trade.php 文件
<?php
namespace App\Table;

class Trade
{

    public function getId() {
        return $this->getColumn('id');
    }

    public function setId($id) {
        $this->setColumn('id', $id);
    }
    ...
}

// lib/Model/Trade.php 文件
<?php
namespace App\Model;
use App\Table\Trade as Base

class Trade extends BaseTable
{
    ...
}

// 这样我们就可以愉快的使用下面代码了
$trade = new Trade();
$trade->setId(1); // set

$trade = Trade::find(1);
$trade->getId(); // get
```

上面的 `lib/Table/Trade.php` 文件使用一个 php 脚本, 读取 mysql 中 `information_schema.COLUMNS` 的记录, 然后处理字符串生成的. 但是, 缺点也非常明显:

- 多了一个脚本需要维护
- 字段修改了, 需要重新生成
- 代码结构中, 多了一层 `Table` 层, 而这层其实就只干了 `get / set`

虽然有了代码提示了, 这样做真的好么? 那好, 我们来按照上面的套路改造一下:

```
// lib/Models/BaseModel.php
<?php
namespace App\Models;

use Hyperframework\Db\DbActiveRecord;


class BaseModel extends DbActiveRecord
{
    // 获取 model 对应的数据库 table 名
    public static function getTableName()
    {
        // 反射, 这个后面会讲到
        $class = new \ReflectionClass(static::class);
        return strtolower(preg_replace('/((?<=[a-z])(?=[A-Z]))/', '_', $class->getShortName()));
    }

    public function __get($key) {
        return $this->getColumn($key);
    }

    public function __set($key, $value) {
        $this->setColumn($key, $value);
    }
}

// lib/Models/User.php
<?php
namespace App\Models;

class User extends BaseModel
{
    ...
}
```

好了, 问题又来了, 代码提示怎么办? 这样常见的问题, 当然有成熟的解决方案:

> [laravel-ide-helper](https://github.com/barryvdh/laravel-ide-helper): laravel package, 用来生成 ide helper

上面 Lesson model 的问题, 就可以这样解决了, 只要执行 `php artisan ide-helper:models`, 就会帮我们生成这样的文件:

```php
<?php
namespace App{
/**
 * App\Lesson
 *
 * @property int $id
 * @property string $title
 * @property string $intro
 * @property string $image
 * @property string $published_at
 * @property \Carbon\Carbon|null $created_at
 * @property \Carbon\Carbon|null $updated_at
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson whereCreatedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson whereId($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson whereImage($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson whereIntro($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson wherePublishedAt($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson whereTitle($value)
 * @method static \Illuminate\Database\Eloquent\Builder|\App\Lesson whereUpdatedAt($value)
 * @mixin \Eloquent
 */
    class Lesson extends \Eloquent {}
}
```

通过注释, 我们的代码提示, 又回来了!

## Facade 设计模式 / 单例设计模式

了解 laravel 的话, 对 Facede 一定不陌生, 不熟悉的同学, 可以通过这篇博客 [设计模式（九）外观模式Facade（结构型）](http://blog.csdn.net/hguisu/article/details/7533759) 了解一下.

现在来看看, 如果我们需要使用 redis, 在 laravel 中, 我们可以这样写:

```php
Redis::get('foo');
Redis::set('foo', 'bar');
```

底层依旧是通过 ext-redis 扩展来实现, 而实际上, 我们使用 ext-redis, 需要这样写:

```php
$cache = new \Redis();
$cache->connect('127.0.0.1', '6379');
$cache->auth('woshimima');

$redis->get('foo');
$redis->set('foo', 'bar');
```

2 个明显的区别: 1. new 不见了(有时候会不会感觉 new 很烦人); 2. 一个是静态方法, 一个是普通方法

如果稍微了解一点设计模式, **单例模式** 肯定听过了, 因为使用场景实在是太普遍了, 比如 db 连接, 而且实现也非常简单:

```php
// 简单实现
class User {
    private static $_instance = null; // 静态变量保存全局实例

    // 私有构造函数，防止外界实例化对象
    private function __construct() {}

    // 私有克隆函数，防止外办克隆对象
    private function __clone() {}

    //静态方法，单例统一访问入口
    public static function getInstance() {
        if (is_null ( self::$_instance ) || isset ( self::$_instance )) {
            self::$_instance = new self ();
        }
        return self::$_instance;
    }
}

// 使用
$user = User::getInstance();
```

好了, 关于 new 的问题解决了. 接下来再看看静态方法. 在我们的另一个使用 [hyperframework](http://hyperframework.com)框架的项目中, 我们也实现了自己的 Redis service 类:

```
// lib/Services/Redis.php 文件
<?php
namespace App\Services;

use Hyperframework\Common\Config;
use Hyperframework\Common\Registry;

class Redis
{
    /**
     * 将 redis 注册到 Hyperframework 的容器中
     * 容器这个概念先留个坑, 下次讲 laravel 核心的时候, 再一起好好讲讲
     * 这里只要简单理解我们已经实现了 redis 的单例模式就好了
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

    // 重点来了
    public static function __callStatic($name, $arguments)
    {
        return static::getEngine()->$name(...$arguments);
    }

    // k-v
    public static function get($key)
    {
        return static::getEngine()->get($key);
    }
}
```

拍黑板划重点: `__callStatic()`, 就是这个魔术方法了. 另外再看看 `...$arguments`, 知识点!

仔细看的话, 我们下面按照 ext-redis 中的方法, 再次实现了一次 `$redis->get()` 方法, 有 2 点理由:

- 魔术方法会有一定性能损失
- 我们又有代码提示可以用了, 只是要用啥, 就要自己把 ext-redis 里的方法封装一次

好了, 来看看我们的老朋友, laravel 是怎么实现的吧:

- laravel: Illuminate\Support\Facades\Facade

```php
// 获取 service 的单例
protected static function resolveFacadeInstance($name)
{
    if (is_object($name)) {
        return $name;
    }

    if (isset(static::$resolvedInstance[$name])) {
        return static::$resolvedInstance[$name];
    }

    return static::$resolvedInstance[$name] = static::$app[$name];
}

// 魔术方法实现静态函数调用
public static function __callStatic($method, $args)
{
    $instance = static::getFacadeRoot();

    if (! $instance) {
        throw new RuntimeException('A facade root has not been set.');
    }

    return $instance->$method(...$args);
}
```

然后, 使用上面的 package, 执行 `php artisan ide-helper:generate`, 就可以得到代码提示了:

```php
namespace Illuminate\Support\Facades {
    ...

    class Redirect {
        /**
         * Create a new redirect response to the "home" route.
         *
         * @param int $status
         * @return \Illuminate\Http\RedirectResponse
         * @static
         */
        public static function home($status = 302)
        {
            return \Illuminate\Routing\Redirector::home($status);
        }

        /**
         * Create a new redirect response to the previous location.
         *
         * @param int $status
         * @param array $headers
         * @param mixed $fallback
         * @return \Illuminate\Http\RedirectResponse
         * @static
         */
        public static function back($status = 302, $headers = array(), $fallback = false)
        {
            return \Illuminate\Routing\Redirector::back($status, $headers, $fallback);
        }
        ...
    }
    ...
}
```

## 通过反射实现 swoole 代码提示

通过反射实现 swoole 代码提示来自此项目 [flyhope/php-reflection-code](https://github.com/flyhope/php-reflection-code), 核心代码其实很简单, 如下

```php
static public function showDoc($class_name) {

    try {
        // 初始化反射实例
        $reflection = new ReflectionClass($class_name);
    } catch(ReflectionException $e) {
        return false;
    }

    // 之后都是字符串处理之类的工作了

    // Class 定义
    $doc_title = ucfirst($class_name) . " Document";
    $result = self::showTitle($doc_title);

    $result .= self::showClass($class_name, $reflection) . " {\n\n";

    // 输出常量
    foreach ($reflection->getConstants() as $key => $value) {
        $result .= "const {$key} = " . var_export($value, true) . ";\n";
    }

    // 输出属性
    foreach ($reflection->getProperties() as $propertie) {
        $result .= self::showPropertie($propertie) . "\n";
    }

    //输出方法
    $result .= "\n";
    foreach($reflection->getmethods() as $value) {
        $result .= self::showMethod($value) . "\n";
    }

    // 文件结尾
    $result .= "}\n";
    return $result;
}
```

再回到上面我们使用反射的例子:

```php

// 获取 model 对应的数据库 table 名
public static function getTableName()
{
    // 反射, 这个后面会讲到
    $class = new \ReflectionClass(static::class);
    return strtolower(preg_replace('/((?<=[a-z])(?=[A-Z]))/', '_', $class->getShortName()));
}
```

注意, 这里要使用 `static`, 如果你使用 `self` 得到的就是 `BaseModel` 了. 至于一个简单的理解 static & self 的方式: static 是指当前内存中运行的实例, 所以永远都是 **所见即所得**.

## 魔术方法的性能损失


本来我也想做一下 profile 的, 还折腾起了 xhprof 和 xdebug, 但是其实可以简单的测试:

```php
$start = microtime();
dosomething();
echo microtime() - $start; // 单位: 微秒
```

感谢这位仁兄做的测试 [PHP 魔术方法性能测试](http://tigerb.cn/2017/03/04/php-magic-function/), 实测结果下来性能损失在 10us 内, 这个数量级, 我个人认为除非少数极端要求性能的场景, 完全是可以接受的.

最后, 补充一下 *单例模式* 的优缺点:

优点：

1. 改进系统的设计
2. 是对全局变量的一种改进

缺点：

1. 难于调试
2. 隐藏的依赖关系
3. 无法用错误类型的数据覆写一个单例
