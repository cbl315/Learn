package numberOfIslands

import "testing"

func TestNumberOfIslands(t *testing.T) {
	type TestUnit struct {
		Grid            [][]byte
		NumberOfIslands int
	}
	var testDatas = []TestUnit{
		{[][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}, 1},
		{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, 3},
	}
	for _, unit := range testDatas {
		got := NumIslands2(unit.Grid)
		if got != unit.NumberOfIslands {
			t.Fatalf("Grid=%v, Expect=%d, Got=%d", unit.Grid, unit.NumberOfIslands, got)
		}
	}
}
