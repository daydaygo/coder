package leetcode

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
    n := len(nums)
    nums = append(nums, nums...)
    ans := MonotoneStack503(nums)
    return ans[:n]
}
func MonotoneStack503(a []int) []int {
    var stack []int // 单调递减: 用来求解下一个更大的元素
    n := len(a)
    ans := make([]int, n, n)
    for i := n*2 - 1; i >= 0; i-- { // 假设此时数组已翻倍
        for len(stack) > 0 && stack[0] <= a[i%n] { // 注意边界
            stack = stack[1:] // pop
        }
        if len(stack) == 0 {
            ans[i%n] = -1
        } else {
            ans[i%n] = stack[0]
        }
        stack = append([]int{a[i%n]}, stack...) // push
    }
    return ans
}

// @lc code=end
