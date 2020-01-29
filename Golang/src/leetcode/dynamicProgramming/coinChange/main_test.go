package coinChange

import "testing"

func TestCoinChange(t *testing.T) {
	type data struct {
		Coins    []int
		Amount   int
		MinCoins int
	}
	var datas = []data{
		{[]int{1, 2, 5}, 11, 3},
	}
	for _, one := range datas {
		got := CoinChange(one.Coins, one.Amount)
		if got != one.MinCoins {
			t.Fatalf("Amount: %d, Coins: %v, Expect %d, Got %d", one.Amount, one.Coins, one.MinCoins, got)
		}
	}
}
