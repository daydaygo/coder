<?php
require_once 'ds.php';

use PHPUnit\Framework\TestCase;

/*
 * @lc app=leetcode.cn id=146 lang=php
 *
 * [146] LRU缓存机制
 */

// @lc code=start
class LRUCache
{
    public $hash = []; // key => DoubleListNode
    public $cache;
    public $size = 0;
    public $cap; // 容量

    /**
     * @param Integer $capacity
     */
    function __construct($capacity)
    {
        $this->cap = $capacity;
        $this->cache = new DoubleList();
    }

    /**
     * @param Integer $key
     * @return Integer
     */
    function get($key)
    {
        if (!isset($this->hash[$key])) return -1;

        // 更新节点
        /** @var DoubleListNode $node */
        $node = $this->hash[$key];
        $this->cache->remove($node);
        $this->cache->append($node);

        return $node->val;
    }

    /**
     * @param Integer $key
     * @param Integer $value
     * @return NULL
     */
    function put($key, $value)
    {
        if (isset($this->hash[$key])) { // key 存在
            // 更新节点
            /** @var DoubleListNode $node */
            $node = $this->hash[$key];
            $this->cache->remove($node);
            $this->cache->append($node);
            $node->val = $value;
        } else { // key 不存在, 新增节点
            $node = new DoubleListNode($key, $value);
            $this->cache->append($node);
            $this->hash[$key] = $node;

            if ($this->size == $this->cap) {
                // 删除原节点
                $oldNode = $this->cache->tail->pre;
                $this->cache->remove($oldNode);
                unset($this->hash[$oldNode->key]);
            } else {
                $this->size++;
            }
        }
    }
}

class DoubleListNode
{
    public $key;
    public $val;
    public $pre;
    public $next;

    public function __construct($key = null, $val = null)
    {
        if ($key) $this->key = $key;
        if ($val) $this->val = $val;
    }
}

class DoubleList
{
    public $head;
    public $tail;

    public function __construct()
    {
        $this->head = new DoubleListNode();
        $this->tail = new DoubleListNode();
        $this->head->next = $this->tail;
        $this->tail->pre = $this->head;
    }

    public function append(DoubleListNode $node)
    {
        $node->pre = $this->head;
        $node->next = $this->head->next;
        $this->head->next->pre = $node;
        $this->head->next = $node;
    }

    public function remove(DoubleListNode $node)
    {
        $node->pre->next = $node->next;
        $node->next->pre = $node->pre;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * $obj = LRUCache($capacity);
 * $ret_1 = $obj->get($key);
 * $obj->put($key, $value);
 */
// @lc code=end

