<?php

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=768 lang=php
 *
 * [768] 最多能完成排序的块 II
 */

// @lc code=start
class Solution
{

    /**
     * @param Integer[] $arr
     * @return Integer
     */
    function maxChunksToSorted($arr)
    {
        $stack = []; // 单调递增栈, stack[-1] 栈顶
        foreach ($arr as $a) {
            // 遇到一个比栈顶小的元素，而前面的块不应该有比 a 小的
            // 而栈中每一个元素都是一个块，并且栈的存的是块的最大值，因此栈中比 a 小的值都需要 pop 出来
            if ($stack && $stack[count($stack) - 1] > $a) {
                $cur = $stack[count($stack) - 1];
                // 维持栈的单调递增
                while ($stack && $stack[count($stack) - 1] > $a) array_pop($stack);
                array_push($stack, $cur);
            } else array_push($stack, $a);
        }
        // 栈存的是块信息，因此栈的大小就是块的数量
        return count($stack);
    }
}

// @lc code=end

class SolutionTest extends TestCase
{
    public function testMaxChunksToSorted()
    {
        $ans = (new Solution())->maxChunksToSorted([5, 4, 3, 2, 1]);
        $this->assertEquals(1, $ans);

        $ans = (new Solution())->maxChunksToSorted([2, 1, 3, 4, 4]);
        $this->assertEquals(4, $ans);
    }
}