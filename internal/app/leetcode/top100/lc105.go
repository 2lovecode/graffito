package top100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) <= 0 {
		return nil
	}
	rootVal := preorder[0]

	root := &TreeNode{Val: rootVal, Left: nil, Right: nil}

	idx := -1
	for j := 0; j < len(inorder); j++ {
		if inorder[j] == rootVal {
			idx = j
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:idx])+1], inorder[:idx])
	root.Right = buildTree(preorder[len(inorder[:idx])+1:], inorder[idx+1:])
	return root
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}

	return nil
}
