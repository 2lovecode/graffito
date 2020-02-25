package main

import (
	"fmt"
	"time"
)

const cnt int = 45

var mem [cnt]uint64

func main() {
	start := time.Now()
	for i := 0; i < cnt; i++ {
		fibonacci(i)
	}
	end := time.Now()
	fmt.Println("run time is:", end.Sub(start))

	s := time.Now()
	for i := 0; i < cnt; i++ {
		fibonacciUseMem(i)
	}
	e := time.Now()
	fmt.Println("run time is:", e.Sub(s))
}

func fibonacci(n int) (num uint64) {
	if n <= 1 {
		num = 1
	} else {
		num = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacciUseMem(n int) (num uint64) {
	if mem[n] != 0 {
		num = mem[n]
		return
	}

	if n <= 1 {
		num = 1
	} else {
		num = fibonacciUseMem(n-2) + fibonacciUseMem(n-1)
	}
	mem[n] = num
	return
}
