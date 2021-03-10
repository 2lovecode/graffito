package list_x

import (
	"container/list"
	"fmt"
)

func Run1() {
	l := list.New()

	l.PushFront(2)
	l.PushFront("a")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
