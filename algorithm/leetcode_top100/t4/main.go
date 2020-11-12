// 4. 寻找两个正序数组的中位数

// 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

// 示例 1：

// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2

// 示例 2：

// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

// 示例 3：

// 输入：nums1 = [0,0], nums2 = [0,0]
// 输出：0.00000

// 示例 4：

// 输入：nums1 = [], nums2 = [1]
// 输出：1.00000

// 示例 5：

// 输入：nums1 = [2], nums2 = []
// 输出：2.00000

// 提示：

//     nums1.length == m
//     nums2.length == n
//     0 <= m <= 1000
//     0 <= n <= 1000
//     1 <= m + n <= 2000
//     -106 <= nums1[i], nums2[i] <= 106
package main

import "fmt"

func main() {
	nums1 := []int{3, 6, 9, 10, 20}
	nums2 := []int{5, 8, 11, 12}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// nums1Len := len(nums1)
	// nums2Len := len(nums2)
	// s1 := 0
	// s2 := 0

	// m := 0
	// n := 0

	// for i := 0; i < (nums1Len+nums2Len)/2; i++ {
	// 	if m < nums1Len && n < nums2Len {
	// 		s1 = nums1[m]
	// 		s2 = nums2[n]

	// 	}
	// }
	return 0
}
