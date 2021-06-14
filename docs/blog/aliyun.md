# aliyun

- [产品列表](aliyun_product.md): 本文档一律使用产品简写
- [产品 - bigdata](aliyun_big_data.md)
- aliyun 网站矩阵 -> SE: `site:*.aliyun.com`
- [阿里云大学](https://edu.aliyun.com)
- [aliyun mvp](https://mvp.aliyun.com)
- 阿里云开发者社区
  - [阿里云总监课](https://yq.aliyun.com/promotion/689)
  - [藏经阁](https://developer.aliyun.com/topic/ebook)

## mark

- ecs: 送的免费云盾有一层白名单
- cs-k8s: 专有版: 3master+worker; 托管版: worker; serverless: 控制台+命令控制
- fc: 支持 php
- [企业级数据库服务选型](https://promotion.aliyun.com/ntms/act/dbservice.html)
- rds
  - 实例规格 <https://help.aliyun.com/document_detail/26312.html>
  - 慢查: sql洞察(搜索) sql日志(汇总/详情)
  - 读写分离: 事务 `/*force_master*/`
  - dblink
- drds
  - 多个 rds 的前置服务器
  - 查询: 分区键>索引>常规查询
  - 并发型vs分析型
  - prepare协议支持: 主库支持, 从库不支持
- redis: 规格性能: <https://help.aliyun.com/document_detail/26350.html>
- mongo: 版本与存储引擎: <https://help.aliyun.com/document_detail/61906.html>
- ads: mysql -> dts -> ads -> quickBI/DataV
- tsdb: ts, influx, 时空(时序+时空)
- elasticsearch: 预置插件
- ossfs: <https://help.aliyun.com/document_detail/32196.html>
- nas: mount NFS <https://help.aliyun.com/document_detail/90529.html>
- vpc&region: <http://static-aliyun-doc.oss-cn-hangzhou.aliyuncs.com/assets/img/80559/155324038634448_zh-CN.png>
- vpc - 路由器 - 路由表 - 交换机 - ecs/rds
- sls
  - logtail 采集原理: <https://help.aliyun.com/document_detail/89928.html>
  - logtail 采集限制: 每次512kb->导致json截断->json解析失败 修改时间后需要重启
  - search syntax: <https://help.aliyun.com/document_detail/29060.html>
  - [search example](aliyun_sls.md)
  - analyze syntax: <https://help.aliyun.com/document_detail/53608.html>
  - grafana: <https://help.aliyun.com/document_detail/60952.html>
  - 日志接入: 文件/容器stdout -> json -> logtail json 解析
  - 配置索引: 索引配置后验证
  - 查询 + 保存快查
  - 面板: 根据 web 服务 SLA 制定数据指标
  - 监控: 监控验证
- pts
  - 阿里云性能测试 PTS: <https://help.aliyun.com/product/29260.html>
  - apache jmeter: <http://jmeter.apache.org/>
- acm
- ![配置结构示意图](https://aliware-images.oss-cn-hangzhou.aliyuncs.com/acms/dg_acm_mq_scheme.png)
- tracing: zipkin
- cloud shell: <https://help.aliyun.com/product/89853.html>
- toolkit: <https://help.aliyun.com/product/29966.html>
- api: <https://api.aliyun.com>

## aliyun sls example

```sh
# ngix access log
# pv/uv
* | select approx_distinct(remote_addr) as uv, count(1) as pv, date_format(date_trunc('hour', __time__), '%m-%d %H:%i') as time group by date_format(date_trunc('hour', __time__), '%m-%d %H:%i') order by time limit 1000
# ip distribution
* | select count(1) as c, ip_to_province(remote_addr) as address group by ip_to_province(remote_addr) limit 100
# top page
* | select split_part(request_uri,'?',1) as path, count(1) as pv group by split_part(request_uri,'?',1) order by pv desc limit 10
# top refer
* | select count(1) as pv, http_referer group by http_referer order by pv desc limit 10
# http method
 * | select count(1) as pv, request_method group by request_method
# http status
* | select count(1) as pv, status group by status
# UserAgent
* | select count(1) as pv, case when http_user_agent like '%Chrome%' then 'Chrome'when http_user_agent like '%Firefox%' then 'Firefox'when http_user_agent like '%Safari%' then 'Safari'else 'unKnown' end as http_user_agent group by  http_user_agent order by pv desc limit 10
# latency
* | select from_unixtime(__time__ -__time__% 300) as time, avg(request_time) as avg_latency , max(request_time) as max_latency group by __time__ -__time__% 300
* | select from_unixtime(__time__ - __time__% 60) , max_by(request_uri,request_time) group by __time__ - __time__%60
* | select numeric_histogram(10,request_time)
* | select max(request_time,10)
request_uri:"/url2" | select count(1) as pv, approx_distinct(remote_addr) as uv, histogram(method) as method_pv, histogram(status) as status_pv, histogram(user_agent) as user_agent_pv, avg(request_time) as avg_latency, max(request_time) as max_latency

* | select time_series(__time__, '1m', '%H:%i' ,'0') as time, count(1) as PV group by time order by time limit 100
host: mms.dayday.tech
# status
select count(1) as pv ,status group by status
# pv/uv
select approx_distinct(remote_addr) as uv ,count(1) as pv , date_format(date_trunc('hour', __time__), '%m-%d %H:%i')  as time group by date_format(date_trunc('hour', __time__), '%m-%d %H:%i')  order by time limit 1000
# 访问前十地址
* | select count(1) as pv, split_part(request_uri,'?',1) as path  group by split_part(request_uri,'?',1) order by pv desc limit 10
# 访问前十来源
* | select count(1) as pv , http_referer  group by http_referer order by pv desc limit 10
# 访问时间
select avg(request_time) as response_time, avg(upstream_response_time) as upstream_response_time  , date_format( from_unixtime(__time__ -__time__%3600),'%m-%d %H:%i' )   as time group by __time__ - __time__% 3600 limit 10000
# 访问时间前十地址
* | select request_uri  as top_latency_request_uri ,request_time order by request_time  desc limit 10
```
