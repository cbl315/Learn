/*
Say you have an array for which the i-th element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
*/
package bestTimeToBuyAndSellStockIV

func myMax(arr ...int) int {
	res := arr[0]
	for _, curr := range arr {
		if curr > res {
			res = curr
		}
	}
	return res
}

func greddy(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func maxProfit(k int, prices []int) int {
	if prices == nil || len(prices) == 0 || k <= 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return greddy(prices)
	}
	/*
		dp solution
			dp[i][l][j]
			dp代表第i天 交易l次 手头是否有股票时的最大收益
			j=1表示有股票 j=0表示无股票;
			l < k, i < len(prices)

			转移方程：
			dp[i][l][0] = max(dp[i-1][l][0], dp[i-1][l][1] + prices[i])
			dp[i][l][1] = max(dp[i-1][l-1][0] - prices[i], dp[i-1][l][1])
	*/
	var mp = make(map[int]map[int]map[int]int, len(prices)+1)
	// init
	for i := 0; i <= len(prices); i++ {
		mp[i] = make(map[int]map[int]int, k+1)
		for j := 0; j <= k; j++ {
			mp[i][j] = make(map[int]int, 2)
			mp[i][j][1] = -9999
		}
	}
	res := 0
	for i := 1; i <= len(prices); i++ {
		price := prices[i-1]
		for l := 1; l <= k; l++ {
			mp[i][l][0] = myMax(mp[i-1][l][0], mp[i-1][l][1]+price)
			mp[i][l][1] = myMax(mp[i-1][l-1][0]-price, mp[i-1][l][1])
			res = myMax(res, mp[i][l][0])
		}
	}
	return res
}
