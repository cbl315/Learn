/*
Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

链接：https://leetcode-cn.com/problems/word-search-ii
*/
package WordSearchII

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

type Result struct {
	Find []string
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	length, width := len(board[0]), len(board)
	var visited = make([][]bool, width)
	for i := 0; i < width; i++ {
		visited[i] = make([]bool, length)
	}
	result := Result{}
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			dfs(board, i, j, trie, visited, "", &result)
		}
	}
	return result.Find
}

func isIn(s []string, word string) bool {
	for _, one := range s {
		if one == word {
			return true
		}
	}
	return false
}

//
func dfs(board [][]byte, i, j int, trie Trie, visited [][]bool, subStr string, result *Result) {
	// 判断输入是否越界
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	// 判断是否已经访问过
	if visited[i][j] == true {
		return
	}
	// 通过字典树 判断查找路径是否正确
	subStr += string(board[i][j])
	if !trie.StartsWith(subStr) {
		return
	}
	if trie.Search(subStr) {
		if !isIn(result.Find, subStr) {
			result.Find = append(result.Find, subStr)
		}
	}

	visited[i][j] = true
	dfs(board, i+1, j, trie, visited, subStr, result)
	dfs(board, i-1, j, trie, visited, subStr, result)
	dfs(board, i, j+1, trie, visited, subStr, result)
	dfs(board, i, j-1, trie, visited, subStr, result)
	visited[i][j] = false
}
