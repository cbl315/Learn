package kthLargestElementInAStream_test

import (
	"container/heap"
	"github.com/hfutcbl/learn/golang/leetcode/priorityQueue/kthLargestElementInAStream"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	type op struct {
		in   int
		want int
	}
	var items = []struct {
		k        int
		initNums []int
		ops      []op
	}{
		{7, []int{-10, 1, 3, 1, 4, 10, 3, 9, 4, 5, 1}, []op{
			{3, 3},
		}},
		{2, []int{0}, []op{
			{-3, -3},
			{-1, -1},
			{1, 0},
			{-2, 0},
			{-4, 0},
			{3, 1},
		}},
		{k: 3, initNums: []int{4, 5, 8, 2}, ops: []op{
			{3, 4},
			{5, 5},
			{10, 5},
			{9, 8},
			{4, 8},
		},
		},
		{1, []int{}, []op{
			{-3, -3},
			{-2, -2},
			{-4, -2},
			{0, 0},
			{4, 4},
		}},
	}
	for _, item := range items {
		obj := kthLargestElementInAStream.Constructor(item.k, item.initNums)
		for _, one := range item.ops {
			got := obj.Add(one.in)
			if got != one.want {
				t.Fatalf("k= %d, initNums=%v, in=%d: got %d, want %d", item.k, item.initNums, one.in, got, one.want)
			}
		}

	}
}

func TestIntHeap_Push(t *testing.T) {
	obj := kthLargestElementInAStream.IntHeap{}
	heap.Init(&obj)
	obj.Push(0)
	t.Logf("min: %d", obj[0])
	obj.Push(-1)
	t.Logf("min: %d", obj[0])
	minVal := obj.Pop()
	t.Logf("min: %d", minVal)
}
