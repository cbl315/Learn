package main

import (
	"fmt"
)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balances() int      { return <-balances }

func teller() {
	var balance int
	for {
		fmt.Println("start select...")
		select {
		case amount := <-deposits:
			fmt.Println("deposits")
			balance += amount
		case balances <- balance: /* 同步发送channel 只有执行接收操作时才会执行发送 否则阻塞 */
			fmt.Println("send balances...")
		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit(100)
	println("begin Balances()")
	Balances()
}
