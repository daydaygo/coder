# php| 现代 PHP 框架浅析
> 给新人熟悉 hyperf 框架的一个全景分享

现代 php 框架浅析:

- 核心: container 容器
- 应用 & 应用生命周期管理: app 应用

## container 容器

现代 php 框架的核心其实都是一致的: container 容器(和 DI 依赖注入, IoC 控制反转 是同一语境), 解决的问题如下:

```php
// 如果类 A 的实例 a 需要使用到类 B 的实例 b
$a->b = new B();

// 如果有 container
$a->b = $container->get(B::class);

// 不仅仅是 A 依赖 B, 其他依赖的地方, 都可以直接用, 比如类 C 的实例 c
$c->b = $container->get(B::class);
```

## app 应用(服务)

用 **生命周期** 来理解 app 会比较好理解

- 入口文件
  - ini
  - 常量
  - composer autoload
  - container
  - app
- app 执行流程
  - command
  - server
    - mvc
    - route -> middleware -> controller -> service / model / view
  - 常用组件
    - config
    - log
    - redis mysql

## php 基础

- PHP 的配置
  - 查看: `php -i |grep xxx'
  - 设置
    - 可以使用 `php --ini` 查看配置文件, 通过修改配置文件设置 
    - 可以在代码中用 `ini_set()` 进行配置
- composer
  - 我想用别人的代码, 怎么用
    - 别人的代码 package: 
      - 设置自动加载: composer.json -> autoload -> ps-4: namespace(命名空间) <=> path(文件路径)
      - 设置名字: composer.json -> name
    - 依赖管理
      - composer.json -> require -> package name
      - composer install/require
        - package -> vendor 目录
        - 生成好 `vendor/autoload.php`
      - 入口添加自动加载文件 `require BASE_PATH . '/vendor/autoload.php';`

## 网络基础

- tcp/ip 4层
- server / client
- http 协议: url path method header body request/response

## co 协程基础

- 进程 线程 协程
  - 进程 process: 运行的代码 = 进程(不开多进程的情况) 
    - 查看进程: `ps aux | grep xxx`
    - 进程包含的资源
      - 环境变量 `env`: **进程皆有环境**
        - `PATH`
  - 线程 thread: 进程的执行单元, 交给操作系统管理
    - `pstree -p pid`
    - thread -> os 调度 -> cpu
  - 协程: 用户态的线程, 用户自己管理, 之后交给操作系统
    - 程序调度 vs os 调度
    - CPU密集型 vs io密集型

## 思维方式

- 分层
- 生命周期 = 执行顺序
- 思维导图
- 画图/实操 来理解概念性的内容

## 职业发展

- 技术 - 和机器交流: 编程基础 -> 胜任业务 -> 技术 blog -> 阅读源码 -> 开源 -> 架构设计
- 业务 - 和团队交流: 了解业务 -> 领域/行业 -> owner/团队

#### mac PHP环境一键配置

```sh
## homebrew https://brew.sh/
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
## 安装 homebrew 慢, 先下载 install 文件, 然后替换为 https://gitee.com/mirrors/homebrew

## homebrew 加速源: https://mirror.tuna.tsinghua.edu.cn/help/homebrew/
git -C "$(brew --repo)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/brew.git
brew tap homebrew/core
git -C "$(brew --repo homebrew/core)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-core.git
brew tap homebrew/cask
git -C "$(brew --repo homebrew/cask)" remote set-url origin https://mirrors.tuna.tsinghua.edu.cn/git/homebrew/homebrew-cask.git
brew update
## 复原
## git -C "$(brew --repo)" remote set-url origin https://github.com/Homebrew/brew.git
## git -C "$(brew --repo homebrew/core)" remote set-url origin https://github.com/Homebrew/homebrew-core.git
## git -C "$(brew --repo homebrew/cask)" remote set-url origin https://github.com/Homebrew/homebrew-cask.git
## brew update

## homebrew-bottles
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles' >> ~/.bash_profile
source ~/.bash_profile

## iterm2
brew cask install iterm2

## vim: 可以命令行使用 code 命令, 使用 vscode 编辑文件
## brew install vim

## fish shell: 推荐使用
brew install fish
## sudo echo '/usr/local/bin/fish' >> /etc/shells
## chsh -> 切换默认shell
## fish_config -> 配置 fish
## fish 配置目录: ~/.config/fish

## php: 项目中目前统一使用 7.2
## 踩到的坑: 本地 7.3 更新 composer 后, 有的包更新后强制要求使用 7.3, 导致其他环境 composer i 失败
brew install php@7.2
## 安装后注意根据提示添加 PATH
## composer: https://developer.aliyun.com/composer
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
## pecl: http://pecl.php.net/
## 可以在 pecl 上下载好包后直接安装, 避免重复网络请求
## swoole
## 确保环境变量里能找到 openssl 的编译库
pecl install swoole-4.4.13.tgz
## 使用 `php --ini` 查找 php.ini 文件位置, 编辑添加 `swoole.use_shortname = 0`, 使用 `php --ri swoole` 确认是否生效
pecl install redis-5.1.1.tgz
pecl install mongodb-1.6.1.tgz
## protobuf
brew install protobuf
pecl install protobuf-3.11.1.tgz
## kafka
brew install librdkafka
pecl install rdkafka-4.0.0.tgz
```

## centos php

```sh
yum install -y epel-release
yum install -y http://rpms.remirepo.net/enterprise/remi-release-7.rpm
yum install -y yum-utils
yum install -y php73-php-cli php73-php-swoole php73-php-redis php73-php-simplexml php73-php-bcmath php73-php-gd php73-php-json php73-php-mbstring php73-php-mcrypt php73-php-mysqlnd php73-php-opcache php73-php-pdo php73-php-pecl-crypto php73-php-pecl-mcrypt php73-php-pecl-geoip php73-php-recode php73-php-snmp php73-php-soap php73-php-xmll
ln -s /bin/php73 /bin/php
curl -sS https://install.phpcomposer.com/installer  | php

server {
    root /data/wwwroot/;
    server_name local.swoole.com;

    location / {
        proxy_http_version 1.1;
        proxy_set_header Connection "keep-alive";
        proxy_set_header X-Real-IP $remote_addr;
        if (!-e $request_filename) {
             proxy_pass http://127.0.0.1:9501;
        }
    }
}
```

## composer
- 命令使用前缀匹配, 可以只输入前几个字符

## phpunit

```bash
# 全局安装
composer global require hyperf/testing

# 使用
composer test
co-phpunit -c phpunit.xml --colors=always --bootstrap=./test/bootstrap.php
# 测试单个用例
co-phpunit --filter=testFoo
```

## swoole

- linux system program
  - co: Process thread goroutine
  - aysnc: AIO epoll io复用
  - net: TCP/UDP IPV4/IPV6/UNIXsocket
  - fs: pagecache mmap IO
  - IPC: mq sign 共享内存/数据同步/lock
  - protocol: HTTP/HTTP2/ws mysql/redis/MongoDB grpc/brpc/tars