package leetcode

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	k = k % n
    // p 为快指针, q 为慢指针
	p = head
	q := head
    for p.Next!=nil {
        p = p.Next
        if k>0 {
            k--
        } else {
            q = q.Next
        }
    }
    // 更新指针
    p.Next = head
    head = q.Next
    q.Next = nil

	return head
}

// @lc code=end
