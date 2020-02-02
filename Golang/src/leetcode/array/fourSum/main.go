/*
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

链接：https://leetcode-cn.com/problems/4sum
*/
package fourSum

import "sort"

func isUnique(res [][]int, curr []int) bool {
	notDupCnt := 0
	for _, one := range res {
		for i, num := range one {
			if num != curr[i] {
				notDupCnt++
				break
			}
		}
	}
	return notDupCnt == len(res)
}

func fourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	numsLen := len(nums)
	res := [][]int{}
	for i := 0; i < numsLen-3; i++ {
		for j := i + 1; j < numsLen-2; j++ {
			// 剪枝
			currMin := nums[i] + nums[j] + nums[j+1] + nums[j+2]
			if currMin > target {
				continue
			}
			currMax := nums[i] + nums[j] + nums[numsLen-2] + nums[numsLen-1]
			if currMax < target {
				continue
			}

			p1, p2 := j+1, numsLen-1
			twoSumTarget := target - nums[i] - nums[j]
			for p1 < p2 {
				dif := twoSumTarget - nums[p1] - nums[p2]
				if dif > 0 {
					p1++
				} else if dif < 0 {
					p2--
				} else {
					currAns := []int{nums[i], nums[j], nums[p1], nums[p2]}
					if isUnique(res, currAns) {
						res = append(res, currAns)
					}
					p1++
				}
			}
		}
	}
	return res
}
