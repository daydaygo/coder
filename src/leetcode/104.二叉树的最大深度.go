package leetcode

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	q := []*TreeNode{root, nil} // queue
	var node *TreeNode
	for len(q) > 0 {
		node, q = q[0], q[1:] // pop
		if node != nil {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		} else if len(q) > 0 { // 注意要判断队列是否只有一个 nil
			q = append(q, nil)
			depth++
		}
	}
	return depth
}

// @lc code=end

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max104(maxDepth(root.Left), maxDepth(root.Right))
}
func max104(a, b int) int {
	if a > b {
		return a
	}
	return b
}
