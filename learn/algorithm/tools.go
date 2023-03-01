package algorithm

import (
	"fmt"
	"graffito/algorithm/link"
)

func GenSingleLinkedList[T comparable](l []T) *link.Node[T] {
	var head *link.Node[T]
	var current *link.Node[T]

	for _, each := range l {
		newNode := link.NewNode(each)
		if head == nil {
			head = newNode
		} else {
			current.SetNext(newNode)
		}
		current = newNode
	}

	return head
}

func LinkPrint(now *link.Node[int]) {
	if now != nil {
		fmt.Printf("%s,", now.Value())
		now = now.Next()
		LinkPrint(now)
	}
}

func LinkReverse(head *link.Node[int]) *link.Node[int] {
	var pre *link.Node[int]
	current := head

	for current != nil {
		next := current.Next()
		current.SetNext(pre)

		pre = current
		current = next
	}
	return pre
}
