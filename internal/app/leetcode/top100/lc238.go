package top100

func productExceptSelf(nums []int) []int {
	answer := make([]int, 0, len(nums))
	if len(nums) <= 0 {
		return answer
	}
	prefixMulMap := make(map[int]int)
	suffixMulMap := make(map[int]int)

	prefixMulMap[0] = nums[0]
	suffixMulMap[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		prefixMulMap[i] = prefixMulMap[i-1] * nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		suffixMulMap[i] = suffixMulMap[i+1] * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		an := 1
		if i > 0 {
			an *= prefixMulMap[i-1]
		}
		if i < len(nums)-1 {
			an *= suffixMulMap[i+1]
		}
		answer = append(answer, an)
	}

	return answer
}
