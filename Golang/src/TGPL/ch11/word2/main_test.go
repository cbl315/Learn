package word_test

import (
	word "TGPL/ch11/word2"
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.IsPalindrome("A man, a plan, a canel: Panama")
	}
}
