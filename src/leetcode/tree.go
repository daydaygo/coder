package leetcode

import (
	"encoding/json"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 构建: 前序+中序 / 后序+中序
func NewTree(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preOrder[0]}
	i := 0
	for inOrder[i] != root.Val {
		i++
	}
	root.Left = NewTree(preOrder[1:i+1], inOrder[:i])
	root.Right = NewTree(preOrder[i+1:], inOrder[i+1:])

	return root
}

// 构建: 前序+nil, 0 表示 nil
func newTree2(a *[]int) *TreeNode {
	if len(*a) == 0 {
		return nil
	}

	var v int
	v, *a = (*a)[0], (*a)[1:]
	if v == 0 {
		return nil
	}
	root := &TreeNode{Val: v}
	root.Left = newTree2(a)
	root.Right = newTree2(a)

	return root
}

// 构建: BFS + leetcode: https://support.leetcode-cn.com/hc/kb/article/1194353/
func NewTree3(s string) *TreeNode {
	var a []int
	json.Unmarshal([]byte(s), &a) // null -> 0
	if len(a) == 0 {
		return nil
	}

	var v int
	v, a = a[0], a[1:]
	root := &TreeNode{Val: v}
	q := []*TreeNode{root}
	var cur *TreeNode
	for len(a) > 0 {
		cur, q = q[0], q[1:] // 当前节点
		v, a = a[0], a[1:]   // 左子树
		if v != 0 {
			node := &TreeNode{Val: v}
			cur.Left = node
			q = append(q, node)
		}
		if len(a) > 0 {
			v, a = a[0], a[1:] // 右子树
			if v != 0 {
				node := &TreeNode{Val: v}
				cur.Right = node
				q = append(q, node)
			}
		}
	}

	return root
}

// 遍历
func TreeTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序 preOrder
	TreeTraversal(root.Left)
	// 中序 inOrder
	TreeTraversal(root.Right)
	// 后序
}

// 序列化
// 字符串: , 分隔 # 表示 null
// go json: 0 表示 null
func treeEncode(root *TreeNode) string {
	var a []int
	TreeSerialize(root, &a)
	// []int -> []string
	s := make([]string, len(a))
	for i, _ := range a {
		if a[i] == -1 {
			s[i] = "null"
		} else {
			s[i] = strconv.Itoa(a[i])
		}
	}
	return strings.Join(s, ",")
}
func treeDecode(s string) *TreeNode {
	a := strings.Split(s, ",")
	b := make([]int, 0, len(a))
	for i, _ := range a {
		if a[i] == "null" {
			b[i] = -1
		} else {
			b[i], _ = strconv.Atoi(a[i])
		}
	}
	return TreeDeserialize(&b)
}

func TreeSerialize(root *TreeNode, a *[]int) {
	if root == nil {
		*a = append(*a, -1)
		return
	}
	*a = append(*a, root.Val)
	TreeSerialize(root.Left, a)
	TreeSerialize(root.Right, a)
}

func TreeDeserialize(a *[]int) *TreeNode {
	if len(*a) == 0 {
		return nil
	}

	v := (*a)[0]
	*a = (*a)[1:]
	if v == -1 {
		return nil
	}
	root := &TreeNode{Val: v}
	root.Left = TreeDeserialize(a)
	root.Right = TreeDeserialize(a)
	return root
}

func TreeBFS(root *TreeNode) {
	q := []*TreeNode{root} // queue
	var cur *TreeNode
	for len(q) > 0 {
		cur, q = q[0], q[1:] // pop
		if cur.Left != nil {
			q = append(q, cur.Left) // push
		}
		if cur.Right != nil {
			q = append(q, cur.Right) // push
		}
	}
}

func TreeBFS2(root *TreeNode) {
	p := []*TreeNode{root}
	for len(p) > 0 {
		var q []*TreeNode
		for _, node := range p { // 一层层遍历
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		p = q
	}
}

// BST, binary search tree
func BST(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	mid := len(a) / 2
	root := &TreeNode{Val: a[mid]}
	root.Left = BST(a[:mid])
	root.Right = BST(a[mid+1:])
	return root
}
