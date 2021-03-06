# appendix

- 格式说明: `缩写 英文 中文;`
- codeSource 源码: file/name, name -> fn/type

---

- pl, programming language 编程语言
  - fmt format 格式化; k key v value kv; h header; l len length
  - b bit 位; B byte 字节; LSB 最低有效位; MSB 最高有效位; 2's complement 补码(原码/反码)
  - dsl domainSpecificLanguage; repl readEvalPrintLoop; cas compareAndSwap 原子操作
  - ext extension 扩展名; dat data
  - Concurrent 并发; parallel 并行; csp communicatingSequentialProcess 并发模型; mpg Machine(内核线程) Processor(上下文 GOMAXPROCS()) Go; shm sharedMemory; sem semaphore 信号量 Amdahl's Law 改进部分->性能改进
- fn func function 函数/功能
  - fp, functional programming 函数式编程
- oo object-oriented 面向对象; oop
  - Polymorphism 多态
  - IoC InversionOfControl 控制反转 TheHollywoodPrinciple; DI DependencyInjection 依赖注入
  - dp designPattern 设计模式; AOP AspectOrientedProgramming 面向切面编程
- fe, frontEnd 前端; be, backEnd 后端; cs, clientSide 客户端; test 测试
  - html hyperTextMarkupLanguage; dom documentObjectModel; attr attribute 属性
  - se searchEngine 搜索引擎; sitemap; spider/crawler 爬虫
  - webRTC webRealTimeCommunication 网页及时通信
  - mvc modelViewController; mvvm viewModel; spa singlePageApp;
- ds datastruct 数据结构
  - op operation 操作; CRUD Retrieve 增删改查; RandomAccess 随机访问; SequentialAccess 顺序访问; iteration 迭代; recursion 递归; traversal 遍历
    - FIFO FistInFirstOut; LIFO; LFU LeastFrequentlyUsed; LRU LeastRecentlyUsed
  - arr array 数组; vec vector 向量/顺序线性表; prefixSum 前缀和
  - list 列表; linklist 链表; skiplist 跳表; ring 环
  - stack 栈; ms monotoneStack 单调栈
  - queue 队列; pq priorityQueue 优先级队列;mq monotoneQueue 单调队列
  - tree 树; bt binaryTree 二叉树; bst binarySearchTree 二叉搜索树; bss binaryTreeSerialization 二叉树序列化; cbt completeBinaryTree 完全二叉树; trie trieTree 字典树/前缀树; binaryIndexTree 二叉索引树; segmentTree 线段树; HuffmanTree 哈夫曼树; b+(数据库索引)
    - bbt binaryBalanceTree 二叉平衡树; AVL; rbt readBlackTree 红黑树
    - pre/in/afterOrder 前中后序 HierarchicalTraverse层次遍历
    - lca leastCommonAncestors 最近公共祖先
  - heap 堆; bh binaryHeap 二叉堆
  - hash 哈希; dict 字典; symbol table 符号表; assocArray 关联数组; map 映射; kv, key-value pair 键值对
  - string 字符串; kmp
  - graph 图; graphTheory 图论
    - algo: Dijkstra最短 Floyd 最小生成树(Kruskal&Prim) A星寻路 二分图(染色法) 拓扑排序
- algo algorithm 算法
  - dp DynamicProgramming 动态规划; memo 备忘录 -> status 状态压缩; fsm finiteStateMachine 有限状态机; trim 剪枝
    - Knapsack 背包; subKnapsack 子集背包; completeknapsack 完全背包; 01knapsack
    - ss subsequence 子序列; lcs LongestCommonSubsequence 最长公共子序列
    - greed 贪心; scheduleProblem 调度问题; jumpGame 跳跃游戏; gameProblem 博弈问题
  - dc DivideConquer 分治; recursion 递归
  - dct, doubleChainedTree 双链树
  - find 查找; bs binarySearch 二分查找; uf unionFind 并查集
  - sort 排序; quickSort 快速排序; selectSort 选择排序; qss quickSelectSort 快速选择排序
  - search 搜索
    - dfs deepFirstSearch 深度优先; backtracking 回溯法
    - BFS BreadthFirstSearch 广度优先
  - point 指针; 2p 2point 双指针; sw SlidingWindow 滑动窗口
