package top100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 三种解法
// 1. 递归
// 2. 迭代 todo
// 3. Morris 遍历 空间复杂度为 O(1) todo

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root != nil {
		vals = append(vals, inorderTraversal(root.Left)...)
		vals = append(vals, root.Val)
		vals = append(vals, inorderTraversal(root.Right)...)
	}

	return vals
}
