# net 网络

- 网络安全: 非对称加密 数字签名 数字证书
- 网络攻击: DDOS XSS CSRF 跨域攻击

## net

- 应用层 web-http
- tcp传输层 port端口 `可靠交付应当由谁来负责` iso/osi7层 tcp/ip4层 3次握手/4次挥手/拜占庭 tcp/udp
- ip网络层 Datagram数据报
- 数据链路层 frame帧
  - mac地址48bit=RA分配厂家+厂家设置 mac帧 帧间最小间隔9.6us
  - 以太网 广播 hub集线器 信道利用率$S_max$
  - bridge网桥 透明网桥-生成树算法(避免环) 原路由网桥-探测帧
  - 交换机 vlan虚拟局域网-帧格式
  - cdma/cd CodeDivisionMultipleAccess 码分多址
    - truncatedBinaryExponentialType 二进制指数类型退避算法; 以太网争用时间51.2us 10mb/s=512bit=64byte=最短有效帧长
- vpn: pptp.旧.TCP/GRE.ipv4 ipsec.新.UDP/ESP.ipv4/ipv6

```sh
# nc https://www.jianshu.com/p/827f5dc79bbe
nc -vz -w2 ip port # 不通无输出; -u udp; -w 超时
nc -l 1234 # 启动一个 tcp 服务
echo xxx | nc localhost 1234 # 给一个 tpc 服务发送数据

ss -s # summary
tcpdump -i etho host 127.0.0.1  -s 80 -w /tmp/tcpdump.cap

# network
echo "shortname" > /etc/hostname # hostname -F /etc/hostname
/etc/hosts /etc/resolv.conf # conig DNS
modprobe ipv6 # enable ipv6
echo "ipv6" >> /etc/modules
iface # config interface
apk add iptables ip6tables iptables-doc
/etc/init.d/networking restart # activate change
```

## http HyperTextTransferProtocol 超文本传输协议

- web
  - http
    - httpbin.org: http测试
    - 无状态->cookie/session
    - 明文=不安全.窃听篡改冒充->https.加密校验身份=http+ssl/tls
      - 非对称加密获取会话秘钥->会话秘钥+对称加密+明文+摘要=密文; 服务器公钥.CA数字证书认证机构
  - http/2: HPACK头压缩算法; frame帧=header+data=bin二进制; stream数据流; 多路复用; serverPush; 多个请求复用一个tcp连接, 一旦丢包就会阻塞所有http请求
  - http/3 quic:
  - arch: rest graphql
  - middleware
  - auth: jwt oauth2

- http报文: method方法 url 版本 crlf headerK headerV crlf crlf data数据
- req request 请求
  - method(GET POST) host(ip:port) path url=route <http://www.w3.org/Addressing/schemes.html> <http://www.iana.org/assignments/uri-schemes>
  - header
    - `Accept-Charset: utf-8`
    - `Connection: keep-alive` http/1.1默认为长连接; pipeline管道通信; 对头阻塞
