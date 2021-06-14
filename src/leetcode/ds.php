<?php

class ListNode
{
    public $val = 0;
    public $next = null;

    function __construct($val = 0, $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }
}

function newLink($arr)
{
    if (!$arr) return null;
    $head = new ListNode($arr[0]);
    for ($i = count($arr) - 1; $i >= 1; $i--) {
        $p = new ListNode($arr[$i], $head->next);
        $head->next = $p;
    }
    return $head;
}

class TreeNode
{
    public $val = null;
    public $left = null;
    public $right = null;

    function __construct($val = 0, $left = null, $right = null)
    {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }
}