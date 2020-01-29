/*
Given an unsorted array of integers, find the length of longest increasing subsequence.

链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
*/
package longestIncreasingSequence

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	/*
		dp solution
			dp[i] 表示选择第i个数的最长子序列
			dp[i] = max(dp[1], dp[1], ..., dp[j]) + 1 其中 0<j<i-1 且 dp[j] < dp[i]
	*/
	var LIS = make(map[int]int, len(nums))
	res := 0
	for i, num := range nums {
		currMaxLIS := 1
		for j := 0; j <= i; j++ {
			if nums[j] < num && currMaxLIS <= LIS[j] {
				currMaxLIS = LIS[j] + 1
			}
		}
		LIS[i] = currMaxLIS
		if res < currMaxLIS {
			res = currMaxLIS
		}
	}
	return res
}

func solution2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	/*
		Binary Search
		维护一个LIS数组 用于存储最终得到的LIS的长度;
		遍历nums，并用二分搜索更新LIS数组，更新策略： 如果当前数比LIS中所有的值都大 append到LIS中；否则替换LIS中首个比他小的数
		返回LIS的长度
	*/
	var LIS = []int{}
	for i, num := range nums {
		if i == 0 || num > LIS[len(LIS)-1] {
			LIS = append(LIS, num)
			continue
		}
		start, end := 0, len(LIS)
		for start <= end {
			mid := (start + end) / 2
			if LIS[mid] > num {
				end = mid - 1
			} else if LIS[mid] < num {
				start = mid + 1
			} else {
				start = mid
				break
			}
		}
		LIS[start] = num
	}
	//fmt.Printf("LIS=%v", LIS)
	return len(LIS)
}
