/*
interface的零值： 类型和值的部分都是nil

*/
package main

import (
	"bytes"
	"fmt"
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

func testAseert() {
	var w io.Writer
	w = os.Stdout

	rw := w.(io.ReadWriter) // rw 具有Read方法， w没有， 都有Write方法

	fmt.Printf("%q", rw)
	w = new(ByteCounter)
	rw = w.(io.ReadWriter)
}

func main() {
	//testByteCounter()
	//testNewReader("test")
	//sleep()
	//interfaceValue()
	//printTracks(tracks)
	//sortTracks()
	//testMultiSortTable()
	//testIsPalindrome()
	//http1()
	//testAseert()
	xmlSelect()
}
