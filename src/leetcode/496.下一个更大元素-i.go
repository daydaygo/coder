package leetcode

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := MonotoneStack2(nums2)
	for i := 0; i < len(nums1); i++ {
		nums1[i] = m[nums1[i]]
	}
	return nums1
}
func MonotoneStack2(a []int) map[int]int {
	// 单调递减: 用来求解下一个更大的元素
	// 变形一: stack 存索引, 用来计算距离
	var stack []int
	ans := make(map[int]int) // 根据问题选择: [] / map
	for i := len(a) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[0] <= a[i] { // 注意边界
			stack = stack[1:] // pop
		}
		if len(stack) == 0 {
			ans[a[i]] = -1
		} else {
			ans[a[i]] = stack[0]
		}
		stack = append([]int{a[i]}, stack...) // push
	}
	return ans
}

// @lc code=end
