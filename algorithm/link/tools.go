package link

import (
	"fmt"
	"math/rand"
)

func GenTwoIntersectLink(a int, b int, c int) (*SingleLink[int], *SingleLink[int]) {
	if a <= 0 {
		panic("错误")
	}
	if b <= 0 {
		panic("错误")
	}
	if c > a || c > b {
		panic("错误")
	}

	// 构造两个链表
	randMax := (a + b) * 10

	s1 := make([]int, a-c)
	s2 := make([]int, b-c)
	s3 := make([]int, c)
	for i := 0; i < (a - c); i++ {
		s1[i] = rand.Intn(randMax)
	}
	for i := 0; i < (b - c); i++ {
		s2[i] = rand.Intn(randMax)
	}
	for i := 0; i < c; i++ {
		s3[i] = rand.Intn(c)
	}

	sl1 := NewSingleLink[int]()
	sl2 := NewSingleLink[int]()
	sl3 := NewSingleLink[int]()

	for _, each := range s1 {
		sl1.Append(NewNode(each))
	}
	for _, each := range s2 {
		sl2.Append(NewNode(each))
	}
	for _, each := range s3 {
		sl3.Append(NewNode(each))
	}

	sl1.Merge(sl3)
	sl2.Merge(sl3)

	sl1.Print()
	fmt.Println("-------")
	sl2.Print()
	return sl1, sl2
}
