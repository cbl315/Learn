package implementQueueUsingStacks_test

import (
	"leetcode/stack/implementQueueUsingStacks"
	"testing"
)

func TestQueue(t *testing.T) {
	obj := implementQueueUsingStacks.Constructor()
	obj.Push(1)
	obj.Push(2)
	param_2 := obj.Peek()
	want1 := 1
	if param_2 != want1 {
		t.Fatalf("Peek: got: %d, want: %d", param_2, want1)
	}
	param_3 := obj.Pop()
	want2 := 1
	if param_3 != want2 {
		t.Fatalf("Pop: got %d, want %d", param_3, want2)
	}
	param_4 := obj.Empty()
	want3 := false
	if param_4 != want3 {
		t.Fatalf("Empty: got %v, want: %v", param_4, want3)
	}
}
