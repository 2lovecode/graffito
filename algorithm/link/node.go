package link

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func NewNode[T comparable](v T) *Node[T] {
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

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) SetValue(v T) {
	n.value = v
}
