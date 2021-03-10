package ginx

type node struct {
	path string
	children []*node
	handlers HandlersChain
}

type methodTree struct {
	method string
	root *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method  {
			return tree.root
		}
	}
	return nil
}