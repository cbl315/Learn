package nQueens

import (
	"fmt"
	"testing"
)

func TestNQueens(t *testing.T) {
	n := 4
	fmt.Println(solveNQueens(n))
}
