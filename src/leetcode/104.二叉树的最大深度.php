<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=104 lang=php
 *
 * [104] 二叉树的最大深度
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
    function maxDepth($root)
    {
        if (!$root) return 0;

        $depth = 1;
        $arr = [$root, null];
        while ($arr) {
            /** @var TreeNode $node */
            $node = array_shift($arr);
            if ($node) {
                if ($node->left) array_push($arr, $node->left);
                if ($node->right) array_push($arr, $node->right);
            } elseif ($arr) {
                $depth++;
                array_push($arr, null);
            }
        }
        return $depth;
    }
}
// @lc code=end

// function maxDepth($root)
// {
//     if (!$root) return 0;
//     return 1 + max($this->maxDepth($root->left), $this->maxDepth($root->right));
// }