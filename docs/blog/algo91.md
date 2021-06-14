# lucifer 力扣加加 91algo

## 91algo

> <https://github.com/leetcode-pp/91alg-2>

- 先导: 概述(ds crud -> algo 经典算法) 物理结构/逻辑结构 bigO
- 基础
  - 枚举: 暴力 状态/不重不漏/效率
  - 数组/栈/队列: 随机访问 LIFO
  - 链表
  - 树
  - 哈希表
  - 双指针
  - 图
- 进阶: UF Trie KMP/RK skiplist 剪枝 高频面试题 堆
- 专题: 二分 滑动窗口 位运算 搜索 背包问题 动态规划 分治 贪心

## 力扣加加

> [Introduction - 力扣加加 - 努力做西湖区最好的算法题解](https://leetcode-solution-leetcode-pp.gitbook.io/leetcode-solution/)

- 算法专题: ds 链表 树 堆 二叉树 dp 霍夫曼编码/游程编码 布隆过滤器 string 前缀树 贪婪 dfs 回溯 滑动窗口 位运算 设计题 小岛问题 最大公约数 并查集 平衡二叉树 蓄水池抽样 单调栈
  - ds
    - tree: k-dTree(游戏中碰撞检测 4维/8维) DOM/HTML/AST/XML `特殊的图`
    - BT: traversal
- 91algo
- 精选题解: 字典序列删除 前缀和 字节跳动算法题 我是你妈妈呀-核心母题 二叉树序列化 最长上升子序列 最长公共子序列 最大子序列和
- 高频考题: 简单-中等-困难

## issue 打卡模板

```md
## 思路
从末位往前开始计算(+1), 不产生进位就结束, 产生进度就继续; 计算完还有进位, 那就是 `99 -> 100` 这种情况, 单独处理即可

## 代码
> https://github.com/daydaygo/coder/tree/master/src/leetcode

- 支持 go+php, 并 pr 到 https://github.com/azl397985856/leetcode
- 支持单测

## 复杂度
- 时间: O(n)
- 空间: O(1)
```

## 参与贡献多语言

- pr 参考: <https://github.com/azl397985856/leetcode/pull/452>
- contributing 说明: <https://github.com/azl397985856/leetcode/blob/master/CONTRIBUTING.md>

## mark

- 基础数据结构
  - 数组array
    - 分桶 & 计数
  - 链表link: 环-快慢指针 指针操作+不要有环
    - 基本概念: 虚拟节点(pre dummy) 尾结点 静态链表
    - 分类: 循环/非循环 单链表/双链表
    - 基本操作: 插入 删除 遍历
    - 常见问题: 反转 合并 相交/环形
  - 栈stack: 最小栈-备胎
  - 队列queue
    - 数组队列(ring buffer, 环形数组)
  - 树tree
    - 基本概念: root/child/leaf/node 高度/深度/层 二叉树/N叉树
    - [二叉树的遍历](https://leetcode-solution-leetcode-pp.gitbook.io/leetcode-solution/thinkings/binary-tree-traversal)
    - 二叉树: left/right
      - 分类: 完全/满/搜索/平衡/红黑 二叉堆=优先级队列
      - 表示/存储: 数组/链表
      - 遍历: 前序/中序/后序/层序 DFS/BFS
      - 构建 BST heap(maxHeap minHeap) 递归
  - 哈希表(hash table, 散列表)
    - 实现: 数组+链表 -> 数组+红黑树; jdk1.8 HashMap
    - 哈希函数: key->hashcode; jdk1.8 hashcode 高16bit和低16bit进行异或
    - 哈希冲突: 抽屉原理; 开放地址法 链式地址发; java HashSet 底层使用 HashMap 实现
    - 常见题: 统计次数频率 O(1)复杂度 图数据结构(如 并查集) 存储之前的状态减少开销(dp)
  - 双指针: 二分法左右双指针 滑动窗口-快慢指针/固定间距指针
  - 图
    - basic: 有向无向 有权无权 入度出度 path/ring 连通/强连通 生成树
    - create: 邻接矩阵 邻接表
    - traverse: BFS DFS
  - 查找
  - 排序
    - 交换排序 冒泡排序-鸡尾酒排序
    - 快速排序-单边循环法-非递归实现
    - 堆排序
    - 计数排序 桶排序 基数排序
  - 基础算法
    - 最大公约数: 辗转相除法=欧几里得算法 更相减损术(九章算术)-位运算(奇偶)
    - bitmap: 用户-标签
    - LRU
    - 启发式搜索-A星寻路
    - 线段切割法-抢红包
- 进阶
  - 并查集
  - Trie
  - KMP & RK
  - 跳表
  - 剪枝
  - 高频面试题
  - 堆
- 专题
  - 二分法
  - 滑动窗口
  - 位运算
  - 搜索
  - 背包问题
  - 动态规划
  - 分治
  - 贪心
