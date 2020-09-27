package validparentheses_test

import (
	"github.com/hfutcbl/learn/golang/leetcode/stack/validparentheses"
	"testing"
)

func TestIsValid(t *testing.T) {
	var datas = []struct {
		In   string
		want bool
	}{
		{"{}", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, data := range datas {
		got := validparentheses.IsValid(data.In)
		if got != data.want {
			t.Fatalf("run %s: got %v, got %v", data.In, got, data.want)
		}
	}
}
