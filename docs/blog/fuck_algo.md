# 算法有套路

- <https://github.com/labuladong/fucking-algorithm>

## 目录

- 必读
  - 方法论=开篇词: **一定要明白自己的目标是什么** 工具人 可量化目标
  - 框架思维: ds 存储=链式+顺序 op=crud 遍历=迭代+递归; 大部分算法技巧本质上都是树的遍历问题
  - 框架/套路: dp-子集/背包 backtracking bfs bs sw 股票买卖 打家劫舍 interval nSum bt git
- ds
  - 思路: 书(算法ed4✅ 算法导论❎) 这是啥/有啥用/如何看书/如何刷题/坚持(激起欲望)
  - 递归思维: 链表 bt/bst/bss 最近公共祖先 cbt
  - 设计: uf lru lfu median 朋友圈时间线 ms
  - array: bs 2p sw 2sum
- dp
  - 技巧: 暴力递归->带memo递归->dp 状态->memo状态压缩->状态间关系
  - ss: 编辑距离 信封嵌套 subArr lcs
  - knapsack: sk ck 01k
  - greed: 调度问题 跳跃游戏
  - other: regex 高楼扔🥚 戳🎈 博弈问题 四键键盘 股票买卖 打家劫舍 fsm-kmp 回文
- 必知必会
  - dfs=回溯: 子集排列组合 数独 括号生成
  - bfs: 智力题
  - math: bit 阶乘/素数
  - other: prefixSum diffArr qss
- 高频面试
- 技术文章: linux(os Process/thread/fd shell) session/cookie 加密算法 git
  - [Git/SQL/正则表达式的在线练习平台 · labuladong的算法小抄](https://labuladong.gitee.io/algo/%E6%8A%80%E6%9C%AF/%E5%9C%A8%E7%BA%BF%E7%BB%83%E4%B9%A0%E5%B9%B3%E5%8F%B0.html)

## mark

- 只有 2 种存储方式: 数组-顺序存储 链表-链式存储
  - 栈 队列
  - 图: 邻接表-稀疏 邻接矩阵-矩阵运算
  - 散列表: 散列函数-数组 冲突-拉链法/线性探查法
  - 树: 堆-数组 链表-二叉搜索树/AVL/红黑/区间/B
    - 二叉搜索树(binary search tree, BST) -> AVL/红黑/B+/线段树
- 数据结构
  - 操作: CRUD
    - 遍历: 数组-线性迭代 链表-线性/递归 环-visited数组
    - 递归(DP 回溯法) -> 树(二叉树->N叉树 二叉树的前序/中序/后续遍历)
- DP: 求最值 找「状态」+做「选择」
  - 思路(正向/反向): 明确 base case -> 明确「状态」 -> 明确「选择」 -> 定义 dp 数组/函数 函数
  - Fib-重叠子问题优化: 备忘录memo-自顶向下 「dp table」-自下向上 -> 优化递归树 -> 压缩 memo/dp
  - 凑零钱: 优化状态转移方程
  - 股票买卖: `dp[i][k][0/1]` 第i天 至多可以交易k次 是否持有股票 -> 状态转移方程 -> 压缩 memo/dp
  - 打家劫舍
  - 高楼扔鸡蛋: 二分查找+重新定义状态转移
  - 背包问题: 01背包 子集背包 完全背包
  - 最长递增子序列SIP
  - 贪心类
- 回溯=DFS 遍历决策树
  - 全排列: 路径-选择列表-结束条件
  - 八皇后问题-高斯
  - 只要一个解
- BFS 图起点到终点的最短路径
  - 周围节点 -> 队列; 路径一定是最短
  - 双向BFS-trick: 知道终点
- 二分查找 思路很简单细节是魔鬼
  - 闭区间+边界检查
- 双指针
  - 滑动窗口
  - nSum: 排序+跳过重复值
- 区间合并: 起点升序+终点升序 -> 枚举相邻区间的所有可能
- 思维
  - 计算机解决问题 -> 穷举 -> 剪枝
  - 分析问题: 递归  自顶向下/自下向上 抽象到具体
