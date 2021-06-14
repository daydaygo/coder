# devops| nginx unit 入门小试

nginx unit 入门小试以及 PHPer 的一点浅见

> [devops| nginx unit 入门小试](https://www.jianshu.com/p/d674b53cecac): 服务器 dev 又多了一个玩具，也算是开心事一件

上周几乎被 nginx unit 的消息给霸屏了, 大致看了看这个产品的 **野心**, 感觉还是挺有意思的:

- Multi-language support  = 方便部署, 方便了喜欢多语言折腾的 dev
- Programmable = 基于API的配置方式, 提供基于 http 的 **配置/监控** 会是未来的首选
- Service mesh = 支持Istio的现有功能, mesh 已经是行业内公认的微服务配套技术方案

对于 phper 而言, 能直接使用 `nginx + php`, 而不用 `nginx + php-fpm`, 也算是一个小小的进步(变化).

## nginx unit 快速上手

nginx unit 在 docker hub 上提供了官方镜像 [nginx/unit](https://hub.docker.com/r/nginx/unit), 所以使用 docker 上手会非常容易:

### nginx unit Dockerfile

```dockerfile
FROM nginx/unit
LABEL maintainer="1252409767@qq.com"

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
    apt update && \
    apt install busybox net-tools

CMD /usr/sbin/unitd --no-daemon --control 0.0.0.0:8010
```

Dockerfile 主要做了以下几件事:

- `FROM nginx/unit`: 获取最新的 nginx/unit 镜像, 官方分为 **full/minimal/各语言版(PHP, Python, Go, Perl, and Ruby)**, 默认为 full 版
- `sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list`: 基础镜像是 Debian9, 这里使用了 [中科大镜像源](http://mirrors.ustc.edu.cn/) 来加速
- `apt install busybox`: Debian9 默认没有 `ps/top` 等命令, 安装 busybox 工具后使用 `busybox ps`
- `apt install net-tools`: Debian9 默认没有 `netstats` 等工具, 另一个选择是 `iproute2`
- `CMD /usr/sbin/unitd --no-daemon --control 0.0.0.0:8010`: 推荐这样的方式使用 `CMD` 指令, 比参数形式更简洁直观, 这里开启了 nginx unit 的 api 控制接口, 可以通过 http put 请求进行配置

### nginx unit docker-compose

```yaml
    unit:
        build:
            context: nginx
            dockerfile: unit.Dockerfile
        ports:
            - "8010:8010"
            - "8011:8011"
        volumes:
            - ../:/var/www
```

这里只配置了 8010 和 8011 端口:

- 8010 进行配置更新
- 8011 用来测试启动的服务

### nginx unit for php

这里使用 vscode 的 [rest client](https://github.com/Huachao/vscode-restclient) 进行测试:

```http
PUT http://localhost:8010 HTTP/1.1
content-type: application/json

{
    "listeners": {
        "*:8011": {
            "application": "php"
        }
    },
    "applications": {
        "php": {
            "type": "php",
            "processes": 20,
            "root": "/var/www/doc",
            "index": "test.php",
        }
    }
}
```

![nginx-unit-php](https://upload-images.jianshu.io/upload_images/567399-22a65590f7af64d8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


也可以在 `/etc/unit/` 目录下新建 json 格式的配置文件来进行配置, **不推荐**这种方式.

test.php 文件只是测试内容:

```php
<?php
phpinfo();
>
```

![nginx-unit-phpinfo](https://upload-images.jianshu.io/upload_images/567399-cdeaa025df4dc724.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


不用 php-fpm, 现在也能跑起 php web 应用了

访问 `http://localhost:8010` 查看服务器配置信息, 和 put 请求中的信息一致:

![nginx-unit-config](https://upload-images.jianshu.io/upload_images/567399-916bee0c855e806e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## nginx unit 小评

nginx unit 目前已经支持 PHP, Python, Go, Perl, and Ruby, java/nodejs 也在计划中, 基本涵盖了 web 应用的主流编程语言. 对于喜欢折腾服务器这块的 dev, 又多了一个 **玩具**, 也是开心事一件.

试用 php 中的主流框架 yii, 会报错 `mb_strlen() not fund`, 会发现 nginx/unit 的docker 镜像中使用 Debian9 软件源中默认的 php7.0 版本, 并没有开启 `mbstring` 扩展. 所以想要对 **业务服务器** 有更精细的控制, 还是要具体到各个语言本身的环境配置上.

其次, docker compose 现有的服务编排技术是 `nginx` 镜像和 `php-fpm` 镜像分开, 各自配置自己的环境, 然后通过网络 `links` 起来:

- nginx+fpm: docker-compose

```yaml
    nginx:
        build:
            context: nginx
            dockerfile: Dockerfile
        volumes:
            - ../:/var/www
            - ./logs/nginx/:/var/log/nginx
        links:
            - fpm
            # - swoft
        ports:
           - "80:80"
           - "443:443"
```

- nginx+fpm: nginx conf

```conf
# yii
server {
    listen 80;
    server_name yii.daydaygo.me yii.dayday.tech;
    index index.php;
    root /var/www/project/tools/yii/frontend/web;
    # access_log /var/log/nginx/yii_access.log main;
    location / {
        try_files $uri $uri/ /index.php?$args;
    }
    location ~ \.php$ {
        fastcgi_pass fpm:9000; # 通过 docker compose 编排时定义的服务名称
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
}
```

nginx unit 怎么实现这样服务编排的效果, 会是下一步探讨(折腾)的方向.

## 写在最后

目前网上关于 nginx unit 的资料比较少, 官网文档通常是解决问题的最快路径.

推荐资源:

> [nginx unit 官网](https://www.nginx.com/products/nginx-unit/)
> [知乎 - 怎么看 NGINX 团队新出的 NGINX-UNIT?](https://www.zhihu.com/question/65126862/answer/388393918)
> [Linux网络管理常用命令：net-tools VS iproute2](https://www.cnblogs.com/wonux/p/6268134.html)