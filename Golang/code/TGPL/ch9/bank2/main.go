package main

var (
	sema    = make(chan struct{}, 1)
	balance int
)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balances() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

func main() {
	Deposit(100)
	println("begin Balances()")
	Balances()
}
