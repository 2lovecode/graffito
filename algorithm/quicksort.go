package main

import (
	"fmt"
	"graffito/algorithm/lib"
)

func main() {
	s := []int{4, 5, 6, 3, 7, 9}
	fmt.Println(s)
	lib.QuickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
