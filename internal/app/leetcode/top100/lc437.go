package top100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return 0
	}
	val := root.Val

	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return res
}

// 前缀和
func pathSum2(root *TreeNode, targetSum int) int {
	prefixSum := map[int]int{0: 1}

	cnt := 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum += node.Val

		cnt += prefixSum[sum-targetSum]

		prefixSum[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		prefixSum[sum]--
	}

	dfs(root, 0)

	return cnt
}
