package numOf1bit

import (
	"testing"
)

func TestNumberOf1Bit(t *testing.T) {
	type data struct {
		Value  uint32
		Number int
	}
	dataSet := []data{
		//{Value: 8, Number: 1},
		//{11, 3},
		{1<<32 - 1 - 2, 31},
	}
	for _, one := range dataSet {
		got := HammingWeight2(one.Value)
		if got != one.Number {
			t.Fatalf("Value: %d, Expect: %d, Got: %d", one.Value, one.Number, got)
		}
	}
}
