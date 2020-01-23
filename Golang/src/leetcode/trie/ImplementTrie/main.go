/*
Implement a trie with insert, search, and startsWith methods.

Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.

链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
*/
package ImplementTrie

type Trie struct {
	val      string
	isWord   bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		val:      "",
		isWord:   false,
		children: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	var node = t
	for wordIndex, char := range word {
		currVal := word[:wordIndex+1]
		index := char - 'a'
		if node.children[index] != nil {
			node = node.children[index]
		} else {
			child := &Trie{
				val:      currVal,
				children: [26]*Trie{},
			}
			node.children[index] = child
			node = child
		}

	}
	node.isWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	var node = t
	for _, char := range word {
		index := char - 'a'
		if node.children[index] != nil {
			node = node.children[index]
		} else {
			return false
		}
	}
	return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	var node = t
	for _, char := range prefix {
		index := char - 'a'
		if node.children[index] != nil {
			node = node.children[index]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
