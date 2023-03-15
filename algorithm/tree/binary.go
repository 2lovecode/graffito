package tree

type BinaryNode[T comparable] struct {
	Data  T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

type Binary[T comparable] struct {
	root *BinaryNode[T]
}

// 二叉树有几种构建方式 - 如何确定一颗二叉树
// 1. 先序序列 + 中序序列
// 2. 后序序列 + 中序序列
// 3. 层序序列 + 中序序列
// 这里选择第一种构建方式

func NewBinary[T comparable](preList []T, midList []T) *Binary[T] {
	root := buildByPreMid(preList, midList)
	return &Binary[T]{
		root: root,
	}
}

func buildByPreMid[T comparable](preList []T, midList []T) *BinaryNode[T] {
	if len(preList) != len(midList) {
		// 长度不一样，有问题
		return nil
	}
	listLen := len(preList)

	// 前序序列和中序序列都为空
	if listLen == 0 {
		return nil
	}

	if listLen == 1 {
		if preList[0] == midList[0] {
			return &BinaryNode[T]{Data: preList[0]}
		} else {
			return nil
		}
	}

	leftCount := 0
	rootData := preList[0]
	findFlag := false
	for _, v := range midList {
		if rootData == v {
			findFlag = true
			break
		} else {
			leftCount++
		}
	}
	if !findFlag {
		return nil
	}

	root := &BinaryNode[T]{
		Data: rootData,
	}
	// 左子树
	if leftCount > 0 {
		root.Left = buildByPreMid(preList[1:leftCount+1], midList[0:leftCount])
	}

	// 右子树
	if (listLen - leftCount) > 0 {
		root.Right = buildByPreMid(preList[leftCount+1:], midList[leftCount+1:])
	}

	return root
}

func (b *Binary[T]) Root() *BinaryNode[T] {
	return b.root
}
