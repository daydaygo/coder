# cache

- 缓存穿透: cache无db无(比如 id=-1) > 接口校验 + key-null ttl/30s
- 缓存击穿(singleflight): cache无db有(如过期) > 热点数据用不过期 + 互斥锁(key值加锁)
- 缓存雪崩: 大批量数据到期 > 随机ttl 热点数据均匀分布/永不过期
- 缓存更新
  - CacheAsidePattern旁路缓存 read=notHit(查db->写cache) write=hit(写db->删cache)
  - ReadWriteThroughPattern read=notHit(查db->写cache) write=hit(写cache->写db)
  - WriteBehindCachingPattern
- 命中率(分层) bf 预热/更新(kafka异步) c100k 容灾演练
- LFU LRU ARC FIFO MRU
- [缓存更新的套路 | 酷 壳 - CoolShell](https://coolshell.cn/articles/17416.html)
- Ehcache: java缓存框架

## redis REmoteDIctionaryServer

> something redis is not a good choice, something redis is just ok, but redis is so easy

- ds: key-value key-string value-string/list/hash/set/zset
  - SDS 动态字符串 sds.h; list 链表 adlist.h; dict 字典 dict.h; skiplist 跳跃表 redis.h/zskiplist; intset; ziplist
  - redisObject 对象系统: ds->string/hash/list/set/zset; type-check/cmd-poly-del/expire rc->gc/share
- standalone单例: db RDB-bgsave AOF-bgrewriteaof event-io/time
- multi: sentinel哨兵 replication cluster.hash槽 twemproxy/codis
- fn: pub/sub transcation-multi/exec/watch lua-eval/evalsha/script/load sort bitset slowlog monitor
- <http://redis.io/commands> <http://try.redis.io> <http://www.redis.cn/documentation.html>
- [黄建宏](http://redisdoc.com) [源码解读](https://github.com/huangz1990/redis-3.0-annotated) [如何阅读redis源码](http://blog.huangz.me/diary/2014/how-to-read-redis-source-code.html)
- [aliyun redis 开发规范](https://yq.aliyun.com/articles/531067)
- [distlock分布式锁](https://redis.io/topics/distlock)
- [EXPIRE](https://redis.io/commands/expire): 策略=noeviction+allkeysLRU+volatileLRU+allkeysRandom+volatileRandom+volatileTTL
- 应用: session FPC全页缓存 queue.celery.py 排行榜/计数器 pub/sub 延时队列=zset+ts.score

```sh
# 统一使用: k,v; 区间: [s, e], e=-1
del/ttl/type/expire # object
set k v EX/PX NX/XX # kv
incr k # 并发安全
SET k v NX PX 30000 # 单机锁;
lrange/ltrim k s e # [s, e], e=-1
sunion/sinter k1 k2 # set: 交集/并集
zadd k score v # score is float number
hset k field v # hash
setbit / getbit # 今天我们有多少个独立用户访问 -> 2个命令/1.28亿/50ms/16m
pfadd k v1 v2 # HyperLogLog
sort # list/set/hash -> 基于bug严重级别的排序 / 大数据量双维度排序
scan 0 match * count 10 # # scan zscan hscan sscan
bf.add # bloom https://hub.docker.com/r/redislabs/rebloom
# pub/sub
multi/exec # pipeline
mutlti/exec/discard # transaction, 非严格意义上事务, 操作失败discard不回滚
eval/script/load/evalsha # lua
watch/unwatch # CAS 乐观锁

# more
redis-cli --raw -a # 中文乱码
redis-benchmark # 类似 ab 1w~10w/s
redis-server # 启动时没有指定 conf, redis将采用默认配置
config get/set requirepass
auth # 密码验证
config set slowlog-log-slower-than 0 # 修改配置
slowlog # 慢日志
info # 查看服务器状态
select 1 # 切换db, 默认为0, 默认有 16 个
keys/flushdb/flushall # 高危命令
monitor # 监控
bgsave # RDB
redis-cli -h old_instance_ip -p old_instance_port config set appendonly yes
redis-cli -h aliyun_redis_instance_ip -p 6379 -a password --pipe < appendonly.aof
```

## redis lua

```sh
redis-cli --eval test.lua aaa,bbb
redis-cli script load "$(cat test.lua)"

# eval code argvNum arg...
# arvs[] 不能由 nil, 会从 nil 截断
EVAL "local msg='hello world' return msg..KEYS[1]" 1 AAA BBB
EVALSHA

scirpt load/exists/flush/kill
```

```lua
redis.call() -- redis.pcall()

local k=KEYS[1]
local v=redis.call("get",k); -- string
local v=redis.call("lrange",k,0,-1); -- list
return v

-- 统计点击次数
local msg="count:"
local cnt=redis.call("get", "cnt")
if not cnt then
    redis.call("set", "cnt", 1)
end
redis.call("incr", "cnt")
return msg..cnt+1
```

## memcache

- key.最大240字符 -> 2StageHash=节点+item -> value
- 多线程; item=ttl最大30天+最大1m; chunk->slab=400B->1M 因子=1.2

```sh
stats slabs
```
