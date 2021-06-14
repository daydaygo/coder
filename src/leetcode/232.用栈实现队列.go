package leetcode
/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
    StackPush []int
    StackPop  []int
}

/** Initialize your data structure here. */
func Constructor232() MyQueue {
    return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
    this.StackPush = append(this.StackPush, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    this.Transfer()
    var x int
    x, this.StackPop = this.StackPop[0], this.StackPop[1:]
    return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
    this.Transfer()
    return this.StackPop[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.StackPop) == 0 && len(this.StackPush) == 0
}

// StackPush 不为空的时候, 转移到 StackPoP 中
func (this *MyQueue) Transfer() {
    var x int
    for len(this.StackPush)>0 {
        x, this.StackPush = this.StackPush[0], this.StackPush[1:] // pop
        this.StackPop = append(this.StackPop, x) // push
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

