package reverselinkedlist_test

import (
	"fmt"
	"leetcode/linkedlist/reverselinkedlist"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	in := &reverselinkedlist.ListNode{
		Val: 1,
		Next: &reverselinkedlist.ListNode{
			Val: 2,
			Next: &reverselinkedlist.ListNode{
				Val: 3,
				Next: &reverselinkedlist.ListNode{
					Val: 4,
					Next: &reverselinkedlist.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	out := reverselinkedlist.ReverseList(in)
	expected := 5
	for {
		fmt.Printf("%v\n", out)
		if out.Val != expected {
			t.Fatal("not pass")
		}
		if out.Next != nil {
			out = out.Next
		} else {
			break
		}
		expected--
	}
	if expected == 1 {
		t.Log("pass")
	}
}
