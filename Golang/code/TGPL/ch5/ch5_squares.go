package main

import "fmt"

// square 返回一个匿名函数
// 该匿名函数每次被调用都会返回下一个数的平方
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func testSquares() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
