package removeNthNodeFromEndOfList

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	result := RemoveNthFromEnd(&ListNode{
		Val:  1,
		Next: nil,
	}, 1)
	fmt.Printf("%v", result)
}
