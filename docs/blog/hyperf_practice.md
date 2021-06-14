# hyperf| 编码实践

整个团队使用 hyperf 开发已经超过半年, 积累了一些最佳实践和规约, 方便团队后续开发, 提供给大家参考~

## hyperf 编程注意事项

禁止项

- `$path = $_SERVER['HTTP_REFERER'];` 不可以使用超全局变量, swoole server 是常驻进程 epoll 方式处理请求, 不是 fpm 依靠多进程一对一处理请求

配置管理 config

- 所有配置都在 config 文件中, 项目中只使用 config() 就可以获取配置
- .env 只写环境相关配置, 最后都需要使用 env() 配置到 config 文件中, env() 只允许在 config 文件中使用

修改 PHP 配置

- `container.php` 保持和官方同步, `int_set() define()` 等可以配置到 hyperf.php 文件中

容器 container

- 尽量使用框架提供的功能, 这样才能发挥出 container 的效果
- 手动 new 的对象, 没有交给 container 管理, 无法使用框架提供的注解功能, 比如 `@Inject`

协程 go/coroutine

- 使用协程有2个必须条件: 开协程 `go() + callback 中执行非阻塞代码`
- 常用的库都提供了非阻塞调用: mysql(基于laravel ORM + 连接池), redis(ext-redis + 连接池), http(guzzle clientFactory)

http client

- http 请求这样的功能很常见, 几乎所有的项目都会自己的封装, 但是要使用 hyperf 提供的 `hyperf/guzzle`, 才能实现 `非阻塞调用`
- 项目自己封装的 http 请求库, 基本可以断定没有 `guzzle/http` 好用

基于 hyperf 的封装
基于 hyperf 的封装 `务必` 要考虑 `反哺` 社区, 有以下几点:

- hyperf 的迭代速度, 包括功能支持上, 速度有目共睹, 所谓的封装, 也许就在下次发版支持了
- hyperf 的代码提交以及整个社区共建, 保证了代码风格和相当高的 `代码质量`
- `反哺社区` 也是 `依赖解决` 的最优解之一
- 最后再来一句灵魂拷问: `自己再封装一层真的有必要么?`

建议项:

- 大段注释的代码基本建议删掉, 哪怕真的有需要, 也可以从 git log 里找回, 我认为这是 `clean code` 的要求之一
- 不建议使用常量, hyperf 只设置了一个常量 `BASE_PATH`, 使用常量通常是没想好要 `怎么放`, 建议多看看 hyperf 的组件设计

## example

