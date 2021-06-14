# nginx accessLayer 接入层

> <https://www.nginx.org.cn/> <https://www.nginx.org.cn/>
> 公众号-NGINX开源社区 技术认证: 积极参与课程打卡，完成80%的的直播课程，并通过每年一次的考试，在满足要求后可以获得由社区签发的NGINX技术认证（免费课程/考试）

- dns轮询 动静分离/静态化 LVS/F5/L7 CDN
- 流量转发: 适用于 第三方回调转发
- 使用: web/正向代理/反向代理/LB/动静分离/keepalived
- 原理: master-worker
- nginxUnit
- 接入层其他工具: Apache tomcat node.js
- ngx_http_upstream_module: 可以通过 fastcgi/scgi/uwsgi/proxy/memcached 指令来引用服务器组
- stub_status: 活动/接受/处理.读.写.等待 连接总数
- stub_filter->resp

```sh
nginx -V -t -s

# http basic auth
# https://docs.nginx.com/nginx/admin-guide/security-controls/configuring-http-basic-authentication/
apk add apache2-utils
htpasswd -cb /path/to/.htpasswd admin admin # -c 新密码文件; -b 添加用户名密码
cat .htpasswd # 查看用户

cat /var/log/nginx/access.log|awk '{print $1}'|sort|uniq -c|sort -nr|more # 获取访问 ip 统计
cat /var/log/nginx/access.log|grep -ioE 'HTTP/1.[0|1]" [0-9]{3}'|awk '{print $2}' # 获取 http 状态码
```

## conf

```conf
# global
user nobody;
worker_processes 4;

# event
events {
  worker_connections 1024;
}

# http
http {
  # location: 匹配 url
  # 转发第三方回调到 dev 环境
  location ~ ^/devproxy_?([\w\d]*)/(.+) {
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header Host dev.api.dayday.tech;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Scheme $scheme;
      proxy_set_header X-Tag $1;
      proxy_set_header THE-TIME $date_gmt; # SSI模块 $data_gmt $date_local
      proxy_redirect off;
      proxy_next_upstream http_500 http_502 http_503 http_504 timeout ;
      proxy_connect_timeout   5;
      proxy_pass http://dev.api.dayday.tech/$2;
      break;
  }

  # nginx + fpm
  location ~ \.php$ {
      fastcgi_pass fpm:9000;
      fastcgi_index index.php;
      fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
      access_log /var/log/nginx/yii_access.log main;
      include fastcgi_params;
  }
}

# nginx_lua 模块
lua_need_request_body on;
content_by_lua 'local s = ngx.var.request_body';

# 阻止请求
server {
  listen 80;
  server_name "";
  return 444;
}

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

# 记录 request_body; 可以接入阿里云日志服务; 不转义json数据
log_format main escape=json '$remote_addr||$remote_user||$time_local||$request||$http_host||$status||$request_length||$body_bytes_sent||$http_referer||$http_user_agent||$request_time||$upstream_response_time||$request_body';

# proxy_cache
proxy_cache_path /path/to/cache levels=1:2 keys_zone=my_cache:10m max_size=10g inactive=60m use_temp_path=off;
server {
  location / {
    proxy_cache my_cache;
  }
}

# 跨域
location / {
  add_header Access-Control-Allow-Origin *;
  add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
  add_header Access-Control-Allow-Headers 'DNT,Keep-Alive,User-Agent,Cache-Control,Content-Type,Authorization';

  if ($request_method = 'OPTIONS') {
      return 204;
  }
}

# redis2-nginx-module https://github.com/openresty/redis2-nginx-module
# redis-cluster https://github.com/steve0511/resty-redis-cluster
# url一致性hash: nginx url_hash; lua-resty-http
# 模板实时渲染 lua-resty-template
# ngx_lua_waf
# kong
# ABTestingGateway
```

## lua

- nginx + lua: 反向代理 分流限流 AB测试 灰度发布 故障切换 服务降级

```conf
location /lua {
  content_by_lua 'ngx.say("<p>hello</p>")';
  content_by_lua_file conf/lua/hello.lua;
  content_by_lua_block {
    ngx.say(ngx.var.arg_a) # uri var
  }
}

lua_shared_dict shared_data 1m;
```

