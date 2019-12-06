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

func untrace(s string) {
	fmt.Println("esc:", s)
}

func min(a int, b ...int) (min int) {
	defer untrace(trace("min"))
	min = a
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	return
}
