package main

import "fmt"

func main() {
	naturals := make(chan int)
	squres := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squres <- x * x
		}
		close(squres)
	}()

	for x := range squres {
		fmt.Println(x)
	}
}
