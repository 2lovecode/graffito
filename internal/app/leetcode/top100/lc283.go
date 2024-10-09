package top100

// 283. 移动零
// 将 0 移动到数组的末尾且保证其他非零元素相对位置不变
func moveZeroes(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
