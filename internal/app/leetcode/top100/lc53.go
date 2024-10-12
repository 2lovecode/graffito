package top100

// 53. 最大子序和
// 1.贪心
// 2.动态规划
// 3.分治
func maxSubArray(nums []int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i == 0 {
			maxSum = sum
		}
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func maxSubArray2(nums []int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i == 0 {
			maxSum = sum
		}
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}

func maxSubArray3(nums []int) int {
	maxSum := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i == 0 {
			maxSum = sum
		}
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
