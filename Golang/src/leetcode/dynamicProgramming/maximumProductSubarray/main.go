/*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

链接：https://leetcode-cn.com/problems/maximum-product-subarray
*/
package maximumProductSubarray

func myMax(arr ...int) int {
	res := arr[0]
	for _, curr := range arr {
		if curr > res {
			res = curr
		}
	}
	return res
}

func myMin(arr ...int) int {
	res := arr[0]
	for _, curr := range arr {
		if curr < res {
			res = curr
		}
	}
	return res
}

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	/*
		if nums[i] > 0:
			dp[i][0] = dp[i-1][0] * nums[i]
			dp[i][1] = dp[i-1][1] * nums[i]
		else
			dp[i][0] = dp[i-1][1] * nums[i]
			dp[i][1] = dp[i-1][0] * nums[i]
	*/
	var dp = make(map[int]map[int]int, 2)
	dp[0], dp[1] = make(map[int]int, 2), make(map[int]int, 2)
	dp[0][0], dp[0][1] = nums[0], nums[0]
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		index, index2 := i%2, (i-1)%2 // sliding window
		num := nums[i]
		tempStage := num * dp[index2][0]
		tempStage2 := num * dp[index2][1]
		dp[index][0] = myMax(num, tempStage, tempStage2)
		dp[index][1] = myMin(num, tempStage, tempStage2)
		res = myMax(res, dp[index][0])
	}
	return res
}
