package btree

type node struct {
	items []int
	nodes []*node
}

type BTree struct {
	degree int
	length int
	root   *node
}

func NewBTree() *BTree {
	return &BTree{}
}
