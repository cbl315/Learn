package bestTimeToBuyAndSellStockIV

import "testing"

func TestFunc(t *testing.T) {
	type input struct {
		Prices     []int
		TransLimit int
		Expect     int
	}
	var dataSet = []input{
		{[]int{1, 2, 4}, 2, 3},
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 2, 6},
		{[]int{2, 4, 1}, 2, 2},
	}
	for _, data := range dataSet {
		got := MaxProfit(data.TransLimit, data.Prices)
		if got != data.Expect {
			t.Fatalf("Prices: %v, got %d but expect %d", data.Prices, got, data.Expect)
		}
	}
}
