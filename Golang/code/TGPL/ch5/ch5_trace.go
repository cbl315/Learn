package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ...code
	time.Sleep(3 * time.Second)  // Simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("eneter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}


func double(x int) (result int) {
	// defer 在return之后执行;
	// 在double中定义的lambda 函数可以访问double中的所有变量
	defer func() {fmt.Printf("double(%d) = %d\n", x, result)}()
	return x + x
}
