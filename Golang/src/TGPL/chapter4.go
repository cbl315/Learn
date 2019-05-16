package main

import "unicode"

func AppendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[: zlen]
	} else{
		zcap := zlen
		if zcap < len(x) * 2{
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
	return strings[: i]
}


func Nonempty2(strings []string) []string {
	out := strings[: 0]  // zero length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1: ])
	return slice[: len(slice)-1]
}


// practice
func rotate(in []int) []int{
	var out = make([]int, len(in))
	for index, v := range in {
		if index == len(in) - 1 {
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
	return NoRepeat
}


func removeRepeatSpace(in []byte) (out []byte) {
	var last rune
	for _, v := range in {
		rv := rune(v)
		if unicode.IsSpace(rv) && unicode.IsSpace(last) {

		} else{
			out = append(out, v)
		}
		last = rv
	}
	return out
}