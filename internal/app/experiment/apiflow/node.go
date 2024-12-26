package apiflow

import "context"

type Node struct {
	name    string
	handler Handler
}

func NewNode(name string) *Node {
	return &Node{
		name: name,
	}
}

func (n *Node) WithHandler(h Handler) *Node {
	n.handler = h
	return n
}

func (n *Node) WithUpstreams(ns ...*Node) *Node {

	return n
}

func (n *Node) Execute(ctx context.Context) {
	n.handler(ctx)
}
