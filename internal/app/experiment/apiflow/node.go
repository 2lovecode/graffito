package apiflow

import "context"

type Node struct {
	handler Handler
}

func NewNode() *Node {
	return &Node{}
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
