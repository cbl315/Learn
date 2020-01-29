/*
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character

链接：https://leetcode-cn.com/problems/edit-distance
*/
package editDistance

func myMin(ints ...int) int {
	currMin := ints[0]
	for _, currInt := range ints {
		if currInt < currMin {
			currMin = currInt
		}
	}
	return currMin
}

func minDistance(word1 string, word2 string) int {
	/*
		dp solution
			dp[i][j] 表示从word1的前i个字符变成word2的前j个字符需要的最少操作数
			状态转移方程
			if word[i] == word[j]:
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			else:
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
	*/
	var substrMinDistance = make(map[int]map[int]int, len(word1)+1)
	m, n := len(word1), len(word2)
	// init
	for i := 0; i <= m; i++ {
		substrMinDistance[i] = make(map[int]int, n+1)
		substrMinDistance[i][0] = i
	}
	for j := 0; j <= n; j++ {
		substrMinDistance[0][j] = j
	}
	// dp process
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			distanceBeforeInsert := substrMinDistance[i-1][j] + 1
			distanceBeforeDelete := substrMinDistance[i][j-1] + 1
			var distanceBeforeReplace int
			if word1[i-1] == word2[j-1] {
				distanceBeforeReplace = substrMinDistance[i-1][j-1]
			} else {
				distanceBeforeReplace = substrMinDistance[i-1][j-1] + 1
			}
			substrMinDistance[i][j] = myMin(distanceBeforeDelete, distanceBeforeInsert, distanceBeforeReplace)
		}
	}
	return substrMinDistance[m][n]
}
