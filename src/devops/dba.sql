-- schema
use information_schema;

show tables;
show create table `db`; desc db;
show index from db;
show status like '%buffer%';
show variables like '%buffer%';
show session status like 'select';
show slave status;
SHOW ENGINE innodb STATUS; -- buffer pool hit rate -> LRU

-- 表大小
select table_name, engine, concat(round(table_rows / 1000000, 2), 'M') as rows,
  concat(round(data_length / 1024 / 1024 / 1024, 2), 'G') as data,
  concat(round(index_length / 1024 / 1024 / 1024, 2), 'G') as idx,
  concat(round((data_length + index_length) / 1024 / 1024 / 1024, 2), 'G') as total_size,
  round(index_length / data_length, 2) as idx
from information_schema.tables
where table_schema in ('mt_billing') -- 填写自己的库
group by table_name, table_schema, data_length + index_length
order by data_length + index_length desc;

-- coon
SHOW FULL PROCESSLIST; -- mysql 上当前连接的进程; state: sleep.wait_timeout
SHOW FULL PROCESSLIST; -- 表锁 state: wait for table metadata lock / wait for table flush
SELECT * FROM information_schema.`PROCESSLIST`; -- information_schema 非常有用
kill 1; -- 通过上面查询出来的 id 杀掉进程, 推荐 pt-kill, 干掉执行超过50s的sql, 避免数据库被hang住
select concat('KILL ',id,';') from information_schema.processlist where user='root';

-- mysql抖一下 脏页75%
select VARIABLE_VALUE into @a from information_schema.global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_dirty';select VARIABLE_VALUE into @b from information_schema.global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_total';select @a/@b;

use performance_schema;
SELECT EVENT_NAME,MAX_TIMER_WAIT FROM performance_schema.file_summary_by_event_name; -- 健康检查

-- prof
set profiling=1; -- 当前session
show profiles ; -- 执行的query都会prof
show profile cpu, block io for query 6; -- 选择先查看的prof

-- orderBy: optimizer_trace
SET optimizer_trace='enabled=on'; -- 当前session
SELECT * FROM `information_schema`.`OPTIMIZER_TRACE`;
select VARIABLE_VALUE into @a from performance_schema.session_status where variable_name = 'Innodb_rows_read'; -- 执行语句, 使用变量保存执行语句前后的 Innodb_rows_read 值

use sys;
SELECT blocking_pid FROM sys.schema_table_lock_waits; -- 表锁
select * from sys.innodb_lock_waits; -- 行锁

use mysql;
show tables ; -- 权限表: user db tables_priv columns_priv host

select version(); -- @@version;
select last_insert_id();
select 0xAE, format(1123.45,1), inet_aton('220.181.38.148');
-- math: abs(),floor(),ceil()
select char_length("中国"), length("中国"); -- string: insert(),upper(),lower(),left(),right(),substring(),reverse()
select current_date,current_time,now(),unix_timestamp(), from_unixtime(unix_timestamp()); -- time: month(),hour()