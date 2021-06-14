package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

var listNodeP *ListNode

func LinkToArray(head *ListNode) []int {
	var a []int
	// p := head // 记录 head 初始位置
	for head != nil { // 线性遍历
		a = append(a, head.Val)
		head = head.Next
	}
	return a
}

// 链表遍历: 线性 + 递归
func LinkTraverseLinear(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		// p.Val
	}
}
func LinkTraverseRecursion(head *ListNode) {
	// 前序遍历
	LinkTraverseRecursion(head.Next)
	// 后续遍历
}

// 头插法
func NewLink(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val: a[0]}
	for i := len(a) - 1; i >= 1; i-- {
		p := &ListNode{
			Val:  a[i],
			Next: head.Next,
		}
		head.Next = p
	}
	return head
}

// 尾插法
func NewLink2(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val: a[0]}
	p := head
	for i := 1; i < len(a); i++ {
		t := &ListNode{Val: a[i]}
		p.Next = t
		p = t
	}
	return head
}

// 翻转: 递归
func Reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := Reverse(head.Next) // 剩下的部分翻转后, last -> head
	head.Next.Next = head
	head.Next = nil
	return last
}

// 翻转: 迭代法; 翻转 [a, b)
func reverse(a, b *ListNode) *ListNode {
	pre := &ListNode{}
	cur := a
	for cur != b { // 如果 b=nil, 就是整个链表
		next := cur.Next
		// 逐个节点反转
		cur.Next = pre
		// 更新指针位置
		pre = cur
		cur = next
	}
	return pre
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	a := head
	b := head
	for i := 0; i < k; i++ { // 不足 k 个元素, 直接返回
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)     // 翻转前 k 个元素
	a.Next = ReverseKGroup(b, k) // 后续链表递归实现 + 和前 k 个元素连接起来
	return newHead
}

// 翻转前 N 个节点
func ReverseN(head *ListNode, n int) *ListNode {
	if n == 1 { // 此时 head 为第 n 个节点, 记录第 n+1 个节点
		listNodeP = head.Next
	}
	last := ReverseN(head.Next, n-1) // 剩下的部分翻转后, last -> head
	head.Next.Next = head
	head.Next = listNodeP
	return last
}

// 翻转 [m, n]
func ReverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return ReverseN(head, n)
	}
	head.Next = ReverseBetween(head.Next, m-1, n-1)
	return head
}

func LinkMerge(a, b *ListNode) *ListNode {
	var ans ListNode
	cur := ans
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
	}
	if a != nil {
		cur.Next = a
	}
	if b != nil {
		cur.Next = b
	}
	return cur.Next
}

// 链表交点
func LinkIntersection(a, b *ListNode) *ListNode {
	// 哈希法

	// a = A + C, b = B + C, a+b = A+C+B+C
	aHead := a
	bHead := b
	for a != b {
		if a == nil {
			a = bHead
		} else {
			a = a.Next
		}
		if b == nil {
			b = aHead
		} else {
			b = b.Next
		}
	}
	return a
}

// 环的起点
func LinkRing(head *ListNode) *ListNode {
	p := head // 快指针
	q := head // 慢指针
	for p != q && p != nil {
		if p.Next == nil || p.Next.Next == nil {
			return nil
		}
		p = p.Next.Next
		q = q.Next
	}
	return p
}

func NewRingLink(a []int, b int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val: a[0]}
	if len(a) == 1 {
		return head
	}
	last := &ListNode{Val: a[len(a)-1]}
	head.Next = last
	if last.Val == head.Val {
		last.Next = head
	}
	for i := len(a) - 2; i >= 1; i-- {
		p := &ListNode{
			Val:  a[i],
			Next: head.Next,
		}
		head.Next = p
		if p.Val == last.Val {
			last.Next = p
		}
	}
	return head
}

// todo: doubleList