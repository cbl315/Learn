/*
interface的零值： 类型和值的部分都是nil

*/
package main

import (
	"bytes"
	"io"
	"os"
)

func interfaceValue() {
	var w io.Writer
	//w.Write(([]byte)("Hello"))  // panic: nil pointer dereference
	w = os.Stdout // equals to io.Writer(os.Stdout)
	w.Write(([]byte)("Hello"))
	w = new(bytes.Buffer)
	w = nil
}

func main() {
	//testByteCounter()
	//testNewReader("test")
	//sleep()
	//interfaceValue()
	//printTracks(tracks)
	//sortTracks()
	//testMultiSortTable()
	testIsPalindrome()
}
