# elastic

- stack技术栈: elastic kibana logstash
- why应用
  - 全文搜索: 电商商品关键词 企业内置知识库
  - 大数据: 计算 聚合统计 自助报表
  - 业务: 海量数据 大宽表
  - log metric apm security安全预警
- noun名词: shard分片 replicate副本 segment分段 document文档 field字段 term词项
  - cluster集群 raft分布式共识 掌控(扩容升级 数据安全 管理监控) 资源/容量规划 性能调优 日志分析&问题诊断
  - node节点: 角色 职责
  - index索引
    - create: 数据模型(json 平铺 对象 关联) 类型 属性(index/doc_values/null_value) 动态映射 别名 模板
    - update: reindex ingest特性
    - query
      - 方式: id与search url与dsl search与filter rep与resp
      - 复杂: bool(should mustNot must) mathFn
      - text: 分词 短语 最小匹配度
      - term: 精确 容错 模糊 前缀 后缀
      - span 跨度->近邻
      - suggest特性 启发式
      - sort 默认/自定义/二次重排序
      - page分页
      - 重度查询->异步
      - 查询模板
    - 管理: 数据分布与分段合并 可靠性 ILM自动化索引
    - aggregation聚合统计: 指标->快速 bucket分桶->日常分组 pipeline->二转聚合(固定窗口 滑动窗口)
    - template
  - sql: sql<->DSL result结果(json/csv/文本) jdbc math 局限(二次查询 二次聚合 分组 复杂DSL)
  - trans数据变换: rollup减少数据 transform转换数据
  - 分词: 分词器 默认/自定义/中文(IK other)
- algo算法: invertedindex倒排索引 DocValues列式数据 FST有限状态转换 BKDTree分块多维度空间树索引 roaringBitmap压缩位图
  - TF/IDF/BM25 词频/逆文本频率/概率模型
- tool
  - [mobz/elasticsearch-head: A web front end for an elastic search cluster](https://github.com/mobz/elasticsearch-head)

```sh
curl http://localhost:9200
./bin/elasticsearch-plugin list
./bin/elasticsearch-plugin install analysis-ik # 中文分词
```
