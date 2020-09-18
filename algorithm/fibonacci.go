package main

import (
	"fmt"
	"graffito/algorithm/lib"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < lib.FibonacciCnt; i++ {
		fmt.Print(lib.Fibonacci(i), ",")
	}
	end := time.Now()
	fmt.Println("not use mem cache run time is:", end.Sub(start))

	s := time.Now()
	for i := 0; i < lib.FibonacciCnt; i++ {
		fmt.Print(lib.FibonacciUseMem(i), ",")
	}
	e := time.Now()
	fmt.Println("use mem cache run time is:", e.Sub(s))
}
