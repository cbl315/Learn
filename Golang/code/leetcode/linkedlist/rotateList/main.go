/*
Given a linked list, rotate the list to the right by k places, where k is non-negative.

链接：https://leetcode-cn.com/problems/rotate-list
*/
package rotateList

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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}
	var p1, p2 *ListNode
	p1, p2 = head, head
	distance := 0
	// 通过间隔为k的双指针获取到最后k+1个结点
	for p1.Next != nil {
		if distance < k {
			distance++
		} else {
			p2 = p2.Next
		}
		p1 = p1.Next
	}
	if distance < k {
		return rotateRight(head, k%(distance+1))
	}
	result := p2.Next
	p2.Next, p1.Next = nil, head
	return result
}
