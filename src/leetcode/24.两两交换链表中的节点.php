<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=24 lang=php
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution
{

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function swapPairs($head)
    {
        if (!$head || !$head->next) return $head;

        /** @var ListNode $next */
        $next = $head->next;
        $head->next = (new Solution())->swapPairs($next->next); // 递归已经将后面链表处理好, 拼接到前面的元素上
        $next->next = $head;
        return $next;
    }
}

// @lc code=end

class SolutionTest extends TestCase
{
    public function testSwapPairs()
    {
        $ans = (new Solution())->swapPairs(newLink([1, 2, 3, 4]));
        $this->assertEquals(newLink([2, 1, 4, 3]), $ans);
    }
}