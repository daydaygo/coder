<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=142 lang=php
 *
 * [142] 环形链表 II
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
     * @return ListNode
     */
    function detectCycle($head)
    {
        // x 到环的距离; y 环的长度; z p/q还上相遇的位置
        // k=x+ay+z; 2k=x+by+z => x+z=cy, x=cy-z
        $p = $q = $head; // $p 快指针, 2倍速; $q 慢指针, 1倍速
        $k = true; // 第一次执行
        while ($p != $q || $k) {
            $k = false;
            if (!$p || !$p->next) return null;
            $p = $p->next->next;
            $q = $q->next;
        }
        // 由于 x=cy-z, p 重置为 head, q 此时在 z 处, 则正好在环起点相遇
        $p = $head;
        while ($p != $q) {
            $p = $p->next;
            $q = $q->next;
        }
        return $p;
    }
}
// @lc code=end

