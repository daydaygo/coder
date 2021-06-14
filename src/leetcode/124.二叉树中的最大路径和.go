package leetcode
/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func max124(a, b int) int {
	if a>b {
		return a
	} 
	return b
}
var ans = ^int(^uint(0) >> 1)
func oneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max124(0, oneSideMax(root.Left))
	right := max124(0, oneSideMax(root.Right))
	ans = max124(ans, left + right + root.Val)
	return max124(left, right) + root.Val
}
func maxPathSum(root *TreeNode) int {
	oneSideMax(root)
	return ans
}
// @lc code=end

