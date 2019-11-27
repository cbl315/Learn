/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

链接：https://leetcode-cn.com/problems/3sum
*/
package threeSum

import "sort"

type Cmp struct {
	data []int
}

func threeSum(nums []int) [][]int {
	targetSum := 0
	var third = make(map[int]bool)
	numsLen := len(nums)
	var result [][3]int
	var result2 [][]int
	sort.Ints(nums)
	for i := 0; i < numsLen-2; i++ {
		third = map[int]bool{} // clear third
		for j := i + 1; j < numsLen; j++ {
			thirdValue := targetSum - nums[i] - nums[j]
			if _, ok := third[nums[j]]; ok {
				temp := [3]int{nums[i], nums[j], thirdValue}
				exist := false
				for _, one := range result {
					if one == temp {
						exist = true
						break
					}
				}
				if exist == false {
					result = append(result, temp)
					result2 = append(result2, temp[:])
				}
			} else {
				third[thirdValue] = true
			}
		}
	}
	return result2
}

/*
双指针法求解
*/
func threeSum2(nums []int) [][]int {
	targetSum := 0
	numsLen := len(nums)
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < numsLen-2; i++ {
		if nums[i] > targetSum {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := numsLen - 1
		for l < r {
			temp := nums[i] + nums[l] + nums[r]
			if temp > targetSum {
				r--
			} else if temp < targetSum {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}
