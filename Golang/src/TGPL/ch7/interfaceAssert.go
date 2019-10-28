package main

import "io"

func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); /* 类型断言 */ ok {
		return sw.WriteString(s) // avoid a copy of memory
	}
	return w.Write(([]byte)(s))
}