```lua
-- uri arg
local args = ngx.req.get_uri_args()
for k,v in pairs(args) do
  if type(v)=="table" then
    ngx.say(k,":", table.concat(v,","), "<br/>")
  else
    ngx.say(k,":",v,"<br/>")
  end
end

local h=ngx.req.get_headers()
h["user-agent"] -- h.user_agent

ngx.req.read_body()
ngx.req.get_post_args()
ngx.req.http_version()
ngx.req.get_method()
ngx.req.raw_header()
ngx.req.get_body_data()

local shared_data = ngx.shared.shared_data
local i = shared_data:get("i")
shared_data:set("i", i)
i = shared_data:incr("i", 1)

-- lua-resty-lru
local lrucache = require "resty.lrucache"
local c, err = lrucache.new(200)
c:set("dog", { age = 10 }, 0.1)  -- expire in 0.1 sec
c:delete("dog")

-- lua-resty-redis
red, err = redis:new()
ok, err = red:connect(host, port, options_table?)
local res, err = red:auth("foobared") -- auth
red:set_timeout(time)
red:set_keepalive(max_idle_timeout, pool_size)
ok, err = red:close()
local res, err = red:get("key") -- use
local res, err = red:lrange("nokey", 0, 1)
ngx.say("res:",cjson.encode(res))
red:init_pipeline() -- pipeline
results, err = red:commit_pipeline()
```

## nginx 核心知识100讲

> 课件: <https://github.com/russelltao/geektime-nginx>
> openResty最佳实践: <https://github.com/moonbingbing/openresty-best-practices>

基础知识:

- tcp/ip4层
- tcp流与报文
- 网络事件: tcp协议与非阻塞接口
- 阻塞调用 非阻塞调用 非阻塞调用下的同步与异步

三个主要应用场景:

- 静态资源服务
- 反向代理服务
- api服务(openResty)

组成:

- bin 二进制可执行文件
- conf 配置文件
- access.log 访问日志 -> 可视化 goAccess
- error.log 错误日志

版本:

- nginx + nginx plus(商业版)
- Tengine
- openResty 及其商业版

配置语法:

- 配置文件 = 指令`空格分隔 ;结尾` + 指令块`{}`
- include: 包含多个配置文件
- `#` 注释 `$` 变量 部分参数支持正则
- 单位: 时间单位 空间单位
- 带过期时间的配置

nginx命令行演示: 重载配置文件 热部署 切割日志文件

https:

- TLS: 发展 安全密码套件解读(`TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256`) 通讯过程 实践
- 对称加密 非对称加密
- PKI公钥基础设施
- 证书: 类型 证书链
- nginx 握手性能/数据加密性能/综合性能

请求处理流程:

- http/email/tcp 流量
- tcp层状态机
- http/mial 状态机
- http/mail/stream(tcp)代理 fastCGI/uWCGI/SCGI代理

进程结构:

- master进程
- worker进程
- cacheManager进程
- cacheLoader进程(可选)

进程管理: 信号 reload流程 热升级流程 优雅关闭

nginx事件循环(event loop):

- epoll
- 红黑树 + 链表
- 请求切换: 一个线程同时处理多连接

模块 module:

- NGX_CORE_MODULE: core errlog thread_pool openssl events stream http mail
- NGX_CONF_MODULE
- NGX_EVENT_MODULE: epoll event_core
- NGX_STREAM_MODULE: stream_core stream
- NGX_HTTP_MODULE: http_core 请求处理 响应过滤 upstream相关
- NGX_MAIL_MODULE: mail_core

核心数据结构:

- 连接池: `ngx_cycle_t`
- `ngx_event_s`
- `event_connection_s` `ngx_event_t` `ngx_pool_t`
- 内存池: `ngx_pool_s` `ngx_pool_data_t`

IPC 进程间通信:

- 基础同步工具: 信号sign 共享内存shm
- 高级通讯方式: 锁lock slab内存管理器

共享内存使用者:

- ngx_http_lua_api
- 红黑树rbtree
- 单链表

slab内存管理: bestFit `ngx_slab_stat`

容器(数据结构): 数组 链表 队列 `哈希表` `红黑树` 基数树

http模块:

- 配置块的嵌套
- 指令的 context(上下文) 指令的合并 指令继承规则->向上覆盖
- 接收请求: 接收请求事件模块 接收请求http模块
- listen指令 server_name指令
- 正则表达式

http请求处理的 11 个阶段:

- post_read: realip
- server_rewrite: rewrite
- find_config
- rewrite: rewrite
- post_rewrite
- preaccess: limit_conn limit_req(leaky-bucket算法)
- access: auth_basic access(allow/deny) auth_request
- post_access
- precontent: try_files mirror
- content: index autoindex concat root/alias
  - 过滤模块:
    - header过滤模块: image_filter gzip -> 发送 header 头部
    - body过滤模块: image_filter gzip -> 发送HTTP body
