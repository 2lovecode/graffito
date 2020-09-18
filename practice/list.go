package practice

import (
	"container/list"
	"fmt"
)

func ListRun() {
	l := list.New()

	l.PushFront(2)
	l.PushFront("a")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
