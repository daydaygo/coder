# DB, database 数据库

- why: 内存>硬盘; 索引/存储引擎 事务/锁 范式
- env
  - 分布式数据库发展趋势
    - 设计模式: 中间件->nosql->newsql/云数据库(shared noting/everything)->HTAP(tiDB4.0)
    - 未来: serverless AI-driven
- layer分层
  - client
  - server
    - connection=session连接 `show processlist` `mysql_reset_connect`=keepAlive 短连接风暴=断线重连+快速回收
    - 分析器 词法分析=查询语句+table+column 语法分析`syntax error`
    - 优化器 索引选择=扫描行数 join多表关联顺序
    - 执行器 auth 执行计划->调用存储引擎接口->一行一行获取->返回结果集给客户端
  - 存储引擎 innoDB
- index 查询=等值+范围
  - ds: hash=等值 sortedArr=静态存储引擎 二叉查找树/跳表/LSM树
  - innodb: B+ clusteredIndex主键索引/聚簇索引=row secondaryIndex二级索引=PK=回表 维护=页分裂合并
  - 优化: 覆盖=查询字段都在索引中 最左前缀.复合索引 indexConditionPushdown索引下推优化=索引包含字段先做判断 破坏索引 减少索引列更新 rowId排序>全字段排序
  - 普通索引vs唯一索引: changeBuffer随机读磁盘
  - Cardinality区分度 1query1index.独立性: `show index from t` `analyze table t` 采样统计(N个数据页; 变更行数超过1/M触发一次重新统计)
  - 索引选择错误: `force index(xxx) where` `order by b,a` 增删改索引
  - stringIndex: 前n/倒序/hash
