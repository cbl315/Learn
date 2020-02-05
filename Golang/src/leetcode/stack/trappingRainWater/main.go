/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

链接：https://leetcode-cn.com/problems/trapping-rain-water
*/
package trappingRainWater

func myMinInt(ints ...int) int {
	min := ints[0]
	for i := 1; i < len(ints); i++ {
		if min > ints[i] {
			min = ints[i]
		}
	}
	return min
}

func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	var stack = []int{0}
	water := 0
	for i := 1; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			leftIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop
			if len(stack) == 0 {
				continue
			}
			distance := i - 1 - stack[len(stack)-1]
			currHeight := myMinInt(height[stack[len(stack)-1]], height[i]) - height[leftIndex]
			water += distance * currHeight
		}
		stack = append(stack, i)
	}
	return water
}
