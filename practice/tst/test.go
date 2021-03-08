package main

import (
	"fmt"
)

func main() {
	a := 4 << (^uintptr(0) >> 63)
	fmt.Println(a)
}
