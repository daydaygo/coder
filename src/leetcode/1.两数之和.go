package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{i, j}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

// @lc code=end
