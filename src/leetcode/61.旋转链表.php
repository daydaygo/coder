<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=61 lang=php
 *
 * [61] 旋转链表
 */

// @lc code=start

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution
{

    /**
     * @param ListNode $head
     * @param Integer $k
     * @return ListNode
     */
    function rotateRight($head, $k)
    {
        if (!$head || !$head->next) return $head;

        $p = $head;
        $n = 0;
        while ($p) {
            $n++;
            $p = $p->next;
        }
        $k = $k % $n;
        $p = $q = $head; // $p 快指针; $q 慢指针
        while ($p->next) {
            $p = $p->next;
            if ($k > 0) $k--;
            else $q = $q->next;
        }
        $p->next = $head;
        $head = $q->next;
        $q->next = null;

        return $head;
    }
}

// @lc code=end

class SolutionTest extends TestCase
{

    public function testRotateRight()
    {
        $ans = (new Solution())->rotateRight(newLink([1, 2, 3, 4, 5]), 2);
        $this->assertEquals(newLink([4, 5, 1, 2, 3]), $ans);

        $ans = (new Solution())->rotateRight(newLink([0, 1, 2]), 4);
        $this->assertEquals(newLink([2, 0, 1]), $ans);
    }
}
