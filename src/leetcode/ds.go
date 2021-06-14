package leetcode

// 单调栈
func MonotoneStack(a []int) []int {
	// 单调递减: 用来求解下一个更大的元素
	// 变形一: stack 存索引, 用来计算距离
	var stack []int
	ans := make([]int, len(a), len(a)) // 根据问题选择: [] / map
	for i := len(a) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[0] <= a[i] { // 注意边界
			stack = stack[1:] // pop
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[0]
		}
		stack = append([]int{a[i]}, stack...) // push
	}
	return ans
}
