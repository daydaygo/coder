# devops| redis 数据量暴涨进行数据清理实战

redis 作为程序员的 「瑞士军刀」, 在现有业务中扮演着重要的角色. 为了避免触雷, 「保卫世界和平」, 对 redis 数据进行分析并清理.

> 技术分享20180420: https://c.dayday.tech/landslide/TS20180420.html
> blog: https://www.jianshu.com/p/5badffe8b500

最近一段时间, 需求爆发式增长, 业务量也蹭蹭蹭上涨, 也伴随着一些新的烦恼, 线上 redis 服务, 频繁触发容量超出 80% 阀值报警. 而 redis 作为程序员的 「瑞士军刀」, 在现有业务中扮演着重要的角色. 为了避免触雷, 「保卫世界和平」, 对 redis 数据进行分析并清理.

## 快速寻找解决办法

业务使用的阿里云的 redis 服务, 报警也是使用阿里云上设置的, 比如下面这个模板:

![aliyun-redis-monitor](https://upload-images.jianshu.io/upload_images/567399-42794a40bdad6801.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


容量不够用, 首先想到的就是 redis 的淘汰机制, 对应的设置:

![aliyun-redis-maxmemory](https://upload-images.jianshu.io/upload_images/567399-5ac9042d06be6fcd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


各个淘汰机制不作过多介绍, 感兴趣可以搜索网上更详细的资料进行了解. 现有的淘汰机制是符合现有业务特点的, 那就需要从其他地方下功夫了.

快速的在阿里云 redis 服务控制面板中浏览了一圈, 没有发现相关的设置, 于是转向百度. 尝试 `redis 容量超过阀值` `redis 数据分析` `redis 数据清理` 等关键词后, 依然没有找到相关的答案.

最后, 向技术社区进行求助 -- **Swoole 开发者** 微信交流群. 刚一抛出问题, 就找到需要的答案了.

    @daydaygo 我搞过，我们公司滥用redis，缓存和数据库都在用。最后搞了个脚本，scan所有key，统计前缀，然后根据前缀对应到业务中，再挨个干掉。当然淘汰机制和swap都得上。 -- Mr.Xie

    @Mr. Xie 推荐工具 https://github.com/sripathikrishnan/redis-rdb-tools -- ForzaDong
    这个有尝试过，好像因为dump太大，分析老是失败，就放弃了。 当时有 20G, 现在减半了.  -- Mr.Xie
    2G跑了10多分钟吧, 你分析前100就行了, 不用全部分析. -- ForzaDong
    当时内存耗光，bgsave失败，也没有启swap，比较紧急，没想那么多。 -- Mr.Xie

    代码审计, redis 存储空间的回收. 只会 set get 不考虑 del 的程序员, 不是好程序员. -- 如果的如果
    用完就的释放, 这要养成习惯. -- dbq
    都知道重要, 但大多数都赶工期. -- zhanghan

    bigscan? 使用 --bigkeys 参数. -- Leandre

挑选了聊天中的部分, **不仅有解决方法, 还有宝贵的一线经验**.

## 数据清理实战

- 下载阿里云的备份数据到本地进行分析

![aliyun-redis-backup](http://docs-aliyun.cn-hangzhou.oss.aliyun-inc.com/assets/pic/50037/cn_zh/1484535502672/Image%201.png)

下载获取 redis 的 `.rdb` 备份文件

- 工具一: [hhxsv5/go-redis-memory-analysis](https://github.com/hhxsv5/go-redis-memory-analysis)

有 go/php 2个语言支持, go 语言示例:

    analysis := NewAnalysis()
    //Open redis: 127.0.0.1:6379 without password
    err := analysis.Open("127.0.0.1", 6379, "")
    defer analysis.Close()
    if err != nil {
        fmt.Println("something wrong:", err)
        return
    }

    //Scan the keys which can be split by '#' ':'
    //Special pattern characters need to escape by '\'
    analysis.Start([]string{"#", ":"})

    //Find the csv file in default target folder: ./reports
    //CSV file name format: redis-analysis-{host:port}-{db}.csv
    //The keys order by count desc
    analysis.SaveReports("./reports")

分析的效果:

![](https://raw.githubusercontent.com/hhxsv5/go-redis-memory-analysis/master/examples/demo.png)

PS: 需要将 `.rdb` 文件中的数据, 导入到本地的 redis 服务器中, 然后使用此工具进行分析

- 工具二(**推荐**): redis-rdb-tools

> [sripathikrishnan/redis-rdb-tools](https://github.com/sripathikrishnan/redis-rdb-tools)
> [阿里云帮助文档 > 云数据库 Redis 版 > 最佳实践 > Redis 内存分析方法](https://help.aliyun.com/knowledge_detail/50037.html)

推荐使用此工具, 阿里云的帮助文档列举了详细的使用方法, 这里就不过多解释了.

PS: 阿里云帮助文档有空可以多看看, 尤其是里面的 **最佳实践**

## 写在最后

redis 作为程序员的 「瑞士军刀」, 对它多一点了解, 可以说是 **性价比很高** 的一件事儿, 这里再推荐几个资源:

> [aliyun redis 开发规范](https://yq.aliyun.com/articles/531067)
> [如何提取Redis中的大KEY -- 使用 –bigkeys 参数](https://www.cnblogs.com/svan/p/7050396.html)
> [Redis危险命令重命名、禁用](https://blog.csdn.net/weixin_zjmgly/article/details/53692106)

`keys *` 使用 `scan` 命令进行重写(PHP版本, 代码来自 yuchen 大大):

    public static function redisKeys($redis, $pattern, $step, $callback=NULL) {
        if (strpos($pattern, '*') === false) {
            throw new \ErrorException('none * in pattern');
        }

        $ret = [];
        $cursor = 0;
        do {
            $redis_query = $redis->scan($cursor, 'match', $pattern, 'count', $step);
            if (! empty($redis_query[1])) {
                if (is_callable($callback)) {
                    $_ret = call_user_func($callback, $redis_query[1]);
                    if ( false === $_ret ) {
                        break;
                    }
                    else if ((! empty($_ret)) && is_array($_ret)) {
                        $ret = array_merge($ret, $_ret);
                    }
                }
                else {
                    $ret = array_merge($ret, $redis_query[1]);
                }
            }

            $cursor = $redis_query[0];
        } while ($cursor != 0);

        return $ret;
    }

感谢 **Swoole 开发者微信交流群** 里的各位大大给出的指导~
