package LRUCache

import "testing"

func TestLRUCache(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	v1 := c.Get(1)
	if v1 != 1 {
		t.Fatalf("1")
	}
	c.Put(3, 3)
	v2 := c.Get(2)
	if v2 != -1 {
		t.Fatalf("2")
	}
	c.Put(4, 4)
	v3 := c.Get(1)
	if v3 != -1 {
		t.Fatalf("3")
	}

}
