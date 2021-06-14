package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    l, r := 0, len(height)-1
    lMax, rMax := height[l], height[r]
    ans := 0
    for l < r {
        if height[l] < height[r] {
            if height[l] < lMax {
                ans += lMax - height[l]
            } else {
                lMax = height[l]
            }
            l++
        } else {
            if height[r] < rMax {
                ans += rMax - height[r]
            } else {
                rMax = height[r]
            }
            r--
        }
    }
    return ans
}
// @lc code=end
