# php| 问题排查: 数据库代理服务报错

前言说完了, 来看问题. 这个问题从发现到最后解决, 前后历时 2 天:

- 排期好了, 业务等着使用, 既是压力, 也是动力
- 尝试各种突破口, 前一晚折腾到了凌晨 2 点, 这种解决问题的 **心流状态**, 很难得了. 

## 问题现场

新开了一个 `数据库代理服务`, 用来屏蔽使用的数据库资源的细节(rds-关系型数据库, drds-关系型数据库), 给业务方带来一致的使用体验.

新服务在测试环境跑了 2 周, 都没有问题, 切到线上环境使用, 使用 phpunit 跑单测报错.

报错原文:

```
TypeError:Argument 1 passed to Hyperf\Database\Connection::prepared() must be an instance of PDOStatement, boolean given, called in /data/vendor/hyperf/database/src/Connection.php on line 294(0) in /data/vendor/hyperf/database/src/Connection.php:977
Stack trace:
#0 /data/vendor/hyperf/database/src/Connection.php(294): Hyperf\Database\Connection->prepared(false)
#1 /data/vendor/hyperf/database/src/Connection.php(1079): Hyperf\Database\Connection->Hyperf\Database\{closure}('select id, type...', Array)
#2 /data/vendor/hyperf/database/src/Connection.php(1044): Hyperf\Database\Connection->runQueryCallback('select id, type...', Array, Object(Closure))
#3 /data/vendor/hyperf/database/src/Connection.php(301): Hyperf\Database\Connection->run('select id, type...', Array, Object(Closure))
#4 /data/vendor/hyperf/database/src/Query/Builder.php(2670): Hyperf\Database\Connection->select('select id, type...', Array, true)
#5 /data/vendor/hyperf/database/src/Query/Builder.php(1838): Hyperf\Database\Query\Builder->runSelect()
#6 /data/vendor/hyperf/database/src/Query/Builder.php(2810): Hyperf\Database\Query\Builder->Hyperf\Database\Query\{closure}()
#7 /data/vendor/hyperf/database/src/Query/Builder.php(1839): Hyperf\Database\Query\Builder->onceWithColumns(Array, Object(Closure))
#8 /data/app/Controller/DbController.php(154): Hyperf\Database\Query\Builder->get()
#9 /data/vendor/hyperf/http-server/src/CoreMiddleware.php(103): App\Controller\DbController->aftersale(Object(Hyperf\HttpServer\Request), Object(Hyperf\HttpServer\Response))
#10 /data/vendor/hyperf/http-server/src/CoreMiddleware.php(77): Hyperf\HttpServer\CoreMiddleware->handleFound(Array, Object(Hyperf\HttpMessage\Server\Request))
#11 /data/vendor/hyperf/dispatcher/src/AbstractRequestHandler.php(66): Hyperf\HttpServer\CoreMiddleware->process(Object(Hyperf\HttpMessage\Server\Request), Object(Hyperf\Dispatcher\HttpRequestHandler))
#12 /data/vendor/hyperf/dispatcher/src/HttpRequestHandler.php(27): Hyperf\Dispatcher\AbstractRequestHandler->handleRequest(Object(Hyperf\HttpMessage\Server\Request))
#13 /data/app/Middleware/AuthMiddleware.php(33): Hyperf\Dispatcher\HttpRequestHandler->handle(Object(Hyperf\HttpMessage\Server\Request))
#14 /data/vendor/hyperf/dispatcher/src/AbstractRequestHandler.php(66): App\Middleware\AuthMiddleware->process(Object(Hyperf\HttpMessage\Server\Request), Object(Hyperf\Dispatcher\HttpRequestHandler))
#15 /data/vendor/hyperf/dispatcher/src/HttpRequestHandler.php(27): Hyperf\Dispatcher\AbstractRequestHandler->handleRequest(Object(Hyperf\HttpMessage\Server\Request))
#16 /data/app/Middleware/HttpLogMiddleware.php(17): Hyperf\Dispatcher\HttpRequestHandler->handle(Object(Hyperf\HttpMessage\Server\Request))
#17 /data/vendor/hyperf/dispatcher/src/AbstractRequestHandler.php(66): App\Middleware\HttpLogMiddleware->process(Object(Hyperf\HttpMessage\Server\Request), Object(Hyperf\Dispatcher\HttpRequestHandler))
#18 /data/vendor/hyperf/dispatcher/src/HttpRequestHandler.php(27): Hyperf\Dispatcher\AbstractRequestHandler->handleRequest(Object(Hyperf\HttpMessage\Server\Request))
#19 /data/vendor/hyperf/dispatcher/src/HttpDispatcher.php(43): Hyperf\Dispatcher\HttpRequestHandler->handle(Object(Hyperf\HttpMessage\Server\Request))
#20 /data/vendor/hyperf/http-server/src/Server.php(103): Hyperf\Dispatcher\HttpDispatcher->dispatch(Object(Hyperf\HttpMessage\Server\Request), Array, Object(Hyperf\HttpServer\CoreMiddleware))
#21 {main}
```

