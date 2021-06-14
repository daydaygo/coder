package leetcode

const UINT_MIN uint = 0
const UINT_MAX uint = ^uint(0)
const INT_MAX int = int(^uint(0) >> 1) // 首位 0, 其余 1
const INT_MIN int = ^INT_MAX           // 首位 1, 其余 0

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
