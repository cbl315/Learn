/*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

链接：https://leetcode-cn.com/problems/validate-binary-search-tree
*/
package validateBST

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if min >= node.Val || max <= node.Val {
		return false
	}
	// 往左边递归时 更新min的值为当前父结点的值 保证左子树中 所有结点的值需要比根节点小
	// 往右边递归时 更新max的值为当前父节点的值 保证右子树中 所有节点的值需要比根节点大
	return valid(node.Left, min, node.Val) && valid(node.Right, node.Val, max)
}

// is in-order result desc
func isValidBST2(root *TreeNode) bool {
	// in-order
	inOrdered := inOrderSearch(root)
	expected := make([]int, len(inOrdered))
	copy(expected, inOrdered)
	sort.Ints(expected)
	for i, v := range inOrdered {
		if v != expected[i] {
			return false
		}
		if i > 0 && v <= expected[i-1] {
			return false
		}
	}
	return true
}

func inOrderSearch(root *TreeNode) (ordered []int) {
	if root == nil {
		return
	}
	lOrdered := inOrderSearch(root.Left)
	rOrdered := inOrderSearch(root.Right)
	ordered = append(lOrdered, root.Val)
	ordered = append(ordered, rOrdered...)
	return ordered
}
