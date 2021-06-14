# hyperf| hyperf 源码解读 1: 启动

- [hyperf 源码解读 1: 启动](https://www.jianshu.com/p/c167f6e130df)

hyperf 的准备工作做好后, 就开始运行启动命令了: 

```
php bin/hyperf
```

可以看到如下输出:

```
root@820d21e61cd8 /d/hyperf-demo# php bin/hyperf.php
Scanning ...
Scan completed.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\Di\Listener\BootApplicationListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\Config\Listener\RegisterPropertyHandlerListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\RpcClient\Listener\AddConsumerDefinitionListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\Paginator\Listener\PageResolverListener listener.
[DEBUG] Event Hyperf\Framework\Event\BootApplication handled by Hyperf\JsonRpc\Listener\RegisterProtocolListener listener.
Console Tool

Usage:
  command [options] [arguments]

Options:
  -h, --help            Display this help message
  -q, --quiet           Do not output any message
  -V, --version         Display this application version
      --ansi            Force ANSI output
      --no-ansi         Disable ANSI output
  -n, --no-interaction  Do not ask any interactive question
  -v|vv|vvv, --verbose  Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug

Available commands:
  help               Displays help for a command
  info               Dump the server info.
  list               Lists commands
  migrate
  start              Start swoole server.
  t                  Hyperf Demo Command
 db
  db:model
 di
  di:init-proxy
 gen
  gen:amqp-consumer  Create a new amqp consumer class
  gen:amqp-producer  Create a new amqp producer class
  gen:aspect         Create a new aspect class
  gen:command        Create a new command class
  gen:controller     Create a new controller class
  gen:job            Create a new job class
  gen:listener       Create a new listener class
  gen:middleware     Create a new middleware class
  gen:migration
  gen:process        Create a new process class
 migrate
  migrate:fresh
  migrate:install
  migrate:refresh
  migrate:reset
  migrate:rollback
  migrate:status
 queue
  queue:flush        Delete all message from failed queue.
  queue:info         Delete all message from failed queue.
  queue:reload       Reload all failed message into waiting queue.
 vendor
  vendor:publish     Publish any publishable configs from vendor packages.
```

今天要看这么多内容么? 不, 只看这部分: 

```
root@820d21e61cd8 /d/hyperf-demo# php bin/hyperf.php
Scanning ...
Scan completed.

...
```

这部分就是整个框架的核心, 这部分搞清楚了, 后面都是搭积木了, `随用随取`.

> PS: 看源码, 尤其是优秀开源项目的源码, 是程序员进阶的「终南捷径」.

## 入口: `bin/hyperf.php`

```php
#!/usr/bin/env php
<?php

use Hyperf\Contract\ApplicationInterface;

// php ini 设置
ini_set('display_errors', 'on');
ini_set('display_startup_errors', 'on');

error_reporting(E_ALL);

// 定义常量 BASE_PATH, 所有路径相关都会使用这个常量
!defined('BASE_PATH') && define('BASE_PATH', dirname(__DIR__, 1));

// composer 自动加载
require BASE_PATH . '/vendor/autoload.php';

// Self-called anonymous function that creates its own scope and keep the global namespace clean.
(function () {
    // container
    /** @var \Psr\Container\ContainerInterface $container */
    $container = require BASE_PATH . '/config/container.php';

    // application
    $application = $container->get(ApplicationInterface::class);
    $application->run();
})();
```

很简单的几部分:
- PHP ini 设置, 按需设置即可, 比如这里还可以设置时区
- 常量 `BASE_PATH`, hyperf 只设置了这个一个常量, 用来所有 `路径` 相关的场景
- `config/container.php`, container 的初始化, 重中之重的内容
- `Application->run()`, 完整的是 `Symfony\Component\Console\Application`, 用来跑 cli 应用

> PS: 有轮子, 而且还很好用, 干嘛非要自己造. 这也是要读源码的理由之一.

## 重点: `config/container.php` 

到重点内容了, `重要的事情说三遍`

```php
use Hyperf\Config\ProviderConfig;
use Hyperf\Di\Annotation\Scanner;
use Hyperf\Di\Container;
use Hyperf\Di\Definition\DefinitionSource;
use Hyperf\Utils\ApplicationContext;

// 使用 composer 提供的工具 ProviderConfig 
$configFromProviders = ProviderConfig::load();

// dependency
$definitions = include __DIR__ . '/dependencies.php';
$serverDependencies = array_replace($configFromProviders['dependencies'] ?? [], $definitions['dependencies'] ?? []);

// annotation
$annotations = include __DIR__ . '/autoload/annotations.php';
$scanDirs = $configFromProviders['scan']['paths'];
$scanDirs = array_merge($scanDirs, $annotations['scan']['paths'] ?? []);

// scan
$ignoreAnnotations = $annotations['scan']['ignore_annotations'] ?? ['mixin'];

// container 初始化
$container = new Container(new DefinitionSource($serverDependencies, $scanDirs, new Scanner($ignoreAnnotations)));

if (! $container instanceof \Psr\Container\ContainerInterface) {
    throw new RuntimeException('The dependency injection container is invalid.');
}

// 设置后, 方便全局获取 container 实例
return ApplicationContext::setContainer($container);
```

### 使用 composer 提供的工具 ProviderConfig

```php
// \Hyperf\Config\providers
$providers = Composer::getMergedExtra('hyperf')['config'] ?? [];
```

关键是这句, 对应获取到的 `composer.json` 文件中的配置:

```json
    "extra": {
        "branch-alias": {
            "dev-master": "1.1-dev"
        },
        // 对应这里的配置
        "hyperf": {
            "config": "Hyperf\\Amqp\\ConfigProvider"
        }
    },
```

对应的 `ConfigProvider` 内容:

```php
namespace Hyperf\Amqp;

use Hyperf\Amqp\Packer\Packer;
use Hyperf\Utils\Packer\JsonPacker;

class ConfigProvider
{
    public function __invoke(): array
    {
        return [
            'dependencies' => [
                Producer::class => Producer::class,
                Packer::class => JsonPacker::class,
                Consumer::class => ConsumerFactory::class,
            ],
            'commands' => [
            ],
            'scan' => [
                'paths' => [
                    __DIR__,
                ],
            ],
            'publish' => [
                [
                    'id' => 'config',
                    'description' => 'The config for amqp.',
                    'source' => __DIR__ . '/../publish/amqp.php',
                    'destination' => BASE_PATH . '/config/autoload/amqp.php',
                ],
            ],
        ];
    }
}
```

这里有 4 部分内容:
- dependencies: 依赖关系, 解耦神器
- commands: 部分 hyperf 组件有有自定义的 command, `php bin/hyperf.php` 看到的命令, 配置就是这里来的
- scan: 设置扫描目录, hyperf 组件是默认是组件源码目录 `src/`
- publish: 通常用来加载组件提供的默认配置文件, 或者其他一些组件提供的 demo 文件

### container 初始化

```php
// \Hyperf\Di\Definition\DefinitionSource::__construct
$container = new Container(new DefinitionSource($serverDependencies, $scanDirs, new Scanner($ignoreAnnotations)));
```

别看只有一行, 这里干的事情可真不少, 要得到 container 这个 `缺啥都找它要` 的神器, `当然没那么简单`(哼起来~)

关键代码是这里:

```php
// \Hyperf\Di\Definition\DefinitionSource::scan
    private function scan(array $paths): bool
    {
        if (empty($paths)) {
            return true;
        }
        $pathsHash = md5(implode(',', $paths));
        if ($this->hasAvailableCache($paths, $pathsHash, $this->cachePath)) {
            $this->printLn('Detected an available cache, skip the scan process.');
            [, $annotationMetadata, $aspectMetadata] = explode(PHP_EOL, file_get_contents($this->cachePath));
            // Deserialize metadata when the cache is valid.
            AnnotationCollector::deserialize($annotationMetadata);
            AspectCollector::deserialize($aspectMetadata);
            return false;
        }
        $this->printLn('Scanning ...');
        // 关键在这里
        $this->scanner->scan($paths);
        $this->printLn('Scan completed.');
        if (! $this->enableCache) {
            return true;
        }
        // enableCache: set cache
        if (! file_exists($this->cachePath)) {
            $exploded = explode('/', $this->cachePath);
            unset($exploded[count($exploded) - 1]);
            $dirPath = implode('/', $exploded);
            if (! is_dir($dirPath)) {
                mkdir($dirPath, 0755, true);
            }
        }
        $data = implode(PHP_EOL, [$pathsHash, AnnotationCollector::serialize(), AspectCollector::serialize()]);
        file_put_contents($this->cachePath, $data);
        return true;
    }
```

看起来有点复杂呀, 别慌, 一言以蔽之, scan 是为了给我们想要的数据:

```php
AnnotationCollector::serialize()
AspectCollector::serialize()
```

没错, `注解(Annotation) + Aspect(切面)`

### container 使用

基于 hyperf 的应用中, 缺啥都找 container 就对了, 具体的文档可以参考 [hyperf doc - 依赖注入](https://doc.hyperf.io/#/zh/di)

这里补充 2 点, 一个是 container 的补充说明:

```php
// \Hyperf\Di\Container::get
    public function get($name)
    {
        // If the entry is already resolved we return it
        if (isset($this->resolvedEntries[$name]) || array_key_exists($name, $this->resolvedEntries)) {
            return $this->resolvedEntries[$name];
        }
        $this->resolvedEntries[$name] = $value = $this->make($name);
        return $value;
    }
```

上看 scan 看似复杂, 最终都会处理到 container 的 `$this->resolvedEntries[$name]` 变量里, 不明白的话, 可以把这变量打印一下看一看

第二点对转到 `依赖注入` 下的小伙伴说的:

> 自己 new 出来的变量是无法用到强大的 container 的, 以及之后各类好用的方法, 真爱生命, 不要瞎 `new` 哦~

## 写到最后

hyperf 最核心的部分我们已经看到了, 没错, 就是 container, container 在手, 天下我有.

> 下篇预告: 依旧从安装 hyperf 就会执行的命令 `php bin/hyperf.php start` 入手, 强大的 swoole 在呼唤着我们 !