# hyperframework| WebClient 源码解读

先说句题外话, 我在每篇 blog 上都会先加上 `date`, 然后一直把 blog 放到编辑器中, 之后不断做类似「提纲」类的记录, 最后找一个大段的时间书写.

这篇 blog 的起源, 来自于上周一个工作任务过程中的「坎坷」 -- 接某支付的 sdk, 返回一直报参数错误. 因为自己也接过不少支付的 sdk, 所以一直怀疑是 「签名错误」. 直到详细阅读 sdk 的源码和 `hyperframework webclient` 的源码, 才解开这个谜题. 应该有很多程序员大大和我一样, 会经常和 http 打交道, 希望这篇文章能有所帮助.

## php 中 3 种 http 请求工具对比

这里对比的 3种 http 请求工具:

- cURL: http://php.net/manual/en/book.curl.php
- 非常火的 composer package - Guzzle: http://docs.guzzlephp.org/en/stable/overview.html
- hyperframework 中的 webclient 工具: http://hyperframework.com/cn/manual/common/web_client_basics

get 请求:

```php
$url = 'www.example.com/curl.php?option=test';

// cURL
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $url);
// 可以合并成: $ch = curl_init($url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
$output = curl_exec($ch);
curl_close($ch);

// Guzzle
$client = new \GuzzleHttp\Client();
$response = $client->get($url);

// hyperframework webclient
$c = new \Hyperframework\Common\WebClient();
$r = $c->get($url);
```

post 请求:

```php
$url = 'http://httpbin.org/post';

// cURL
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_HEADER, 0);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
$output = curl_exec($ch);
curl_close($ch);
echo $output;

// Guzzle
$response = $client->request('POST', $url, [
    'form_params' => [
        'field_name' => 'abc',
        'other_field' => '123',
        'nested_field' => [
            'nested' => 'hello'
        ]
    ]
]);

// hyperframework webclient
$c = new \Hyperframework\Common\WebClient();
$r = $c->post($url, [
    'field_name' => 'abc',
    'other_field' => '123',
]);
```

通过对比, 希望你能从 3 种「风格」中感受到工具各自的设计思想:

- 都采用 「对象」 来完成一次 http 请求, 为什么? 因为从一次 http 请求的生命周期来看, 非常适合使用对象这个概念来处理
- 抽象程度依次增加, cURL 需要你设置更多(意味着你需要知道更多细节), 对比 Guzzle 和 hyperframework webclient 可以发现 get 相差无几, 但是 post 上, Guzzle 多了一层 key 值来设置, 而 webclient 则把这层隐藏掉了

而我这次踩到的坑, 就和 post 的这层隐藏有关.

