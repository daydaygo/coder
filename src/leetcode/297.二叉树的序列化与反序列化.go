package leetcode

import (
    "strconv"
    "strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    ans := ""
    q := []*TreeNode{root} // queue
    var cur *TreeNode
    for len(q) > 0 {
        cur, q = q[0], q[1:]
        if cur != nil {
            ans += strconv.Itoa(cur.Val) + ","
            q = append(q, cur.Left)
            q = append(q, cur.Right)
        } else {
            ans += "#,"
        }
    }
    return ans[:len(ans)-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    if data == "#" {
        return nil
    }

    a := strings.Split(data, ",")
    var s string
    s, a = a[0], a[1:]
    v, _ := strconv.Atoi(s)
    root := &TreeNode{Val: v}
    q := []*TreeNode{root} // queue
    var cur, newNode *TreeNode
    for len(a) > 0 {
        cur, q = q[0], q[1:] // pop

        s, a = a[0], a[1:] // 左子树
        if s != "#" {
            v, _ := strconv.Atoi(s)
            newNode = &TreeNode{Val: v}
            cur.Left = newNode
            q = append(q, newNode)
        }

        if len(a) == 0 {
            return root
        }

        s, a = a[0], a[1:] // 右子树
        if s != "#" {
            v, _ := strconv.Atoi(s)
            newNode = &TreeNode{Val: v}
            cur.Right = newNode
            q = append(q, newNode)
        }
    }
    return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
