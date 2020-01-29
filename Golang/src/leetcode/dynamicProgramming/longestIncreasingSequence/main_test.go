package longestIncreasingSequence

import "testing"

func TestLIS(t *testing.T) {
	type data struct {
		Nums []int
		LIS  int
	}
	var dataSet = []data{
		{[]int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}, 6},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, unit := range dataSet {
		got := LIS2(unit.Nums)
		if got != unit.LIS {
			t.Fatalf("Nums: %v Expect LIS: %d, Actually Got: %d", unit.Nums, unit.LIS, got)
		}
	}
}
