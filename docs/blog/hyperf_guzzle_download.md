# hyperf| hyperf/guzzle 下载文件问题详解

使用 `hyperf/guzzle` 遇到下载文件的问题, 过程虽波折, 却颇有意思, 分享给大家.

## 业务上有一个下载文件的任务, 太简单啦

业务上要下载一个 oss 文件, 假设地址为 `oss_url`. 先撸一遍 guzzle 的文档, 下载需要在 request option 中设置 [`sink`](http://docs.guzzlephp.org/en/stable/request-options.html#sink) 参数:

```php
$oss_url = 'oss_url';
$file_path = 'xxx';
$client = new \GuzzleHttp\Client([
    'verify' => false,
    'decode_content' => false,
]);
$client->get($oss_url, [
    'sink' => $file_path,
]);
```

换成 hyperf/guzzle 来写:

```php
// 使用 clientFactory 获取到协程版 client 即可
$container = ApplicationContext::getContainer();
$clientFactory = $container->get(ClientFactory::class);
$client = $clientFactory->create([
    'verify' => false,
    'decode_content' => false,
]);
```

开心的执行, 剧情按照预期的方向发展~

**并没有!!! 下载没反应!!!**

## 下载地址的问题?

使用 curl/wget 等命令行工具验证

```
wget oss_url
curl -O oss_url
```

下载地址没问题

## 文件有 300M, 会不会是超时了?

```php
$client = $clientFactory->create([
    'timeout' => 600, // 超时设置为 10 分钟
    'verify' => false,
    'decode_content' => false,
]);
```

噗, 还是一样没效果?

## 翻 hyperf 文档

> hyperf/guzzle: https://doc.hyperf.io/#/zh/guzzle

哦, 原来使用 swoole 配置要这样:

```php
<?php
use GuzzleHttp\Client;
use Hyperf\Guzzle\CoroutineHandler;
use GuzzleHttp\HandlerStack;

$client = new Client([
    'base_uri' => 'http://127.0.0.1:8080',
    'handler' => HandlerStack::create(new CoroutineHandler()),
    'timeout' => 5,
    'swoole' => [ // 看这里看这里
        'timeout' => 10,
        'socket_buffer_size' => 1024 * 1024 * 2,
    ],
]);

$response = $client->get('/');
```

这次应该行了吧 !?

噗, 还是一样没效果.

## 直接用 guzzle 行不行

```php
$oss_url = 'oss_url';
$file_path = 'xxx';
$client = new \GuzzleHttp\Client([
    'verify' => false,
    'decode_content' => false,
]);
$client->get($oss_url, [
    'sink' => $file_path,
]);
```

终于看了需要下载的问题. 是时候刷锅一波给 hyperf, `什么渣渣框架, 文件下载居然都不支持`.

## 我们看看锅到底在哪

老规矩, 看源码, 既然是使用 `\Hyperf\Guzzle\ClientFactory->create()` 新建的, 看看源码涨啥样:

```php
public function create(array $options = []): Client
{
    $stack = null;
    if (Coroutine::getCid() > 0) {
        $stack = HandlerStack::create(new CoroutineHandler());
    }

    $config = array_replace(['handler' => $stack], $options);

    if (method_exists($this->container, 'make')) {
        // Create by DI for AOP.
        return $this->container->make(Client::class, ['config' => $config]);
    }
    return new Client($config);
}
```

协程环境下使用的 `new CoroutineHandler`, 来看看庐山真面目(文件有点长, 不要轻言放弃):


```php
// \Hyperf\Guzzle\CoroutineHandler
// 关键在这句, 这里其实是 \Swoole\Coroutine\Http\Client
$client = new Client($host, $port, $ssl);
```

原来这里用的 `\Swoole\Coroutine\Http\Client`. 这时候我灵机一动, 会不会是?

来看 swoole 的文档, [下载方法在这里](https://wiki.swoole.com/wiki/page/938.html)

```php
$host = 'www.swoole.com';
$cli = new \Swoole\Coroutine\Http\Client($host, 443, true);
$cli->set(['timeout' => -1]);
$cli->setHeaders([
    'Host' => $host,
    "User-Agent" => 'Chrome/49.0.2587.3',
    'Accept' => '*',
    'Accept-Encoding' => 'gzip'
]);
$cli->download('/static/files/swoole-logo.svg', __DIR__ . '/logo.svg');
```

这 api 和 guzzle 完全不一样呀 !!! 坑爹呢这是, 用 `swoole + hyperf`, 连个文件下载都搞不定 ?!

## swoole + hyperf 表示我这么强大, 你居然不会用 !

在 [swoole v4.4.6](https://wiki.swoole.com/wiki/page/p-4.4.6.html) 的版本更新中, 就增加了对 curl hook 的支持, 添加 `SWOOLE_HOOK_FLAGS` 即可, hyperf v1.1.0 版本中已经提供了支持:

```php
// bin/hyperf.php:11
! defined('SWOOLE_HOOK_FLAGS') && define('SWOOLE_HOOK_FLAGS', SWOOLE_HOOK_ALL);
```

开启后, swoole 就会无缝将 curl hook 为协程版.

> 无形加速, 最为致命 !

可是 hyperf/guzzle 默认还是使用的 `\Swoole\Coroutine\Http\Client` 怎么办? 这就太简单了:

- 如果你喜欢静态类风格:

```php
<?php

namespace App\Util;

use GuzzleHttp\Client;

class Guzzle
{
    /**
     * @param array $config
     * @return Client
     */
    public static function create(array $config)
    {
        return make(Client::class, ['config' => $config]);
    }
}
```

- 如果你喜欢 ClientFactory 风格

```php
<?php

namespace App\Util;

use GuzzleHttp\Client;

class ClientFactory
{
    /**
     * @param array $config
     * @return Client
     */
    public function create(array $config)
    {
        return make(Client::class, ['config' => $config]);
    }
}
```

## 写在最后

没事不要乱刷锅, 多问问问题, 多找找原因, 收获就在这不经意间~

---

web 开发就绕不开 http 请求, PHPer 做过几个项目之后, 就能看到各种对 http 的封装, 其中最多的就是大名鼎鼎的 curl. 

> 轮子这么多, 怎么从各种各样的轮子中解脱出来? 或者说, 什么才是轮子最核心的部分?

绕来绕去, 绕不开 HTTP 协议, 绕不开协议对应的 RFC. 在 PHP 中, 有很好的规范进行支持 -- [PSR-7](https://www.php-fig.org/psr/psr-7/). PSR-7 根据 RFC 定义了一组接口, 实现这组接口的库, 就和 RFC 是保持一致的. 在实现 PSR-7 的库中, 推荐使用 `guzzlehttp/guzzle`.

- guzzle 基础使用: get/post
- guzzle 高阶使用: try-catch stack-handler-middleware
- 文件下载: clientFactory 默认使用 Swoole\Http\Client, [下载方法是独立的](https://wiki.swoole.com/wiki/page/938.html), 和 guzzle 下载方式不兼容

```php
use GuzzleHttp\Client;

require __DIR__ . '/../vendor/autoload.php';

$client = new Client([
    'base_uri' => 'http://httpbin.org', // test
    'handler' => '', // handler is a constructor only option that cannot be overridden in per/request options
    'timeout' => 2.0,
]);
$r = $client->get('/');
var_dump($r->getStatusCode());

try {
    $r = $client->post('/', [
        'headers' => ['a' => 'b'],
        'query' => ['a' => 'b'], // 'query' => 'a=b',
        // for post: body json raw form_params(application/x-www-form-urlencoded) multipart(multipart/form-data)
        'body' => json_encode(['a' => 'b']),
        'sink' => '/path/to/file', // file download
        'cookies' => new \GuzzleHttp\Cookie\CookieJar(),
        'allow_redirects' => true, // default: true
    ]);
    var_dump($r->getBody()->getContents());
} catch (\GuzzleHttp\Exception\RequestException $e) {
    echo GuzzleHttp\Psr7\str($e->getRequest());
    // 4xx 5xx 时获取数据
    echo \GuzzleHttp\Psr7\str($e->getResponse());
}
```