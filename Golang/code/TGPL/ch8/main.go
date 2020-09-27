package main

func main() {
	ch := make(chan int)
	defer close(ch) // close channel; and still can get data received before closed
	x := 1
	ch <- x  // send x to channel
	x = <-ch // receive data from channel
}
