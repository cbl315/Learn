package trappingRainWater

import "testing"

func TestTrap(t *testing.T) {
	type TestUnit struct {
		Height []int
		Water  int
	}
	var datas = []TestUnit{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}
	for _, data := range datas {
		got := Trap(data.Height)
		if got != data.Water {
			t.Fatalf("Height=%v, water=%d got=%d", data.Height, data.Water, got)
		}
	}
}
