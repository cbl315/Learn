/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

链接：https://leetcode-cn.com/problems/add-two-numbers
*/
package AddTwoNumbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return &ListNode{}
	}
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	carry := 0
	var currNode *ListNode
	for l1 != nil || l2 != nil || carry > 0 {
		bitResult := carry
		if l1 != nil {
			bitResult += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			bitResult += l2.Val
			l2 = l2.Next
		}
		bitResult, carry = bitResult%10, bitResult/10
		if currNode == nil {
			currNode = res
		} else {
			currNode.Next = &ListNode{}
			currNode = currNode.Next
		}
		currNode.Val = bitResult
	}
	return res
}
