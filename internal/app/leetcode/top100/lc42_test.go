package top100

import (
	"fmt"
	"testing"
)

func Test_Trap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))

	fmt.Println("----")

	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap2([]int{4, 2, 0, 3, 2, 5}))

	fmt.Println("----")

	fmt.Println(trap3([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap3([]int{4, 2, 0, 3, 2, 5}))
}
