package link

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	o := new(Node[T])
	o.value = v
	return o
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) SetNext(next *Node[T]) *Node[T] {
	n.next = next
	return n
}

func (n *Node[T]) Print() {
	fmt.Print(n.value)
}

type SingleLink[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func NewSingleLink[T any]() *SingleLink[T] {
	o := new(SingleLink[T])
	return o
}

func (sl *SingleLink[T]) Append(n *Node[T]) {
	if sl.head == nil {
		sl.head = n
		sl.tail = n
	} else {
		sl.tail.next = n
		sl.tail = n
	}
}

func (sl *SingleLink[T]) Merge(al *SingleLink[T]) {
	sl.tail.next = al.head
	sl.tail = al.tail
}

func (sl *SingleLink[T]) Head() *Node[T] {
	return sl.head
}

func (sl *SingleLink[T]) Print() {
	now := sl.head
	for now != nil {
		now.Print()
		fmt.Print(",")
		now = now.next
	}
}