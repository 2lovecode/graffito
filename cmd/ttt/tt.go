package main

import (
	"fmt"
	"unsafe"
)


func main() {
	a := []string{"1", "2", "3"}

	fmt.Println(unsafe.Sizeof(a))

	b := a[4:]
	fmt.Println(b)
}

