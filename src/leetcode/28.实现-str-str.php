/*
 * @lc app=leetcode.cn id=28 lang=php
 *
 * [28] 实现 strStr()
 */

// @lc code=start
class Solution {

    /**
     * @param String $haystack
     * @param String $needle
     * @return Integer
     */
    function strStr($haystack, $needle) {
        if ($needle === '') {
            return 0;
        }
        $a = strpos($haystack, $needle);
        if ($a === false) {
            return -1;
        }
        return $a;
    }
}
// @lc code=end

