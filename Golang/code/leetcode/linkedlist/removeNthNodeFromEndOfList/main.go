/*
Given a linked list, remove the n-th node from the end of list and return its head.

链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
*/
package removeNthNodeFromEndOfList

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return &ListNode{}
	}
	// 遍历一遍得到链表长度n
	node := head
	length := 0
	for node != nil {
		length++
		node = node.Next
	}
	nthFromBegin := length - n
	vNode := &ListNode{
		Val:  0,
		Next: head,
	}
	node = vNode
	for i := 0; i < nthFromBegin; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return vNode.Next
}