```php
<?php

declare(strict_types=1);

namespace App\Command;

use App\Amqp\Producer\TestProducer;
use App\Model\Logistic;
use Carbon\Carbon;
use Hyperf\Amqp\Producer;
use Hyperf\Command\Command as HyperfCommand;
use Hyperf\Command\Annotation\Command;
use Hyperf\DbConnection\Db;
use Hyperf\Logger\Logger;
use Hyperf\Logger\LoggerFactory;
use Mt\Util\Log;
use PDO;

/**
 * @Command
 */
class HyperfDemoCommand extends HyperfCommand
{
    protected $name = 'hd';

    public function configure()
    {
        parent::configure();
        $this->setDescription('Hyperf Demo Command');
    }

    public function handle()
    {
        $this->line('Hello Hyperf!', 'info');
    }

    /**
     * 生命周期
     * @link https://doc.hyperf.io/#/zh-cn/lifecycle
     */
    public function lifecycle()
    {
        // 生命周期分析法是非常常用的一种分析手法, 可以帮忙快速理解一个新系统
        // 分析方式一: 流程图, 入口->出口, 一次 http 请求经过了哪些步骤?
        // 分析方式二: 分层, 分解/拆分问题

        // hyperf 生命周期: container application swoole http coroutine
    }

    /**
     * 协程
     * @link https://www.jianshu.com/p/12d645ac02b2 swoole-wiki 笔记, 全面梳理 swoole 基础知识
     * @link https://doc.hyperf.io/#/zh-cn/coroutine
     */
    public function coroutine()
    {
        // demo: basic
        go(function () {
            sleep(1);
            echo 'go' . PHP_EOL;
        });
        echo 'main' . PHP_EOL;

        // demo: 性能测试
        // 使用命令行 time 命令查看耗时, 可以看到 sys/user 分别耗时
        // 推荐使用 for 循环, 明确使用的协程的数量, 避免把下游服务(db/第三方接口)打爆
        for ($i = 0; $i < 10; $i++) {
            go(function () use ($i) {
                // do some **io task**
                sleep(1);
                echo "go $i \n";
            });
        }

        // swoole runtime
        // 暂时不要使用 SWOOLE_HOOK_CURL, curl api 目前只兼容了大部分常用的
        !defined('SWOOLE_HOOK_FLAGS') && define('SWOOLE_HOOK_FLAGS', SWOOLE_HOOK_ALL | SWOOLE_HOOK_CURL);

        // todo: 协程更多特性 demo
        // 特性一定要使用其 **最佳实践**, 不能为了使用而使用
    }

    /**
     * 配置
     * 配置哲学: 约定大于配置, 不必要的配置导致不必要的灵活性
     * @link http://deploy-dev.mengtuiapp.com/doc#/ms?id=config-%e6%a8%a1%e5%9d%97%e8%af%a6%e8%a7%a3
     * @link https://doc.hyperf.io/#/zh-cn/config
     */
    public function config()
    {
        // 推荐统一使用使用 config()
        var_dump(config('app_name'));

        // 不推荐使用 Config 对象重新设置配置

        // 只允许在配置文件中使用 env()
    }

    /**
     * 容器 container
     * 依赖注入 DI
     * @link https://doc.hyperf.io/#/zh-cn/di
     */
    public function container()
    {
        // Config 加载后, container 会根据 config 实例化, 并处理类之间的依赖关
        // container 在手, 天下我有
        // 封装了非常好用 container() 的方法, 推荐统一使用此方法

        // 想要 logger
        /** @var Logger $logger */
        $logger = container(LoggerFactory::class)->get('test');
        $logger->info('test');

        // 当然, 这样常用的类已经封装好了
        // 效果和上面等同
        Log::get('test')->info('test2');

        // 一个常犯的错误: new + @Inject
        // new 出来的类里面使用了注解, 那么这个类必须交给 container 管理, 注解才能生效
    }

    /**
     * 事件机制
     * @link https://doc.hyperf.io/#/zh-cn/event
     * @link https://doc.hyperf.io/#/zh-cn/event?id=%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9 事件中的循环依赖问题
     */
    public function listener()
    {
        // 典型场景: swoole event
        // \Hyperf\Server\SwooleEvent
        // config/autoload/server.php server中配置 swoole event 的回调

        // 典型场景: application event
        // \Mt\Listener\BootAppConfListener::process kms 处理等

        // 典型场景: db event
        // \Mt\Listener\DbQueryExecutedListener::process db 日志/监控都是在这里实现

        // 事件机制使用场景非常多, 并且可以有效扩展系统能力, 务必熟悉并掌握
    }


    /**
     * 路由
     * @link https://doc.hyperf.io/#/zh-cn/router
     */
    public function routes()
    {
        // 统一在 routes.php 中定义路由
        // apidog 组件已经设置 @AutoController 注解失效

        // 可以使用 route 文件, 定义更多路由
        // 实现: \Mt\Util\RoutesDispatcher

        // 还有一种简单的方式: 在 routes.php 中 require
    }

    /**
     * 中间件
     * @link https://doc.hyperf.io/#/zh-cn/middleware/middleware
     */
    public function middleware()
    {
        // 这里是狭义的中间件, 只处理 request->response
        // 执行顺序: 洋葱模型 全局->类级别->方法级别

        // demo1: http log, 记录 http 请求日志
        // \Mt\Middleware\HttpLogMiddleware

        // demo2: 鉴权
        // \Mt\Middleware\AuthMiddleware

        // demo3: sso, 单点登录
        // \Mt\Middleware\SsoMiddleware

        // demo4: 跨域中间件, 跨域配置也可以直接挂在 nginx 上
        // CorsMiddleware
    }

    /**
     * 控制器
     * @link https://doc.hyperf.io/#/zh-cn/controller
     */
    public function controller()
    {
        // 最佳实践: 封装好 BaseController, 约定好 response 数据格式等常用功能
        // \Mt\Util\AbstractController

        // 知识点一: swoole 会自动为每个请求分配好协程
        // 知识点二: 贡献数据要使用 **协程上下文(Context)**, 不可使用 类属性/类常量
    }

    /**
     * 请求
     * @link https://doc.hyperf.io/#/zh-cn/request
     */
    public function request()
    {
        // 统一使用 \Hyperf\HttpServer\Request 来处理请求
        // 或者基于 \Hyperf\HttpServer\Request 对象进行封装

        // psr-7 标准 api
        // \Hyperf\HttpServer\Request 提供: 请求路径 输入预处理 json cookie file

        // 注意: header 相关方法要注意返回值, 可能嵌套一层 array
    }

    /**
     * 响应
     * @link https://doc.hyperf.io/#/zh-cn/response
     */
    public function response()
    {
        // 统一使用 Hyperf\HttpServer\Response 对象
        // 响应格式以及自动设置 `content-type`: json xml raw view
        // 其他: redirect file 等
    }

    /**
     * 异常处理器
     * @link https://doc.hyperf.io/#/zh-cn/exception-handler
     */
    public function exceptionHandler()
    {
        // 如果 swoole worker 进程遇到未捕获的异常, 会导致进程退出, 所以必须要有异常处理器

        // demo1: 统一处理 user exception
        // \Mt\Handler\HttpExceptionHandler 输出日志并根据环境返回 response

        // demo2: error_reporting() 错误
        // Hyperf\ExceptionHandler\Listener\ErrorExceptionHandler
        // 还有一种方式: php bin/hyperf.php > runtime/run.log 中
    }

    /**
     * 日志
     * @link https://doc.hyperf.io/#/zh-cn/logger
     */
    public function log()
    {
        // 已经封装好了, 直接使用
        // 注意日志的结构: channel / level / message / context
        // 不同环境: dev 会直接打到 stdout, 其他环境打到日志文件->日志服务
        Log::get('test')->info('test');
    }

    /**
     * 命令行
     * @link https://doc.hyperf.io/#/zh-cn/command
     */
    public function command()
    {
        // 脚本有制作好的脚手架, 使用脚手架可以减少大量重复性的开发工作
        // mslib/core/Util/ScriptJob/AbstractScriptScaffold.php
        // doc/script_scaffold.md

        // 常见错误一: 脚本中有语法错误, 导致 `Command "hd" is not defined.`
//        echo 'error 1'

        // 常见错误二: 使用 amqp producer 报错
        // 这是因为 amqp 连接池在 command 执行完后没有比较好的方式进行关闭, 实际脚本逻辑是正常执行的
        /** @var Producer $producer */
        $producer = container(Producer::class);
        $producer->produce(new TestProducer(Carbon::now()));
    }

    /**
     * 单元测试
     * @link https://doc.hyperf.io/#/zh-cn/testing
     */
    public function test()
    {
        // 调试代码
        // 传统方式: 修改 -> 重启 server -> 浏览器/其他工具 测试接口
        // 单元测试: 通过配合 testing，来快速调试代码，顺便完成单元测试

        // 测试替身 mock
        // 有时候对 被测系统(SUT) 进行测试是很困难的，因为它依赖于其他无法在测试环境中使用的组件
    }

    /**
     * 视图
     * @link https://doc.hyperf.io/#/zh-cn/view
     */
    public function view()
    {
        // 不建议使用, 需要使用单独的进程完成视图的渲染
        // 前后端分离 + 组件化平台构建
    }

    /**
     * 验证器
     * @link https://doc.hyperf.io/#/zh-cn/validation
     */
    public function validation()
    {
        // 推荐使用, 可以将请求校验从 Controller 逻辑中解耦出来
    }

    /**
     * session
     * @link https://doc.hyperf.io/#/zh-cn/session
     */
    public function session()
    {
        // 兼容必须使用 session 的场景
        // 微服务中不建议使用
    }

    /**
     * db
     * @link https://doc.hyperf.io/#/zh-cn/db/quick-start
     * @link https://doc.hyperf.io/#/zh-cn/db/querybuilder
     * @link https://doc.hyperf.io/#/zh-cn/db/model
     */
    public function db()
    {
        // 配置最佳实践: 已 db 为粒度, 哪怕在同一个 db 实例上, 也分开进行配置
        // 读写分离: 优先使用 db 层的读写分离, 业务中显式使用 可写连接/只读连接
        // 不允许使用 default, 显式命名并使用

        // 连接池配置
        $pool = [
            'min_connections' => 1,
            'max_connections' => 10,
            'connect_timeout' => 10.0,
            'wait_timeout' => 3.0,
            // 心跳检查
            'heartbeat' => -1,
            'max_idle_time' => (float) env('DB_MAX_IDLE_TIME', 60),
        ];

        // pdo 配置
        $option = [
            // 框架默认配置
            PDO::ATTR_CASE => PDO::CASE_NATURAL,
            PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            PDO::ATTR_ORACLE_NULLS => PDO::NULL_NATURAL,
            PDO::ATTR_STRINGIFY_FETCHES => false,
            // 不支持 MySQL prepare 协议的 db, 需要配置为 true
            PDO::ATTR_EMULATE_PREPARES => false,
        ];

        // 事务: 必须显式指定连接
        $db = Db::connection('pt_logistic');

        // 自动管理事务
        $db->transaction(function () {
            // do something
        });

        // 手动管理事务
        $db->beginTransaction();
        try {
            // do something
            $db->commit();
        } catch (\Throwable $ex) {
            $db->rollBack();
        }

        // query builder
        // 所有连接从 Model 中获取, 不允许直接使用 Db::connection()
        // 查询结果, 统一转化为数组
        // 原则上不允许使用 join

        // 返回一行
        Logistic::query()->first();
        // 返回单个值
        Logistic::query()->value('id');
        // 返回一列值
        Logistic::query()->pluck('name', 'id');
        // 返回多行
        Logistic::query()->limit(10)->get();

        // 悲观锁
        Logistic::query()->sharedLock()->limit(10)->get();
        Logistic::query()->lockForUpdate()->limit(10)->get();

        // 批量插入/更新: 传入数组即可
        Logistic::query()->insert([['id' => time()]]);
        Logistic::query()->update([['id' => 1]]);

        // 软删除: use SoftDeletes;

        // 极简 db 组件: https://doc.hyperf.io/#/zh-cn/db/db
        // 性能更好, 不推荐使用在较重业务中使用
    }
}
```
