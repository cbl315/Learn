/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

链接：https://leetcode-cn.com/problems/majority-element
*/
package majorityElement

import "sort"

func majorityElement(nums []int) int {
	cnt := make(map[int]int, len(nums))
	for _, num := range nums {
		cnt[num] += 1
	}
	result, currMaxCnt := 0, 0
	for num, currCnt := range cnt {
		if currCnt > currMaxCnt {
			result = num
			currMaxCnt = currCnt
		}
	}
	return result
}

func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