- math
  - numberTheory 数论; Geometry 几何
  - add/incr/increase 加; minus/decr/decrease/subtract 减; Multiply 乘; divide 除; power 幂; factorial 阶乘; Interval 区间; median 中值/中位数; prime 素数/质数; rand/random 随机; diff 差分; sudoku 数独; oddEven 奇偶
    - binary 二进制; decimal 十进制; hex 十六进制; octal 8进制
    - set 集合; intersection union complement 交并补; subset 子集
    - fib Fibonacci 斐波那契数列 $f(n)=f(n-1)+f(n-2)$
  - prob probability 概率; permutation 排列; combination 组合
- gc, garbage collection 垃圾回收
  - rc, reference count 引用计数
- store 存储
  - persistence 持久化: snapshot append-only slaveDB backup
  - 数据量
  - cache 缓存; 缓存雪崩 缓存击穿(singleflight) 缓存穿透; bitmap 位图; bf bloomFilter 布隆过滤器
  - db database 数据库; mvvc multiVersionConcurrentControl 多版本并发管理; staleData 脏数据; warm 数据预热; pk primaryKey 主键; 2pc prepareCommit 2阶段提交
- da, data analysis 数据分析
- ai, artificial intelligence 人工智能
  - nn neuralNetwork 神经网络; cnn 卷积; r-cnn
  - dl deepLearning 深度学习
  - cv computerVsion 计算机视觉: imageClassification objectLocalization semanticSegmentation instanceSegmentation
  - av, audioVideo 音视频
  - nlp natureLanguageProcess 自然语言处理
  - svm, support vector machines 支持向量机
  - recommendSystem 推荐系统
  - train test validate label
- lock 锁
  - DLM DistributedLockManager
- service 服务
  - io-bound cpu-bound io/cpu密集型; sync 同步; async 异步; coroutine 协程; req request 请求; resp response 响应
  - server 服务端
    - ratelimit tokenBucket; QoS qualityOfService 服务质量; qps queryPerSecond
    - RPC remoteProcedureCall 远程过程调用
  - client 客户端
    - circuitBreaker 熔断; downgrade 降级
  - soa serviceOrientedArch
  - ms, microService 微服务
    - serviceCenter 服务中心; serviceRegisterDiscovery 服务注册发现
    - configCenter 配置中心
    - gateway 网关; LB loadBalance 负载均衡; auth 认证/授权
    - tracing 调用链追踪; log err metric monitor
- project 项目
  - code/src 代码; comment 注释; doc/document 文档; apidoc api文档
  - coding 编码; refactor 重构; reuse 复用
  - 测试 test; 单元测试 单测 unittest; 压力测试 压测 prof/pprof/profile/profiler bench/benchmark
- foundation 基石
  - Formation 体系结构 structure 构造 interpretation 解释 Composition 组成原理 theory 理论
  - os operatingSystem 操作系统
  - ds distributedSystem 分布式系统
  - compile/languageEngineering 编译原理
    - asm assembly 汇编; jit justInTime 即时编译
  - net, network 计算机网络
  - Graph 图形学
  - code/codec 编码; base64 utf8
  - security 安全; crypto/cryptology 密码学; informationSecurity 信息安全; hash 哈希/散列; AsymmetricEncryption 非对称加密
  - db database 数据库
  - embed 嵌入式
  - ai artificialIntelligence 人工智能; ml machineLearning 机器学习
- manage 管理
  - mba MasterOfBusinessAdmin; mem MasterOfEngineManage; mpa, MasterOfPublicAdmin; mpacc, MasterOfProfessionalAccount
  - ChiefExecutiveOfficer 执行 finance.财务 information.信息 operation.运营 tech.技术 human.人才 market.营销 product.产品 [CXO_百度百科](https://baike.baidu.com/item/CXO)
  - business 商业; mvv missionVisionValue 使命 愿景 价值观
  - tl, team leader; tech 技术; mvp miniumViableProduct 最简可行产品
  - clv customerLifetimeValue 客户生命周期价值; acquisition获取 activation激活 retention留存 referral推荐 revenue收益; theMagicNumber神奇数字
- know knowledge 知识
  - linguistic 语言学; semantic 语义; syntax 语法; English 英语
- tool 工具
  - md, markdown; mpe, markdown preview enhance
- other
  - fineTune 微调
  - TuringAward 图灵奖
