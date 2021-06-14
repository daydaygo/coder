package leetcode

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// x 到环的距离; y 环的长度; z p/q环上相遇的位置; a/b/c 正整数, 表示绕环整数圈
	// k=x+ay+z; 2k=x+by+z => x+z=cy, x=cy-z
	p := head // 快指针, 2倍速
	q := head // 慢指针, 1倍速
	k := true // 第一次执行
	for p != q || k {
		k = false
		if p == nil || p.Next == nil {
			return nil
		}
		p = p.Next.Next
		q = q.Next
	}
	// 由于 x=cy-z, p 重置为 head, q 此时在 z 处, 则正好在环起点相遇
	p = head
	for p != q {
		p = p.Next
		q = q.Next
	}
	return p
}

// @lc code=end
