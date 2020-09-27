package validAnagram_test

import "testing"
import "github.com/hfutcbl/learn/golang/leetcode/hashtable/validAnagram"

func TestIsAnagram(t *testing.T) {
	var testSet = []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "anagrma", true},
		{"tar", "rat", true},
		{"bill", "kill", false},
		{"a", "ab", false},
	}
	for _, one := range testSet {

		got := validAnagram.IsAnagram(one.s, one.t)
		if got != one.want {
			t.Fatalf("s: %s t: %s got: %v want: %v", one.s, one.t, got, one.want)
		}

	}
}
