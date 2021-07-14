# TS| 技术分享: 刷题资料与技巧

## 如何刷题

- ds数据结构
  - 存储/物理=链式+顺序
  - 逻辑=线性(arr/link/stack)+非线性(tree/graph)
  - op=crud
  - 遍历=迭代+递归
    - 递归:
      - 不要跳进递归, 而是利用明确的定义来实现算法逻辑
      - baseCase(各种边界, 也可以通过边界来找找子问题) + 递归处理子问题
- algo算法
  - bigO 时间/空间
  - 大部分算法技巧本质上都是树的遍历问题

```go
// TreeNode 二叉树
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func TreeTraverse(root *TreeNode) {
	// 前序遍历
	TreeTraverse(root.Left)
	// 中序遍历
	TreeTraverse(root.Right)
	// 后序遍历
}

// NTree N 叉树
type NTree struct {
	Val  int
	Node []*NTree
}

func NTreeTraverse(root *NTree) {
	// root.Val
	for _, tree := range root.Node {
		NTreeTraverse(tree)
	}
}
```

- algo算法
  - interval区间: 左闭(0/1开始)右开? => 二分法的11种变形(>还是>= mid向上取还是向下取 是否有相等值)
  - boundary边界=Layer层=std规范=protocol协议=$O(mn) > O(m+n)$

> All problems in computer science can be solved by another level of indirection. --David Wheeler

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
  - 思路: 枚举/暴力(规范/协议)->子问题/状态压缩/剪枝->最优解

## 工具leetcode

- leetcode web
  - leetbook学习: 面试+数据结构 -> 我的书架; 学习计划 => 可以直接链接到 leetcode 题库
  - problemset题库
    - 分类: 算法 剑指offer 程序员面试经典 hot100 top200
    - **标签tag**: 学了一个知识后, 按标签刷, 快速巩固
    - [设置-进度管理](https://leetcode-cn.com/session/)
  - circle社区: 题解
  - code题目页面: 模拟面试 快捷键 使用上一次的编辑 还原 playground调试 **测试用例** 可视化 收藏中心
- [力扣刷题插件](https://lucifer.ren/blog/2020/08/16/leetcode-cheat/)
  - 题解模板/查看题解
  - 代码: **一键复制所有测试用例** 禅定模式
  - 代码模板: 前缀和(一维 二维) 二分法(baisc+4变种) BFS(是否带层信息) 堆(最小堆) 滑动窗口(固定窗口/可变窗口) 回溯(标准 笛卡尔积优化) 前缀树 并查集(是否带权) 线段树(区间和 计数)
  - 数据结构可视化
  - 复杂度速查
  - 学习路线: 动态规划 树 链表 二分

## 代码code

> https://github.com/daydaygo/coder/blob/main/src/go

- go语言基础

> [go语言圣经](http://coder.dayday.tech/plbook/gopl.html)

```go
// basic/demo_test.go#L32
// 使用 Test_ 方便测试
func Test_type(t *testing.T) {
	fmt.Println(-10%3, -10%-3) // -1 -1, same sign with dividend

	// s string
	s := "hello"
	for k, v := range s { // k int; v int32
		fmt.Println(k, v)
	}
	// c := 'c' // int32
	// var c2 byte = 'c' // uint8
	fmt.Println('0') // '0'-'9'
	// s := `[1, null, 2, 3]` // json -> [1, 0, 2, 3]

	// arr array: fixed-length; Using pointer
	arr1 := [...]int{1, 2, 3}
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Printf("%x %x %t %T\n", arr1, arr2, arr1 == arr2, arr1) // %x array/slice %t bool %T type

	// a slice: pointer+len()+cap(); not comparable; nil
	// make([]int, 0, 3) // make([]int, 3) 会导致前 3 个已赋值, append() 从第 4 个开始
	// https://ueokande.github.io/go-slice-tricks/
	a := arr1[:]
	a[2] = 4
	// var a []int // a=nil; a = []int{} // a!=nil
	// make([]T, len, cap) = make([]T, cap)[:len]
	// append() cap扩容
	a = append(a[:10], a[10+1:]...) // 删除第i个值
	var x int
	a = append(a, x)                 // push
	x, a = a[len(a)-1], a[:len(a)-1] // stack
	x, a = a[0], a[1:]               // queue

	// m map[K]V: unordered/random; K == comparable; cannot &V
	m := map[string]int{"a": 1, "b": 2} // m := make(map[string]int)
	delete(m, "a")
	if _, ok := m["b"]; ok {
	}
	// graph := make(map[string]map[string]bool) // graph[from][to] edge
	seen := make(map[string]bool)
	if !seen["i"] {
		seen["i"] = true
		// do
	}

	// p struct
	// p := &Point{1, 2} // %#v struct; p=new(Point); %p
	// w = Wheel{Circle{Point{8, 8}, 5}, 20} // embed %#v
	// w.X = 8 // equivalent to w.circle.point.X = 8
	// Year int `json:"released"` // field tag
}
```

- 数据结构

```go
// algo/array.go#L16
func MonotoneStack(a []int) []int {
	// 单调递减: 用来求解下一个更大的元素
	// 变形: stack 存索引, 可以得到距离
	var stack []int
	ans := make([]int, len(a), len(a)) // 根据题目要求使用 slice/map
	for i := len(a) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[0] <= a[i] { // 注意边界
			stack = stack[1:] // pop
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[0]
		}
		stack = append([]int{a[i]}, stack...) // push
	}
	return ans
}
```

- 题解

```go
// algo/bit.go#L189
// 函数名和 leetcode 一致
func countBits(num int) []int {
	ans := make([]int, num+1)
	for i := 1; i <= num; i++ {
		ans[i] = ans[i&(i-1)] + 1 // 空间压缩
	}
	return ans
}
```

## 资料

- [ds&algo 资料汇总, 持续更新](https://coder.dayday.tech/a/ds_algo.html)
- [力扣加加 - 努力做西湖区最好的算法题解](https://coder.dayday.tech/blog/algo91.html) 91天算法提升计划
- [算法有套路 fucking algo](https://coder.dayday.tech/blog/fuck_algo.html) 刷题模板
- [leetcode cookbook](https://coder.dayday.tech/blog/leetcode_cookbook.html) 用go语言刷leetcode
