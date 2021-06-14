package leetcode

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int) // 元素:元素出现次数
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
		} else {
			m[num] = 1
		}
	}
	m2 := make(map[int]int) // 元素出现次数:元素
	var a []int
	for i, j := range m {
		m2[j] = m2[i]
		a = append(a, j)
	}
	// 对 a 进行快速选择
	l := 0
	r := len(a) - 1
	for l <= r {
		for a[l] < a[k-1] {
			l++
		}
		for a[r] > a[k-1] {
			r--
		}
		a[l], a[r] = a[r], a[l]
	}

	return nil
}

// @lc code=end
