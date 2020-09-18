package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 [4]int
	a2 := new([4]int)
	var s1, s2 []int

	var v *[4]int = new([4]int)
	var v1 = make([]int, 4)

	fmt.Println(v)
	fmt.Println(v1)
	s1 = a1[0:3]
	s2 = a1[2:4]
	s2[0] = 4
	fmt.Println(len(a1))
	fmt.Println(a2)
	fmt.Println(s1)
	fmt.Println(s2)

	//
	a3 := [...]int{10, 20, 30, 40, 50}
	for _, item := range a3 {
		item *= 2
	}
	fmt.Println(reflect.TypeOf(a3).String())
	fmt.Println(a3)

	fmt.Println(len(a3[2:2]))
	fmt.Println(len(a3[2:3]))
}
