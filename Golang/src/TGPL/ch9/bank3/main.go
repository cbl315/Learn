package main

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balances() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balances() < 0 {
		deposit(amount)
		return false
	}
	return true
}

func deposit(amount int) {
	balance += amount
}

func main() {
	Deposit(100)
	println("begin Balances()")
	Balances()
}
