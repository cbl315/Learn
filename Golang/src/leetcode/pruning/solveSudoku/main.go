/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

链接：https://leetcode-cn.com/problems/sudoku-solver
*/
package solveSudoku

func solveSudoku(board [][]byte) {
	_, board = DFS(board)
}

func isValid(board [][]byte, i, j int, value byte) bool {
	for index := 0; index < len(board); index++ {
		if board[i][index] == value { // check row
			return false
		}
		if board[index][j] == value { // check col
			return false
		}
		if board[i/3*3+index/3][j/3*3+index%3] == value {
			return false
		}
	}
	return true
}

func DFS(board [][]byte) (bool, [][]byte) {
	for rowIndex, row := range board {
		for colIndex, v := range row {
			if v == '.' {
				for i := 1; i <= 9; i++ {
					value := byte(i + 48) // 0的ascii码是48
					if isValid(board, rowIndex, colIndex, value) {
						board[rowIndex][colIndex] = value
						if ok, tempBoard := DFS(board); ok {
							board = tempBoard
							return true, board
						} else {
							board[rowIndex][colIndex] = '.'
						}
					}
				}
				return false, board
			}
		}
	}
	return true, board
}
