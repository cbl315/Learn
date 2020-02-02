package canMakePalindromeFromSubstring

import "testing"

func TestFunc(t *testing.T) {
	type Unit struct {
		S       string
		Queries [][]int
		Out     []bool
	}
	var datas = []Unit{
		{"abcda", [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}, []bool{true, false, false, true, true}},
	}
	for _, data := range datas {
		got := CanMakePaliQueries(data.S, data.Queries)
		for i, flag := range got {
			if flag != data.Out[i] {
				t.Fatalf("%v, %v", data.Out, got)
			}
		}
	}
}
