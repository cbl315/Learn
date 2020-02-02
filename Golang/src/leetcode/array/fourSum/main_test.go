package fourSum

import "testing"

func TestFourSum(t *testing.T) {
	type Unit struct {
		Nums   []int
		Target int
	}
	var testDatas = []Unit{
		{[]int{1, 0, -1, 0, -2, 2}, 0},
	}
	for _, data := range testDatas {
		got := FourSum(data.Nums, data.Target)
		t.Logf("nums: %v, target: %d, got: %v", data.Nums, data.Target, got)
	}
}
