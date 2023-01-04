package slice_x

import "fmt"

func myAppend(s []int) []int {
	s = append(s, 10)
	return s
}

func changeSlice(s []int) {
	s[0] = 10
}

// 作为函数参数传递实验
// golang中，还书参数传递，只有值传递，没有引用传递
// 如果改变了slice底层数组的数据，会反映到实参slice的底层数据

func FuncParams() {

	s := []int{1, 2, 3}

	newS := myAppend(s)

	fmt.Println(s, newS)

	base := []int{0, 1, 2, 3, 4, 5, 6}

	o := base[2:4:5]
	newO := myAppend(o)
	//newO = myAppend(newO)
	fmt.Println(base, o, newO)

	base1 := []int{0, 1, 2, 3, 4, 5, 6}
	o1 := base1[2:4]
	newO1 := myAppend(o1)
	//newO1 = myAppend(newO1)
	fmt.Println(base1, o1, newO1)

	base3 := []int{0, 1, 2}
	changeSlice(base3)
	fmt.Println(base3)
}
