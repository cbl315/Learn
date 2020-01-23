package sqrt

import "testing"

func TestSqrt(t *testing.T) {
	type TestUnit struct {
		In     int
		Expect int
	}
	var testSet = []TestUnit{
		{In: 8, Expect: 2},
	}
	for _, unit := range testSet {
		got := Sqrt(unit.In)
		if got != unit.Expect {
			t.Fatalf("Input: %d, Expect: %d, Got: %d", unit.In, unit.Expect, got)
		}
	}
	t.Log("finish")
}

func TestSqrt2(t *testing.T) {
	type TestUnit struct {
		In     float64
		Expect float64
	}
	var testSet = []TestUnit{
		{In: 7, Expect: 2.645751},
	}
	for _, unit := range testSet {
		got := Sqrt2(unit.In)
		if got != unit.Expect {
			t.Fatalf("Input: %f, Expect: %f, Got: %f", unit.In, unit.Expect, got)
		}
	}
	t.Log("finish")
}
