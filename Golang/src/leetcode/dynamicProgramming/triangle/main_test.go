package triangle

import "testing"

func TestMini(t *testing.T) {
	var data = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	expect := 11
	got := MinimumTotal(data)
	if got != expect {
		t.Fatalf("triangle=%v, expect: %d, got: %d", data, expect, got)
	}
}