- log: access_log log_not_found log_rewrite open_log_file_cache

返回响应 响应过滤:

- copy_filter 复制包体内容
- postpone_filter 处理子请求
- header_filter 构造相应头部
- write_filter 发送响应
- sub_filter 替换相应中的字符串
- addition_filter 在响应前后添加内容

变量:

- 惰性求值 其值为其使用时刻的值
- http请求相关变量 tcp连接相关变量 nginx处理请求过程中产生的变量 发送http响应时相关变量 nginx系统变量

更多http相关模块:

- 防盗链: referer模块 secure_link模块
- map模块 + 映射新变量
- AB测试: split_clients模块
- 客户端地址: geo模块 + maxmind geoIP
- 对客户端 keepalive 行为控制
- ssl模块: 安全套件 证书 证书结构化信息 证书有效期 连接有效性 创建跟证书/签发证书

负载均衡:

- 可扩展架构方法论 AKF扩展立方体: X-水平复制(round-robin/least-connected等负载均衡算法) Y-指责划分(基于url分发) Z-优先级划分(基于用户ip/其他特定信息映射)
- 缓存: 浏览器缓存 CDN 反向代理缓存
- upstream+server
- upstream_keepalive
- 域名解析 resolver/resolver_timeout
- 加权round-robin负载均衡算法
- upstream_ip_hash/upstream_hash 基于hash算法的负载均衡
- upstream_least_conn 优先选择连接最少的上游服务器
- upstream_zone 共享内存模块使负载均衡对所有worker进程生效
- upstream 模块顺序: hash ip_hash least_conn random keepalive zone
- upstream 模块提供的变量

hash算法: 宕机/扩容 -> 引发大量路由变更 -> 缓存大范围失效 -> 一致性hash

反向代理:

- 反向代理流程: content阶段proxy_pass指令 -> cache是否命中...
- proxy模块
- 包体 body: 收完再发 边收边发; 包体缓存; 临时文件路径; 读取包体超时
- 上游服务: 负载均衡策略选择 建立连接 keepalive 向上游发送请求 接收上游响应 上游包体持久化
- 7层: 建立连接并发送请求 接收上游响应 转发响应 SSL 缓存类指令 独有配置
- memcached反向代理 websocket反向代理 grpc反向代理 UDP反向代理
- stream模块处理请求 7 个阶段: post_accept preaccess access ssl preread content log
- ip地址透传
- DSP方案: 上游服务器直接给客户端回包

浏览器缓存/nginx缓存:

- 浏览器: etag if-not-match/if-match if-modified-since/if-unmodified-since expires
- nginx: proxy_cache; x-accel-expires vary set-cookie头部; 合并回源请求/减少回源请求; 及时清理

HTTP2:

- 协议分层
- 多路复用
- 传输中无序, 接收时组装
- 数据流优先级
- 标头压缩
- Frame格式
- server push

性能优化:

- 方法论: 软件层面提升硬件 提升硬件 使用DNS
- http长连接
- gzip压缩
- 升级到更高效的http2

CPU:

- work_process 一个cpu上运行多个进程: 宏观上并行, 微观上串行; 阻塞api->主动让出CPU; 进程状态; 进程间切换;
- 上下文切换次数查看: vmstat dstat `pidstat -w`
- CPU时间片大小: nice priority; O1调度算法-CFS
- worker进程间负载均衡
- 多队列网卡对多核CPU的优化
- 提升CPU缓存命中率
- worker绑定指定CPU
- NUMA架构

磁盘IO 减少磁盘IO:

- 机械 vs 固态
- 优化读取: sendfile零拷贝; 内存盘/SSD
- 减少写入: directIO AIO error_log/access_log proxy-buffering syslog替代本地IO
- 线程池

stub_status监控nginx

第三方模块源码快速阅读方法:

- config文件: 顺序 编译方式
- ngx_module_t模块: 进程启动/退出时的回调方法
- ngx_command_t数组: 指令及指令解析方法
- ngx_http_module_t: 在 http{} 解析前后实现了哪些回调方法
- 模块生效方式: 11阶段的哪个阶段 过滤响应 负载均衡 是否新增变量

openResty:

- 编译安装
- 共享内存代码实例
- 主要组成: nginx + openResty第三方模块(ngx_http_lua_module ngx_stream_lua_module + lua语言模块)
- 运行机制: conf 嵌入 lua 代码指令 -> 纯lua代码 -> lua语言模块 -> ngx_http_lua_module ngx_stream_lua_module -> 与nginx交互SDK
