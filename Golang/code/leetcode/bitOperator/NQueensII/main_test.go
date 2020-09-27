package NQueensII

import "testing"

func TestTotalQueens(t *testing.T) {
	type data struct {
		Level  int
		Expect int
	}
	dataSet := []data{
		{4, 2},
	}
	for _, one := range dataSet {
		got := TotalNQueens(one.Level)
		if got != one.Expect {
			t.Fatalf("%d Queens got: %d, expect: %d", one.Level, got, one.Expect)
		}
	}
}
