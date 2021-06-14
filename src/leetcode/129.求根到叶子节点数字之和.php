<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=129 lang=php
 *
 * [129] 求根到叶子节点数字之和
 */

// @lc code=start

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution
{

    /**
     * @param TreeNode $root
     * @return Integer
     */
    function sumNumbers(TreeNode $root)
    {
        return helper($root, 0);
    }
}

/**
 * @param TreeNode $root
 * @param int $cur
 * @return int
 */
function helper(TreeNode $root, int $cur)
{
    if (!$root) return $cur;
    $next = $cur * 10 + $root->val;
    if (!$root->left && !$root->right) return $next;

    $l = helper($root->left, $next);
    $r = helper($root->right, $next);
    return $l + $r;
}
// @lc code=end

