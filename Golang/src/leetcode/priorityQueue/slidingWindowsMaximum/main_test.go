package slidingWindowsMaximum_test

import (
	"leetcode/priorityQueue/slidingWindowsMaximum"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	var items = []struct {
		in  []int
		k   int
		out []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1, -1}, 1, []int{1, -1}},
	}
	for _, item := range items {
		got := slidingWindowsMaximum.MaxSlidingWindow(item.in, item.k)
		for i, v := range got {
			if v != item.out[i] {
				t.Fatalf("in=%v k=%d got: %v, want: %v", item.in, item.k, got, item.out)
			}
		}
	}
}
