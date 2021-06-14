/*
 * @lc app=leetcode.cn id=75 lang=php
 *
 * [75] 颜色分类
 */

// @lc code=start
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function sortColors(&$nums) {
        $cnt0 = $cnt1 = $cnt2 = 0;
        foreach ($nums as $num) {
            if ($num === 0) {
                $cnt0++;
            }
            if ($num === 1) {
                $cnt1++;
            }
            if ($num === 2) {
                $cnt2++;
            }
        }
        for ($i=0; $i<$cnt0; $i++) {
            $nums[$i] = 0;
        }
        for ($i=0; $i<$cnt1; $i++) {
            $nums[$cnt0+$i] = 1;
        }
        for ($i=0; $i<$cnt2; $i++) {
            $nums[$cnt0+$cnt1+$i] = 2;
        }
    }
}
// @lc code=end

