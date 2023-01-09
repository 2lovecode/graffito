package graph

type Node struct {
	ID   uint
	Next *Node
}

type Graph struct {
	adj []*Node
}

func NewGraph(size uint) {

}
