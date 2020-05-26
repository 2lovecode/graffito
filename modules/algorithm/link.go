package algorithm

import (
	"errors"
)

type MLinkNode struct {
	Value MObject
	Next *MLinkNode
}

func NewMLinkNode () *MLinkNode {
	return &MLinkNode{
		Value: MObject{},
		Next:  nil,
	}
}

type MSingleLink struct {
	Head *MLinkNode
	Tail *MLinkNode
	Len int
}

func NewMSingleLink() *MSingleLink {
	head := NewMLinkNode()
	return &MSingleLink{
		Head: head,
		Tail: head,
		Len: 0,
	}
}

func (link *MSingleLink) Find(pos int) (linkNode *MLinkNode, e error) {
	if pos < 0 || pos >= link.Len {
		e = errors.New("exceed length")
	} else {
		nNode := link.Head
		for i := 0; i < pos; i++ {
			nNode = nNode.Next
		}
		linkNode = nNode
	}
	return
}

func (link *MSingleLink) Insert(value MObject, pos int) error {
	if pos > link.Len + 1 || pos < 0 {
		return errors.New("exceed length")
	} else {
		newNode := NewMLinkNode()
		if link.Head.Next == nil {
			link.Head.Next = newNode
			link.Tail = newNode
		} else if pos == link.Len {
			link.Tail.Next = newNode
			link.Tail = newNode
		} else {
			frontNode, _ := link.Find(pos - 1)
			newNode.Next = frontNode.Next
			frontNode.Next = newNode
		}
		link.Len++
	}

	return nil
}

func (link *MSingleLink) Delete(pos int) {

}

func LinkRun() {
	var tmp MInterface
	tmp = 2
	singleLink := NewMSingleLink()
	singleLink.Insert(MObject{Value: tmp}, 0)
}
