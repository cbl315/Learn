package main

import "fmt"

func testAppendInt() {
	var x, y []int
	for i:= 0;i<10;i++{
		y = AppendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func testNonempty() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", Nonempty(data))
	fmt.Printf("%q\n", data)
}


func testRemove() {
	s := []int{5,6,7,8,9}
	fmt.Println(remove(s, 2))
}

func testRotate() {
	s := []int{1,2,3,4}
	fmt.Println(rotate(s))
}


func testRemoveRep() {
	s := []string{"hello", "hello", "world", "hello"}
	fmt.Printf("Before: %q\n", s)
	fmt.Printf("After: %q\n", removeRepeat(s))
}


func testRemoveRepSpace() {
	s := ([]byte)("test  test    as")
	fmt.Println(s)
	fmt.Println(removeRepeatSpace(s))
}

// `Not Equal` expected, get `Equal`
func testMapKeyExists() {
	a := map[string]int{"A": 0}
	b := map[string]int{"B": 42}
	for k, xv := range a{
		// should be:
		//if yv, ok := b[k]; !ok || yv != xv{
		if xv != b[k]{
			fmt.Println("Not equal")
		} else {
			fmt.Println("Equal")
		}
	}
}

func testCharCount() {
	charcount()
}

func main() {


	// 4.2.1 append func
	//testAppendInt()

	// 4.2.2
	// slice replace
	//testNonempty()
	// slice remove
	//testRemove()
	// practice test rotate
	//testRotate()
	//testRemoveRep()
	//testRemoveRepSpace()

	// 4.3
	//testMapKeyExists()
	//testCharCount()
	wordFreq()
	return
}