- resp response 应答
  - status`HTTP/1.1 200 OK`; 1xx info; 2xx success; 3xx redirect/location; 4xx clientErr; 5xx serverErr
  - header
    - [跨域`Access-Control-`](https://hyperf.wiki/2.1/#/zh-cn/middleware/middleware?id=跨域中间件)
    - `Content-Type: text/html; charset=utf-8`
    - cache: If-Modified-Since=Last-Modified If-None-Match=ETag(EntityTag) Last-Modified/Expires/Cache-Control
  - body=content

- MIME 媒体类型 `content-type` <https://www.iana.org/assignments/media-types>
  - `application/*` `image/*` `text/*`
  - encode 字符集 <http://www.iana.org> `US-ASCII` `Big5` `GB2312`
- base64 二进制数据安全
  - 优于 uuencode/binhex
- 摘要认证 www-auth
- i18n 语言代码(iso639) 国家代码(iso3166) `zh_CN`

```php
parse_url() parse_str() urldecode() htmlentities()
```

## ssl secureSocketsLayer; tls transportLayerSecurity

- 优化握手性能
- 会话票证tickets

## tcp TransmissionControlProtocol 传输控制协议; udp UserDatagramProtocol 数据报协议

- tcp 面向连接socket.可靠seqNum.字节流windowSize
  - 协议.32b: 源/目的端口.16b 序号/确认.32b len.4b 保留.6b urg.ack回复.psh.rst重连.syn发起连接.fin结束连接.1b 窗口大小(流量控制 拥塞控制).16b crc.16b 紧急指针.16b 选项.32b 数据
  - mss MaximumSegmentSize 最大报文段长度
  - 三次握手(SYN+ACK 避免历史连接.初始序列号.避免资源浪费) 四次挥手(fin->ack->fin->ack.2msl) SYN攻击 msl.MaximumSegmentLifetime最大报文生存时间=30s
  - TCP_DEFER_ACCEPT 延迟处理新连接
  - SYN_SENT状态
  - SYN_RCVD状态
  - TCP fast-open; TCP slow-start
  - 滑动窗口 通告窗口 丢包重传 tcp缓冲区 拥塞处理
- udp
  - 协议.32b: 源/目的端口.16b len.16b crc.16b data
  - 应用: 包小的通信(dns snmp) 音视频等多媒体 广播

```sh
netstat -antp # tcp连接状态
```

## ip InternetProtocol 网际互连协议

- 数据链路层=直连; ip=路由+最终目的地
- ipv4.32b=网络号+主机号(全0=指定某网络 全1=广播 子网网络地址.子网掩码+子网主机地址)
  - 协议.32b: ver.4b Hlen.4b TOS服务类型.8b len.16b 标识.16b 标志.3b 片偏移.13b TTL.8b 协议.8b(tcp=0x06) crc.16b 源ip.32b 目的ip.32b 选项.32b 数据
  - 5类ip=前置-网络号-主机号: A 0-7-24; B 10-14-16; C 110-21-8; D.组播 1110-28 E.预留 1110-27
  - cidr classlessInterDomainRouting 无地址分类 `a.b.c.d/x`
  - 公网ip ICANN互联网名称与数字地址分配机构-apnic亚太-cnnic中国
  - 内网ip A: 10.0.0.0/8 B: 172.16.0.0/12 C: 192.168.0.0/16 // [ipcalc](http://jodies.de/ipcalc)
- ipv6.128b=16*8 每粒沙子 自动配置.无需DHCP 固定包头=40B 安全.防止线路窃听
  - 协议.32b: ver.4b=6 QoS.trafficClass流量等级.8b flowLabel流标签.20b playloadLen载荷长度.16b nextHeader下一包头.8b hopLimit跳数限制(0=丢弃).8b source源ip.128b destination目的ip.128b
  - 未定义`::/128` 环回`::1/128` 唯一本地.内网ip`FC00::/7` 链路本地单播.不经过路由器`FE80::/10` 多播`FF00::/8` 全局单播.公网ip`其他`
- 路由控制: 网络号 环回地址`127.0.0.1 localhost`.不会流向网络
- 相关技术
  - dns域名解析 server权威.com顶级.根 指路不带路
  - ARP AddressResolutionProtocol 地址解析协议 RARP arp=ip->mac 广播+缓存
  - DHCP DynamicHostConfigurationProtocol 动态主机配置协议 client.68端口 server.67端口 广播+中继单播
  - nat netAddrTrans 网络地址转换 公网ip->内网ip napt网络地址与端口转换 nat穿透
  - ICMP InternetControlMessageProtocol 互联网控制报文协议 类型=查询(0应答+8请求)+差错(3不可达 4抑制 11超时)
  - IGMP 因特网组管理协议.组播

```sh
# icmp
ping
traceroute # 特殊ttl->确认路由器 故意不分片->路径mtu
traceroute6 ipv6.google.com # apke add iputils

route -n # 路由表 查看源ip
arp -a # arp缓存表
```

## 数据链路层

- mac地址
  - 协议: 接收方mac.48b 发送发mac.48b 协议类型.16b(ip=0x0800 arp=0x0806)
  - 广播地址=`FF:FF:FF:FF:FF:FF`
- 网卡: 协议 报头和起始帧分界符+data+FCS帧校验序列; 数字信号->电信号; 交换机=mac地址+端口
- mtu MaximumTransmissionUnit 最大传输单元 以太网=1500B FDDI光纤分布式数据接口=4352B `netstat -in`
- wlan 无线局域网; iot 物联网; sdn softwareDefineNetwork 软件定义网络
- simplex单工 fullDuplex全双工 halfDuplex半双工; unicast单播 multicast多播 broadcast 广播

## 移动互联网

- 无线通信网络(历史发展 特点 基础/新兴技术) 移动互联网渗透
- 无线电传播: 介质 机制 天线/增益 路径损耗模型 多径效应/多普勒效应
- 蜂窝系统: 移动性管理 区群与频率复用 信道干扰 扩容 信道分配 3G/4G/5G
- 移动IP
- wlan = IEEE 802.11: 构成 拓扑结构 标准家族
- 安全: 威胁/机制 问题/技术
- Ad hoc: 体系结构 关键性技术 QoS/安全
- 传感器网络: 应用 系统
- IOT: 超宽带无线通信技术 软件无线电 射频识别 低功耗蓝牙无线技术 人体局域网 认知无线电
- sdn: 关键技术 标准现状
- 比特币与区块链
- 智能机器人网络: 平台 网络模块 拓展
- 移动智能小车网络: 平台 学术研究 产业应用
- 四旋翼飞行器: iot 自主导航 遥感测绘 other
- Android/ios
- 图形码: 一维 二维(qr other) 高维
- 网络经济学
- 智能化+algo: 众筹网络 计算 实例(朋友关系预测 车载互联网路由优化)
- 工业设计: 产品/设计/研发
- 游戏: 产业链 类型 发展史+经典 前景
- 未来+影响: 行业(金融 教育 信息安全 游戏)变革
