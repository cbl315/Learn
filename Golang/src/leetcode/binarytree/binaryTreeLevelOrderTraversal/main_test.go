package binaryTreeLevelOrderTraversal_test

import (
	"leetcode/binarytree/binaryTreeLevelOrderTraversal"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	in := &binaryTreeLevelOrderTraversal.TreeNode{
		Val: 3,
		Left: &binaryTreeLevelOrderTraversal.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &binaryTreeLevelOrderTraversal.TreeNode{
			Val: 20,
			Left: &binaryTreeLevelOrderTraversal.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &binaryTreeLevelOrderTraversal.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	binaryTreeLevelOrderTraversal.LvlOrder(in)
}
