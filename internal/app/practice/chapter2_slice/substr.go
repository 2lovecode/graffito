package chapter2_slice

import "fmt"

// 截取实验
func Substr() {
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := original[2:5]
	s2 := s1[2:6:7]

	fmt.Println("before:")
	fmt.Println(len(original), cap(original), original)
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	fmt.Println("after:")
	fmt.Println(len(original), cap(original), original)
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)
}
