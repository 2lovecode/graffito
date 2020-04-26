package algorithm

import "fmt"

type HeapInterface interface {
	Push(value interface{})
	Pop() interface{}
}

// 数组实现堆
type IntHeap []int

func  (heap IntHeap) Push(value int) {
	a := make(map[string]string, 1)
	fmt.Println(a)
}
g
func (heap IntHeap) Pop() (value int) {
	return 0
}
