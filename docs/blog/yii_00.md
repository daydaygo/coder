# yii| 最佳实践之黑箱思维

> [yii| 最佳实践之黑箱思维](https://www.jianshu.com/p/8f52b75055d5)

工作中需要重度使用 yii, 逐渐积累 **最佳实践** 的过程中, 加上自己一些编程方面的所思所得, 希望通过 **yii最佳实践 blog 系列**, 可以给使用 yii 的同学提供帮助.

本文速览:

- 黑箱思维概念解读: 关注输入和输出; 输出是不是你想要的
- 黑箱一: PHP中的包管理 -- 站在巨人的肩膀上
- 黑箱二: 自动加载 -- PHP 代码的自动加载进制
- 黑箱三: 改功能 = 改配置 -- yii 底层机制实现功能灵活修改
- 黑箱四: 组件 component -- yii 中代码复用

## 黑箱思维

> 黑箱思维: 关注输入和输出; 输出是不是你想要的.

黑箱思维是在 **工具思维** 上延伸出来的, 用来解决 **怎么使用好工具?** 这一问题. 这和 yii 框架有什么关系呢?

> 框架: 一类特殊库/可重用组件, 提供一个能被进一步开发的通用应用程序功能

yii框架, 就是这样一个帮助我们开发大型软件的工具. 黑箱思维可以很好的帮助我们用好这个工具

## 黑箱一: 包管理

提到工具, 这句话一定听过: **站在巨人的肩膀上**. 具体到软件开发, 就是代码复用的问题了.

> 代码复用: 这是软件工程的难题, 需要慢慢积累, 有些地方可以通过遵循规范走走捷径

具体到语言的生态圈, 就是 **包管理**. PHP因为其发展历史之久, 包管理也有些复杂. 不过好在有 [php the right way(PHP之道)](http://laravel-china.github.io/php-the-right-way/), 帮助大家整理出 **the modern way**

> PHP 之道 收集了现有的 PHP 最佳实践/编码规范/权威学习指南, 方便 PHP 开发者阅读和查找

PHP包管理现状:

- 扩展: 使用 `php -m` 就可以查看到扩展, 扩展使用 c 语言写, 使用 [pecl](http://pecl.php.net/) 管理

```bash
php -m # 查看扩展
php --ri swoole # 查看扩展的信息

pecl install swoole
pecl install swoole-1.9.8 # 安装指定版本

# 既然是 c 语言写的, 也可以编译安装
phpize
./configure # 这里修改默认的编译参数
make && make install
```

- 包([packagist](https://packagist.org/)): 包使用 php 语言写, 使用 [composer](http://docs.phpcomposer.com/) 管理, 配置文件为 `composer.json`

```bash
composer require yiisoft/yii2 # 添加新包
composer install # 根据 composer.json 文件安装包, 初始化(第一次)时使用
composer update # 如果修改了 composer.json 添加/更改包, 执行这个命令来生效
```

**扩展(ext)** 和 **包(package)** 是 **必须** 要知道和掌握的技能.

可能还会听到一些比较古老的方式:

- pear: composer 之前的包管理工具, 只做了解
- phar: 类似 java 的 jar, 将 php 程序打包, 然后直接 `php xxx.phar` 运行

```bash
php composer.phar install
```

composer 这个黑箱里面, 其实就是 `composer.phar`, 命令行下多封装了一层:

```bash
# composer 可执行文件的内容
php "${dir}/composer.phar" "$@" # 其实还是转为 php composer.phar xxx 来执行
```

类似的工具还有 2 个值得推荐: `php-cs-fixer-v2.phar` `phpDocumentor.phar`

```bash
# 使用代码规范格式化指定目录下的 php 代码, 默认使用 psr-2
php-cs-fixer fix src/

# 通过 php 代码中的注释, 来生成文档
php phpDocumentor.phar --title="SLS_PHP_SDK" --defaultpackagename="SLS_PHP_SDK" --template="responsive" -d Aliyun -t docs
```

继续 **代码复用** 的话题, **有些地方可以通过遵循规范走走捷径**, 所以 [PSR: php standard recommend](http://www.php-fig.org/psr/) 一定要清楚

> 关键词: 扩展(ext + pecl) 包(package + composer) PSR

## 黑箱二: 自动加载

通过 **扩展(ext)** 添加的功能, 只要扩展开启了, php 代码中就可以直接使用:

```php
$lock = new \Swoole\Lock(SWOOLE_MUTEX); // 直接使用类
```

但是 **包(package)** 使用 php 语言编写的, 想要做到扩展这样 **直接使用**, 就需要用到 **自动加载**.

关于自动加载的原理, 这个教程 [5-1 SPL使用spl_autoload_register函数装载类 (10:03)](http://www.imooc.com/video/2620) 非常好, 值得一看.

**自动加载** 发展的历程:

- 使用 `require() / include()` 等方法, 这个方法现在也经常使用, 比如加载配置文件
- `__autoload()` 魔术方法, 已经被下面的方法取代
- `spl_autoload_register()` 来注册自动加载方法, **自动加载方法定义了怎么帮你找到需要的类**

是不是 **闻所未闻**? 很正常, composer 把这件事做掉了:

- 先看项目根目录的 composr.json:

```json
{
    "require": { // 这里表明需要的包和版本
    "php": ">=5.4.0",
    "yiisoft/yii2": "~2.0.13",
    "yiisoft/yii2-bootstrap": "~2.0.0",
    "yiisoft/yii2-mongodb": "^2.1@dev"
  }
}
```

- 安装好后, 再来看 **yiisoft/yii2** 中的 composer.json 文件:

```json
{
    ...
    "autoload": { // 定义自动加载
        "psr-4": { // psr-4 标准
            "yii\\": "" // 命名空间 -> 文件路径, 一般使用相对于当前 composer.json 的相对路径
        }
    }
    ...
}
```

真正起作用的地方还是 `autoload`, 最终通过 psr-4 标准来加载类. psr-4 标准规定 **路径/类名 要和命名空间一一对应**, 所以在使用 `yii\web\Request` 时, 对应的就是 `yii2/web/Request.php` 文件.

到这里其实并没有解决问题, 反而多了2个问题, 对比一下 laravel 框架的 composer.json:

```
...
"autoload": {
    "classmap": [
        "database/seeds",
        "database/factories"
    ],
    "psr-4": {
        "App\\": "app/"
    }
}
...
```

- 问题1: autoload 有几种方式?

四种, psr-4 和 psr-0 区别不大, 请尽量使用 psr-4; `classmap` 用来直接加载类, 可以用来兼容历史代码(不带命名空间的类); `function` 用来加载自定义的 PHP 函数, 一般用来放 help function

- 问题2: yii 框架中, basic模板业务代码在 `app\` 命名空间下, advanced模板业务代码分别在 `backend\` `frontend\` `common\` 命名空间下, 但是, 并没有再 autoload 中定义, 那它是怎么自动加载的呢?

答案在 `Yii.php` 文件中:

```php
class Yii extends \yii\BaseYii
{
}

// 看这里
spl_autoload_register(['Yii', 'autoload'], true, true);

Yii::$classMap = require __DIR__ . '/classes.php';
Yii::$container = new yii\di\Container();
```

看看, 又回到 `spl_autoload_register()` 了, yii 自己定义了一个 [autoload](http://www.yiichina.com/doc/guide/2.0/concept-autoloading). **我不提倡这种自动加载的方式**, 推荐使用 composer 统一管理

> 关键词: 自动加载的发展历程 自动加载的种类 yii中的自动加载

## 黑箱三: 改功能 = 改配置

先说一个 **颇为浪费时间** 的经历, 在 coding `yii best practice` 的过程中, 需要精简 console 应用下的命令, 默认实在有点多而且我根本用不上:

![yii 默认提供的命令](http://qiniu.dayday.tech/yii-cli-origin.png)

过程有些曲折, 避免变成流水账, 简述一下步骤:

- 确定需求: 默认命令用不上, 精简掉 -> 这一步是对的, 如果一开始需求就错了, 大部分情况不会是 「无心插柳柳成荫」 的结果
- 尝试一: 查看 [官方文档 - 控制台命令](http://www.yiichina.com/doc/guide/2.0/tutorial-console), 可惜没有提到 -> 这一步也是正确的, 文档基本是 **common-base(最常用)** 的内容慢慢叠加起来的
- 尝试二: 百度一下, 可惜也没有找到 -> 这一步也是正确的, 随着信息的爆炸, 越有价值的问题, 就越容易找到答案
- 尝试三: 因为之前有 [yii源码解读](https://www.jianshu.com/p/fd85383783eb) 的经验, 趁着 「热乎劲」 直接读源码, 结果花了半个小时各种 `var_dump()+die()`, 还是没有解决 -> 这一步开始错了, 一部分是出于 **自信**, 一部分是常用方法(文档 + 百度)用完了, 不动脑筋想其他方法
- 尝试四: 一步挣扎后, 发现 `yii\console\Applicate` 中有 `coreCommands()` 方法, 一番窃喜, 可惜这个是修改 yii 框架源码, 不可取 -> 这一步运气成分居多, 有点像 **苹果砸中了牛顿**
- 尝试五: 暂时搁置了问题一段时间, 再一次读源码的过程中, 发现 `yii\console\Applicate` 有 `public $enableCoreCommands = true;` 属性, 添加这个配置到 app 中即可 -> 其实一开始就该想到的

总结: 其实一开始就知道的方法 -- **改功能 = 改配置**, 却因为没有多动一下脑筋而失之交臂, 可惜呀. 同时也能说明 yii 框架的灵活性, 这也是 yii 框架设计上的优势.

来看看具体怎么实现的: `yii\base\BaseYii`

```php
public static function createObject($type, array $params = [])
{
    if (is_string($type)) {
        return static::$container->get($type, $params);
    } elseif (is_array($type) && isset($type['class'])) {
        $class = $type['class'];
        unset($type['class']);
        return static::$container->get($class, $params, $type);
    } elseif (is_callable($type, true)) {
        return static::$container->invoke($type, $params);
    } elseif (is_array($type)) {
        throw new InvalidConfigException('Object configuration must be an array containing a "class" element.');
    }

    throw new InvalidConfigException('Unsupported configuration type: ' . gettype($type));
}
```

Yii 框架中使用到的类, 都会通过这个方法, 然后经 `yii\di\Container` 类来生成. 配置中的配置项, 对应这个生成类中的成员属性. 所以, 当需要什么功能时, **可以找找相关的类, 看看类中有哪些属性**. 而且 yii 的注释写的很好, 有的还有代码示例.

> 好的框架, 也许就是你发现你需要什么, 不需要动底层源码来解决

## 黑箱四: 组件(component)

组件(component) 其实很简单, 就是一个类 -- **类 = 功能**. 联系到上面的内容, 这个类的配置, 就对应 App 配置中的 `component` 项, 通过 `Yii::$app->xxx` 就可以访问. 组件是很好的代码复用的方式. 下面用一个具体的例子来说明.

有一个接入阿里云日志系统的需求, 最开始使用 logtail(阿里云提供的基于文件的日志收集工具), 不过对日志的处理要使用到正则(yii框架的多行日志, 其他日志文件解析模式只能处理单行), 写起来实在不舒服, 所以把视角转到 [阿里云提供的日志服务SDK](https://github.com/aliyun/aliyun-log-php-sdk), 看了一下 `sample.php` 文件

```php
// 自动加载
require_once realpath(dirname(__FILE__) . '/../Log_Autoload.php');

// 配置
$endpoint = '<log service endpoint';
$accessKeyId = 'your access key id';
$accessKey = 'your access key';
$project = 'your project';
$logstore = 'your logstore';
$token = "";

$client = new Aliyun_Log_Client($endpoint, $accessKeyId, $accessKey,$token);

// 需要的功能
putLogs($client, $project, $logstore);

function putLogs(Aliyun_Log_Client $client, $project, $logstore) {
    $topic = 'TestTopic';

    $contents = array( // key-value pair
        'TestKey'=>'TestContent'
    );
    $logItem = new Aliyun_Log_Models_LogItem();
    $logItem->setTime(time());
    $logItem->setContents($contents);
    $logitems = array($logItem);
    $request = new Aliyun_Log_Models_PutLogsRequest($project, $logstore,
            $topic, null, $logitems);

    try {
        $response = $client->putLogs($request);
        var_dump($response);
    } catch (Aliyun_Log_Exception $ex) {
        var_dump($ex);
    } catch (Exception $ex) {
        var_dump($ex);
    }
}
```

抛去代码风格, 这个SDK要使用是不是hin简单?

> 好用的 SDK, 只用看一下 sample 或者 quick start 就能分辨出来.

好了, 再来看看怎么把这个 SDK 复用到 yii 的项目中:

- 项目中建一个目录, 专门用来这样的 SDK(不在 composer packgist 中), 引入自动加载

```php
// 以 console 应用的入口 yii 为例
require __DIR__ . '/../vendor/autoload.php';
require __DIR__ . '/../vendor/yiisoft/yii2/Yii.php';
require __DIR__ . '/../app/sdk/aliyun-log-php-sdk/Log_Autoload.php'; // 同样 require 进来就好
```

- 编写 component 来使用这个 SDK

新建 `AliyunLog` 类继承自 `Component`, 然后在 `init()` 进行日志服务客户端的初始化

```php
<?php

namespace app\service;

use yii\base\Component;

/**
 *  https://github.com/aliyun/aliyun-log-php-sdk
 */
class AliyunLog extends Component
{
    /**
     * 服务入口: https://help.aliyun.com/document_detail/29008.html
     * @var string
     */
    public $region = 'cn-shanghai-intranet.log.aliyuncs.com';
//    public $region = 'cn-shanghai.log.aliyuncs.com'; // 公网
    public $ak;
    public $sk;
    public $token = '';
    public $project;
    public $logStore;
    public $topic = 'TestTopic';
    /** @var \Aliyun_Log_Client $client */
    public $client;

    public function init()
    {
        parent::init();
        $this->client = new \Aliyun_Log_Client(
            $this->region,
            $this->ak,
            $this->sk,
            $this->token
        );
    }

    public function putLogs(array $logs)
    {
        $logitems = [];
        foreach ($logs as $log) {
            $logItem = new \Aliyun_Log_Models_LogItem();
            $logItem->setTime(time());
            $logItem->setContents($log);
            $logitems[] = $logItem;
        }

        $request = new \Aliyun_Log_Models_PutLogsRequest(
            $this->project,
            $this->logStore,
            $this->topic,
            null,
            $logitems
        );

        $this->client->putLogs($request);
    }
}
```

- 加到配置中, 就可以通过 `Yii::$app->aliyunLog` 使用它了:

```php
[
    'components' => [
        ...
        'aliyunLog' => [
            'class'    => \app\service\AliyunLog::class,
            'region'   => $env['aliyunLog']['region'],
            'ak'       => $env['aliyunLog']['ak'],
            'sk'       => $env['aliyunLog']['sk'],
//            'token' => '',
            'project'  => 'daydaygo-aliyun',
            'logStore' => 'aliyun',
//            'topic' => '',
        ],
    ]
]
```

到这里, 其实关于 **组件** 的内容已经讲完了, 不过我们的需求是 **日志接入阿里云日志服务**, 在 yii 中就还需要一个 `AliyunLogTarget` 类, 这个参考 `MongoDbTarget` 实现即可.

- 最后, 配置上 `AliyunLogTarget`, 打日志的业务代码完全不用动, 改一下配置就可以决定哪些放到阿里云日志服务里

```php
[
    'components' => [
        'log' => [
            'targets' => [
                [
                    'class'   => \app\service\AliyunLogTarget::class,
                    'levels'  => ['info', 'error', 'warning'],
                    'logVars' => [],
                    'except'  => [
                        'yii\db\*',
                        'yii\mongodb\*',
                        'yii\web\session',
                        'yii\base\UserException',
                        'yii\web\HttpException:403',
                    ],
//                    'project' => '',
//                    'logStore' => '',
                    'topic'   => 't2',
                ],
            ]
        ]
    ]
]
```

如果使用 yii 框架, 非常推荐使用 component 来实现代码复用. 这一节也提到 **怎么判断一个好用的 SDK**, 作为 yii 生态, 当然是官方支持的扩展, 比如 yii2-mongodb 是第一优先选择, 其次是 PHP 生态的包管理(扩展/包), 再其次就是上面提到的阿里云日志服务SDK, 一个简单的 sample, 完全不用关心实现细节, 几行代码就可以用起来.

## 写在最后

**黑箱思维**, 多少有点 **业务优先, 技术次之** 的感觉在里面, 但是却是符合实际的:

> 业务需求虽然无穷无尽千变万化, 但越往底层走, 需求改动的可能性越小, 需要推倒重来的情况非常有限.

**黑箱** 可以看做是面向对象中 **封装** 这一概念的延伸, 可以适用的范围更广.
