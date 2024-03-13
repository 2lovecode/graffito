package qa1

import (
	"context"
	"fmt"
)

type A struct {
}

func NewA() *A {
	return &A{}
}

func (a *A) String() string {
	return "是否对实参产生影响，取决于是否对实参的底层数组进行了修改。第一种情况: 形参和实参指向同一个底层数组，则会受到影响；第二种情况：形参因为扩容等操作，指向新的底层数组，再次更改则不受影响"
}

func (a *A) Run(ctx context.Context) string {

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}

	t(s1, s2)

	fmt.Println("第一种：", s1)
	fmt.Println("第二种：", s2)

	return ""
}

func t(s1 []int, s2 []int) {
	// 第一种情况: 形参和实参指向同一个底层数组，则会受到影响
	s1[1] = 33

	// 第二种情况：形参因为扩容等操作，指向新的底层数组，再次更改则不受影响
	s2 = append(s2, 9, 10, 11, 12, 13, 14, 15)
	s2[1] = 33
}
