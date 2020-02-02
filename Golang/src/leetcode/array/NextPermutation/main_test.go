package NextPermutation

import "testing"

func TestNextPermutation(t *testing.T) {
	type TestUnit struct {
		Nums     []int
		Arranged []int
	}
	var datas = []TestUnit{
		{[]int{4, 3, 2}, []int{2, 3, 4}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
	}
	for _, data := range datas {
		NextPermutaion(data.Nums)
		for i, num := range data.Arranged {
			if data.Nums[i] != num {
				t.Fatalf("%v, %v", data.Nums, data.Arranged)
			}
		}
	}
}
