/*
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

链接：https://leetcode-cn.com/problems/sliding-window-maximum
*/
package slidingWindowsMaximum

import (
	"container/heap"
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 最大堆实现滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		window := make(IntMaxHeap, k, k)
		copy(window, nums[i:i+k])
		heap.Init(&window)
		currMax := window[0]
		result = append(result, currMax)
	}
	return result
}

// 双向队列实现
func maxSlidingWindowByDeque(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := []int{}
	deq := deque.New()
	for i := 0; i < len(nums); i++ {
		if i >= k && deq.Left().(int) <= i-k { // sliding window
			deq.PopLeft()
		}
		// keep max in the very left of deque
		for deq.Size() > 0 && nums[deq.Left().(int)] <= nums[i] {
			deq.PopLeft()
		}
		deq.PushRight(i)
		if i >= k-1 {
			result = append(result, nums[deq.Left().(int)])
		}
	}
	return result
}
