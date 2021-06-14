package leetcode

import (
	"reflect"
)

/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	// 层次遍历
	pa := []*TreeNode{p}
	qa := []*TreeNode{q}
	var curp, curq *TreeNode
	for len(pa) > 0 && len(qa) > 0 {
		curp, pa = pa[0], pa[1:] // pop
		curq, qa = qa[0], qa[1:]
		if curp == nil && curq == nil { // 继续比较队列中下一个
			continue
		}
		// 比较当前节点
		if curp == nil && curq != nil {
			return false
		}
		if curp != nil && curq == nil {
			return false
		}
		if curp.Val != curq.Val {
			return false
		}

		// 添加子节点
		pa = append(pa, curp.Left)
		qa = append(qa, curq.Left)
		pa = append(pa, curp.Right)
		qa = append(qa, curp.Right)
	}

	if len(pa) > 0 || len(qa) > 0 {
		return false
	}

	return true
}

// @lc code=end

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree3(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}

	// 前序遍历
	var pPre, qPre []int
	pPre = preOrder(p, pPre)
	qPre = preOrder(q, qPre)
	if !reflect.DeepEqual(pPre, qPre) {
		return false
	}

	// 中序遍历
	var pIn, qIn []int
	pIn = inOrder(p, pIn)
	qIn = inOrder(q, qIn)
	if !reflect.DeepEqual(pIn, qIn) {
		return false
	}

	return true
}

func preOrder(root *TreeNode, a []int) []int {
	if root == nil {
		return a
	}
	a = append(a, root.Val)
	a = append(a, preOrder(root.Left, a)...)
	a = append(a, preOrder(root.Right, a)...)
	return a
}

func inOrder(root *TreeNode, a []int) []int {
	if root == nil {
		return a
	}
	a = append(a, inOrder(root.Left, a)...)
	a = append(a, root.Val)
	a = append(a, inOrder(root.Right, a)...)
	return a
}
