/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/
package binaryTreeLevelOrderTraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// bfs
	currLvlNodes := []*TreeNode{root}
	// visted = []*TreeNode{root}
	for len(currLvlNodes) > 0 {
		currLvlVal := make([]int, len(currLvlNodes))
		currSize := len(currLvlNodes)
		for i := 0; i < currSize; i++ {
			node := currLvlNodes[i]
			currLvlVal[i] = node.Val
			if node.Left != nil {
				currLvlNodes = append(currLvlNodes, node.Left)
			}
			if node.Right != nil {
				currLvlNodes = append(currLvlNodes, node.Right)
			}
		}
		currLvlNodes = currLvlNodes[currSize:]
		result = append(result, currLvlVal)
	}
	return result
}

func levelOrder2(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	// dfs
	result = dfs(root, 0, result)
	return result
}

func dfs(root *TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return result
	}
	var currLvlVal []int
	if len(result) < level+1 {
		result = append(result, currLvlVal)
	}
	result[level] = append(result[level], root.Val)
	result = dfs(root.Left, level+1, result)
	result = dfs(root.Right, level+1, result)
	return result
}
