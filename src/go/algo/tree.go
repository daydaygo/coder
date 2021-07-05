package algo

import "fmt"

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	preorder(root, &ans)
	return ans
}

func preorder(root *TreeNode, a *[]int) {
	if root != nil {
		*a = append(*a, root.Val)
		preorder(root.Left, a)
		preorder(root.Right, a)
	}
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	inorder(root, &ans)
	return ans
}

func inorder(root *TreeNode, a *[]int) {
	if root != nil {
		inorder(root.Left, a)
		*a = append(*a, root.Val)
		inorder(root.Right, a)
	}
}

func postorderTraversal(root *TreeNode) []int {
	var ans []int
	postorder(root, &ans)
	return ans
}

func postorder(root *TreeNode, a *[]int) {
	if root != nil {
		postorder(root.Left, a)
		postorder(root.Right, a)
		*a = append(*a, root.Val)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root != nil {
		return [][]int{}
	}
	ans, q := [][]int{}, []*TreeNode{root}
	for len(q) > 0 {
		a, n := []int{}, len(q)  // 也可以用2个变量记录当前层和下一层
		for i := 0; i < n; i++ { // 遍历当前层
			t := q[i]
			a = append(a, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
			q, ans = q[n:], append(ans, a)
		}
	}
	return ans
}

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

// 回溯法 RM
var rm_res []int

func backtrack(path int, path_choice []int) {
	if path == 0 { // 如果满足结束条件
		rm_res = append(rm_res, path)
		return // return true 可以在找到一个解时结束
	}
	for _, i := range path_choice {
		// 做选择=前序遍历
		backtrack(i, path_choice)
		// 撤销选择=后序遍历
	}
}

// BFS 广度优先搜索
func BFS(start, end *NTree) int {
	var q []*NTree       // 核心数据结构: queue
	var visited []*NTree // 避免走回头路: set; 不会走回头路就不需要
	step := 0

	q = append(q, start)
	visited = append(visited, start)
	for len(q) != 0 {
		for i := 0; i < len(q); i++ {
			cur := q[i]     // 注意是 dequeue 操作
			if cur == end { // 重点: 到达终点
				return step
			}
			for _, tree := range cur.Node {
				if tree != nil { // 注意是判断 tree 是否在 visited 中
					q = append(q, tree)
					visited = append(visited, tree)
				}
			}
		}
		step++ // 重点: 更新步数在这里
	}
	return step
}

// BinarySearch 二分查找
func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { // 统一使用闭区间, 统一思维逻辑
		mid := left + (right-left)>>1 // 防止溢出-trick
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 技巧: 不使用 else, 使用 else if 把所有条件列清楚
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

func LeftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			right = mid - 1 // 别返回, 锁定边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target { // 检查边界
		return -1
	}
	return left
}

func RightBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			left = mid + 1 // 别返回, 锁定边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target { // 检查边界
		return -1
	}
	return right
}

// SlidingWindow 滑动窗口
func SlidingWindow(s, t string) {
	var need, window map[byte]int
	for _, i := range t {
		need[byte(i)]++
	}

	left, right := 0, 0
	// valid := 0
	for right < len(s) {
		c := s[right]  // c 是将移入窗口的字符
		fmt.Println(c) // 移入
		right++
		updateWindow(window)
		fmt.Println(left, right) // debug
		for checkWindowNeedShrink(window) {
			d := s[left]
			fmt.Println(d) // 移出
			left++
			updateWindow(window)
		}
	}
}

// 判断窗口左侧窗口是否要收缩
func checkWindowNeedShrink(window map[byte]int) bool {
	return true
}

// 更新窗口内数据
func updateWindow(window map[byte]int) {

}
