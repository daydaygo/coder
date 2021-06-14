package algo

// PreSum 前缀和
func PreSum(a []int) []int {
	var nums []int
	preSum := make([]int, len(nums)+1)
	preSumI := 0 // 空间压缩
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
		preSumI = preSumI + nums[i] // 空间压缩要注意何时更新下一个值
	}
	return preSum
}

// MonotoneStack 单调栈
func MonotoneStack(a []int) []int {
	// 单调递减: 用来求解下一个更大的元素
	// 变形: stack 存索引, 可以得到距离
	var stack []int
	ans := make([]int, len(a), len(a)) // 根据题目要求使用 slice/map
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

func MonotoneStack2(a []int) map[int]int {
	var stack []int          // 单调递减: 用来求解下一个更大的元素
	ans := make(map[int]int) // 当前值: 下一个更大的元素
	for i := 0; i < len(a); i++ {
		for len(stack) > 0 && a[i] > stack[0] {
			ans[stack[0]] = a[i]
			stack = stack[1:] // pop
		}
		stack = append([]int{a[i]}, stack...) // push
	}
	for i := 0; i < len(stack); i++ { // 还在 stack 中的没有下一个更大的元素
		ans[stack[i]] = -1
	}
	return ans
}
