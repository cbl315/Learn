/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
*/
package minimumDepthOfTree

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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)
	if leftMinDepth == 0 || rightMinDepth == 0 {
		return leftMinDepth + rightMinDepth + 1
	} else if leftMinDepth > rightMinDepth {
		return 1 + rightMinDepth
	} else {
		return 1 + leftMinDepth
	}
}
