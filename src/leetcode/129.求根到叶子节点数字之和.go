package leetcode

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, cur int) int {
	if root == nil {
		return 0 // 当前非叶子节点, 不计算
	}

	next := cur*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return next // 当前为叶子节点, 计算
	}

	l := helper(root.Left, next)
	r := helper(root.Right, next)
	return l + r
}

// @lc code=end
