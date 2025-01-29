package top100

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return check2(root, math.MinInt64, math.MaxInt64)
}

func check2(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}

	return check2(node.Left, lower, node.Val) && check2(node.Right, node.Val, upper)
}
