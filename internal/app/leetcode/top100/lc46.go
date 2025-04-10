package top100

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		sub := make([]int, 0, len(nums))
		sub = append(sub, nums[:i]...)
		sub = append(sub, nums[i+1:]...)
		sr := permute(sub)
		for _, v := range sr {
			result = append(result, append([]int{nums[i]}, v...))
		}
	}
	return result
}
