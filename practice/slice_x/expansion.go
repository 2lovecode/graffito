package slice_x

import "fmt"

// 扩容实验 - 每次添加一个元素
func ExpansionOneByOne() {
	ms := make([]int, 0)

	oldCap := cap(ms)
	newCap := oldCap

	fmt.Printf("初始容量: %d\n", oldCap)

	for i := 0; i < 2048; i++ {
		ms = append(ms, i)
		newCap = cap(ms)
		if oldCap != newCap {
			fmt.Printf("length: %d, old cap: %d, new cap: %d, rate: %f\n", i, oldCap, newCap, float64(newCap)/float64(oldCap))
			oldCap = newCap
		}
	}
}

// 扩容实验 - 一次添加多个元素
func ExpansionMultiple() {
	ms := []int{1, 2}

	fmt.Printf("初始容量: %d\n", cap(ms))
	ms = append(ms, 3, 4, 5)
	fmt.Printf("扩容容量: %d\n", cap(ms))
}
