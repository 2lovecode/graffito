package top100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

	list := [][]int{}
	if root == nil {
		return list
	}

	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {

		p := make([]*TreeNode, 0)
		tmp := make([]int, 0)
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		list = append(list, tmp)
		queue = p
	}
	return list
}
