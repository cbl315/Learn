package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func testAppendInt() {
	var x, y []int
	for i := 0; i < 10; i++ {
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
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func testRotate() {
	s := []int{1, 2, 3, 4}
	fmt.Println(rotate(s))
}

func testRemoveRep() {
	s := []string{"hello", "hello", "world", "hello"}
	fmt.Printf("Before: %q\n", s)
	removeRepeat(s)
	fmt.Printf("After: %q\n", s)
}

func testRemoveRepSpace() {
	s := ([]byte)("test  test    as")
	fmt.Printf("%q\n", s)
	removeRepeatSpace(s)
	fmt.Printf("%q\n", s)
}

// `Not Equal` expected, get `Equal`
func testMapKeyExists() {
	a := map[string]int{"A": 0}
	b := map[string]int{"B": 42}
	for k, xv := range a {
		// should be:
		//if yv, ok := b[k]; !ok || yv != xv{
		if xv != b[k] {
			fmt.Println("Not equal")
		} else {
			fmt.Println("Equal")
		}
	}
}

func testCharCount() {
	charcount()
}

func changeLocal(num [5]int) {
	num[0] = 55

	fmt.Println("Inside function", num)
}

func testJsonMarshal() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("Json marshaling failed: %v", err)
	}
	fmt.Printf("%s\n", data)

	data2, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatal("Json marshaling failed: %v", err)
	}
	fmt.Printf("%s\n", data2)
}

func testJsonUnmarshal() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("Json marshaling failed: %v", err)
	}
	fmt.Printf("%s\n", data)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatal("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%#q\n", titles)
}

func testGithub() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
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
	// todo decrease len of slice
	//testRemoveRep()
	//testRemoveRepSpace()

	// 4.3
	//testMapKeyExists()
	//testCharCount()
	//wordFreq()
	// sort
	//a := []int{3, 1, 2}
	//fmt.Printf("%v\n", a)
	//Sort(a)
	//fmt.Printf("%v\n", a)
	//testJsonMarshal()
	//testJsonUnmarshal()
	testGithub()
	return
}
