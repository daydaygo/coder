# php psr

[PSR](https://www.php-fig.org/psr/), PHP Standards Recommendations, php 标准建议, 俗称 **PHP 道德规劝委员会**

遵循标准的好处不再赘述, 简单说说 swoft 中实现的 PSR 标准

## PSR4: 自动加载

自动加载是现代 PHP 的基础, 非常有必掌握的基础知识, 涉及到的 PHP 类/命名空间 等相关基础知识

- [PHP the right way](http://laravel-china.github.io/php-the-right-way/) 上关于自动加载的内容
- 自动加载的原理 [大话PHP设计模式](https://www.imooc.com/learn/236) - 第3章 命名空间与Autoload: 从 `__autoload()` 魔法函数 -> `spl_register_autoload()` -> composer 中从 PSR0 发展到 PSR4
- [自动加载的几种方式](https://docs.phpcomposer.com/04-schema.html#autoload): PSR4 classmap file, 这里列举了三个最常用的

## PSR1 PSR2: 代码规范

想要写好代码, 代码规范算是 **基础中的基础**, 至少不要让其他人接手你的代码时, 动不动就 **这乱代码谁写的**. 可以参考的工具:

- `php-cs-fixer`: swoft 中使用的此工具, 设计理念是 **only fixer**
- `php_CodeSniffer`: [GitHub上有 2 者的对比](https://github.com/FriendsOfPHP/PHP-CS-Fixer/issues/3459), 设计理念最初是 **linting**, 只是后来慢慢加入了 fix 相关功能
- [phpro/grumphp](https://github.com/phpro/grumphp): 如果想要混合使用多个工具

swoft 中使用 `php-cs-fixer`:

- 安装使用参考 [composer.json](https://github.com/swoft-cloud/swoft/blob/master/composer.json)
- php-cs 配置请参考 [.php_cs](https://github.com/swoft-cloud/swoft/blob/master/.php_cs)

上手很简单:

```sh
php-cs-fixer fix $dir
```

默认使用 PSR1/PSR2, 其他配置可以参考官网文档, 配置到 `.php_cs` 文件中

## PSR3: 日志规范

想要用好日志, 还需要做更多的事, 相关参考:

- [devops| 日志服务实践](https://www.jianshu.com/p/9dae3ba679e6): 云平台的日志服务值得一试
- [yii 源码解读](https://www.jianshu.com/p/fd85383783eb) - 日志模块架构: `Logger - Dispatcher - Target` 三层结构, 将日志输入与输出进行解耦
- docker 实战ELK

## PSR6: 缓存规范
