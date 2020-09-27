package main

import "fmt"

func main() {
	naturals := make(chan int)
	squres := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squres <- x * x
		}
	}()

	for {
		fmt.Println(<-squres)
	}
}
