package validateBST_test

import (
	"leetcode/binarytree/validateBST"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	in := validateBST.TreeNode{
		10,
		&validateBST.TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		&validateBST.TreeNode{
			Val: 15,
			Left: &validateBST.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &validateBST.TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}
	got := validateBST.IsValidBST(&in)
	if got != false {
		t.Fatalf("not pass")
	}
}
