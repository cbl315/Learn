package threeSumClosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	type Unit struct {
		Nums    []int
		Target  int
		Closest int
	}
	datas := []Unit{
		{[]int{1, 2, 4, 8, 16, 32, 64, 128}, 82, 82},
		{[]int{0, 2, 1, -3}, 1, 0},
		{[]int{-1, 2, 1, -1}, 1, 2},
	}
	for _, data := range datas {
		got := ThreeSumClosest(data.Nums, data.Target)
		if got != data.Closest {
			t.Fatalf("nums: %v, target: %d, closest: %d, got: %d", data.Nums, data.Target, data.Closest, got)
		}
	}
}
