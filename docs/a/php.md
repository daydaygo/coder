# php

- env
  - cli / command line: fpm conf
  - ini: runtime conf; `PHP_INI_*` modes; ini_set() phpino()
  - [pecl](https://pecl.php.net/): dl()
  - debug: yasd=ext+`php -e`+`phpstorm > debug > start listening`
  - 环境配置: https://github.com/heiyeluren/heiyeluren-tools/tree/master/linux-env-init
- type
  - `string`: 4 ways; numberic strings
  - array: ordered map; key(int/string autoindex)
  - iterable: array/object+Traversable; foreach/yield
  - resource: `$handle` get_resource_type() get_resource_id()/fd gc
  - null: is_null()/unset()
  - callback/`callable`: passing
  - type declaration: mixed self
  - type juggling
- var
  - init-default
  - predefined: `$argv`等
  - scope: global/`$GLOBALS`/superglobal static `symbol_table`
  - var-var: `$$a`
  - external: `$_GET`等 `php_globals.h`
  - const: define()/defined()
    - magic const: `__xxx__` `ClassName::class`
- op: `?? ?:` `@` error_reporting(); `` shell_exec(); array `$arr1 + $arr2`; instanceof is_a()
- control: match; declare; `require/include/require_once/include_once`
- function
  - func > var/type > func handle
  - language construct: echo/print/unset/isset/empty/include/require
- OO
  - visiable: var=public
  - property type: self/parent/class/interface/`?type`
  - autoload: `spl_autoload_register() vs __autoload()`
  - trait: method reuse
  - anonymous
  - iterate: visible comparison
  - serialize session
  - Covariance/Contravariance: type redeclare
  - OOP changelog
- namespace
  - `__NAMESPACE__` namespace use as
  - name resolution rule
- error
  - `E_ERROR` error_reporting display_errors/display_startup_errors log_error error_log set_error_handler()
- exception
  - register_shutdown_function()
- attribute/annotation
  - `#[Attribute]` `#[SetUp]`
- reference
  - spot: global
- protocol & wrapper
  - URL-style protocols `php://stderr`(`STDERR`) `file://`
- feature
  - http auth basic/digest
  - header setcookie session file-upload
  - gc: zval xdebug_debug_zval `gc_collect_cycles` 清理内存碎片`gc_mem_caches()`
- function reference/definition/prototype
  - php behavior: opcache; error handle; runtime conf; outputControl(flush ob_xxx output_xxx); phpdbg; php option/info
  - crypto: hash/openssl/password_hash sodium
  - db database: abstract layer(PDO) -> vendorSpecific(pdo_mysql/mysqli) -> client(mysqlnd/libmysqlclient) -> server
  - datatime
  - fs, file system: mine
  - i18n / char encoding: iconv; multibyte string
  - image: exif GD Gmagick imagemagick
  - other: SPL(DS iterator interface exception) DS
    - type & type related: array class
- project
  - code: php-cs-fixer phpstan
  - dep: composer autoload psr4 依赖管理可视化?
  - phar
- security php安全之道
- core
  - zen engine(v4 php8): zval zend_string opcode
    - scan(lexing) -> parse(tokens) -> compile -> exe; src->tokens->exp->opcode->bin
  - ext extension: build system; struct; counter example
  - working with: var func class resource ini stream
  - embedSAPI: c/cpp php/zen
- FAQ/appendix
  - ini
  - ext list&cate: membership
  - func alias
  - type comparison table
  - parser token list

---

- context option & parameter
- connection handle: connection_status set_time_limit ignore_user_abort connection_aborted
- Persistent connections: are good if the overhead to create a link to your SQL server is high
- dtrace
- sodium/taint
- Trader
- seaslog wkhtmltox(html -> pdf/img) apc/apcu

## code

- base algo coder(小工具)

## mark

```sh
php -i|ag xxx # php --ini
php -S localhost:8000 -t src # built-in web server
php -r # run code
php -n -e -d extension=/path/to/ext

php --ri ext-name
php-config
# install extension
pecl install xxx # 可以带 version; 可以是本地文件
phpize && ./configure && make && make install # phpize clean; make clean
```

```ini
max_execution_time = 0 # set_time_limit(0);
; error_reporting()
E_ALL & ~E_NOTICE
E_ALL ^ E_NOTICE
E_ERROR | E_RECOVERABLE_ERROR
```

---

- 语言基础
  - [php manual](http://www.php.net/manual) 大致过一遍；随用随查
  - [php the right way](http://www.phptherightway.com/)：面向现代化的 php 开发者
  - [php 学习路线图](https://www.zhihu.com/question/27170424)
- 进阶
  - [Swoole：重新定义PHP](http://wiki.swoole.com/)：扩展 php 在服务器编程领域的边界
    - [swoole/phpx: 💗 C++ wrapper for Zend API](https://github.com/swoole/phpx)
    - [yasd: swoole debugger](https://github.com/swoole/yasd)
    - [swow: 纯协程版 swoole](https://github.com/swow/swow) [swow doc](https://wiki.s-wow.com/)
    - swoole tracker
      - [Swoole Tracker3.1发布，支持完善的内存泄漏检测！ - Swoole](https://wenda.swoole.com/detail/107688)
  - 项目/工程/框架: [hyperf](https://hyperf.wiki)
- devops
  - [php 标准规范 psr](https://psr.phphub.org/)：php道德规劝委员会
  - [composer - php 包管理](http://docs.phpcomposer.com/)：[packagist - The PHP Package Repository](https://packagist.org/)
  - 单测: PHPunit
  - debug: 原理非常简单
    - IDE 开启 debug 监听: `start listen for php debug connection`
    - debug 扩展: ini 中配置 remote host+port
- 生态
  - `league/csv` `simps/mqtt` `psy/psysh` deployer/walle phpdocument
- 资源
  - [learnku-php 社区](https://learnku.com/php): 社区文档很给力
  - 关注大咖: [rango](http://rango.swoole.com/) [鸟哥](http://www.laruence.com/) [phpconchina](http://www.phpconchina.com/)
- learn
  - [大话PHP设计模式 rango](http://www.imooc.com/learn/236): 自动加载 SPL 魔术方法 面向对象基础 11种设计模式
  - [php性能优化](http://www.imooc.com/learn/205)
  - [PHP LifeCycle演讲幻灯片(PHP LifeCycle Slides)](http://www.laruence.com/2008/08/15/283.html)

## debug

```php
echo xdebug_time_index(), "\n"; // 放在最后, 查看运行时间
```

## composer

```sh
# https://developer.aliyun.com/composer
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/ # -g global
composer config -g --unset
composer -vvv require alibabacloud/sdk # -vvv debug
composer install/update/show/dumpautoload
composer i --no-dev -vvv # u
composer show | ag hyperf
composer why-not php:8
composer create-project hyperf/swow-skeleton:dev-master

composer install --optimize-autoloader --no-dev # optimize autoload

# global: ~/.composer
```

```json
"extra": {
    "hyperf": {
        "plugin": {
            "sort-autoload": {
                "hyperf/utils": -1 // 对自动加载进行排序
            }
        }
    }
},
```

## swoole

```sh
swoole --enable-thread-context # mac m1
```

## swow

- >=php7.3
- ide helper: `lib/src/Swow.php`
- [swow/README-CN.md](https://github.com/swow/swow/blob/develop/README-CN.md)
  - 高性能: c协程+php协程 纯c协程=大部分单栈切换
  - 高可控: php虚拟机->协程->超细粒度控制这些协程 WatchDog组件
  - 易兼容: swow所及之处皆是协程
  - 事件驱动: swow->libcat->libuv.Proactor模型.linux新特性io_uring
  - php可编程性: 基于swow+php实现c扩展实现的网络功能
  - 现代化: OO exception 绿色增强=几乎不改原代码
  - 编程理念: csp=纯协程
- [SDB](https://wiki.s-wow.com/#/zh-cn/tools/sdb)
- example
  - buffer
  - channel select/callback
  - coroutine debug/interrupt
  - debug
  - http: ab echo mixed=http+ws+chat.html
  - tcp: echo heartbeat
  - runtime: nanosleep/epoll
  - signal: wait
  - watchDog
  - amazing: usleep 10k.file.pdo.mysqli stream_socket_client()/stream_socket_server() udp

## laravel
