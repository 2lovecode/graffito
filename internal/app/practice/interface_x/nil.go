package interface_x

import (
	"fmt"
	"unsafe"
)

// 关于接口与nil是否相等的实验
// 知识点
// 接口值的零值是指：动态类型和动态值都为nil
func compareWithNil() {
	type person interface{}
	type man interface{}
	type student struct{}

	var p person
	var m man
	var s *student

	fmt.Println("init p is nil: ", p == nil)
	fmt.Println("init m is nil: ", m == nil)
	fmt.Println("init s is nil: ", s == nil)
	p = s
	fmt.Println("set s to p, p is nil: ", p == nil)
	p = m
	fmt.Println("set m to p, p is nil: ", p == nil)
}

// 打印接口的动态类型和值

type iface struct {
	tab  *itab
	data unsafe.Pointer
}
type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

func printDynamicTypeAndValue() {
	var a interface{} = nil
	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)
	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}
