package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func AppendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < len(x)*2 {
			zcap = len(x) * 2
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func Nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func Nonempty2(strings []string) []string {
	out := strings[:0] // zero length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// practice
func rotate(in []int) []int {
	var out = make([]int, len(in))
	for index, v := range in {
		if index == len(in)-1 {
			out[0] = v
		} else {
			out[index+1] = v
		}
	}
	return out
}

func removeRepeat(s []string) []string {
	var last string
	var NoRepeat []string
	for _, s := range s {
		if s != last {
			NoRepeat = append(NoRepeat, s)
		}
		last = s
	}
	copy(s[:len(NoRepeat)], NoRepeat)
	for i := len(NoRepeat); i < len(s); i++ {
		s[i] = ""
	}
	return NoRepeat
}

func removeRepeatSpace(in []byte) (out []byte) {
	var last rune
	for _, v := range in {
		rv := rune(v)
		if unicode.IsSpace(rv) && unicode.IsSpace(last) {

		} else {
			out = append(out, v)
		}
		last = rv
	}
	copy(in[:len(out)], out)
	for i := len(out); i < len(in); i++ {
		in[i] = 0
	}
	return out
}

func charcount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	var stat = map[string]int{} // stat input rune's type count

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		// better solution?
		if unicode.IsLetter(r) {
			stat["Letter"]++
		}
		if unicode.IsNumber(r) {
			stat["Number"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\ncount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Printf("type\tcount\n")
	for c, n := range stat {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

func wordFreq() {
	wordCounts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordCounts[word]++
	}

	// whether err
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("word\tcount\n")
	for w, c := range wordCounts {
		fmt.Printf("%q\t%d\n", w, c)
	}
}
