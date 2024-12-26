package apiflow

type tree struct {
	heads []*leaf
}

type leaf struct {
	val  *Node
	prev []*leaf
	next []*leaf
}
