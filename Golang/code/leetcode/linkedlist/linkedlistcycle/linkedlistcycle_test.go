package linkedlistcycle_test

import (
	"github.com/hfutcbl/learn/golang/leetcode/linkedlist/linkedlistcycle"
	"testing"
)

func TestHasCycly(t *testing.T) {
	var dataSet = []struct {
		in   *linkedlistcycle.ListNode
		want bool
	}{
		{&linkedlistcycle.ListNode{
			Val: 1,
			Next: &linkedlistcycle.ListNode{
				Val:  2,
				Next: nil,
			},
		}, false},
	}
	for i, data := range dataSet {
		result := linkedlistcycle.HasCycle(data.in)
		if result != data.want {
			t.Errorf("%d, got %t, want %t", i, result, data.want)
		}
	}
}
