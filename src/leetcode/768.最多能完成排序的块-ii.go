package leetcode

/*
 * @lc app=leetcode.cn id=768 lang=golang
 *
 * [768] 最多能完成排序的块 II
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
	var stack []int // 单调递增栈, stack[-1] 栈顶
	for _, a := range arr {
		// 遇到一个比栈顶小的元素，而前面的块不应该有比 a 小的
		// 而栈中每一个元素都是一个块，并且栈的存的是块的最大值，因此栈中比 a 小的值都需要 pop 出来
		if len(stack) > 0 && stack[len(stack)-1] > a {
			// 我们需要将融合后的区块的最大值重新放回栈
			// 而 stack 是递增的，因此 stack[-1] 是最大的
			cur := stack[len(stack)-1]
			// 维持栈的单调递增
			for len(stack) > 0 && stack[len(stack)-1] > a {
				stack = stack[:len(stack)-1] // pop
			}
			stack = append(stack, cur) // push
		} else {
			stack = append(stack, a) // push
		}
	}
	// 栈存的是块信息，因此栈的大小就是块的数量
	return len(stack)
}

// @lc code=end
