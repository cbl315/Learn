package easy

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
func rangeSumBST(root *TreeNode, L int, R int) int {
	// refer to others
	var result = 0
	if root == nil {
		return 0
	}
	if root.Val < L {
		result += rangeSumBST(root.Right, L, R)
	} else if root.Val > R {
		result += rangeSumBST(root.Left, L, R)
	} else {
		result += root.Val + rangeSumBST(root.Right, L, R) + rangeSumBST(root.Left, L, R)
	}
	return result

	/**
	first create
	  if root.Val <= R && root.Val >= L {
	      result += root.Val
	  }
	  if root.Left != nil && root.Val > L{
	  result += rangeSumBST(root.Left, L, R)
	  }
	  if root.Right != nil && root.Val < R{
	      result += rangeSumBST(root.Right, L, R)
	  }
	  return result*/
}
