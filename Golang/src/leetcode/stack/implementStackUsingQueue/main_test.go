package implementStackUsingQueue_test

import (
	"leetcode/stack/implementStackUsingQueue"
	"testing"
)

func TestStack(t *testing.T) {
	obj := implementStackUsingQueue.Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Pop()
	if param_2 != 2 {
		t.Fatalf("first pop, got %d, want %d", param_2, 2)
	}
	param_3 := obj.Top()
	if param_3 != 1 {
		t.Fatalf("top, got %d, want %d", param_3, 1)
	}
	param_4 := obj.Empty()
	if param_4 != false {
		t.Fatalf("top, got %v, want %v", param_4, false)
	}
}
