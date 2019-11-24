/*
Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
*/
package kthLargestElementInAStream

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type KthLargest struct {
	myHeap *IntHeap
	k      int
}

func Constructor(k int, nums []int) KthLargest {
	//initLen := math.Min(float64(k), float64(len(nums)))
	var myIntHeap IntHeap
	if len(nums) == 0 {

	} else {
		initCap := math.Max(float64(k), float64(len(nums)))
		myIntHeap = make(IntHeap, len(nums), int(initCap))
	}

	if len(nums) > 0 {
		copy(myIntHeap, nums)
	}
	heap.Init(&myIntHeap)
	for myIntHeap.Len() > k {
		heap.Pop(&myIntHeap)
	}
	return KthLargest{myHeap: &myIntHeap, k: k}
}

func (this *KthLargest) Add(val int) int {
	if this.myHeap.Len() < this.k {
		heap.Push(this.myHeap, val)
	} else {
		currMinVal := (*this.myHeap)[0]
		if currMinVal >= val {
			return currMinVal
		} else {
			heap.Pop(this.myHeap)
			heap.Push(this.myHeap, val)
		}
	}
	currMinVal := (*this.myHeap)[0]
	return currMinVal
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
