package algo

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums { // 一次遍历即可
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{i, j}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}