## 排查第一步: 源码

一般报错, 都发生在自己写的代码里, 这样会形成一个心理(这里面隐藏着一个 **二八法则**, 不过多展开, 感兴趣可以继续翻书 -- **墨菲定理**):

- 自己写的代码出错更常见 -> 解决的更多 -> 心理上会感觉更轻松
- 框架层的代码出错较少见 -> 解决的较少 -> 心理上会感觉更困难

> 告诉自己, 都是 PHP 代码, 有什么难的?! PHP is best !

数据库代理服务基于微服务框架 [hyperf](http://doc.hyperf.io) 来实现.

到了框架层, 代码往往耦合较少, 结构拆分很清晰, 虽然调用看起来很多, 但是核心代码就是 `trace#1` 的地方:

```php
// 原函数
public function select(string $query, array $bindings = [], bool $useReadPdo = true): array
{
    return $this->run($query, $bindings, function ($query, $bindings) use ($useReadPdo) {
        if ($this->pretending()) {
            return [];
        }

        // For select statements, we'll simply execute the query and return an array
        // of the database result set. Each element in the array will be a single
        // row from the database table, and will either be an array or objects.
        $statement = $this->prepared($this->getPdoForSelect($useReadPdo)
            ->prepare($query));

        $this->bindValues($statement, $this->prepareBindings($bindings));

        $statement->execute();

        return $statement->fetchAll();
    });
}
```

继续抽丝剥茧:

```php
// trace 中有行号
$statement = $this->prepared($this->getPdoForSelect($useReadPdo)
    ->prepare($query));

// 根据 exception message 进行确定范围
$statement = $this->prepared(false); // 报错来自这里
$this->getPdoForSelect($useReadPdo)
    ->prepare($query); // 这行代码返回了 false

// 这行代码等效于
$pdo->prepare($query);
```

这是关键的一步, 报错的来自 [Pdo::prepare](https://www.php.net/manual/en/pdo.prepare.php)

## 排查第二步: 查

果然, 我们不太可能成为那个只有 70亿(地球人口)分之一的幸运儿, 这个坑果然有不少人踩过, [Stack Overflow](https://stackoverflow.com/questions/3671237/php-pdo-prepare-in-a-function-returns) 有人提了相同的问题.

查文档, Stack Overflow 里给的回答, 就来自官方的文档 [Pdo::prepare](https://www.php.net/manual/en/pdo.prepare.php).

查的关键词:
- 查搜索引擎: 百度/谷歌
- 查文档

## 排查第三部: 加日志

目前只知道 `pdo->prepare()` 返回了 false, 还需要更多信息. 

> 怎么获得更多信息? 加日志!


```php
Log::get('sql')->info($query);
try {
    $pdo = $this->getPdoForSelect($useReadPdo);
    Log::info(var_export($pdo, true));
    $r = $pdo->prepare($query);
    Log::info(var_export($r, true));
    Log::info('errCode: '. $pdo->errorCode() . '|errInfo: '. json_encode($pdo->errorInfo()));
} catch (\Throwable $exception) {
    Log::get('sql')->info(format_throwable($exception));
}

$statement = $this->prepared($this->getPdoForSelect($useReadPdo)
```

加上日志后:

```
errCode: 00000|errInfo: ["00000",null,null]

false

PDO::__set_state(array(
))

select id,aftersale_id from `aftersale_step` where `aftersale_id` = ? limit 2
```

除了拿到 `$query` 的值以外, 好像没有拿到更有用的信息.

## 排查第四步: 交流

单打独斗许久之后, 尤其是打了日志还没拿到有用信息后, 确实有点 **没头脑**. 这个时候:

- 不要放弃, 拖着拖着, 可能就真的放弃了
- 集思广益: 和技术团队交流, 和技术社区交流

交流的好处:

- 更多的尝试, 更多的突破口
- 更多的知识, 更多技术细节

## 科学方法论: 找不同

正常态 -> 异常态, 而且还是必现, 那么肯定有 `固定原因`, 这个时候抛弃 `量子跃迁` `见鬼了` 等等想法, 选择 **科学方法**:

> 科学实验的方法: 控制变量法. 换言之, 找不同.

明显的不同, 环境不一样:

- 测试环境是好的: 测试环境使用的 rds(读写) + drds(读写)
- 线上有问题: 线上使用正式的 rds(读写+只读) + drds(读写+只读)

添加测试代码来比较不同(**方法来自于社区**):

```php
$dsn = 'xxx'; // 分别使用线上的使用的链接信息
$pdo = new \PDO("mysql:host={$dsn};dbname=xxx", 'xxx', 'xxx');
$sql = 'xxx'; // 使用日志中打出的 query
$stmt = $pdo->prepare($sql);
var_dump($stmt);
```

- 测试代码正常返回 `PDOStatement` 对象, 不会返回 false

现在写出来, 只有关键的 2 点, 实际排查过程其实走了很多弯路, mark 一下, 引以为戒!

## 解决

既然有了 **科学的方法**, 那么就可以大胆的得出可靠的结局:

- 环境的锅!!! 和 aliyun drds 技术人员确认, drds只读实例暂不支持 `mysql prepare`
- 测试代码表现和框架不一致, PDO 一定有配置控制相关的表现

框架层基于 laravel ORM, 默认覆盖了 PDO 的一些属性(由 [hyperf 社区](http://doc.hyperf.io) 提供):

```php
// vendor/hyperf/database/src/Connectors/Connector.php
protected $options = [
    PDO::ATTR_CASE => PDO::CASE_NATURAL,
    PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
    PDO::ATTR_ORACLE_NULLS => PDO::NULL_NATURAL,
    PDO::ATTR_STRINGIFY_FETCHES => false,
    PDO::ATTR_EMULATE_PREPARES => false,
];
```

很可能就是 `PDO::ATTR_EMULATE_PREPARES' 属性, 使用测试代码验证:

```php
$dsn = 'xxx'; // 分别使用线上的使用的链接信息
$pdo = new \PDO("mysql:host={$dsn};dbname=xxx", 'xxx', 'xxx', [PDO::ATTR_EMULATE_PREPARES => false,]);
$sql = 'xxx'; // 使用日志中打出的 query
$stmt = $pdo->prepare($sql);
var_dump($stmt);
```

- 测试代码果然返回 false

## 写在最后

梳理涉及到的技术知识:

- (prepare sql: mysql prepare 协议使用说明](https://help.aliyun.com/document_detail/71326.html)
- php 使用 PDO 访问 mysql, 可以通过 `pdo->prepare()` 和 mysql prepare 协议交互
- PDO 有很多属性可以设置, 包括 `prepare()` 时的行为: `PDO::ATTR_EMULATE_PREPARES'

总结重要的几点:

- 单测很重要, 上线后跑 phpunit, 立刻就发现了问题, 然后马上开始填坑
- 心理很重要: 不要怕问题, `都是 PHP 代码, 有什么好怕的`
- 科学方法很重要: 看似做了 **各种尝试**, 但是没有科学的方法支撑, 反而在获取到越来越多的信息后, 更容易迷茫, 不敢下结论
- 事上练: 增加和周围世界的联系, 技术也可以做到, 多和 团队/社区 交流想法和知识

历史类似经历:

- [alipay ILLEGAL_SIGN 错误解决](https://www.jianshu.com/p/28585a6454b2)