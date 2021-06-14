# php| Nginx+PHP-fpm 出现 Primary script unknown 错误解决

这次的问题耗时非常长，持续了一周，最终在周末（2个整天）才拿下，虽然最后的解决方案非常傻瓜，不过过程还是蛮有意思的，在此一并梳理下。

## 起源

出于对 **编程方法学** 的研究，名词盗用自 [网易公开课 - Stanford - 编程方法学（cs106A）](http://open.163.com/special/sp/programming.html) ，强烈推荐！！！

其实只是单纯的思考：

- 怎么少写 bug？
- 怎么样代码上线了能少出问题？

这里着重说一下第二个问题，也是很容易遇到的问题 ：

- 我本地明明好好的，为什么测试那里就有问题？测试环境没有问题呀，为什么上线了又出问题了？
- 之前明明好好的，怎么突然就坏了？

下面依次说明这 2 个问题。

## docker docker docker

这里引用 docker 社区里的一句话：

> 程序员认为自己需要交付的是代码，只要代码逻辑正确就好了，但实际上项目需要交付的是 代码 + 代码运行所需的环境。

那么，问题就来了，按照「单一职责原则」，只用交付代码才是符合这一原则的，但现实却并不允许。

这个问题可以说困扰了整个程序界，也困扰了我许久，欢迎访问我的 [wiki](http://wiki.dayday.tech) 查看我的「项目开发部署变迁史」。

好消息是，docker 终于来了。关于 docker，我简单概括一下：

> 终于，代码运行的环境，也可以写到代码里了。

我现在使用的环境 `docker-compose.yaml`：

```yaml
version: '3'

services:
  a: # 命令行工具
    build:
      context: ./fpm
      dockerfile: alpine-Dockerfile
    volumes:
      - ../:/var/www
    dns:  # 加速
      - 223.5.5.5
      - 223.6.6.6
    tty: true

  nginx:
    build: ./nginx
    volumes:
      - ../:/var/www  # 挂载项目根目录
      - ./logs/nginx/:/var/log/nginx  # 日志
    links:
      - fpm
    dns:
      - 223.5.5.5
      - 223.6.6.6
    # network_mode: "service:php-fpm" # 尝试过，但是有点问题，又切回去了
    extra_hosts:
      - "aliyun:106.14.158.236"
    ports:
       - "80:80"
       - "443:443"

  fpm:
    build:
      context: ./fpm
      dockerfile: Dockerfile
    volumes:
      - ../:/var/www
      - ./logs/fpm/:/var/log/php-fpm
      # - ./data/sessions:/var/lib/php/sessions # 使用 redis 来保存 session 了
    links:
      - mysql
      - redis
    dns:
      - 223.5.5.5
      - 223.6.6.6
    extra_hosts:
      - "aliyun:106.14.158.236"

  mysql:
    build: ./mysql
    volumes:
      - ./data/mysql:/var/lib/mysql
    ports:
      - "3306:3306"
    environment:
      # MYSQL_DATABASE: test
      # MYSQL_USER: test
      # MYSQL_PASSWORD: test
      MYSQL_ROOT_PASSWORD: dayday.tech

  redis:
    build: ./redis
    volumes:
      - ./data/redis:/data
    ports:
      - "6379:6379"
  mongo:
    build: ./mongo
    # volumes:
      # - ./data/mongo:/data/db
    ports:
      - "27017:27017"
```

这里只是 `docker-compose.yaml` 文件，需要查看 dockerfile 欢迎访问的我的 coding 项目 [docker](https://git.coding.net/daydaygo/docker.git)。

其实最开始并没有使用这种方式，还是 lnmp 老一套，直接一个镜像搞定，显然没有这灵活了，思路借鉴自 github 上的 [laradock](https://github.com/laradock/laradock)。在技术社区逛久了，果然充满惊喜。

## keep running

除了开发环境引发的各种「血案」，还有一个非常重要的问题：**可用性**，我认为这是非常重要的计算机思维之一。

- 我们眼下可以有很多种方案，看似每种都可以，但是放到一周后呢？
- 终于做完一个需求了，代码提交，悲剧了，影响到原有功能了。
- 赶紧，网站报 500 了，快看看。尼玛，mysql 挂了，服务器 ssh 不上去

说多了都是经历……

回到正题，我目前发现的一个好的解决办法：**keep running**，有长期维护的项目，并且持续运行。

## Primary script unknown

目前采用 `aliyun ecs + docker compose` 来运行我的环境，可以在我的 [wiki](http://wiki.dayday.tech) 查看现在正在跑的项目。

确定目标后，就开始买 ecs，然后配置 docker 环境，`git clone` 项目，`docker-compose up` 起服务环境：

然后访问：https://laravel.dayday.tech （其实是 http2，目前有 2 套 https 免费方案，又拍云的 cdn 加速 和 腾讯云的免费 ssl 申请），500。

然后查看 nginx 日志 `docker/logs/nginx/error.log`，就发现了 `Primary script unknown` 这个错误。

当时非常非常的郁闷：

- 说好的 docker 大法呢！！！
- 我本地没问题呀
- 我快到期的 腾讯云 上面也部署过呀，没问题呀

怀疑完人生之后，问题还是要解决的，首先是各种搜索，基本都是说的配置文件有问题：

```
# laravel
server {
	listen 80;
	server_name laravel.daydaygo.dev laravel2.dayday.tech;
	index index.php;
	root /var/www/project/laravel/public;
	location / {
		try_files $uri $uri/ /index.php?$args;
	}
	location ~ \.php$ {
		fastcgi_pass fpm:9000;
		fastcgi_index index.php;
		fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
		include fastcgi_params;
	}
}
```

其实并没有问题，特别是根据搜索到的结果添加 nginx 日志查看后：

```
# add to nginx.conf
log_format scripts '$document_root$fastcgi_script_name > $request';

# add to server conf block
access_log /var/log/nginx/scripts.log scripts;
```

anyway，一定还是环境的锅！

下面开始就是错误的示范了，也就是为什么这个问题会持续一周之久：

- ecs 原系统是 Ubuntu，尝试国内更通用的 centos，ok，重装系统、重走一遍环境部署流程，然而并没有用
- 期间重复重装了几次，开始的时候还不知道 ecs 的快照和镜像功能
- 怀疑 docker 的版本，我本地 17.05，腾讯云上面 17.04，ecs 上 17.03，然后又重装系统重装 docker
- 我的 docker 用的基础镜像是 alpine linux，换成 Debian 试试？
- 要不试一下单镜像？

上面全都是环境相关的修改，期间也有想过会不会是 **权限的问题**，php-fpm 是不允许运行到 root 下面的：

```
# fpm/www.pool.conf
user = www-data
group = www-data
```

默认使用 `www-data` 来运行。所以尝试改 laravel 这个项目的权限：

```
chown -R www-data:www-data /var/www/project/laravel
```

尝试单镜像的时候，也考虑使用 unix socket 来连接 nginx + php-fpm：

```
# nginx server conf
location ~ \.php$ {
  fastcgi_pass unix:/var/fpm.sock; # 看这里
  fastcgi_index index.php;
  fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
  include fastcgi_params;
}

# fpm/www.pool.conf
listen = /var/fpm.sock
```

问题依旧没有解决！！！

最后，快要放弃的时候：

```
chmod -R 777 /var/www
```

呵呵、哭笑不得、黑人问号脸。

所以说一下最终的解决方案吧：

- 腾讯云安装官方 Ubuntu 镜像时，默认使用 ubuntu 作为用户名，而 ecs 用的 root
- 腾讯云在使用的过程中，基本操作都是在默认用户下，docker 也是，需要权限才使用 sudo
- 所以给 ecs 建了 www 用户，基本操作都在 www 下，fpm 也使用 www 用户

> 珍爱生命，远离 root