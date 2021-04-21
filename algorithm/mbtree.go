package algorithm

type mBTreeNode struct {
	items []int
	children []*mBTreeNode
}

func newMBTreeNode() *mBTreeNode {
	return &mBTreeNode{}
}

func (mn *mBTreeNode) insert(value int) {

}



type MBTree struct {
	num int
	root *mBTreeNode
}

func NewMBTree() *MBTree {
	return &MBTree{
		num : 4,
		root: nil,
	}
}


func (mt *MBTree) Insert(value int) {
	if mt.root == nil {

	}
	mt.root.insert(value)
}


func (mt *MBTree) Range(sValue int, eValue int) []int {
	return []int{}
}