// 34. 在排序数组中查找元素的第一个和最后一个位置

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 进阶：

//     你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

// 示例 1：

// 输入：nums = [5,7,7,8,8,10], target = 8
// 输出：[3,4]

// 示例 2：

// 输入：nums = [5,7,7,8,8,10], target = 6
// 输出：[-1,-1]

// 示例 3：

// 输入：nums = [], target = 0
// 输出：[-1,-1]

// 提示：

//     0 <= nums.length <= 105
//     -109 <= nums[i] <= 109
//     nums 是一个非递减数组
//     -109 <= target <= 109

package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 7))
}

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	nL := len(nums)
	if nL > 0 {
		idx := search(nums, 0, nL-1, target)
		if idx != -1 {
			start := idx
			end := idx

			for i := idx; i >= 0; i-- {
				if nums[i] == target {
					start = i
				} else {
					break
				}
			}
			for i := idx; i < nL; i++ {
				if nums[i] == target {
					end = i
				} else {
					break
				}
			}
			res = []int{start, end}
		}
	}
	return res
}

func search(nums []int, left int, right int, target int) int {
	rIdx := -1
	if right < left {
		return rIdx
	}
	idx := (left + right) / 2
	if nums[idx] == target {
		rIdx = idx
	} else if nums[idx] > target {
		rIdx = search(nums, left, idx-1, target)
	} else {
		rIdx = search(nums, idx+1, right, target)
	}
	return rIdx
}
