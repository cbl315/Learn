/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?

链接：https://leetcode-cn.com/problems/reverse-linked-list
*/
package reverselinkedlist

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

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur, prev := head, nil
	for cur != nil {
		next := cur.Next
		cur.Next, prev, cur = prev, cur, next
	}
	return prev
}
