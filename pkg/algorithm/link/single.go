package link

import "fmt"

//type SingleLink[T comparable] struct {
//	head *Node[T]
//	tail *Node[T]
//}
//
//func NewSingleLink[T comparable]() *SingleLink[T] {
//	o := new(SingleLink[T])
//	return o
//}
//
//func (sl *SingleLink[T]) Append(n *Node[T]) {
//	if sl.head == nil {
//		sl.head = n
//		sl.tail = n
//	} else {
//		sl.tail.next = n
//		sl.tail = n
//	}
//}
//
//func (sl *SingleLink[T]) Merge(al *SingleLink[T]) {
//	sl.tail.next = al.head
//	sl.tail = al.tail
//}
//
//func (sl *SingleLink[T]) Head() *Node[T] {
//	return sl.head
//}
//
//func (sl *SingleLink[T]) Print() {
//	now := sl.head
//	for now != nil {
//
//		fmt.Printf("%s,", now.Value())
//		now = now.next
//	}
//}

type SingleLinkedList[T comparable] struct {
	head *Node[T]
}

func NewSingleLinkedList[T comparable]() *SingleLinkedList[T] {
	o := new(SingleLinkedList[T])
	return o
}

func (sl *SingleLinkedList[T]) Add(v T) {
	newNode := NewNode(v)
	if sl.head == nil {
		sl.head = newNode
	} else {
		var pre, now *Node[T]
		now = sl.head
		for now != nil {
			pre = now
			now = now.Next()
		}
		pre.SetNext(newNode)
	}
}

func (sl *SingleLinkedList[T]) Del(v T) {
	if !sl.Empty() {
		pre, now := sl.Seek(v)
		if now != nil {
			pre.SetNext(now.Next())
		}
	}
}

func (sl *SingleLinkedList[T]) Seek(v T) (pre *Node[T], now *Node[T]) {
	if !sl.Empty() {
		now = sl.head
		for now != nil && now.Value() != v {
			pre = now
			now = now.Next()
		}
	}
	return
}

func (sl *SingleLinkedList[T]) Empty() bool {
	return sl.head == nil
}

func (sl *SingleLinkedList[T]) Merge(al *SingleLinkedList[T]) {
	if !sl.Empty() {
		var pre, now *Node[T]
		now = sl.head
		for now != nil {
			pre = now
			now = now.Next()
		}
		pre.SetNext(al.Head())
	}
}

func (sl *SingleLinkedList[T]) Head() *Node[T] {
	return sl.head
}

func (sl *SingleLinkedList[T]) Print() {
	now := sl.head
	for now != nil {
		fmt.Printf("%v,", now.Value())
		now = now.next
	}
}
