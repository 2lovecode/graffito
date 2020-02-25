package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where()
	func(a int, b int) { fmt.Println(a + b) }(1, 3)
	where()
	fmt.Println(f())
	where()
	s := []int{4, 5, 6, 3, 7, 9}
	quickSort(s, 0, len(s)-1)
	where()
	fmt.Println(s)
}

func where() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

func f() (ret int) {
	defer func() {
		ret++
	}()

	return 1
}
func quickSort(list []int, left int, right int) {
	start := left
	end := right

	if left >= right {
		return
	}
	flagNum := list[left]

	for left < right {
		for list[right] >= flagNum && right > left {
			right--
		}
		list[left] = list[right]
		for list[left] <= flagNum && right > left {
			left++
		}
		list[right] = list[left]
	}
	list[left] = flagNum

	quickSort(list, start, left-1)
	quickSort(list, right+1, end)
}
