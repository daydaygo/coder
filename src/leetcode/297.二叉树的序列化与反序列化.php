<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=297 lang=php
 *
 * [297] 二叉树的序列化与反序列化
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
class Codec
{
    function __construct()
    {

    }

    /**
     * @param TreeNode $root
     * @return String
     */
    function serialize(TreeNode $root)
    {
        $ans = "";
        $q = [$root]; // queue
        while ($q) {
            /** @var TreeNode $node */
            $node = array_unshift($q); // pop
            if ($node) {
                $ans .= $node->val . ",";
                array_push($node->left, $node->right);
            } else {
                $ans .= "#,";
            }
        }
        return rtrim($ans);
    }

    /**
     * @param String $data
     * @return void
     */
    function deserialize(string $data)
    {
        if ($data == '#') return null;

        $a = explode(',', $data);
        $v = array_unshift($a);
        $root = new TreeNode((int)$v);
        $q = [$root];
        while ($a) {
            /** @var TreeNode $cur */
            $cur = array_unshift($q);
            $v = array_unshift($a); // 左子树
            if ($v != '#') {
                $newNode = new TreeNode((int)$v);
                $cur->left = $newNode;
                array_push($q, $newNode);
            }
            if (!$a) return $root;
            $v = array_unshift($a); // 右子树
            if ($v != '#') {
                $newNode = new TreeNode((int)$v);
                $cur->right = $newNode;
                array_push($q, $newNode);
            }
        }
        return $root;
    }
}

/**
 * Your Codec object will be instantiated and called as such:
 * $ser = Codec();
 * $deser = Codec();
 * $data = $ser->serialize($root);
 * $ans = $deser->deserialize($data);
 */
// @lc code=end

