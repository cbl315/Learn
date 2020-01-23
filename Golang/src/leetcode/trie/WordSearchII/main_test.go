package WordSearchII

import (
	"fmt"
	"testing"
)

func TestFindWords(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	type data struct {
		Board    [][]byte
		Words    []string
		Expected []string
	}
	var dataSet = []data{
		{Board: board, Words: words, Expected: []string{"eat", "oath"}},
	}
	for _, one := range dataSet {
		got := FindWords(one.Board, one.Words)
		fmt.Printf("%v", got)
	}
}

func BenchmarkFindWords(b *testing.B) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	type data struct {
		Board    [][]byte
		Words    []string
		Expected []string
	}
	var dataSet = []data{
		{Board: board, Words: words, Expected: []string{"eat", "oath"}},
	}

	for _, one := range dataSet {
		for i := 0; i < b.N; i++ {
			_ = FindWords(one.Board, one.Words)
			//fmt.Printf("%v", got)
		}
	}
}