> PS: 关于我说 cURL 也是「对象」的方式, 大家可以参考 swoole 的源码: 2层目录, 面对对象风格写 c ([from swoole wiki](https://wiki.swoole.com/))

## hyperframework webclient 源码解读
> 源码在此: https://github.com/hyperframework/hyperframework/blob/master/lib/Common/WebClient.php

hyperframework webclient 源码解读起来非常容易, 也推荐大家也读一读看看, 可以帮助你看到一些 http 请求的细节.

解读源码, 尤其是 webclient 这样的单个类文件, 可以从 「生命周期」 的角度来试试, 会简单很多:

- `$c = new WebClien();` 实例化一个 WebClient 对象, `__construct()` 方法中可以设置初始化 `$options`
- `$r = $c->post($url, $data, $option);` 执行一次 http 请求, 实际操作的方法 `sendHttpRequest() -> send() -> initializeRequest()`, 而这些方法本质做的是同一件事: 设置 `$requestOptions`
- 执行 `curl_setopt_array()` 来使用 `$requestOptions` 中的值, 然后执行 `curl_exec`

这里有 2 个概念要明确: **`$options` 属于实例级别, `$requestOptions` 属于请求级别**. 多这样一层抽象出来, 就是为了方便对象的复用.

大家着重看一下 `processDataRequestOption()` 方法的代码, 这里有 post 请求的一些细节. 这里先说一下我踩到的坑:

```php
private function processDataRequestOption() {
    if ($this->hasRequestOption(self::OPT_DATA) === false) {
        return;
    }
    $data = $this->getRequestOption(self::OPT_DATA);
    $defaultType = 'application/json';
    ...
}
```

可以看到, 这里默认会设置 `Content-Type: application/json`, 而我对接的某支付 sdk, 服务器那边必须要使用 `application/x-www-form-urlencoded` 才可以. **而由于我之前对接支付 sdk 的经历, 我一直纠结在签名错误上, 导致处理这个问题花了很久**. 因为 `Content-Type` 设置错误, 导致服务器接受到的数据解析出错, 那当然会验签失败.

继续阅读 `processDataRequestOption()` 的源码, 下面会处理不同的 `Content-Type`, 而本质上, 就是在 **处理字符串** 而已.(处理字符串也是基本功呀.)

```php
private function processDataRequestOption() {
    if ($this->hasRequestOption(self::OPT_DATA) === false) {
        return;
    }
    $data = $this->getRequestOption(self::OPT_DATA);
    $defaultType = 'application/json';
    $type = $this->hasRequestOption(self::OPT_DATA_TYPE) ?
        $this->getRequestOption(self::OPT_DATA_TYPE) : $defaultType;
    $typeSuffix = null;
    $position = strpos($type, ';');
    if ($position !== false) {
        $typeSuffix = substr($type, $position);
        $type = substr($type, 0, $position);
    }
    $lowercaseType = strtolower(trim($type));
    if (is_string($data)) {
        $this->addRequestHeader('Content-Type: ' . $type . $typeSuffix);
        $this->initializeCurlPostFieldOptions();
        $this->setRequestOption(CURLOPT_POSTFIELDS, $data);
        return;
    }
    if ($lowercaseType === 'multipart/form-data') {
        ...
    } elseif ($lowercaseType === 'application/x-www-form-urlencoded') {
        ...
    } elseif ($lowercaseType === $defaultType) {
        ...
    } else {
        throw new WebClientException(
            "Data type '$type' is not supported."
        );
    }
}
```

## 继续聊聊 `Content-Type`
> Http Header里的Content-Type: https://www.cnblogs.com/52fhy/p/5436673.html

推荐大家读一下上面这篇博客, 结合 [postman](https://www.getpostman.com/) 实操来讲解 http header 中的 `Content-Type`, **理论 + 实践**.

常用的 `Content-Type` 只有几种, 可以参照上面的源码解读:

- `application/json` 对应 postman 中的 `raw -> JSON`, 随着 「大前端」 时代的到来以及文档型存储数据库的兴盛, json 格式普及率越来越高
- `multipart/form-data` 对应 postman 中的 `form-data`, 格式最复杂, 可以用来上传文件
- `application/x-www-form-urlencoded` 对应 postman 中的 `x-www-form-urlencoded`, 默认值, html form 表单提交默认就是这种格式
- 还有 `text/plain` `text/xml` `text/html` 等几种, 对用 postman 中的 `raw -> `, 纯文本 / xml / html 都是常见的格式

通过源码细节可以知道, 不同的 `Content-Type`, 字符串格式是不一样的, 其中 `multipart/form-data` 最复杂, `x-www-form-urlencoded` 其实使用 php 中的 `htp_build_query()` 函数来格式化数据.

> PS: 有没有和我一样, 一直以为 `htp_build_query()` 函数只是用来拼接 get 请求参数的, 所以还是要多读一些源码.

稍等, 到这里还没完, 这里只完成了 **你按照一定格式组装好数据**, 而对方接收到数据, 还需要 **按照格式解析数据**.

对 php 熟悉的小伙伴, 应该知道 php 中有 3 种方式接收 post 来的数据:

- `$_POST` 数组, 也是最常见的方式, 不过大家使用框架的过程中, 会发现框架都会提供 `Request::get('xxx')` 这样类似的方法
- `file_get_contents('php://input')`, 需要读取一些 **原始** 数据的时候, 通常是 `$_POST` 无法解析的数据
- `$HTTP_RAW_POST_DATA`, 读过 php manual 就知道这是个 **旧** 方法, 推荐使用 `$_POST` 来代替

之所以推荐框架中封装的 `Request::get('xxx')` 这样类似的方法, 是因为 `$_POST` 并不是每次都能处理好数据解析, 比如 json 数据. 而框架多了这一层抽象, 其中之一就是为了处理这种问题.

之前也写过比较上面三者的 blog, 当时给出了尽量使用 `file_get_contents('php://input')` 的结论. 这里着重说一下, **千万不要迷信**.

> 迷信, 其实来自无知.

推荐阅读下面这篇 blog, `php://input` 是解析不了 `form-data` 格式的数据的, 这个问题让我使用 `postman 测试 + php://input 设置断点` 时一直返回为空时郁闷了很久

> 深入理解 `php://input` : http://www.nowamagic.net/academy/detail/12220520

## 数一数踩过的坑吧

其实在上面也列举了一部分, 这里总结一下, 方便大家查阅:

- 对接某支付 sdk, 由于 `Content-Type` 错误, 导致签名一直失败, 最后通过阅读支付 sdk 提供的 demo 以及 hyperframework WebClient 的源码解决
- 之前维护过一个 nodejs 的项目, 出现端(IOS/Android)发起的请求均失败, 通过日志发现端这边使用的 `form-data` 格式提交的数据, 而 nodejs 这边使用的 koa2 框架, 默认解析 post 请求是支持 `x-www-form-urlencoded` 的. 事情到这里还没完, 为了支持 `form-data` 格式, 需要 npm 安装一个包, 但是当时在十九大期间, 连淘宝镜像源都无法安装这个包, 而因为深夜执行 npm 操作, 导致整个项目的包管理挂了, 最终服务器宕了. 最后通过以前备份的服务器镜像, 复制项目目录替换解决. **注意: 一定要小心包管理**
- 以前对接支付 sdk 的时候, 经常遇到异步回调没有正常处理的情况. 通过几种方式打日志, 最终发现 `php://input` 最靠谱, 虽然需要自己 `json_decode()` 一下来格式化数据
- 对接某大行的支付 sdk 的时候, 通过打的日志发现, 异步回调的数据格式有 `json_encode()` 和 `http_build_query()` 2种形式, 但是这个问题在测试过程中(接近 10 笔订单)过程中没有出现过, 而且由于这个渠道订单量一直很小, 所以问题也是过了一段时间才发现. **提醒一下: 一定要打好日志; 在敏感数据处理的时候, 没有获取到预期的数据, 最好加一下预警**

还有一些年代可能有点久远了, 或者是当时实在太 sb, 犯了一些低级错误, 就不赘述了.

## 写在最后

关于 http 的学习, 其实我已经在之前的 [blog - alipay ILLEGAL_SIGN 错误解决](http://www.jianshu.com/p/28585a6454b2) 有提到. 当时的问题, 抽丝剥茧一层层下来, 最后到 http 协议这一层, 竟然是非常非常的简单.

所以继续推荐这本书:

> <http 权威指南> - 图灵社区: http://www.ituring.com.cn/book/844

这里给一点阅读的建议:

- 这是一本很纯粹的工具书. 工具书的特点其实和字典非常类似, 你用字典的时候, 只要知道查词的方法(具体的方法就是大名鼎鼎的 「二分查找法」, 可以配合 [网易公开课 - 斯坦福大学公开课：编程方法学](http://open.163.com/special/sp/programming.html), 第一集举的就是这个例子)就好了, 并不需要记住所有细节.
- http 其实是 `tcp/ip` 4层网络体系下一种应用层的协议. 从协议的角度来看待, 它由哪些部分组成, 这些部分之间如何协同, 就是学习 http 需要掌握的方法, 比如我们常说到的 `header body method` 等概念, 都是 http 协议的组成部分.

当然, 光看一本书是不能完全解决问题的, 毕竟基于 http 的基础上, 大家又多了各式形形色色的工具.

> 工具的目的是提供便利, 从编程方法学的角度考虑, 其实是增加抽象.

所以, 多度一些源码吧, 既用来解决实际的业务问题, 也可以用来培养自己对 「编程方法」 的理解.
