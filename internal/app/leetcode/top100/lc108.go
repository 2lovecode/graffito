package top100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}
	mid := (left + right + 1) / 2

	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)
	return root
}
