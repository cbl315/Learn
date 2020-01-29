package editDistance

import "testing"

func TestEditDistance(t *testing.T) {
	type TestUnit struct {
		Word1           string
		Word2           string
		MinEditDistance int
	}
	var testDatas = []TestUnit{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, unit := range testDatas {
		got := MinDistance(unit.Word1, unit.Word2)
		if got != unit.MinEditDistance {
			t.Fatalf("word1=%s, word2=%s, minEditDistance=%d, got %d", unit.Word1, unit.Word2, unit.MinEditDistance, got)
		}
	}
}
