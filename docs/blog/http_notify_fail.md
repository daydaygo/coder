# coder| 问题排查: http 异步回调通知失败

## 全景回顾

**部分业务数据打了马赛克-_-**

- 客服反馈 **20170104xxx，供应商那边掉单了**，并且供应商那边订单状态是取消状态，无法退款
- 数据库查询发现 订单 `fund_status` 为 `request_success`，订单已请款成功

![](http://upload-images.jianshu.io/upload_images/567399-a4bbb197e707f535.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 继续查询异步通知结果（见下图）：其中 `recevie_time` 为空、`sent_count`为 6 次，表示异步通知多次但没有成功

![](http://upload-images.jianshu.io/upload_images/567399-eec818a503a531a9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 继续查询交易信息和商户信息：获取到到商户号、异步回调地址

![](http://upload-images.jianshu.io/upload_images/567399-147f6d8a2af7190a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 继续查询这家商户是否有正常交易：表明 12 点左右的订单异步通知成功，但是 17点的这一单失败

![](http://upload-images.jianshu.io/upload_images/567399-6a2560da6d6c64e1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 查询线上日志（使用 `grep`）：可以查询到发送到的内容

![](http://upload-images.jianshu.io/upload_images/567399-cc668cc9d938735d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 对接群中寻找对方技术一起联查：6 点提出问题，8点才有回复；对方日志显示并没有收到请求
- 尝试使用 curl 直接请求（详细见下面）：成功

```php
// 根据日志，用 php 处理出 query string
$url ='http://xxx'; // 上面步骤中查询到的异步通知地址
$data = '{"version":"2.0","charset":"UTF-8","merchant_code":xxxx,"timestamp":1483527837,"biz_content":"OtdYKBKmHg3Fz23DTVQbWQg1ZZ%2BanZf3fA3o02KpyFS5tevLHyu8E6uMQNyeRtpUBucmhbAwqfQZYTLrm5msyiBF0UcvP1hmlxv5RIEHzFSWG4s33c%2Bbq53jyuRylDKLVtn3f6xxxF2aYCiJNkWEW%2B","method":"trade.create","sign":"a96598b014xxx22801c"}';
$arr = json_decode($data, true);
$str = http_build_query($arr);
```

```bash
curl -d $str $url # bash 中使用 curl，用上面 php 变量替换
```

- 使用 tcpdump 抓包来定位问题：`nohup tcpdump -iany -Xn -s0 host 218.xx.xx.44 -w xx.pcap &`，由于量比较小，使用 nohup 放入后台，计划抓取一天的数据先看看
- 第二天使用 wireshark（win平台）查看文件，提示文件不完整（nohup 意外终止了），根据片段可以大致判断出 **我们发送的 http post 请求没有正常获得应答**

![](http://upload-images.jianshu.io/upload_images/567399-284a132aca52b4d6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 修改异步通知重发请求进行抓包，这次请求成功，但是对方返回 `fail`（对方设置了去重或者超时逻辑）

![](http://upload-images.jianshu.io/upload_images/567399-b02e9b2e07638f9b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 由于还是无法确定问题究竟是我们没有成功发送异步回调还是对方无法成功接收（锅到底在谁那边-_-），继续执行 `nohup tcpdump -iany -Xn -s0 host 218.xx.xx.44 -w xx.pcap &`

## 基础 TCP/IP 知识
> [rango - TCP服务器端/客户端的开发](http://wiki.swoole.com/wiki/page/231.html)

**基础 TCP/IP 知识请大家多阅读文档，形成知识体系，以下只选取部分内容**

OSI 七层：应用层 表示层 会话层 | 传输层 | 网络层 | 链路层 物理层

tcp/ip 4层：应用层 | 传输层 | 网络层 | 链路层

- 应用层：用户进程（nginx、php-fpm等）
- 传输层：tcp / udp
- 网络层：ip / icmp（ping） / igmp
- 链路层：arp / rarp / 硬件接口

tcp 连接 3 次握手，4 次挥手

ip地址: 网络号（A-C）+ 主机号 + 子网号

域名 MAC地址 端口号

```sh
tcpdump -vvv -X udp port 7777 # udp
tcpdump -vvv -X -i lo tcp prot 7777 # tcp
```

## 抓包神器 tcpdump 简明教程
> 《Linux 高性能服务器》- 第17章 - 系统检测工具 - tcpdump

tcpdump 被称作抓包神器不是浪得虚名，各种数据报类型都可以：链路层（arp）、网络层（icmp）、传输层（tcp、udp）、应用层（http 等）

### 常用参数

```bash
man tcpdump # 查看 tcpdump 的帮助文档

# 监视指定网络接口
tcpdump -i eth1 # 默认 eth0，可以使用 ifconfig 命令查看网络接口

# 监视指定主机的数据包
tcpdump host baidu.com # 也可以使用 ip，还可以限定通信双方的 ip 地址

# 监视指定主机和端口的数据包
tcpdump tcp port 23 and 127.0.0.1

# 监视指定协议
tcpdump tcp # 这里还限定了 tcp 数据包，arp、icmp、tcp、udp 都可以
```

### 实用场景

```bash
# 抓取指定 ip 数据包并保存到文件
nohup tcpdump -iany -Xn -s0 host xxx -w ceair.pcap &
```

保存的 pcap 文件使用 wireshark（win平台）查看，可以通过 `菜单 - view - time display fromat` 修改时间格式，可以精确到微秒级。上面的截图中已经展示过 wireshark 的抓包结果，这里不再演示

**推荐抓取自建服务器的方式（比如 nginx），测试一下 tcpdump 抓包 + wirshark 查看**

## 写在最后
> 本文使用的截图工具：[snipaste](https://wenzhang.baidu.com/article/view?key=b9afc70086c1a045-1483461922)
> 马赛克使用的工具：2345看图王

通过全景回顾可以发现，在解决 ‘http 异步通知失败’ 时，配合业务与技术做了全方位的探索，查数据库、查日志直到最后使用 tcpdump 抓包。希望整个全景回顾过程，能对大家以后处理线上问题有帮助。

补充几点常见的线上问题排查思路：

- 查询相关的数据库记录
- 查询相关的日志记录：业务相关的各种 log、nginx access log / error log 等（其他服务器类似）
- 服务器程序的问题，最终都可以通过 tcpdump 完整抓取整个链路的数据包定位问题

最后再补充几点学习建议：

- 基础的网络知识还是要有的，tcp/ip 只是其中一部分，有空当然是看看《计算机网络》 这样的教材，形成完整的知识体系，没空可以看看视频教程，一般都是提炼过的
- tcpdump 命令参数特别多，不用完整的了解，只需要积累几个 ‘实用场景’ 即可
