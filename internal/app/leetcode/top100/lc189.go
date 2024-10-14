package top100

import "slices"

// 189. 旋转数组
func rotate(nums []int, k int) {
	k = k % len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
