/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

链接：https://leetcode-cn.com/problems/coin-change
*/
package coinChange

func myMin(ints ...int) int {
	currMin := ints[0]
	for _, currInt := range ints {
		if currInt < currMin {
			currMin = currInt
		}
	}
	return currMin
}

func coinChange(coins []int, amount int) int {
	if coins == nil || len(coins) == 0 || amount <= 0 {
		return 0
	}
	/*
		dp solution
			dp[i] 代表凑到i块钱时最少的硬币数量
			转移方程 dp[i] = min(dp[i-coin[j]]), 0<j<len(coins)
	*/
	var minCoinNum = map[int]int{}
	// init
	for i := 1; i <= amount; i++ {
		minCoinNum[i] = amount + 1
	}
	minCoinNum[0] = 0
	// calc
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				minCoinNum[i] = myMin(minCoinNum[i], minCoinNum[i-coin]+1)
			}
		}
	}
	if min, _ := minCoinNum[amount]; min <= amount {
		return min
	} else {
		return -1
	}
}
