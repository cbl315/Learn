package pow_test

import "testing"
import "leetcode/recursionAndDivideConquer/pow"

func TestMyPow(t *testing.T) {
	var dataSet = []struct {
		x    float64
		n    int
		want float64
	}{
		{1, 0, 1},
		{0.8, -1, 1.25},
		{2.0000, 10, 1024},
		{2.100, 3, 9.261},
		{2.000, -2, 0.25000},
	}

	for _, data := range dataSet {
		got := pow.MyPow(data.x, data.n)
		if got != data.want {
			t.Fatalf("x=%v, n=%d, got %v != want %v", data.x, data.n, got, data.want)
		}
	}
}
