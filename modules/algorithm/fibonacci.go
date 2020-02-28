package algorithm

// ##### 两种方式
//    1.正常
//    2.使用内存缓存

import (
	"fmt"
	"time"
)

const cnt int = 30

var mem [cnt]uint64

func FibonacciRun() {
	start := time.Now()
	for i := 0; i < cnt; i++ {
		fmt.Print(fibonacci(i), ",")
	}
	end := time.Now()
	fmt.Println("not use mem cache run time is:", end.Sub(start))

	s := time.Now()
	for i := 0; i < cnt; i++ {
		fmt.Print(fibonacciUseMem(i), ",")
	}
	e := time.Now()
	fmt.Println("use mem cache run time is:", e.Sub(s))
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
