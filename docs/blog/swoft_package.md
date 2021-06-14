# swoft| Swoft 框架组件化改造

Swoft 框架从单体应用到组件化改造的架构升级之路

经过一年多的开发, Swoft 框架功能越来越完善, 也越来越复杂. 初创时期的 **单体应用**, 已经无法支撑项目的快速发展, 于是开发组在年前为 1.0-beta 版制定了 **组件化改造** 的重构方案.

内容速览:

- 组件化原理: PHP 包管理基础知识
- 组件化方案: 来自 laravel/symfony 等成熟框架的组件化实现方案
- Swoft 框架组件化实现

## 组件化原理

编程始终要解决的一个问题: **代码复用**. 好的代码, 基本要求是 **正确**, 能拿到预期的结果, 少 bug. 语言层的代码复用解决方案, 通常称之为 **包管理**(或者 **依赖管理**). 流行的编程语言, 都提供了很好的工具链对包管理的支持:

- 一个命令行工具, 用来 获取/管理 包, 比如 php 的 [composer](https://getcomposer.org/), python 的 pip, js 的 npm, java 的 maven, go 的 `go get`
- 一个包管理的配置文件, 用来说明需要用到(依赖)的包, 比如 PHP 中 composer 使用 `composer.josn`, js 的 npm 使用 `package.json`
- 一个浏览包的网站, 用来查看包的信息, 比如 php 的 [packgist](https://packagist.org/), python 的 pypi 等

这样, 当我们需要不同功能的时候, 就可以去查看是否有包已经提供了类似功能, **不用重复造轮子, 站在巨人的肩膀上**.

回到 PHP 中, PHP中的包管理是如何实现的呢?

- 命名空间

首先需要知道的一个基础概念, 是 [命名空间](http://php.net/manual/en/language.namespaces.php). 引入命名空间是为了 **解决同名冲突** -- 2 个包中有名字相同的类, 同时使用时就会出现类重复定义的提示. 使用命名空间后, 因为不同的包有不同的命名空间, 就不会出现冲突.

```php
// 如果需要在同一个文件中使用相同名字的类, 使用别名
use A\Far;
use B\Far as BFar;
```

- 自动加载 & PSR4

第二个需要知道的基础概念, 是 [自动加载](http://php.net/manual/en/language.oop5.autoload.php). PHP 中最基础(或者说最原始)的复用代码的方法: `include/include_one/require/require_once`. 不过得益于 PHP 的 [SPL库](http://php.net/manual/en/book.spl.php) 中的 `spl_autoload_register()` 方法, 现在有了更优雅的方式来复用代码 -- 自动加载. 自动加载的规范也经历了一段时间的升级与打磨, 最新的是 [PSR4标准](https://www.php-fig.org/psr/psr-4).

> 关于自动加载, 有一个很好的教程: [ 5-1 SPL使用spl_autoload_register函数装载类 (10:03)](http://www.imooc.com/video/2620)

## composer 中的包管理

了解了基础知识后, 就可以来掌握工具怎么用了. [composer 中的包管理](http://docs.phpcomposer.com/02-libraries.html) 根据 `composer.json` 文件中的 `autoload / require / require-dev` 配置项来管理.

`autoload` 定义自动加载, 项目自身的代码, 也应该按照包管理的规范, 进行组织, 比如 [Swoft 的 composer.json 配置文件](https://github.com/swoft-cloud/swoft/blob/master/composer.json):

```
...
    "autoload": {
        "psr-4": {
            "App\\": "app/"
        },
        "files": [
            "app/Swoft.php"
        ]
    },
...
```

composer 支持[多种方式的自动加载方式](http://docs.phpcomposer.com/04-schema.html#autoload), 这里面有一定的历史原因, 因为需要兼容一些 **陈旧** 的代码. 现在比较常用的 2 种方式:

- `psr-4`: PSR4 标准, 优先推荐的方式
- `files`: 直接加载文件, 通常用来加载 **帮助函数**, 类似于 PHP 的 `require` 语法来代码复用

`require` 标识需要依赖的包, 格式是 `包名 - 版本限制` 的键值对:

```
...
    "require": {
        "php": ">=7.0",
        "swoft/framework": "^1.0",
        "swoft/rpc": "^1.0",
        "swoft/rpc-server": "^1.0",
        "swoft/rpc-client": "^1.0",
        "swoft/http-server": "^1.0",
        "swoft/http-client": "^1.0",
        "swoft/task": "^1.0",
        "swoft/http-message": "^1.0",
        "swoft/view": "^1.0",
        "swoft/db": "^1.0",
        "swoft/cache": "^1.0",
        "swoft/redis": "^1.0",
        "swoft/console": "^1.0",
        "swoft/devtool": "^1.0",
        "swoft/session": "^1.0",
        "swoft/i18n": "^1.0",
        "swoft/process": "^1.0",
        "swoft/memory": "^1.0",
        "swoft/service-governance": "^1.0"
    },
...
```

关于 **版本控制** 的知识, 以及 `>= ^ ~` 等特殊字符, `alpha beta dev dev-master` 等标识, 只是约定俗成的一些定义, 了解清楚即可.

`require-dev` 标识开发环境需要依赖的包, 即正式环境不需要使用到的包, 比如单元测试等:

```
...
    "require-dev": {
        "eaglewu/swoole-ide-helper": "dev-master",
        "phpunit/phpunit": "^5.7"
    },
...
```

类似的, 还有 `autoload-dev`, 表示测试环境下使用到自动加载.

## 组件化方案: laravel 与 symfony 使用的方案

参考 [symfony 中的 composer.json 配置文件](https://github.com/symfony/symfony/blob/master/composer.json) 和 [laravel 中的 composer.json 配置文件](https://github.com/laravel/framework/blob/master/composer.json), 会发现里面有一个配置项: `replace`.

`replace` 这个配置项, 在普通项目中很难看到, 却是组件化改造中的重要配置, 它的定义如下:

> 使用项目中已有的包, 替换需要依赖的包

比如 [symfony 中的 composer.json 配置文件](https://github.com/symfony/symfony/blob/master/composer.json):

```
...
    "replace": {
        "symfony/asset": "self.version",
        "symfony/browser-kit": "self.version",
        "symfony/cache": "self.version",
        "symfony/config": "self.version",
        "symfony/console": "self.version",
        "symfony/css-selector": "self.version",
        "symfony/dependency-injection": "self.version",
        ...

```

其中 `"symfony/asset"` 包, 有一个单独的github 仓库 [symfony/asset](https://github.com/symfony/asset), [symfony 项目](https://github.com/symfony/symfony) 本身也包含 `"symfony/asset"` 包, 使用 `replace`, symfony 就可以使用自身包含的包, 不用去单独获取.

这样带来的好处:

- 主包包含所有的子包, 使用时使用 `replace` 配置, 所有的修改和提交都在主包中进行
- 其他项目依旧可以使用 `require`, 单独使用子包; 子包只接受来自主包分发来的代码, 不接受在子包上的更改

## Swoft 框架组件化实现

Swoft 在 1.0-beta版中的依赖, [Swoft 项目](https://github.com/swoft-cluod/swoft):

```
...
    "require": {
        "php": ">=7.0",
        "swoft/framework": "^1.0",
        "swoft/rpc": "^1.0",
        "swoft/rpc-server": "^1.0",
        "swoft/rpc-client": "^1.0",
        "swoft/http-server": "^1.0",
        "swoft/http-client": "^1.0",
        "swoft/task": "^1.0",
        "swoft/http-message": "^1.0",
        "swoft/view": "^1.0",
        "swoft/db": "^1.0",
        "swoft/cache": "^1.0",
        "swoft/redis": "^1.0",
        "swoft/console": "^1.0",
        "swoft/devtool": "^1.0",
        "swoft/session": "^1.0",
        "swoft/i18n": "^1.0",
        "swoft/process": "^1.0",
        "swoft/memory": "^1.0",
        "swoft/service-governance": "^1.0"
    },
...
```

改造后, [Swoft 项目](https://github.com/swoft-cluod/swoft), 主项目只用依赖 `"swoft/framework"`:

```
...
    "require": {
        "php": ">=7.0",
        "swoft/framework": "^1.0"
    },
...
```

["swoft/framework" 项目](https://github.com/swoft-cloud/swoft-framework), 包含其他子包:

```
...
    "replace": {
        "swoft/rpc": "self.version",
        "swoft/rpc-server": "self.version",
        "swoft/rpc-client": "self.version",
        "swoft/http-server": "self.version",
        "swoft/http-client": "self.version",
        "swoft/task": "self.version",
        "swoft/http-message": "self.version",
        "swoft/view": "self.version",
        "swoft/db": "self.version",
        "swoft/cache": "self.version",
        "swoft/redis": "self.version",
        "swoft/console": "self.version",
        "swoft/devtool": "self.version",
        "swoft/session": "self.version",
        "swoft/i18n": "self.version",
        "swoft/process": "self.version",
        "swoft/memory": "self.version",
        "swoft/service-governance": "self.version"
    }
...
```

其中子项目声明到主项目提交修改:

- [Report issues](https://github.com/swoft-cloud/swoft-framework/issues) and send [Pull Requests](https://github.com/swoft-cloud/swoft-framework/pulls) in the [main Swoft repository](https://github.com/swoft-cloud/swoft-framework)

整个开发流程如下:

- 在 [daydaygo/swoft-framework 项目](https://github.com/daydaygo/swoft-framework/tree/component2) 新建 component2 分支开发此次组件化改造
- 修改 Swoft 项目的 composer.json 文件, 快速获取所有 Swoft 组件的 master 分支代码:

```
    "require": {
        "php": ">=7.0",
        "swoft/framework": "dev-master",
        "swoft/rpc": "dev-master",
        "swoft/rpc-server": "dev-master",
        "swoft/rpc-client": "dev-master",
        "swoft/http-server": "dev-master",
        "swoft/http-client": "dev-master",
        "swoft/task": "dev-master",
        "swoft/http-message": "dev-master",
        "swoft/view": "dev-master",
        "swoft/db": "dev-master",
        "swoft/cache": "dev-master",
        "swoft/redis": "dev-master",
        "swoft/console": "dev-master",
        "swoft/devtool": "dev-master",
        "swoft/session": "dev-master",
        "swoft/i18n": "dev-master",
        "swoft/process": "dev-master",
        "swoft/memory": "dev-master",
        "swoft/service-governance": "dev-master"
    },
```

- 复制各个组件的代码到 swoft-framework 项目中, 修改 [composer.json](https://github.com/daydaygo/swoft-framework/blob/component2/composer.json) 的中的 `autoload / replace` 配置(具体修改点击链接查看)

Swoft 各组件依赖关系图: http://naotu.baidu.com/file/7f89010cc4b7cd379fbf924f4f0752f1?token=7eadf529c90dcba4

- 提交 swoft-framework 代码.

下面以推送 `swoft-view` 组件到对应仓库中为例:

*出于 github 网速的原因, 测试过程使用 gitee 来加速*

推送子项目到相应的 github 仓库中, 参考:

- [dflydev/git-subsplit](https://github.com/dflydev/git-subsplit): 封装 `git subtree` 为 `git subplite` 命令, 方便使用
- laravel 中使用 git subsplit 的示例](https://github.com/laravel/framework/tree/5.1/build)

```
# 建立 gitee.com:daydaygo/swoft-framework 仓库
git remote add gitee git@gitee.com:daydaygo/swoft-framework.git
git push gitee component2

# 拆分
git subsplit init git@gitee.com:daydaygo/swoft-framework.git
# 更多项目, 一次填写即可
git subsplit publish --heads="component2" --no-tags view:git@gitee.com:daydaygo/swoft-view.git
# 清除生成的临时文件
rm .subsplit
```

这个拆分过程耗时较长, 拆分后的效果: [gitee.com/daydaygo/swoft-view](https://gitee.com/daydaygo/swoft-view/tree/component2/), [gitee.com/daydaygo/swoft-framework](https://gitee.com/daydaygo/swoft-framework/tree/component2/)

可以通过添加 github webhook 来做自动化, 具体请参考: [dflydev/dflydev-git-subsplit-github-webhook](https://github.com/dflydev/dflydev-git-subsplit-github-webhook)

最后, 测试拆分后的代码:

- 修改 swoft 项目的 composer.json 文件, 使用新版的 swoft-framework 项目

```
...
  // 现在只需要依赖 swoft/framework, 版本号要制定分支, composer 会默认给分支名带上 dev- 前缀
  "require": {
    "php": ">=7.0",
    "swoft/framework": "dev-component2"
  },
....
  "repositories": {
    "packagist": {
      "type": "composer",
      "url": "https://packagist.phpcomposer.com"
    },
    // 制定包的地址, 这里指向我的 giee 仓库地址
    "0": {
      "type": "vcs",
      "url": "https://gitee.com/daydaygo/swoft-framework"
    }
  }
...
```

```
# 删除以前的依赖
rm -rf composer.lock vendor
# 更新
composer install --no-dev
```

- 运行测试, 可以参考项目下 [.travis.yml 文件](https://github.com/daydaygo/swoft-framework/blob/master/.travis.yml)

至此, 大功告成.

## 写在最后

**对项目进行组件化拆分, 推送子包到不同 github 仓库** 这样的需求, 也许只有写一个大型框架才会遇到. 但这也是正是动手写一个框架的乐趣所在. PHP 中的包管理的基础知识一直感觉 **游刃有余**, 直到遇到新的问题, 提出新的挑战, 才发现还有更多的天地. 愿你也能感受到这分技术的乐趣.
