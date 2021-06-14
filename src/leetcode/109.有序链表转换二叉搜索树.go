package leetcode

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	var a []int
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	return BST109(a)

}
func BST109(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	mid := len(a) / 2
	root := &TreeNode{Val: a[mid]}
	root.Left = BST109(a[:mid])
	root.Right = BST109(a[mid+1:])
	return root
}

// @lc code=end
