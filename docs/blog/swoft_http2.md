# swoft| Swoft官网全站 HTTP2 实践

拥抱新时代, 全站HTTP2 + 免费泛域名证书, Swoft官网全站 HTTP2 实践 https://www.jianshu.com/p/624fdaee5e38

[Swoft1.0正式来袭](http://www.oschina.net/news/93971/swoft-1-0), Swoft 也迎来自己的一个里程碑, [star数正式突破 1k](https://github.com/swoft-cloud/swoft). Swoft官网作为项目组服务开发者们的重要渠道, 也迎来了自己的一次重大更新:

- 重构, 升级到 Swoft1.0
- 全站实现HTTP2

本篇先介绍 **Swoft官网全站 HTTP2 实践**

先来一张 [Swoft 官网](https://www.swoft.org/) 效果图镇楼:

![swoft 官网: 全站 HTTP2](https://upload-images.jianshu.io/upload_images/567399-21cf5c83f6b6b354.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 静态资源由 nginx 托管, 开启 http2
- 业务代码交由 [Swoft](https://github.com/swoft-cloud/swoft) 执行, 设置 [\Swoole\HttpServer](https://wiki.swoole.com/wiki/page/326.html) 使用 HTTP2 协议

要实现 HTTP2 非常简单:

- nginx 开启 HTTP2
- Swoft 开启 HTTP2
- nginx + Swoft 配合使用
- 福利: 域名证书申请 **轻松指南**

## nginx 开启 HTTP2

首先查看 nginx 中是否开启了 HTTP2 module(模块)

```
# -V: show version and configure options then exit
/var/www # nginx -V

# 新版 nginx 默认开启了 HTTP2: --with-http_v2_module
nginx version: nginx/1.13.8
built by gcc 6.2.1 20160822 (Alpine 6.2.1)
built with OpenSSL 1.0.2n  7 Dec 2017
TLS SNI support enabled
configure arguments: --prefix=/etc/nginx --sbin-path=/usr/sbin/nginx --modules-path=/usr/lib/nginx/modules --conf-path=/etc/nginx/nginx.conf --error-log-path=/var/log/nginx/error.log --http-log-path=/var/log/nginx/access.log --pid-path=/var/run/nginx.pid --lock-path=/var/run/nginx.lock --http-client-body-temp-path=/var/cache/nginx/client_temp --http-proxy-temp-path=/var/cache/nginx/proxy_temp --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp --http-scgi-temp-path=/var/cache/nginx/scgi_temp --user=nginx --group=nginx --with-http_ssl_module --with-http_realip_module --with-http_addition_module --with-http_sub_module --with-http_dav_module --with-http_flv_module --with-http_mp4_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_random_index_module --with-http_secure_link_module --with-http_stub_status_module --with-http_auth_request_module --with-http_xslt_module=dynamic --with-http_image_filter_module=dynamic --with-http_geoip_module=dynamic --with-threads --with-stream --with-stream_ssl_module --with-stream_ssl_preread_module --with-stream_realip_module --with-stream_geoip_module=dynamic --with-http_slice_module --with-mail --with-mail_ssl_module --with-compat --with-file-aio
--with-http_v2_module
```

```
# http2
server {
    listen 80;
    server_name www.dayday.tech;
    # 将 HTTP 请求强制跳转到 HTTPS
    rewrite ^(.*)$ https://${server_name}$1 permanent;
}
server {
    # 开启 HTTP2
    listen 443 ssl http2 default_server;
    server_name www.dayday.tech;

    # 证书极简设置
    ssl on;
    ssl_certificate dayday.tech.crt;
    ssl_certificate_key dayday.tech.key;

    root /var/www/https_test;
    index index.php index.html;
    location / {}
    location ~ \.php$ {
        fastcgi_pass fpm:9000;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
}
```

## Swoft 开启 HTTP2

Swoole 开启 HTTP2, 可以参考 [Swoft 提供的 Dockerfile](https://github.com/swoft-cloud/swoft/blob/master/Dockerfile)

```
# Debian系Linux
apt-get install -y libssl-dev libnghttp2-dev

# Swoole 添加编译参数
./configure --enable-async-redis --enable-mysqlnd --enable-coroutine --enable-openssl --enable-http2
```

Swoft 配置中开启 HTTP2, 参考 [.env.example 文件](https://github.com/swoft-cloud/swoft/blob/master/.env.example)

```
# 默认配置
OPEN_HTTP2_PROTOCOL=false
SSL_CERT_FILE=/path/to/ssl_cert_file
SSL_KEY_FILE=/path/to/ssl_key_file

# 开启 HTTP2: 这里是将证书放到项目 resource/ 目录下
OPEN_HTTP2_PROTOCOL=true
SSL_CERT_FILE=@res/ssl/ssl_cert_file
SSL_KEY_FILE=@res/ssl/ssl_key_file
```

## nginx 配合 Swoft 使用

nginx 配合 Swoft 使用, 类似 `nginx+fpm` 配置即可, 代码示例可以参考 [我的开源项目-docker](https://gitee.com/daydaygo/docker/blob/master/nginx/sites/swoft.bak)

```
# swoft-site
server {
  listen 80;
  server_name swoft.dayday.tech;
  # 将 HTTP 请求强制跳转到 HTTPS
  rewrite ^(.*)$ https://${server_name}$1 permanent;
}
server {
  # 开启 HTTP2
  listen 443 ssl http2;
  server_name swoft.dayday.tech;

  # 证书极简配置
  ssl on;
  ssl_certificate 1_swoft.dayday.tech_bundle.crt;
  ssl_certificate_key 2_swoft.dayday.tech.key;

  root /var/www/swoole/swoft-offcial-site/public;
  index index.php index.html;
  error_log /var/log/nginx/swoft-site.error.log;
  access_log /var/log/nginx/swoft-site.access.log;

  # nginx 转发请求给 swoft
  location / {
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header Host $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Connection "keep-alive";
    proxy_pass https://swoft:9501;
  }
  location ~ \.php(.*)$ {
    proxy_pass https://swoft:9501;
  }

  # nginx 托管静态文件
  location ~* \.(js|map|css|png|jpg|jpeg|gif|ico|ttf|woff2|woff)$ {
    expires       max;
  }
}
```

## 福利: 域名证书申请 **轻松指南**

先确认你知道关于域名的几个基础知识:

- 为什么用域名?
- 什么是子域名?
- 为什么域名要备案?
- 什么是域名证书?

如果这些都不熟悉, 建议申请一个域名体验一下.

域名证书分为 2 种: 单域名证书 泛域名证书, 区别来自于 **什么是子域名**. 比如我拥有域名 `.dayday.tech`, 那么我可以设置任意子域名, 比如 `www.dayday.tech`, `test.www.dayday.tech`. 如果是单域名证书, 那么我每一个子域名都需要一个证书, 泛域名证书则可以对我所有的子域名生效.

域名证书由相关机构发放, 一般需要花钱购买. 既然是 **福利**, 这里介绍 2 个免费好用的途径:

- **动动鼠标, 证书到手**, 腾讯云-申请免费单域名证书
- **终于等到免费泛域名证书**, Let's Encrypt 泛域名证书

### 单域名证书实践
> 腾讯云-申请免费单域名证书: https://console.qcloud.com/ssl

全程只需要动动鼠标即可:

- 到腾讯云官网申请
![腾讯云 - 单域名证书申请](https://upload-images.jianshu.io/upload_images/567399-68faa37e694f7289.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 配置域名解析验证域名所有权
![配置域名解析](https://upload-images.jianshu.io/upload_images/567399-3b54582c913692fb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


然后下载证书, 配置到 nginx 中即可. 详细教程请参考腾讯云官方文档.

不过要注意:

- 证书有效期 1 年
- 同一域名最多只能申请 20 个证书

### 通配符域名证书实践
> Let's Encrypt 终于支持通配符证书了: https://www.jianshu.com/p/c5c9d071e395

`Let's Encrypt` 在免费域名证书领域算是 **家喻户晓**, 现在终于支持 **通配符证书** 了. 不过按照上面 blog 的教程, 很是一番折腾. 虽然一波三折, 但是得益于自己使用 docker 作为开发环境, 在尝试各种解决方案时, 都没有太大阻碍.

这里记录下来最终成功使用的一种方式:

- 来自官网的教程: [ACME v2 Production Environment & Wildcards](https://community.letsencrypt.org/t/acme-v2-production-environment-wildcards/55578): **Remember: You must use an ACME v2 compatible client to access this endpoint**
- [ACME v2 Compatible Clients](https://letsencrypt.org/docs/client-options/#acme-v2-compatible-clients): 结合上面的 blog, 选择 certbot
- 在 [certbot](https://certbot.eff.org) 官网, 选择 [`nginx+centos7` 环境](https://certbot.eff.org/lets-encrypt/centosrhel7-nginx), 出现教程
- 使用自己的 [docker 开发环境 - centos](https://gitee.com/ddaydaygo/docker/centos) 进行实践

```
# 安装 certbot
yum install certbot-nginx

# 稍微修改教程中的命令
certbot certonly -d *.dayday.tech --manual --preferred-challenges dns --server https://acme-v02.api.letsencrypt.org/directory --email 1252409767@qq.com

# 单域名证书
certbot certonly --standalone --agree-tos -d swoft.org -d www.swoft.org -d doc.swoft.org --email daydaygo@swoft.org
```

之后一路确认, 最后添加 **配置域名解析验证域名所有权**, 大功告成!

```
[root@e6be50c34c81 www]# ll /etc/letsencrypt/live/dayday.tech/
total 4
-rw-r--r-- 1 root root 543 Mar 16 16:48 README
lrwxrwxrwx 1 root root  36 Mar 16 16:48 cert.pem -> ../../archive/dayday.tech/cert1.pem
lrwxrwxrwx 1 root root  37 Mar 16 16:48 chain.pem -> ../../archive/dayday.tech/chain1.pem
lrwxrwxrwx 1 root root  41 Mar 16 16:48 fullchain.pem -> ../../archive/dayday.tech/fullchain1.pem
lrwxrwxrwx 1 root root  39 Mar 16 16:48 privkey.pem -> ../../archive/dayday.tech/privkey1.pem
```

查看 `README`, 所得证书与 nginx 配置对应关系如下:

```
ssl_certificate  -> fullchain1.pem
ssl_certificate_key -> privkey1.pem
```

`certbot` 还可以配置 crontab 来 **自动更新证书**, 按照 [官方教程](https://certbot.eff.org/lets-encrypt/centosrhel7-nginx) 配置即可

折腾的过程颇为一波三折, 简单记录一下, 希望能给大家帮助:

- 我本人喜欢使用 alpine linux, 所以直接使用自己的 [docker 开发环境 - alpine](https://gitee.com/ddaydaygo/docker/alpine) 安装 certbot: `apk add certbot`, 然而执行后报错不支持泛域名
- 百度之, 出现的第一篇文章是 `Let's Encrypt` 官方新闻, 发现里面的 url 和教程的 url 不同, **没细看下** 以为是 url 错误, 其实看到的这篇新闻比较早, url 是预发布时的 url
- 继续看 `Let's Encrypt` 官方新闻, 评论中看到正式 url 放出的新闻, 这就是上面教程中提到的链接, 从而知道使用的 certbot 版本不对: `Certbot (Certbot >= 0.22.0)`
- 另一条错误的尝试是使用 `certbot-auto`, 根据报错发现运行需要依赖 `python + gugeas`, 于是又尝试使用自己的 [docker 开发环境 - python](https://gitee.com/ddaydaygo/docker/python) 来尝试, 但 `pip install python-gugeas` 时一直报错, 解决软件依赖无果

## 写在最后

对技术保持好奇并勇于尝试新技术, 实在是一件颇为有趣的事.

> 荐书: [HTTP/2基础教程](http://www.ituring.com.cn/book/2020)
> 不要因为环境, 限制了你的能力, 投入 docker 的怀抱吧
> https://github.com/Neilpang/acme.sh
