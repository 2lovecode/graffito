// 283. 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 说明:

// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
package leetcode

//import "fmt"
//
//func main() {
//	nums := []int{0, 1, 0, 3, 12}
//	moveZeroes2(nums)
//	fmt.Println(nums)
//}

func moveZeroes(nums []int) {
	nLen := len(nums)

	for i := 0; i < nLen; i++ {
		for j := 0; j < nLen-i-1; j++ {
			if nums[j] == 0 {
				t := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = t
			}
		}
	}
}

func moveZeroes1(nums []int) {
	nLen := len(nums)
	zeroIndex := -1

	for i := 0; i < nLen; i++ {
		if nums[i] != 0 {
			if zeroIndex != -1 {
				nums[zeroIndex] = nums[i]
				nums[i] = 0
				
				for j := zeroIndex+1; j <= i; j++ {
					if nums[j] == 0 {
						zeroIndex = j
						break
					}
				}
			}
		} else {
			if zeroIndex == -1 {
				zeroIndex = i
			}
		}
	}
}

func moveZeroes2(nums []int) {
	nLen := len(nums)
	zeroIndex := -1

	for i := 0; i < nLen; i++ {
		if nums[i] != 0 {
			if zeroIndex != -1 {
				nums[zeroIndex] = nums[i]
				nums[i] = 0
				zeroIndex++
			}
		} else if zeroIndex == -1{
			zeroIndex = i
		}
	}
}
