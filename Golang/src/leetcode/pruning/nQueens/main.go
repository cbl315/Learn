/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

链接：https://leetcode-cn.com/problems/n-queens
*/
package nQueens

import "strings"

func solveNQueens(n int) [][]string {
	var currResult [][]int
	currResult = DFS(n, []int{}, []int{}, []int{}, currResult)
	return out(currResult, n)
}

// 判断target是否在slice ints中
func isIn(ints []int, target int) bool {
	for _, v := range ints {
		if v == target {
			return true
		}
	}
	return false
}

// 深度优先遍历+剪枝
// 剪枝策略： 通过queens记录被攻击的列 xyDif记录被攻击的左边斜列 xySum记录被攻击的右边斜列
// xyDif 存储y-x xySum记录y+x 记录同一斜列上的常数值
func DFS(n int, queens []int, xyDif []int, xySum []int, currResult [][]int) [][]int {
	row := len(queens)
	if row == n {
		var single = make([]int, len(queens))
		for col, row := range queens {
			single[row] = col
		}
		currResult = append(currResult, single)
		return currResult
	}
	for col := 0; col < n; col++ {
		if okCol := isIn(queens, col); !okCol { // col not attacked
			if okDif := isIn(xyDif, row-col); !okDif { // row - col not attacked
				if okSum := isIn(xySum, row+col); !okSum { // row + col not attacked
					currResult = DFS(n, append(queens, col), append(xyDif, row-col), append(xySum, row+col), currResult)
				}
			}
		}
	}
	return currResult
}

/*
转换输出格式为leetcode要求的格式
*/
func out(in [][]int, n int) [][]string {
	result := [][]string{}
	for _, single := range in {
		one := make([]string, n)
		for rowIndex, col := range single {
			row := make([]string, n)
			for i := 0; i < n; i++ {
				row[i] = "."
			}
			row[col] = "Q"
			one[rowIndex] = strings.Join(row, "")
		}
		result = append(result, one)
	}
	return result
}
