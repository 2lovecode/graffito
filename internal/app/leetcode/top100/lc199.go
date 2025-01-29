package top100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{root}
	list := make([]int, 0)
	for i := 0; len(queue) > 0; i++ {
		p := make([]*TreeNode, 0)

		for j := 0; j < len(queue); j++ {
			node := queue[j]
			if node == nil {
				continue
			}

			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			if j == len(queue)-1 {
				list = append(list, node.Val)
			}
		}
		queue = p
	}
	return list
}
