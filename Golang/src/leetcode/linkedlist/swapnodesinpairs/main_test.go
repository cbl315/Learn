package swapnodesinpairs_test

import (
	"leetcode/linkedlist/swapnodesinpairs"
	"testing"
)

func TestFunc(t *testing.T) {
	in := &swapnodesinpairs.ListNode{
		Val: 1,
		Next: &swapnodesinpairs.ListNode{
			Val: 2,
			Next: &swapnodesinpairs.ListNode{
				Val: 3,
				Next: &swapnodesinpairs.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	//out := swapnodesinpairs.MySwapPairs(in)
	out := swapnodesinpairs.SwapPairs(in)
	expected := []int{2, 1, 4, 3}
	index := 0
	for {
		t.Log(out.Val)
		if out.Val != expected[index] {
			t.Fatal("not pass")
		}
		index++
		if out.Next == nil {
			break
		}
		out = out.Next
	}
	t.Log("pass")
}
