package reversenodesinkgroup_test

import (
	"fmt"
	"github.com/hfutcbl/learn/golang/leetcode/linkedlist/reversenodesinkgroup"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	in := &reversenodesinkgroup.ListNode{
		Val: 1,
		Next: &reversenodesinkgroup.ListNode{
			Val: 2,
			Next: &reversenodesinkgroup.ListNode{
				Val: 3,
				Next: &reversenodesinkgroup.ListNode{
					Val: 4,
					Next: &reversenodesinkgroup.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	out := reversenodesinkgroup.ReverseKGroup(in, 2)
	expected := []int{2, 1, 4, 3, 5}
	index := 0
	for {
		fmt.Printf("%v\n", out)
		if out.Val != expected[index] {
			t.Fatalf("not pass, got=%d, want=%d", out.Val, expected[index])
		}
		if out.Next != nil {
			out = out.Next
		} else {
			break
		}
		index++
		if index == len(expected)-1 {
			t.Log("pass")
			break
		}
	}
}

func TestReverseLinkedListMd(t *testing.T) {
	in := &reversenodesinkgroup.ListNode{
		Val: 1,
		Next: &reversenodesinkgroup.ListNode{
			Val: 2,
			Next: &reversenodesinkgroup.ListNode{
				Val: 3,
				Next: &reversenodesinkgroup.ListNode{
					Val: 4,
					Next: &reversenodesinkgroup.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	out := reversenodesinkgroup.ReverseKGroupModify(in, 2)
	expected := []int{2, 1, 4, 3, 5}
	index := 0
	for {
		fmt.Printf("%v\n", out)
		if out.Val != expected[index] {
			t.Fatalf("not pass, got=%d, want=%d", out.Val, expected[index])
		}
		if out.Next != nil {
			out = out.Next
		} else {
			break
		}
		index++
		if index == len(expected)-1 {
			t.Log("pass")
			break
		}
	}
}
