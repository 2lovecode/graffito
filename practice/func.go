package main

import "fmt"

func main() {
	a := 2
	b := []int{9, 4, 1, 3, 10}
	fmt.Println(min(a, b...))
}

func trace(s string) string {
	fmt.Println("enter:", s)

	return s
}

func unTrace(s string) {
	fmt.Println("esc:", s)
}

func min(a int, b ...int) (min int) {
	defer unTrace(trace("min"))
	min = a
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	return
}
