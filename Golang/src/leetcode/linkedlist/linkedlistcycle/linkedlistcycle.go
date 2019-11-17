package linkedlistcycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := visited[head]; ok {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}

// slow fast pointer solution
func hasCyclySF(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head      // one step once
	fast := head.Next // two steps once
	for slow != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
