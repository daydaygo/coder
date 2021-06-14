# cloud| 阿里云日志服务实践

- description: 阿里云日志服务实践: nginx access log; yii 框架接入阿里云日志服务

> [devops| 日志服务实践](https://www.jianshu.com/p/9dae3ba679e6)

> [技术分享 - devops| 日志服务实践](https://c.dayday.tech/landslide/TS20180423.html)

提纲:

1. 日志服务功能点一览
2. 阿里云日志服务实践
3. 示例一: nginx access log
4. 示例二: yii 框架接入阿里云日志服务
5. 再探 protobuf

日志服务可以说的上是构建软件项目的基石之一, 系统持续稳定运行必不可少的一部分. 这里从阿里云日志服务入手, 借助云平台带来的技术更新迭代, 聊一聊日志服务实践.

## 以前日志服务的打开方式

日志最常见的方式是写入到文件中. 「小作坊」的情况下, 把服务器的权限给开发, 开发自己 ssh 到服务器上面用 `grep` 查日志. 是的, 我就是这样过来的, 所以常用的几个 grep 命令, 甚至一些稍微高级的命令, 还能默写出来:

```sh
grep xxx xxFile # 正则匹配查询字符串
grep 'xxx xxx' # 查询包含特殊字符, 比如空格的字符串
grep -i xxx # 忽略大小写
grep -n xx # 显示行号
grep -v xxx # 查询不包含字符串的行
grep -r xxx xxDir # 在文件夹中递归查询
ps aux | grep xxx | grep -v 'grep' # -v 常用的一种方式

# 2个复杂些的例子
# 获取访问 ip 统计
cat /var/log/nginx/access.log|awk '{print $1}'|sort|uniq -c|sort -nr|more
# 获取 http 状态码
cat /var/log/nginx/access.log|grep -ioE 'HTTP/1.[0|1]" [0-9]{3}'|awk '{print $2}'
```

grep 查询可以使用多种 **正则** 方式: 基础, 扩展, perl. 支持的正则功能一次增多, 部分细节有些许差异.

```sh
-E, --extended-regexp     PATTERN is an extended regular expression (ERE)
-G, --basic-regexp        PATTERN is a basic regular expression (BRE)
-P, --perl-regexp         PATTERN is a Perl regular expression
```

一句话概括这种方式: 简单直接. 当然有时直接 `vim` 打开, 然后再查看的. 不过数据量一大, `vim` 的速度就不乐观了. 所以通常会对日志文件进行 **切分**, 这样也便于以后 **归档**:

- 按照业务切分: 服务器各项日志, 不同业务模块的日志, 第三方接口的日志
- 按照时间维度: 日切 `xxx_20180421.log`; 月切 `xxx_201804.log`
- 按照文件大小: 比如到 1G 了, 从 `xxx.log.1` 到 `xxx.log.2`, 一次递增
- 多种方式组合使用

文件一多, 查询就变得困难起来了.

数据量大, 还要考虑日志的 **写入性能**, 通常的做法是 **加缓存**: 这里称之为 **刷新(flash)**:

- 一定时间间隔写入一次
- 日志达到多少条写入一次
- 日志超过多大写入一次

开放服务器 ssh 权限出来, 会带来 **安全隐患**, 有开发上去误操作就不好了. 所以有了新的替代方案:

- 运维开日志的 ftp, 需要看日志自己去下载
- 存储到数据库中, 比如 MongoDB, 走数据库查询
- 自建日志中心

当然 **自建日志中心** 是最高级的玩法, 之前鹅厂的分享提到过, 会走 **UDP** 进行日志的上传与统一分析.

关于日志的其他细节:

- 全链路跟踪: 去年很火的方案, 请求进来时生成一个 `trace_id`, 之后的所有调用都会带上这个 `trace_id`, 这样就可以在日志中通过 `trace_id` 查询到整个调用链路
- 安全问题: 日志中可能有用户未经处理的敏感信息, 比如手机号, 甚至没有经过处理的密码
- 日志归档的问题: 打包归档历史数据来降低日志存储的成本

最后, 对大部分使用日志的人(通常是开发, 定位 **犯罪现场**)而言, **好查** **好用** 才是重中之重, 日志的存储/归档都不用自己操心, 由日志系统来解决.

## 上手阿里云日志服务

阿里云的日志服务上手比较容易, 在控制台点点点即可, 大致的分层设计如下:

- 开通日志服务: 总的入口
- project: 项目, 第一级分层, `project + region` 构成 api 的访问地址
- logStore: 日志存储, 每个 project 下可以建立多个 logstore, logStore 可以配置多个 `shared` 和 **过期时间**
- log data source: 需要为每个 logStore 配置数据源
- 日志投递: 日志数据除了供日志服务消费外, 还可以投递给其他云产品, 比如 **OSS 进行归档处理**
- 日志查询: 重点功能, 包含 `search` `analysis` `chart` 3个主要部分

配置数据源常用的方式:

- nginx access log: 下面还会详细提到
- 文本 + logtail 工具收集 + 自定义日志分割
- sdk 接入

关于 logtail: 阿里云提供的日志收集工具, 安装到 ecs 上就可以按照 logStore 配置的日志路径进行搜集

PS: 如果 ecs 和 日志服务是不同的账号下的, [需要配置授权](https://help.aliyun.com/document_detail/49007.html)

日志查询快捷指南:

- search: 支持部分正则的查询语法; 直接点击日志即可查询
- analysis: 使用 `|` 管道对查询结果进行分析; 类 sql 的语法
- chart: 将 analysis 得到的结果转化为图表, 更直观
- 其他小技巧: 保存常用查询

![日志数据源配置](https://upload-images.jianshu.io/upload_images/567399-cf9dca145ac49d00.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 实践一: nginx access log

nginx access log 的接入提供了很好的支持:

- 配置好 logtail 收集 access log
- logStore 中配置 `log_format` 就可以自动分割日志建立索引
- 配置常用 查询/分析/图表

推荐下面的 `log_format`:

```
log_format main '$remote_addr||$remote_user||$time_local||$request||$http_host||$status||$request_length||$body_bytes_sent||$http_referer||$http_user_agent||$request_time||$upstream_response_time||$request_body';
```

PS: **细节出魔鬼**, 之前没有采用 `||` 的方式, 导致部分日志解析出现问题, 字段没有对上

日志记录:

![nginx-access-log 示例](https://upload-images.jianshu.io/upload_images/567399-f1b6df0e3be0139a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


配置好图表:

![image](http://upload-images.jianshu.io/upload_images/567399-968c2c5894fde72c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```sql
-- access log 报表
-- pv/uv
* | select approx_distinct(remote_addr) as uv, count(1) as pv, date_format(date_trunc('hour', __time__), '%m-%d %H:%i') as time group by date_format(date_trunc('hour', __time__), '%m-%d %H:%i') order by time limit 1000
-- ip distribution
* | select count(1) as c, ip_to_province(remote_addr) as address group by ip_to_province(remote_addr) limit 100
-- top page
* | select split_part(request_uri,'?',1) as path, count(1) as pv group by split_part(request_uri,'?',1) order by pv desc limit 10
-- top refer
* | select count(1) as pv, http_referer group by http_referer order by pv desc limit 10
-- http method
 * | select count(1) as pv, request_method group by request_method
-- http status
* | select count(1) as pv, status group by status
-- UserAgent
* | select count(1) as pv, case when http_user_agent like '%Chrome%' then 'Chrome'when http_user_agent like '%Firefox%' then 'Firefox'when http_user_agent like '%Safari%' then 'Safari'else 'unKnown' end as http_user_agent group by  http_user_agent order by pv desc limit 10
-- latency
* | select from_unixtime(__time__ -__time__% 300) as time, avg(request_time) as avg_latency , max(request_time) as max_latency group by __time__ -__time__% 300
* | select from_unixtime(__time__ - __time__% 60) , max_by(request_uri,request_time) group by __time__ - __time__%60
* | select numeric_histogram(10,request_time)
* | select max(request_time,10)
request_uri:"/url2" | select count(1) as pv, approx_distinct(remote_addr) as uv, histogram(method) as method_pv, histogram(status) as status_pv, histogram(user_agent) as user_agent_pv, avg(request_time) as avg_latency, max(request_time) as max_latency
```

关于 `request_body`:

- 推荐加到 `log_format` 中
- 当前 nginx 版本(我的是 1.13) 直接配置上就可以收集 `form-data x-www-form-urlencoded application/json` 等格式的 post 数据

如何解析 `form-data` 格式的数据:

```php
function hextostr($hex) {
    return preg_replace_callback('/\\\x([0-9a-fA-F]{2})/', function($matches) {
        return chr(hexdec($matches[1]));
    }, $hex);
}
echo hextostr('----------------------------400719531552868304622917\x0D\x0AContent-Disposition:');
```

如果 `request_body` 无法记录, 网上提供了 2 种方案(**当前版本并不需要**):

- 将 `access log` 记录到 `fastcgi_pass` 配置处

```conf
location ~ \.php$ {
    fastcgi_pass fpm:9000;
    fastcgi_index index.php;
    fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    access_log /var/log/nginx/yii_access.log main;
    include fastcgi_params;
}
```

- 使用 `nginx_lua` 模块

```sh
lua_need_request_body on;
content_by_lua 'local s = ngx.var.request_body';
```

## 实践二: 接入日志服务sdk(以 yii 框架为例)

使用 logtail 来作为数据源实在是 **简单**, 搜集的数据通过 **分隔符** 或者 **正则** 进行分割, 有时候会很麻烦. 写过 **正则** 的人都知道, 正则这东西并不难, 它就是 **烦**, 稍微有一点点变动, 正则可能就需要调整了. 而且用 logtail 来收集日志, 看似和业务 **比较隔离**, 实际感觉确实 **偏离业务** 更多一点. 接入 **日志服务sdk** 会是一个不错的选择.

在之前 [yii| 最佳实践之黑箱思维](https://www.jianshu.com/p/8f52b75055d5) 提到过 yii 的日志服务, 这里再简要复述一下:

- 分层设计: `logger - dispatch - target` 3层, `logger` 专注于日志功能, `dispatch` 来调度, `target` 来适配不同日志存储
- 日志标记: 包括 level / category / tag / perfix 等多种日志标记方式, 方便对日志更细粒度的控制
- flush, 刷新, 比如 1000 条后再输出(落地). **缓冲(buffer)** 的思想可以说在系统设计中比比皆是.
- 切片, 比如说日志按照时间日切, 或者按照大小 10m 一切

> PS: 这就是成熟框架的威力, 常用功能近乎全面无死角的解决掉

具体 yii 中接入, 其实就是新增一个 target, 通过阿里云日志服务SDK写入日志:
- 引入日志服务SDK: https://gitee.com/daydaygo/yii/blob/master/common/config/bootstrap.php

```php
<?php
Yii::setAlias('@common', dirname(__DIR__));
Yii::setAlias('@frontend', dirname(dirname(__DIR__)) . '/frontend');
Yii::setAlias('@backend', dirname(dirname(__DIR__)) . '/backend');
Yii::setAlias('@console', dirname(dirname(__DIR__)) . '/console');

require __DIR__ . '/../sdk/aliyun-log-php-sdk/Log_Autoload.php';
```

- 封装日志服务sdk到 component: https://gitee.com/daydaygo/yii/blob/master/common/components/AliyunLog.php

```php
?php
namespace common\components;

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
    public $endPoint = 'cn-shanghai-intranet.log.aliyuncs.com';
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
        $this->client = new \Aliyun_Log_Client(
            $this->endPoint,
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

- 添加 AliyunlogTarget: https://gitee.com/daydaygo/yii/blob/master/common/components/AliyunLogTarget.php

```php
<?php
namespace common\components;

use yii\di\Instance;
use yii\helpers\VarDumper;
use yii\log\Logger;
use yii\log\Target;

class AliyunLogTarget extends Target
{
    /** @var AliyunLog $log */
    public $log = 'aliyunLog';
    public $project;
    public $logStore;
    public $topic;

    public function init()
    {
        $this->log = Instance::ensure($this->log);
    }

    public function export()
    {
        $rows = [];
        foreach ($this->messages as $message) {
            list($text, $level, $category, $timestamp) = $message;
            $level = Logger::getLevelName($level);
            if (!is_string($text)) {
                // exceptions may not be serializable if in the call stack somewhere is a Closure
                if ($text instanceof \Throwable || $text instanceof \Exception) {
                    $text = (string) $text;
                } else {
                    $text = VarDumper::export($text);
                }
            }
            $rows[] = [
                'level' => $level,
                'category' => $category,
                'prefix' => $this->getMessagePrefix($message),
                'message' => $text,
            ];
        }

        if ($this->project) {
            $this->log->project = $this->project;
        }
        if ($this->logStore) {
            $this->log->logStore = $this->logStore;
        }
        if ($this->topic) {
            $this->log->topic = $this->topic;
        }
        $this->log->putLogs($rows);
    }
}
```

- 最后, 把日志服务配置到想要的地方: https://gitee.com/daydaygo/yii/blob/master/console/config/main.php

```php
...
    'components' => [
        'log' => [
            'targets' => [
                [
                    'class' => 'yii\log\FileTarget',
                    'levels' => ['error', 'warning'],
                ],
                [
                    'class' => \common\components\AliyunLogTarget::class,
                    'levels' => ['info', 'warning', 'error'],
                    'except' => $_info_except,
                    'logVars' => [],
                    'exportInterval' => YII_ENV_PROD ? 1000 : 1,
                    'topic' => 'console',
                ],
            ],
        ],
    ],
...
```

来张效果图:

![aliyun-log-yii](https://upload-images.jianshu.io/upload_images/567399-db7f171b8d3d9c77.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 题外: 什么是好的SDK

在 [yii| 最佳实践之黑箱思维](https://www.jianshu.com/p/8f52b75055d5) 里我还提到如何判断 **好的sdk**:

> 好用的 SDK, 只用看一下 sample 或者 quick start 就能分辨出来.

不过这次实践下来, 我要收回这句话, 从 「进化论」 的角度来看才更趋于真理:

> 好用的 sdk, 应该是能跟上社区最佳标准与实践, 不断进化的.

说一下项目实践中遇到的问题: 同时使用 阿里云日志服务和OSS服务的sdk, 而 2 这的 sdk 中都定义了 `RequestCore` 来作为 http 请求基类, 导致类冲突

- 日志服务的 sdk 是自己使用 `spl_autoload_register()` 来做类的自动加载: https://github.com/aliyun/aliyun-log-php-sdk/blob/master/Log_Autoload.php

```php
<?php
/**
 * Copyright (C) Alibaba Cloud Computing
 * All rights reserved
 */
$version = '0.6.0';
function Aliyun_Log_PHP_Client_Autoload($className) {
    $classPath = explode('_', $className);
    if ($classPath[0] == 'Aliyun') {
        if(count($classPath)>5)
            $classPath = array_slice($classPath, 0, 5);
        if(strpos($className, 'Request') !== false){
            $lastPath = end($classPath);
            array_pop($classPath);
            array_push($classPath,'Request');
            array_push($classPath, $lastPath);
        }
        if(strpos($className, 'Response') !== false){
            $lastPath = end($classPath);
            array_pop($classPath);
            array_push($classPath,'Response');
            array_push($classPath, $lastPath);
        }
        $filePath = dirname(__FILE__) . '/' . implode('/', $classPath) . '.php';
        if (file_exists($filePath))
            require_once($filePath);
    }
}
spl_autoload_register('Aliyun_Log_PHP_Client_Autoload');
```

- 只能按照 demo 中 require 此 autoload 文件才能使用: https://github.com/aliyun/aliyun-log-php-sdk/blob/master/sample/sample.php#L7

```php
require_once realpath(dirname(__FILE__) . '/../Log_Autoload.php');
```

本来只是想对现有日志功能进行改造, 要是导致原有的 OSS 功能不能用了, 那就不好了. 基于此, 就动了直接接入日志服务 api 的念头:

-  参考官方文档一步一步接的示例: https://gitee.com/daydaygo/yii/blob/master/console/controllers/TController.php#L49

```php
public function actionAliyunlog2()
{
    $ak = 'bq2sjzesjmo86kq35behupbq';
    $sk = '4fdO2fTDDnZPU/L7CHNdemB2Nsk=';

    // 服务入口: https://help.aliyun.com/document_detail/29008.html
    $project = 'test-project';
    $endpoint = 'cn-hangzhou-devcommon-intranet.sls.aliyuncs.com';

    // 请求签名: https://help.aliyun.com/document_detail/29012.html
    // get
    $httpMethod = 'GET';
    $contentMd5 = '';
    $contentType = '';
    $gmDate = 'Mon, 09 Nov 2015 06:11:16 GMT';
    $logHeaders = [
        'x-log-apiversion:0.6.0',
        'x-log-signaturemethod:hmac-sha1',
    ];
    $logHeadersStr = join("\n", $logHeaders);
    $logResource = '/logstores?' . http_build_query(['logstoreName' => '', 'offset' => 0, 'size' => 1000]);
    $signStr = $httpMethod . "\n" . $contentMd5 . "\n" . $contentType . "\n" . $gmDate . "\n" .
        $logHeadersStr . "\n" . $logResource;
    $sign = base64_encode(hash_hmac('sha1', $signStr, $sk, true));

    // 公共请求头: https://help.aliyun.com/document_detail/29010.html
    $headers = [
        "Date: $gmDate",
        "Host: {$project}.{$endpoint}",
        "Authorization:LOG {$ak}:{$sign}",
    ];
    $headers = array_merge($headers, $logHeaders);

    // post
    // 数据编码方式 - protobuf: https://help.aliyun.com/document_detail/29055.html
    $body = [
        'TestKey' => 'TestContent',
    ];
    $contents = [];
    foreach ($body as $k => $v) {
        $content = new \Protobuf\Aliyunlog\Log_Content();
        $content->setKey($k);
        $content->setValue($v);
        $contents[] = $content;
    }
    $log = new \Protobuf\Aliyunlog\Log();
    $log->setTime(1447048976);
    $log->setContents($contents);
    $logGroup = new \Protobuf\Aliyunlog\LogGroup();
    $logGroup->setLogs([$log]);
    $logGroup->setTopic('');
    $logGroup->setSource('10.230.201.117');
    $bodyProto = $logGroup->serializeToString();

    $httpMethod = 'POST';
    $contentMd5 = strtoupper(md5($bodyProto));
    $contentType = 'application/x-protobuf';
    $contentLen = strlen($bodyProto);
    $gmDate = 'Mon, 09 Nov 2015 06:11:16 GMT';
    $logHeaders = [
        'x-log-apiversion:0.6.0',
        'x-log-signaturemethod:hmac-sha1',
    ];
    $logHeadersStr = join("\n", $logHeaders);
    $logResource = '/logstores?' . http_build_query(['logstoreName' => '', 'offset' => 0, 'size' => 1000]);
    $signStr = $httpMethod . "\n" . $contentMd5 . "\n" . $contentType . "\n" . $gmDate . "\n" .
        $logHeadersStr . "\n" . $logResource;
    $sign = base64_encode(hash_hmac('sha1', $signStr, $sk, true));

    $headers = [
        "Date: $gmDate",
        "Host: {$project}.{$endpoint}",
        "Authorization:LOG {$ak}:{$sign}",
        "Content-MD5: $contentMd5",
        "Content-Length: $contentLen",
    ];
    $headers = array_merge($headers, $logHeaders);
}
```

事实证明 **我还是太年轻了, 日志传输用的 protobuf**. 这东西说实话并不难, 之前的服务器系列有protobuf 的入门使用([blog - 服务器开发系列 1](https://www.jianshu.com/p/1633fa196c43)), 无非是安装一个 protobuf 的编译器(protoc), 然后安装一个protobuf的解析器(对应 php 中的 ext-protobuf 扩展)

- 但是阿里云官方是 proto2 syntax 语法: https://gitee.com/daydaygo/yii/blob/master/common/protobuf/aliyunlog.proto2

```protobufbuf
message Log
{
    required uint32 time = 1; // UNIX Time Format
    message Content
    {
        required string key = 1;
        required string value = 2;
    }
    repeated Content contents= 2;
}
message LogGroup
{
    repeated Log logs= 1;
    optional string reserved =2; // 内部字段，不需要填写
    optional string topic = 3;
    optional string source = 4;
}
message LogGroupList
{
    repeated LogGroup logGroupList = 1;
}
```

- 而 protoc 用来生成编译生成PHP文件的命令 `protoc -php_out=./ aliyunlog.protoc` 只支持 proto3 syntax: https://gitee.com/daydaygo/yii/blob/master/common/protobuf/aliyunlog.proto

```protobufbuf
syntax="proto3";
package Protobuf.Aliyunlog;

message Log
{
    uint32 time = 1; // UNIX Time Format
    message Content
    {
        string key = 1;
        string value = 2;
    }
    repeated Content contents= 2;
}
message LogGroup
{
    repeated Log logs= 1;
    string reserved =2; // 内部字段，不需要填写
    string topic = 3;
    string source = 4;
}
message LogGroupList
{
    repeated LogGroup logGroupList = 1;
}
```

导致的结果就是, protobuf序列化的数据大小, 和 demo 对上, api 自然就不通了. 而官方 SDK 中, 是用 `pack()` 自己一点点实现的. 这事我在刚接触服务器开发的时候也干过...

不过好在, OSS的SDK按照 psr-4 标准进行组织了, 引入命名空间后就不会有现在类冲突的尴尬了.

## 写在最后

日志服务是一个深究起来还颇为复杂的话题, 重要实践, 让日志真正起到 **系统保驾护航** 和 **异常时还原犯罪现场的作用**

推荐资源:

- [日志服务 - 最佳实践 - nginx访问日志](https://help.aliyun.com/document_detail/56728.html)
- [日志服务帮助文档](https://help.aliyun.com/document_detail/56728.html?spm=a2c4g.11174283.6.559.bzC4Rp)
- [为了无法计算的价值](https://promotion.aliyun.com/ntms/act/ambassador/sharetouser.html?userCode=3fxolx6j&utm_source=3fxolx6j)