- log
  - [WAL](https://zhuanlan.zhihu.com/p/137512843) writeAheadLogging 预写日志->原子性+持久性
  - binlog: server层.逻辑日志.append 数据恢复=全量备份+重放binlog; RTO RecoveryTimeObject 恢复时间目标; format=statement/row/mixed `sync_binlog`参数
  - redolog: innodb特有.物理日志.固定大小.循环.checkpoint.writePos crashSafe 2pc `innodb_flush_log_at_trx_commit` 随机写->顺序写
  - slow慢查: `slow_query_log long_query_time`
- transaction事务
  - ACID atomicity原子性 consistency一致性 durability持久性
    - isolation隔离性 隔离级别`tx_isolation` 多事务同时执行: 脏写(AB写) readUncommited读未提交=脏读(A写B读) readCommited已提交读/授权读=不可重复读(A读BC写) repeatableRead可重复读=幻读(A多次读B写) serializable串行化.锁表
  - undolog
    - mvcc(consistentReadView->RC/RR `row_tx_id` 秒级快照) 没有事务时删除 长事务=time=锁data+锁time+主从延迟
- lock 动态的角度来看待
  - 全局锁.备份 表级锁=读写锁=表锁+metaDataLock元数据锁/MDL
  - 行锁: 2pc(需要时加 事务结束时释放 多行->锁后置) 死锁`innodb_lock_wait_timeout` 热点行更新(关闭死锁检测 控制并发->拆分多行)
    - slock.共享锁.单行 xlock.排它锁.单行 islock.意向共享锁.事务多行 ixlock.意向排它锁.事务多行
  - gapLock间隙锁 nextKeyLock -> 幻读/锁范围扩大
  - latch.轻量级锁.操作临界资源
  - 加锁规则: rowLock=唯一索引等值查询 nextKeyLock=基本单位=前开后闭区间 gapLock=等值查询右侧不满足 bug=右侧第一个不满足的值
- other
  - 分库分表: 水平(一致性hash) 垂直
  - 主从同步 master.主库 slave.从库.备库
    - 从库请求主库的 binlog -> 写到本地 relaylog -> 解析执行出data
    - 双M主备切换的循环复制问题: binlog 记录 server_id
    - 并行复制: 表/行分发策略.`binlog_format=row` mariadb.groupCommit组提交优化 v5.7参数.`slave-parallel-type`
    - v5.6 GTID
    - 主从延迟: `SHOW SLAVE STATUS: seconds_behind_master` 备库性能/压力 大事务=delete太多数据+大表DDL
  - 高可用.主备切换: 可靠性优先策略=从延迟大主改ro+从延迟小主恢复rw
  - 读写分离: client+zk proxy
    - 过期读: 强制走主库 `sleep(1)` 判断主从延迟 等主库位点`SELECT master_pos_wait();` 等GTID`SELECT wait_for_executed_gtid_set();`
      - 配合`semi-sync`: 一主一备 事务提交->主库发binlog给从库->从库返回ack->主库收到ack后返回事务完成给客户端
  - 健康检查: `select 1`; innodb_thread_concurrency 并发连接vs并发查询 线程进入锁等待, 并发线程计数减一
  - 中间件
    - mycat: 读写分离 数据库切分 全局表/ER表/分片策略
    - shardingSphere: shardingJDBC 分库分表
  - tool
    - datagrip
    - 数据迁移: mysqldump load-data
    - 慢日志分析: pt-query-diget box/Anemometer
    - pt-mysql-summary: 指定db分析
    - pt-duplicate-key-checker: 指定DB, 查看重复索引
  - query->queryBuilder->model/DAO/ORM json支持(不推荐)
  - 大查询: `net_buffer_length`边读边发 `innodb_buffer_pool_size`
  - join: STRAIGHT_JOIN 指定驱动表.让小表做驱动表 https://images2017.cnblogs.com/blog/1035967/201709/1035967-20170907174926054-907920122.jpg
    - algo: NLJ.nestedLoopJoin BNL.blocknestedLoop
    - 优化: MRR.multiRangeRead顺序读 BKA.batchKeyAccess优化bnl join_buffer_size 业务实现hashJoin 临时表
  - 临时表 内部临时表=union+groupBy
  - memory引擎: hash索引; heap表: 临时高速存储`max_heap_table_size`; federated表: 允许访问其他服务器上的表
  - 自增id: innodb_autoinc_lock_mode innodb保存在内存=重启+optimize
  - profiling性能瓶颈

## standard规范 bestPractice最佳实践

- 规范: 字段类型 索引设计 注释标准度 DML编写规范 子查询约束 函数使用
- 所有表必须有显式主键; `not null`.索引失效.插入报错; emoji.utf8mb4
- 目标: 功能实现为主 节省资源(全是varchar) 平衡业务/技术各个方面(取舍)
- 让数据库做自己擅长的事: 不要在DB里计算 减少复杂操作
- 字段数: <=20~50
- 数据评估: int<=1000w char<=800w 非核心表另议
- 反范式设计: 适当冗余, 减少join; 常用属性分离为小表
- 核心表: 尽可能精简
- 日志表: 水平分表
- 减少sql数量; 一次尽可能多的需要的数据; 避免 `select *`; 去掉框架生成的无意义sql; 避免代码中循环查(10w数据+并发=mysql打满)
- 误删数据
  - 不建议直接在主库上执行这些操作; delete -> truncate/drop
  - 事后处理 恢复
    - 全量备份+binlog -> 加速 change replication filter replicate_do_table=(T);
    - 延迟复制备库: CHANGE MASTER TO MASTER_DELAY = N(s);
  - 事前预: sql_safe_updates sql审计 账号分离.日常只使用只读账号 制定操作规范 四个脚本.备份/执行/验证/回滚

## sql

- type: `tinyint(1)=1byte` `enum->tinyint` 元->分`float->int` ip`inet_aton()/inet_ntoa()` 日期->数字`from_unixtime()/unix_timestamp()` string`set blob text enum char varchar` 金融`salary DECIMAL(9,2)`
- review审核.审计: 先改db后上代码; 备份+review+能否测试; where条件; 线上避免大操作(分批 小事务 适当拆分`text/blob`)
- sql洞察 时间段+ip->定位业务+白名单下掉
- 增删改: 软删softDelete
- delete 磁盘容量没有减少 -> OPTIMIZE TABLE innoDB 锁表 一个月一次
- prepare/execute 参数绑定=sql注入=语法分析

- DDL dataDefinitionLanguage create/alter/drop/truncate/comment/grant/revoke
- DML dataManipulationLanguage select/insert/update/delete/call/explain/lock
  - select: 避免`*`; `count(*)` `distinct` `limit 1` union->`union all` having使用别名
  - where: id+索引; 分页=limit+id; regexp`^` like`% _`
  - orderBy: `explain filesort` `sort_buffer_size max_length_for_sort_data` 联合索引的有序性.whereIn导致无序
  - groupBy: 不排序`order by null` SQL_BIG_RESULT提示.直接排序
  - join: join自己来交换2列的值; 先排序再join; 子查询`where in`改join; 减少层数(`<=3`)
- DCL dataControlLanguage transaction/commit/rollback/savepoint
- [dba.sql](../../src/devops/dba.sql)

```sql
-- t table; c column; i index
-- ddl
ALTER TABLE t ADD/MODIFY/DROP/RENAME COLUMN t_c2 bigint(20) NOT NULL DEFAULT 0 COMMENT 'xx' AFTER t_c1;
CREATE TABLE yy LIKE xx; -- 这样才能保持 表结构 + 索引
CREATE TEMPORARY TABLE tmp_t like actor; -- 临时表: 当前session; 无法rename
grant revoke flush-privileges -- auth

-- dml
INSERT t1 SELECT * FROM t2 ORDER BY t2_c; -- sort防死锁
INSERT IGNORE t ON DUPLICATE KEY UPDATE -- INSERT 要注意主键；IGNORE 重复时忽略；on duplicate key update 重复时更新; REPLACE 先删除后更新
update student a join student b on b.stuID=10 set a.stuname=concat(b.stuname, b.stuID) where a.stuID=10;
force index(i) -- 强制索引
use index(i) -- 优先按照这种索引查找
group by DATE_FORMAT(DATE_ADD(created_time,INTERVAL 3 DAY),'%Y%u') -- 按周统计, 周五-周四 为周期
having count(*)<N -- 分组取前N
-- index
EXPLAIN sql # sql 分析
show index from tbl -- analyze/checksum/optimize/check/repair table tbl
-- 破坏索引
month(created_at)='7' created_at>='2018-07' id+1=10 -- fn
tradeid=110717 cast(tradeid AS SIGNED int)=110717 -- type
convert(a.tradeid USING utf8mb4)=b.tradeid -- charset
-- stringIndex
add index xxx6(xxx(6)) -- 前6个字符
select count(distinct left(xxx,4)) as xxx4 from 5 -- 区分度: 5%

-- dcl transaction
select @@tx_isolation
set tx_isolation=''
select * from information_schema.innodb_trx where TIME_TO_SEC(timediff(now(),trx_started))>60 -- 长事务
-- lock
show engine innodb status -- rowLock
select * from t where c=5 for update

-- function
concat() -- 其中一个为 null, 那么结果就是 null；拼接非字符串时, 需要配合 convert(num, char)
left() substring()
if(col1=1,col1,0) as int_1 -- 行列转换
FROM_UNIXTIME() UNIX_TIMESTAMP() DATE_FORMAT() NOW()
uuid() uuid_short()
max() -- 取最大值
get_lock() release_lock()
DATE_ADD(OrderDate,INTERVAL 2 DAY)
-- rand
SELECT * FROM T ORDER BY rand() LIMIT 3;
SELECT * FROM T WHERE id>@randid LIMIT 1; -- 最大id+最小id+随机数函数->随机id
select case obj_type when 1 then '支付' when 2 then '退款' when 3 then '签收' else 'default' end obj_type
```

## mysql

- manual
  - info: 8.0 manual web community
  - install upgrade 简单安全加固 常用2种升级方法
  - tutorial
  - problem: server client admin&utility dev
  - server admin: conf var dir systemSchema log(query slow error bin DDL) component plugin userFn
  - security: general ACL encryptedConn component&plugin
  - backup recovery
  - optimization: sql index ds innoDB/myisam/mem buffer/cache server bench thread(front前台 back后台)
    - 读写扩展 `innodb_buffer_pool_size query_cache_size read_buffer_size`
  - pl
  - character collation Unicode
  - dataType: numberic datetime string spatial json default
  - fn op
  - [sql statement](#sql): DDL DML trans replication prepare compound admin utility
  - dataDir
  - storageEngine存储引擎: innoDB TokuDB(log表) alternative(myisam memory) 前台线程/后台线程
    - storageMechanism lockingLevels indexing capability&fn
  - replication partition 复制技术演进: 格式 数据安全vs复制效率
  - schema
    - performance: 等待事件 锁 sql语句 事务 多线程复制
    - sys: 慢sql 事务锁/MDL锁 innoDB缓冲池/热点数据 冗余/未使用索引 io/磁盘统计 全表扫描/文件排序/临时表
    - information: metaData db/table大小
    - mysql: 权限priv ACL 统计信息 复制信息 日志记录
  - cnnector&api
  - abc: FrequentProblem errors indexes

```sh
mysql -h127.0.0.1 -uroot -proot -P2207 -D test # host user passwd port Database; -q -A 获取表信息做代码补全
mysqladmin -uroot -p password "123456" # 修改密码

mysql -uroot -proot < /path/to/xxx.sql # 导入数据
mysql> source /path/to/xxx.sql
mysqldump /*login*/ database > xxx.sql # 导出
mysqldump /*login*/ database < xxx.sql # 导入
```

## sqlite

> <https://www.runoob.com/sqlite>

- litecli 数据分析 作为单测的内存数据库

```sh
sqlite3 xxx.db # mac 自带 sqlite3; conn, no new

.help # .archive .atuh .backup
.databases
.table
sqlite>.schema user # show create table user
.import FILE TABLE
.head on
.mod csv/insert
```

## sql server

## oracle

## mark

- [叶金荣 专注mysql技术](http://imysql.com) mysql实战45讲
- explain执行计划
  - ID：执行查询的序列号；
  - select_type：使用的查询类型
    - DEPENDENT SUBQUERY：子查询中内层的第一个SELECT，依赖于外部查询的结果集；
    - DEPENDENT UNION：子查询中的UNION，且为UNION 中从第二个SELECT 开始的后面所有SELECT，同样依赖于外部查询的结果集；
    - PRIMARY：子查询中的最外层查询，注意并不是主键查询；
    - SIMPLE：除子查询或者UNION 之外的其他查询；
    - SUBQUERY：子查询内层查询的第一个SELECT，结果不依赖于外部查询结果集；
    - UNCACHEABLE SUBQUERY：结果集无法缓存的子查询；
    - UNION：UNION 语句中第二个SELECT 开始的后面所有SELECT，第一个SELECT 为PRIMARY
    - UNION RESULT：UNION 中的合并结果；
  - table：这次查询访问的数据表；
  - type：对表所使用的访问方式：
    - all：全表扫描
    - const：读常量，且最多只会有一条记录匹配，由于是常量，所以实际上只需要读一次；
    - eq_ref：最多只会有一条匹配结果，一般是通过主键或者唯一键索引来访问；
    - fulltext：全文检索，针对full text索引列；
    - index：全索引扫描；
    - index_merge：查询中同时使用两个（或更多）索引，然后对索引结果进行merge 之后再读取表数据；
    - index_subquery：子查询中的返回结果字段组合是一个索引（或索引组合），但不是一个主键或者唯一索引；
    - rang：索引范围扫描；
    - ref：Join 语句中被驱动表索引引用查询；
    - ref_or_null：与ref 的唯一区别就是在使用索引引用查询之外再增加一个空值的查询；
    - system：系统表，表中只有一行数据；
    - unique_subquery：子查询中的返回结果字段组合是主键或者唯一约束；
  - possible_keys：可选的索引；如果没有使用索引，为null；
  - key：最终选择的索引；
  - key_len：被选择的索引长度；
  - ref：过滤的方式，比如const（常量），column（join），func（某个函数）；
  - rows：查询优化器通过收集到的统计信息估算出的查询条数；
  - Extra：查询中每一步实现的额外细节信息
    - Distinct：查找distinct 值，所以当mysql 找到了第一条匹配的结果后，将停止该值的查询而转为后面其他值的查询；
    - Full scan on NULL key：子查询中的一种优化方式，主要在遇到无法通过索引访问null值的使用使用；
    - Impossible WHERE noticed after reading const tables：MySQL Query Optimizer 通过收集到的统计信息判断出不可能存在结果；
    - No tables：Query 语句中使用FROM DUAL 或者不包含任何FROM 子句；
    - Not exists：在某些左连接中MySQL Query Optimizer 所通过改变原有Query 的组成而使用的优化方法，可以部分减少数据访问次数；
    - Select tables optimized away：当我们使用某些聚合函数来访问存在索引的某个字段的时候，MySQL Query Optimizer 会通过索引而直接一次定位到所需的数据行完成整个查询。当然，前提是在Query 中不能有GROUP BY 操作。如使用MIN()或者MAX（）的时候；
    - Using filesort：当我们的Query 中包含ORDER BY 操作，而且无法利用索引完成排序操作的时候，MySQL Query Optimizer 不得不选择相应的排序算法来实现。
    - Using index：所需要的数据只需要在Index 即可全部获得而不需要再到表中取数据；
    - Using index for group-by：数据访问和Using index 一样，所需数据只需要读取索引即可，而当Query 中使用了GROUP BY 或者DISTINCT 子句的时候，如果分组字段也在索引中，Extra 中的信息就会是Using index for group-by；
    - Using temporary：当MySQL 在某些操作中必须使用临时表的时候，在Extra 信息中就会出现Using temporary 。主要常见于GROUP BY 和ORDER BY 等操作中。
    - Using where：如果我们不是读取表的所有数据，或者不是仅仅通过索引就可以获取所有需要的数据，则会出现Using where 信息；
    - Using where with pushed condition：这是一个仅仅在NDBCluster 存储引擎中才会出现的信息，而且还需要通过打开Condition Pushdown 优化功能才可能会被使用。控制参数为engine_condition_pushdown 。
