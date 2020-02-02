/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

链接：https://leetcode-cn.com/problems/3sum-closest
*/
package threeSumClosest

import "sort"

func myIntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) < 3 {
		return 0
	}
	sort.Ints(nums) // o(n*log(n))
	numsLen := len(nums)
	sum := nums[0] + nums[1] + nums[2]
	for i := numsLen - 1; i >= 2; i-- {
		first := nums[i]
		// double pointer
		p1, p2 := 0, i-1
		for p1 < p2 {
			tempSum := nums[p1] + nums[p2] + first
			if myIntAbs(sum-target) > myIntAbs(tempSum-target) {
				sum = tempSum
			}
			if target > tempSum {
				p1++
			} else if target < tempSum {
				p2--
			} else {
				return target
			}

		}
	}
	return sum
}
