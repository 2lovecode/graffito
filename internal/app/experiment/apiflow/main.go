package apiflow

import "context"

type ApiFlow struct {
	nodes map[string]*Node
	tree  []*Node
}

func NewApiFlow() *ApiFlow {
	return &ApiFlow{
		nodes: make(map[string]*Node),
	}
}

func (af *ApiFlow) Add(n *Node) {
	af.nodes[n.name] = n
}

func (af *ApiFlow) Start(ctx context.Context) {

}
