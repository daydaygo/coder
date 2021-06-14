package leetcode

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := &ListNode{}
	cur := head
	head = head.Next // 如果能够翻转, head.Next 就是新的 head

	for cur != nil && cur.Next != nil {
		next := cur.Next
		// 更新 next 指向
		cur.Next = next.Next
		pre.Next = next
		next.Next = cur
		// 更新节点位置
		pre = cur
		cur = cur.Next
	}

	return head
}

// @lc code=end

// 只交换节点值
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p != nil && p.Next != nil {
		p.Val, p.Next.Val = p.Next.Val, p.Val
		p = p.Next.Next
	}
	return head
}

// 递归
func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next) // 剩下的节点递归已经处理好, 拼接到前 2 个节点上
	next.Next = head
	return next
}
