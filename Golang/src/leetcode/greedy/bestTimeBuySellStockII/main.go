/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
*/
package bestTimeBuySellStockII

func maxProfit(prices []int) int {
	totalDays := len(prices)
	profit := 0
	for i := 0; i < totalDays-1; i++ {
		if prices[i] < prices[i+1] {
			// buy this day and sell tomorrow
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
