/*
24. Swap Nodes in Pairs
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.

链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
*/
package swapnodesinpairs

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

/*
passed but ugly...
*/
func MySwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head // node that point to next node
	result := cur.Next
	next := cur.Next
	for cur != nil && cur.Next != nil {
		nextCur := next.Next
		var nextTop *ListNode
		if nextCur != nil {
			if nextCur.Next != nil {
				nextTop = nextCur.Next
			} else {
				nextTop = nextCur
			}
		}
		cur.Next = nextTop
		next.Next = cur
		cur = nextCur
		if nextCur == nil {
			break
		}
		next = cur.Next
	}
	return result
}

// 使用虚拟头结点进行结点交换
func swapPairs(head *ListNode) *ListNode {
	// 头结点， 用于每次交换node时保存临时的头
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	pHead := &ListNode{
		Next: head,
	}
	for pHead.Next != nil && pHead.Next.Next != nil {
		// 一次循环涉及到3个结点: 辅助头结点pHead、结点a、结点b
		// pHead初始为指向head的node 每次循环结束时 变成pair交换顺序后的第二个结点, 即a结点;
		a := pHead.Next
		b := a.Next
		// pHead.Next = b 解释： 这里要把头结点和交换后的nodes pair中的头node进行绑定；
		// 				第一次循环时没有用处，从第二次循环开始, 每次要把头结点（即a）和后面的pair对连接起来，否则pair之间失去连接
		// b.Next = a 解释: 交换pair对顺序
		// a.Next = b.Next 解释:
		pHead.Next, b.Next, a.Next = b, a, b.Next
		pHead = a
	}
	return result
}
