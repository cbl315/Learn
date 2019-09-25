package main

import (
	"fmt"
	"io"
)

type HtmlParse struct {
	Raw []byte
}

func (h *HtmlParse) Read(b []byte) (n int, err error) {
	return len(b), nil
}

func NewReader(s string) *HtmlParse {
	raw := ([]byte)(s)
	h := &HtmlParse{Raw: raw}
	return h
}

func testNewReader(s string) {
	r := NewReader(s)
	fmt.Printf("%v", r)
}

// Copyed from built-in package
type LimitedReader struct {
	Data []byte
	N    int64
	R    io.Reader
}
