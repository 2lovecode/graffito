package top100

// 11. 盛最多水的容器
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxA := 0
	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > maxA {
			maxA = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxA
}
