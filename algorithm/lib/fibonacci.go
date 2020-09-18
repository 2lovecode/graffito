package lib

// ##### 两种方式
//    1.正常
//    2.使用内存缓存

const FibonacciCnt int = 30

var mem [FibonacciCnt]uint64

func Fibonacci(n int) (num uint64) {
	if n <= 1 {
		num = 1
	} else {
		num = Fibonacci(n-1) + Fibonacci(n-2)
	}
	return
}

func FibonacciUseMem(n int) (num uint64) {
	if mem[n] != 0 {
		num = mem[n]
		return
	}

	if n <= 1 {
		num = 1
	} else {
		num = FibonacciUseMem(n-2) + FibonacciUseMem(n-1)
	}
	mem[n] = num
	return
}
