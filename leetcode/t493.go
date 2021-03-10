// 493. 翻转对
// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。

// 你需要返回给定数组中的重要翻转对的数量。

// 示例 1:

// 输入: [1,3,2,3,1]
// 输出: 2
// 示例 2:

// 输入: [2,4,3,5,1]
// 输出: 3
// 注意:

// 给定数组的长度不会超过50000。
// 输入数组中的所有数字都在32位整数的表示范围内
package leetcode

//import "fmt"

//func main() {
//	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
//}

func reversePairs(nums []int) int {
	l := len(nums)
	cnt := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > (2 * nums[j]) {
				cnt++
			}
		}
	}
	return cnt
}
