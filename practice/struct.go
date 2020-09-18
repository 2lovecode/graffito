package main

import (
	"fmt"
	"math/big"
	"reflect"
	"runtime"
)

type person struct {
	Name string	`json:"name"`
	Age int 	`json:"age"`
}

type a1 struct {
	in1 int
	str1 string
}

type a2 struct {
	in2 int
	a1
	int
}

func typeTag(p person, idx int) {
	t := reflect.TypeOf(p)
	tt := t.Field(idx)
	fmt.Println(tt.Tag)
}

func main() {
	im := big.NewInt(2)
	in := im
	io := big.NewInt(1)
	ip := big.NewInt(5)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)

	person := new(person)
	person.Name = "I"
	person.Age = 10

	fmt.Printf("My name is %s, I am %d years old\n", person.Name, person.Age)
	for i := 0; i < 2; i++ {
		typeTag(*person, i)
	}

	aa2 := a2{1, a1{2, "aa1"}, 2}

	fmt.Printf("%d -- %d -- %s -- %d\n", aa2.in2, aa2.in1, aa2.str1, aa2.int)

	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)
	fmt.Printf("memory %d KB has been used!\n", mem.Alloc / 1024)

}