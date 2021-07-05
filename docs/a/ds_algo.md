# ds&algo

- todo
  - 清理掉 go.md 中的 isdamir/gotype

- 方法
  - name命名: ans ij kv n-len/node t-tmp l-left r-right a-arr m-map q-queue s-stack str
  - dimension维度 scope范围 classify分类 field/domain领域 boundary边界 interval区间: 左闭(0/1开始)右开? 边界=Layer层=规范=协议=$O(mn) > O(m+n)$ 环/假溢出? 连续?
    - 协=2+参与者 议=行为约定与规范
  - 有序?: sort 二分
    - 二分法: 快/慢 左/右 正/逆 前/后; 二分->四分
  - 多画图
  - 递归:
    - 不要跳进递归, 而是利用明确的定义来实现算法逻辑
    - baseCase(各种边界, 也可以通过边界来找找子问题) + 递归处理子问题
  - 计数排序: 计数信息=排序结果
- 思路: 枚举/暴力(规范/协议)->子问题/状态压缩/剪枝->最优解
- 学习算法是非常有趣和令人激动的
- 但是算法, 就是解决问题时的那份「优雅」.

## ds 数据结构

- ds 存储=链式+顺序 逻辑=线性(arr/link/stack)+非线性(tree/graph) op=crud 遍历=迭代+递归
- array/vector
  - 前缀和(连续问题) 空间压缩 环形数组/取模
  - stack 栈 LIFO pop/push
    - 单调栈: 下一个大于/小于 边界-哨兵法; 广义栈
    - op: polishNotation/波兰表达式 前缀/中缀/后缀 括号匹配/栈匹配 栈实现队列
  - queue 队列 FIFO back/front enqueue/dequeue
    - deque 双端队列; pool
- linkList 链表 singleLinkedList 单链表
  - 画图: 迭代操作(pre+cur+next) + 2遍
  - doubleLinkedList双向链表 线性链表 staticList静态链表 symmetricMatrix对称矩阵 sparseMatrix稀疏矩阵
  - op: reverse反转 中值 奇偶 快排
- hash map hashTable 哈希表
  - hashFunction 散列函数
  - 解决碰撞/填充因子
  - conflict冲突: separateChain链地址法/拉链法 开放地址=linearProbing线性探查+线性补偿探测+随机探测
  - rehash 渐进式rehash
  - consistentHash 一致性hash; DHT 分布式hash
    - uint32环 node(真实节点)/replica(虚拟节点) 顺时针查找第一个
    - judge: balance monotonous spread load
  - MurmurHash: 规律输入依然可以给出很好的随机分布+计算速度
- string 字符串
  - KMP 有限状态机 模式匹配有限状态机 BM BM-KMP BF
  - approximate/fuzzy string match
  - suffixArr后缀数组
- tree
  - binaryTree二叉树
    - 满 完全 平衡 查找
    - op: `遍历`(前中后 层次) create(前序+中序 树状数组) depth(min max) lca
  - RB红黑->区间树->线段树 B B+ 2-3 2-3-4
  - unionFind 并查集
  - Huffman merkleTree哈希二叉树.区块链
- heap 堆
  - array实现: 极大堆/极小堆/极大极小堆 双端堆Deap d叉堆
  - tree实现: 左堆 扁堆 二项式堆 fibHeap pairingHeap配对堆
- find 查找
  - hash binarySortTree排序二叉树 AVL/B/B+/B*/AA/RBT/splay/DCT/R binaryHeap
  - Trie 动态路由 `/`分隔
  - skipList 跳跃表/跳表
- other
  - bloomFilter
  - 分桶

## algo 算法

