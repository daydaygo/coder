<?php

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=42 lang=php
 *
 * [42] 接雨水
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function trap($height)
    {
        $n = count($height);
        if (!$n) return 0;

        $l = 0;
        $l_max = $height[$l];
        $r = $n - 1;
        $r_max = $height[$r];
        $ans = 0;
        while ($l < $r) {
            if ($height[$l] < $height[$r]) {
                if ($height[$l] < $l_max) $ans += $l_max - $height[$l];
                else $l_max = $height[$l];
                $l++;
            } else {
                if ($height[$r] < $r_max) $ans += $r_max - $height[$r];
                else $r_max = $height[$r];
                $r--;
            }
        }
        return $ans;
    }
}

// @lc code=end

class SolutionTest extends TestCase
{

    public function testTrap()
    {
        $ans = (new Solution())->trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]);
        $this->assertEquals(6, $ans);

        $ans = (new Solution())->trap([4, 2, 0, 3, 2, 5]);
        $this->assertEquals(9, $ans);
    }
}