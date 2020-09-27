package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union  of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") { // judge whether first int
					buf.WriteByte(' ')
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	return len(s.words)
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	// single ^ means bitwise negative
	s.words[word] &= ^(1 << bit)
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] &= 0
	}
}

// copy object
func (s *IntSet) Copy() (dup IntSet) {
	newWords := make([]uint64, len(s.words), cap(s.words))
	_ = copy(newWords, s.words)
	dup.words = newWords
	return
}

// Add All Int, e.g. s.AddAll(1, 2, 3)
func (s *IntSet) AddAll(lst ...int) {
	for _, item := range lst {
		s.Add(item)
	}
}

// Intersect Operation
func (s *IntSet) IntersecWith(t *IntSet) {
	setLen := len(s.words)
	for i, tword := range t.words {
		if i < setLen {
			s.words[i] &= tword
		} else {
			break
		}
	}
	return
}

// Difference
func (s *IntSet) DifferenceWith(t *IntSet) {
	setLen := len(s.words)
	for i, tword := range t.words {
		if i < setLen {
			// first, xor operation
			mid := s.words[i] ^ tword
			s.words[i] = mid & s.words[i]
		} else {
			break
		}
	}
}

// SymmetricDifference
func (s *IntSet) SymmetricDifference(t *IntSet) {
	setLen := len(s.words)
	for i, tword := range t.words {
		if i < setLen {
			s.words[i] = s.words[i] ^ tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// return all elements in the set
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return
}

func testIntSet() {
	var x, y IntSet
	var z IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))

	// copy
	z = x.Copy()
	fmt.Printf("---Before Clear---\n x: %s\nz: %s\n", x.String(), z.String())
	x.Clear()
	fmt.Printf("---Clear x---\nx: %s\nz: %s\n", x.String(), z.String())
	z.Clear()
	fmt.Printf("---Clear z---\nx: %s\nz: %s\n", x.String(), z.String())
	x.AddAll(7, 8, 9)
	z.AddAll(23, 123, 53, 9)
	fmt.Printf("---x AddAll---\n x: %s, z: %s\n", x.String(), z.String())
	// Set Operation Test
	x.IntersecWith(&z)
	fmt.Printf("---x Intersect y---\n x: %s, z: %s\n", x.String(), z.String())
	x.AddAll(10, 23, 100)
	x.DifferenceWith(&z)
	fmt.Printf("---x Difference y---\n x: %s, z: %s\n", x.String(), z.String())
	// SymmetricDifference
	z.Add(123456)
	x.SymmetricDifference(&z)
	fmt.Printf("---x Symmetric Difference z---\n x: %s, z: %s\n", x.String(), z.String())
	fmt.Printf("---All Elements---\n x: %v\n", x.Elems())
}
