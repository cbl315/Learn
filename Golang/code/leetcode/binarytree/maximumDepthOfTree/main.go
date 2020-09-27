/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
*/
package maximumDepthOfTree

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

func maxDepth(root *TreeNode) int {
	// dfs
	if root == nil {
		return 0
	}
	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)
	if leftMaxDepth > rightMaxDepth {
		return 1 + leftMaxDepth
	} else {
		return 1 + rightMaxDepth
	}
}

/*
BFS 非递归 资源消耗更少
*/
func maxDepthBFS(root *TreeNode) (depth int) {
	if root == nil {
		return
	}
	// bfs
	currLvlNodes := []*TreeNode{root}
	// visted = []*TreeNode{root}
	for len(currLvlNodes) > 0 {
		depth++
		currSize := len(currLvlNodes)
		for i := 0; i < currSize; i++ {
			node := currLvlNodes[i]
			if node.Left != nil {
				currLvlNodes = append(currLvlNodes, node.Left)
			}
			if node.Right != nil {
				currLvlNodes = append(currLvlNodes, node.Right)
			}
		}
		currLvlNodes = currLvlNodes[currSize:]
	}
	return
}
