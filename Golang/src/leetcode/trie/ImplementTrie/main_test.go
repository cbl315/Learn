package ImplementTrie

import "testing"

func TestTrie(t *testing.T) {
	obj := Constructor()
	word := "app"
	word2 := "apps"
	obj.Insert(word)
	obj.Insert(word2)
	exist := obj.Search(word)
	if exist != true {
		t.Fatalf("word %s not exist which should", word)
	}
	exist2 := obj.Search(word2)
	if exist2 != true {
		t.Fatalf("word %s not exist which should", word2)
	}
}
