package main

import (
	"fmt"
	"graffito/algorithm/lib"
)

func main() {
	fmt.Println("最大堆实现：")
	heap := lib.NewIntHeap(5)
	testData := []int{1, 9, 4, 10, 22, 14, 8, 7}
	for _, v := range testData {
		heap.Push(v)
	}

	l := heap.Len
	for i := 0; i < l; i++ {
		fmt.Print(heap.Pop(), ", ")
	}
}
