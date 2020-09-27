/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

链接：https://leetcode-cn.com/problems/next-permutation
*/
package NextPermutation

import "sort"

func nextPermutation(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}
	// 从后面往前遍历 在后面比当前数大的数中找到一个最小的数 互换位置后，将后面的数排序即可; 如果找不到 升序排序
	numsLen := len(nums)
	for i := numsLen - 2; i >= 0; i-- {
		minValueGreaterThanI, index := nums[i], i
		for j := i + 1; j < numsLen; j++ {
			if nums[j] > nums[i] {
				if minValueGreaterThanI == nums[i] {
					minValueGreaterThanI = nums[j]
					index = j
				} else if nums[j] < minValueGreaterThanI {
					minValueGreaterThanI = nums[j]
					index = j
				}
			}
		}
		if index != i {
			nums[index], nums[i] = nums[i], nums[index]
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
	return
}
