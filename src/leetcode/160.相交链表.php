<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=160 lang=php
 *
 * [160] 相交链表
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
     * @param ListNode $headA
     * @param ListNode $headB
     * @return ListNode
     */
    function getIntersectionNode($headA, $headB)
    {
        $a = $headA;
        $b = $headB;
        while ($a !== $b) { // 注意, 这里要用 !==
            $a = $a ? $a->next : $headB;
            $b = $b ? $b->next : $headA;
        }
        return $a;
    }
}
// @lc code=end

