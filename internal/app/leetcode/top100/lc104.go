package top100

// 1. 深度优先搜索
// 2. 广度优先搜索 todo

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	maxD := 0

	if root != nil {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if leftDepth >= rightDepth {
			maxD = leftDepth + 1
		} else {
			maxD = rightDepth + 1
		}

	}
	return maxD
}
