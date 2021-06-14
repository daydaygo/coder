<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=109 lang=php
 *
 * [109] 有序链表转换二叉搜索树
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

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution
{

    /**
     * @param ListNode $head
     * @return TreeNode
     */
    function sortedListToBST($head)
    {
        $arr = [];
        while ($head) {
            $arr[] = $head->val;
            $head = $head->next;
        }
        return $this->BST($arr);
    }

    function BST(array $arr)
    {
        if (!$arr) return null;
        $mid = (int)(count($arr) / 2);
        $root = new TreeNode($arr[$mid]);
        $root->left = $this->BST(array_slice($arr, 0, $mid));
        $root->right = $this->BST(array_slice($arr, $mid + 1));
        return $root;
    }
}
// @lc code=end

