package slice_x

import "fmt"

func Run_1() {

	//实验扩容规则 - 每次添加一个元素

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

func Run_2() {
	// 实验扩容规则 - 一次添加多个元素
	ms := []int{1, 2}

	fmt.Printf("初始容量: %d\n", cap(ms))
	ms = append(ms, 3, 4, 5)
	fmt.Printf("扩容容量: %d\n", cap(ms))
}

func myAppend(s []int) []int {
	s = append(s, 10)
	return s
}

func Run_3() {
	// 实验作为函数参数传递
	s := []int{1, 2, 3}

	newS := myAppend(s)

	fmt.Println(s, newS)

	base := []int{0, 1, 2, 3, 4, 5, 6}

	o := base[2:4:5]
	newO := myAppend(o)
	fmt.Println(base, o, newO)
}
