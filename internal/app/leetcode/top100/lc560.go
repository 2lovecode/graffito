package top100

// 560. 和为K的子数组
func subarraySum(nums []int, k int) int {
	preSumMap := make(map[int][]int)
	sum := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum-k == 0 {
			cnt++
		}
		if _, ok := preSumMap[sum-k]; ok {
			cnt += len(preSumMap[sum-k])
		}
		preSumMap[sum] = append(preSumMap[sum], i)
	}

	return cnt
}
