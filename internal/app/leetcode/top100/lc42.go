package top100

// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 示例 1：

// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

// 示例 2：

// 输入：height = [4,2,0,3,2,5]
// 输出：9

// 提示：

//     n == height.length
//     1 <= n <= 2 * 104
//     0 <= height[i] <= 105

// [0,1,0,2,1,0,1,3,2,1,2,1]

func trap(height []int) int {
	leftMax := 0
	rightMax := 0

	leftMaxMap := make(map[int]int)
	rightMaxMap := make(map[int]int)
	for i := 0; i < len(height); i++ {
		leftMaxMap[i] = leftMax
		if height[i] > leftMax {
			leftMax = height[i]
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		rightMaxMap[i] = rightMax
		if height[i] > rightMax {
			rightMax = height[i]
		}
	}

	sum := 0

	for i := 0; i < len(height); i++ {
		if leftMaxMap[i] == 0 || rightMaxMap[i] == 0 {
			continue
		}
		if height[i] > leftMaxMap[i] || height[i] > rightMaxMap[i] {
			continue
		}
		if leftMaxMap[i] < rightMaxMap[i] {
			sum += leftMaxMap[i] - height[i]
		} else {
			sum += rightMaxMap[i] - height[i]
		}
	}

	return sum
}

func trap2(height []int) int {
	left := 0
	right := len(height) - 1

	lmax := 0
	rmax := 0
	sum := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > lmax {
				lmax = height[left]
			} else {
				sum += lmax - height[left]
			}
			left++
		} else {
			if height[right] > rmax {
				rmax = height[right]
			} else {
				sum += rmax - height[right]

			}
			right--
		}
	}
	return sum
}

func trap3(height []int) int {
	stack := make([]int, 0)

	sum := 0
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bidx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) <= 0 {
				break
			}
			lidx := stack[len(stack)-1]
			h := min(height[lidx], height[i]) - height[bidx]
			// 不是 sum += h 而是 sum += (i - lidx - 1) * h 的原因?
			// 考虑到连续height相等的情况
			sum += (i - lidx - 1) * h
		}
		stack = append(stack, i)
	}
	return sum
}
