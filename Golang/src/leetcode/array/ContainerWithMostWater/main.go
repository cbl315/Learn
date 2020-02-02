/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

链接：https://leetcode-cn.com/problems/container-with-most-water
*/
package ContainerWithMostWater

func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	return doubleP(height)
}

func myMinInt(ints ...int) int {
	min := ints[0]
	for _, v := range ints {
		if min > v {
			min = v
		}
	}
	return min
}

func force(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			currArea := (j - i) * myMinInt(height[i], height[j])
			if currArea > maxArea {
				maxArea = currArea
			}
		}
	}
	return maxArea
}

func doubleP(height []int) int {
	maxArea := 0
	p1, p2 := 0, len(height)-1
	for p1 < p2 {
		currArea := (p2 - p1) * myMinInt(height[p1], height[p2])
		if currArea > maxArea {
			maxArea = currArea
		}
		if height[p1] > height[p2] {
			p2--
		} else {
			p1++
		}
	}
	return maxArea
}
