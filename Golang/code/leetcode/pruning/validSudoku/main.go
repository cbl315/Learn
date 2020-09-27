/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

链接：https://leetcode-cn.com/problems/valid-sudoku
*/
package validSudoku

func isValidSudoku(board [][]byte) bool {
	return isValid(board)
}

func isInSlice(ints []byte, tgt byte) bool {
	for _, v := range ints {
		if v == tgt {
			return true
		}
	}
	return false
}

func isValid(board [][]byte) bool {
	rowUsed := make([][]byte, len(board))       // 记录第n行出现的row
	colUsed := make([][]byte, len(board))       // 记录第n列出现的col
	blockUsed := make([][][]byte, len(board)/3) // 第一层是 row/3 第二层下标是col/3 第三层是这个block出现的数
	for i := 0; i < len(board)/3; i++ {
		blockUsed[i] = make([][]byte, len(board)/3)
	}
	for rowIndex, row := range board {
		for colIndex, v := range row {
			if v == '.' {
				continue
			}
			blockRowIndex, blockColIndex := rowIndex/3, colIndex/3
			if isInSlice(rowUsed[rowIndex], v) || isInSlice(colUsed[colIndex], v) || isInSlice(blockUsed[blockRowIndex][blockColIndex], v) {
				return false
			}
			rowUsed[rowIndex] = append(rowUsed[rowIndex], v)
			colUsed[colIndex] = append(colUsed[colIndex], v)
			blockUsed[blockRowIndex][blockColIndex] = append(blockUsed[blockRowIndex][blockColIndex], v)
		}
	}
	return true
}