- 大部分算法技巧本质上都是树的遍历问题
- bigO 时间/空间
- sort 排序
  - 类型: 交换(冒泡 快排) 插入(直插 shell希尔) 选择(简选 堆) 归并(二路 多路) 基数 计数 桶 线性 自省 间接
  - [十种排序算法](https://blog.csdn.net/coolwriter/article/details/78732728)
  - 外部: k路归并败者树 最佳归并树
- dc分治 recursion 递归
  - 二分查找
- dp动态规划
  - 最优二叉搜索树
- greedy贪心
- backtracking回溯
- search
  - 枚举 dfs bfs 启发式搜索
- 分支定界法
- string字符串匹配
- graph图论
  - 邻接表-稀疏 邻接矩阵-矩阵运算 权 向
  - [dag](https://github.com/hyperf/dag-incubator), Directed Acyclic Graph 有向无环图
  - op: `遍历`(dfs bfs) dijkstra最短路径 primKruskal最小生成树 fordFulkerson最大流 A*搜索 图匹配 网络流
- 大数据
  - 决策树: 信息熵/增益/增益比/基尼系数 ID3/C4.5/CART 剪枝 数据集分类
  - 回归: 逻辑(logistic sigmoid函数 最佳系数) 线性(最佳拟合 局部加权) 树(连续 离散 CART 剪枝)
  - NaiveBayesian朴素贝叶斯: 条件概率->分类
  - KNN=K近邻 聚类
  - SVM支持向量机: 最大间隔 SMO高效优化 复杂数据+核fn 径向基核fn
  - 推荐: 内容关联 协同(领域/模型 过滤) 内容/知识/组合
- other
  - LB, load balance: rand roundRobin weightRoundRobin p2c
  - limit: 滑动窗口(频率/并发上限) tokenBucket(令牌桶 瞬时流量)
  - timingWheel 时间轮: 延迟操作
  - 双指针: 快慢指针+左右指针
    - 滑动窗口: 连续问题

## todo

- 企业题库
  - 程序员面试宝典.pdf
  - 阿里云技术面试红宝书.pdf
- 数据结构与算法之美/算法面试通关40讲(极客时间)
- 8个常用刷题网站: <https://mp.weixin.qq.com/s/WTjzp7NnD2kD7nbhrvnBIQ>
- set: map实现?(java/js/py原生支持)

## tool

- edge collection - algo
- leetcode web
  - leetbook 学习: 面试+数据结构 -> 我的书架; 学习计划
  - problemset 题库
    - 分类: 算法 剑指offer 程序员面试经典 hot100 top200
    - **标签tag**: 学了一个知识后, 按标签刷, 快速巩固
      - [位运算 - 力扣（LeetCode）](https://leetcode-cn.com/tag/bit-manipulation/)
    - [设置-进度管理](https://leetcode-cn.com/session/)
  - circle 社区: 题解
  - code 题目页面: 模拟面试 快捷键 使用上一次的编辑 还原 playground调试 测试用例 可视化 收藏中心
- leetcode app
  - 摇一摇切换中英文描述/随机一题
  - 学习分析
- [力扣刷题插件](https://lucifer.ren/blog/2020/08/16/leetcode-cheat/)
  - 题解模板/查看题解
  - 代码: 一键复制所有测试用例 禅定模式
  - 代码模板: 前缀和(一维 二维) 二分法(baisc+4变种) BFS(是否带层信息) 堆(最小堆) 滑动窗口(固定窗口/可变窗口) 回溯(标准 笛卡尔积优化) 前缀树 并查集(是否带权) 线段树(区间和 计数)
  - 数据结构可视化
  - 复杂度速查
  - 学习路线: 动态规划 树 链表 二分
- 如何刷题
  - 题解: 自己读题 -> 15分钟还没思路 -> 看题解 -> 完全不会才看代码/记笔记 -> coding/debug
  - 按专题(tag)刷 -> 模板/笔记/多题同解
    - 构建知识体系/框架/套路
  - 广度优先而非深度(死磕某一个知识)
  - 掌握编程语言
  - 模拟面试: 时间观念+一次AC
  - 训练目标: AC数量 AC速度 AC通过率 竞赛排名前100 beats-CPU/mem-100%
    - bug-free: 基础数据结构+算法形成肌肉记忆 单测+调试
  - code: 命名name(简洁&一致性) 小心全局变量 提交后查看排名->分段->看100%的代码 直接使用系统api
  - debug: 批量测试 数据可视化
  - 锁定使用哪种算法: 记忆/题感/暴力解(暴力+剪枝 大力出奇迹) 关键字 限制条件+复杂度速查 分治思维
- 九章算法
  - 企业题库 领扣lintcode题库（Java c++ py）
  - 免费课：ai 企业题库/面试 算法
  - 发现：求职面试 学习笔记 上班摸鱼
  - 消息
  - 我：发布 收藏 学分 学习记录

## leetcode help

- [help](https://support.leetcode-cn.com/hc)
  - plus: 付费精选题目和内容; 企业题库; 企业模拟面试(免费可 随机); 题目热度; 极速判题; playground(调试)
    - 模拟面试: 1-3题, 30-90分钟
  - 竞赛: 报名 -> AC +1分, 每次错误 +5min 罚时
  - 个人主页: **做题进度** 收藏/笔记/积分 订单 账号
  - 技术问题
    - [各语言对应版本和环境](https://support.leetcode-cn.com/hc/kb/article/1194343): phpinfo() => 7.2
    - 全局变量和类内静态变量 -> 手动初始化
    - stdout 打印: debug + 增加耗时
    - [二叉树序列化](https://support.leetcode-cn.com/hc/kb/article/1194353)
    - timeout - leetcode 的锅; Time Limit Exceeded, TLE - 代码的锅
- leetbook
  - 数据结构
    - 数组和字符串 数组类算法
      - 基本概念+操作方式; 二维数组; 字符串 概念+特性; KMP算法; 双指针
    - 链表
    - 哈希表 查找表类算法
    - 队列和栈
    - 二叉树 二叉搜索树 前缀树 N叉树
    - 图解数据结构
  - 算法
    - 初级算法 中级算法 高级算法
    - 递归 二分查找
  - 面试
  - AI数学基础
  - 漫画算法 画解剑指offer

## mark

- [React hooks: not magic, just arrays](https://medium.com/@ryardley/react-hooks-not-magic-just-arrays-cd4f1857236e)
