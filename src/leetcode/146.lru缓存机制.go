package leetcode

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
type LRUCache struct {
	Cap   int
	Hash  map[int]*DoubleListNode
	Cache *DoubleList
	Size  int
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		Cap:   capacity,
		Hash:  make(map[int]*DoubleListNode),
		Cache: NewDoubleList(),
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.Hash[key]
	if !ok {
		return -1
	}

	// 设为最近使用过
	this.Cache.remove(n)
	this.Cache.append(n)

	return n.Val
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.Hash[key]
	if !ok { // cache 不存在
		// 增加节点
		newNode := &DoubleListNode{Key: key, Val: value}
		this.Hash[key] = newNode
		this.Cache.append(newNode)

		if this.Cap == this.Size { // 已满
			// 删除尾节点
			tailNode := this.Cache.tail.Pre
			delete(this.Hash, tailNode.Key)
			this.Cache.remove(tailNode)
		} else {
			this.Size++
		}
	} else {
		// cache 存在
		// 设为最近使用过
		this.Cache.remove(n)
		this.Cache.append(n)

		n.Val = value // 更新值
	}
}

type DoubleListNode struct {
	Key, Val  int
	Pre, Next *DoubleListNode
}

type DoubleList struct {
	Head, tail *DoubleListNode
}

func NewDoubleList() *DoubleList {
	dl := &DoubleList{
		Head: &DoubleListNode{},
		tail: &DoubleListNode{},
	}
	dl.Head.Next = dl.tail
	dl.tail.Pre = dl.Head
	return dl
}

func (dl *DoubleList) append(n *DoubleListNode) {
	n.Next = dl.Head.Next
	n.Pre = dl.Head
	dl.Head.Next.Pre = n
	dl.Head.Next = n
}

func (dl *DoubleList) remove(n *DoubleListNode) {
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
