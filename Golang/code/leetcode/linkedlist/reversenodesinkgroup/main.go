/*
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
*/

package reversenodesinkgroup

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	pHead := &ListNode{
		Val:  0,
		Next: head,
	} // virtual head node, for saving next part of head node in exchanging
	cache := []*ListNode{pHead}
	curr := head
	index := 1
	var result *ListNode = head
	for curr != nil {
		cache = append(cache, curr)
		if index == k {
			result = curr
		}
		curr = curr.Next
		if index%k == 0 { // 序号k
			// exchange node's Next
			// firstly, modify virtual node's Next
			nextHead := curr
			pHead.Next = cache[len(cache)-1]
			// secondly, modify nodes in the group
			for i := len(cache) - 1; i > 1; i-- {
				cache[i].Next = cache[i-1]
			}
			// lastly, change virtual head to the current last node in the group
			pHead = cache[1]
			pHead.Next = nextHead
			cache = []*ListNode{pHead}
		}
		index++
	}
	return result
}

// 修改版本 翻转链表时只用头指针进行反转
func reverseKGroupModify(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	pHead := &ListNode{
		Val:  0,
		Next: head,
	} // virtual head node, for saving next part of head node in exchanging
	curr := head
	index := 1
	var result *ListNode = head
	for curr != nil {
		if index == k {
			result = curr
		}
		curr = curr.Next
		if index%k == 0 { // 序号k
			// exchange node's Next
			// firstly, modify virtual node's Next
			modifiedLastNode := pHead.Next
			// secondly, modify nodes in the group
			formerNode, nextNode := pHead.Next, pHead.Next.Next
			for i := 0; i < k-1; i++ {
				formerNode, nextNode, nextNode.Next = nextNode, nextNode.Next, formerNode
			}
			pHead.Next = formerNode
			// lastly, change virtual head to the current last node in the group
			pHead = modifiedLastNode
			pHead.Next = curr
		}
		index++
	}
	return result
}
