/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return the number of distinct solutions to the n-queens puzzle.

链接：https://leetcode-cn.com/problems/n-queens-ii
*/
package NQueensII

func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	cnt := 0
	cnt = dfs(n, cnt, 0, 0, 0, 0)
	return cnt
}

func dfs(n, cnt, row int, col, xySum, xyDif int64) int {
	// 判断是否停止递归
	if row == n {
		cnt++
		return cnt
	}
	// 获取所有未被攻击的空位
	emptyBits := (^(col | xySum | xyDif)) & (1<<n - 1)
	for emptyBits > 0 {
		lastPlace := emptyBits & -emptyBits // 获取最后一位1
		cnt = dfs(n, cnt, row+1, col|lastPlace, (xySum|lastPlace)>>1, (xyDif|lastPlace)<<1)
		emptyBits &= emptyBits - 1 // 末位1置空
	}
	return cnt
}
