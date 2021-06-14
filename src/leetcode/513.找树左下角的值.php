<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=513 lang=php
 *
 * [513] 找树左下角的值
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
    function findBottomLeftValue($root)
    {
        $curLevel = [$root];
        $res = $root->val;
        while (count($curLevel)) {
            $nextLevel = [];
            $res = $curLevel[0]->val;
            foreach ($curLevel as $node) {
                if ($node->left) $nextLevel[] = $node->left;
                if ($node->right) $nextLevel[] = $node->right;
            }
            $curLevel = $nextLevel;
        }
        return $res;
    }
}
// @lc code=end

