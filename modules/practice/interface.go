package practice

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

type shape interface {
	area() float64
}

type square struct {
	length int
	width int
}

type circle struct {
	r float64
}

func (q *square) area() float64 {
	return float64(q.length * q.width)
}

func (c *circle) area() float64 {
	return math.Pi * math.Sqrt(c.r)
}

type myArray []int

func (ma myArray) Len() int {
	return len(ma)
}

func (ma myArray) Less(i, j int) bool {
	return ma[i] < ma[j]
}

func (ma myArray) Swap(i, j int) {
	ma[i], ma[j] = ma[j], ma[i]
}

type Ele interface {}

type Ll struct {
	a [3]Ele
}

func (l *Ll) Set(i int, e Ele) {
	l.a[i] = e
}

func (l *Ll) At(i int) Ele {
	return l.a[i]
}

func InterfaceRun() {
	var s shape

	q := &square{2, 5}
	c := &circle{3}

	s = q

	if _, ok := s.(*square); ok {
		fmt.Println("is square!")
	} else if _, ok := s.(*circle); ok {
		fmt.Println("is circle!")
	} else {
		fmt.Println("other!")
	}

	switch t := s.(type) {
	case *square:
		fmt.Printf("%T is square\n", t)
	case *circle:
		fmt.Printf("%T is circle\n", t)
	case nil:
		fmt.Println("null")
	default:
		fmt.Println("default")
	}

	fmt.Println(c)

	aaa := myArray{3, 4, 2, 1, 6}
	sort.Sort(aaa)
	fmt.Println(aaa)

	ea := 1
	eb := "start"

	e1 := new(Ll)
	e1.Set(0, ea)
	e1.Set(1, eb)

	fmt.Println(e1.At(0))
	fmt.Println(e1.At(1))

	var empty interface{}
	var el1 Ele

	empty = e1

	el1 = empty.(Ele)

	fmt.Println(el1)

	type myInt int

	var m myInt = 1

	v := reflect.ValueOf(m)

	fmt.Println(v)
	fmt.Println(v.Interface())

}
