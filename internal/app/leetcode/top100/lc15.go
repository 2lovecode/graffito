package top100

import "sort"

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	for k := 0; k < (len(nums) - 2); k++ {
		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, len(nums)-1

		for i < j {
			s := nums[k] + nums[i] + nums[j]

			if s < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if s > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}

		}

	}
	return res
}
