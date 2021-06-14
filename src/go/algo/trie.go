package algo

import "fmt"

func testTrieNode() {
	t := &TrieNode{Child: make(map[string]*TrieNode)}
	t.Add("hello")
	t.Add("hi")
	t.Add("he")
	b := t.Find("czl", false)
	t.Delete("hi")
	fmt.Println(t, b)
}

type TrieChild map[string]*TrieNode

type TrieNode struct {
	IsWord bool
	Child  TrieChild
}

func (root *TrieNode) Add(s string) {
	node := root
	for _, i := range s {
		b := string(i)
		if v, ok := node.Child[b]; !ok {
			newNode := &TrieNode{Child: make(TrieChild)}
			node.Child[b] = newNode
			node = newNode
		} else {
			node = v
		}
	}
	node.IsWord = true
}

func (root *TrieNode) Find(s string, isPrefix bool) bool {
	node := root
	for _, i := range s {
		k := string(i)
		if v, ok := node.Child[k]; !ok {
			return false
		} else {
			node = v
		}
	}
	if isPrefix {
		return true
	}
	return node.IsWord
}

func (root *TrieNode) Delete(s string) {
	// 前缀: pan/panda -> n.IsWord = false
	// 无分支: 删除整个单词

	node := root
	var branchKey string
	var branchNode *TrieNode
	for _, i := range s {
		k := string(i)
		if v, ok := node.Child[k]; !ok {
			return
		} else {
			if len(node.Child) > 1 {
				branchKey = k
				branchNode = node
			}
			node = v
		}
	}
	if len(node.Child) > 0 { // 有子节点
		node.IsWord = false
		return
	}
	if branchNode == nil { // 无分支: 删除整个单词
		delete(root.Child, string(s[0]))
	} else {
		delete(branchNode.Child, branchKey) // 有分支: 从分支删除剩余单词
	}
}
