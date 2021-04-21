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

// 节点分裂
func (mn *mBTreeNode) split(i int) (value int, next *mBTreeNode){
	value = mn.items[i]
	next = newMBTreeNode()
	next.items = append(next.items, mn.items[i+1:]...)
	mn.items = mn.items[:i]
	if len(mn.children) > 0 {
		next.children = append(next.children, mn.children[i+1:]...)
		mn.children = mn.children[:i]
	}
	return
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
		mt.root = newMBTreeNode()
		mt.root.items = append(mt.root.items, value)
	} else {
		if len(mt.root.items) >= (mt.num - 1) {
			// 根节点满了，需要分裂
			newRootValue, rightNode := mt.root.split((mt.num - 1) / 2)
			leftNode := mt.root

			newRootNode := newMBTreeNode()
			newRootNode.items = append(newRootNode.items, newRootValue)
			newRootNode.children = append(newRootNode.children, leftNode, rightNode)
			mt.root = newRootNode
		}
		mt.root.insert(value)
	}
}


func (mt *MBTree) Range(sValue int, eValue int) []int {
	return []int{}
}