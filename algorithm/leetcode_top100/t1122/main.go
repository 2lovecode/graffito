// 1122. 数组的相对排序
// 给你两个数组，arr1 和 arr2，

// arr2 中的元素各不相同
// arr2 中的每个元素都出现在 arr1 中
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

// 示例：

// 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// 输出：[2,2,2,1,4,3,3,9,6,7,19]

// 提示：

// arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// arr2 中的元素 arr2[i] 各不相同
// arr2 中的每个元素 arr2[i] 都出现在 arr1 中

package main

import (
	"fmt"
)

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	a2Len := len(arr2)
	a1Len := len(arr1)
	a2Map := make(map[int]int, a2Len)
	t := 0

	for k, v := range arr2 {
		a2Map[v] = k
	}

	for i := 0; i < a1Len; i++ {
		for j := 1; j < a1Len-i; j++ {
			v1, ok1 := a2Map[arr1[j-1]]
			v2, ok2 := a2Map[arr1[j]]
			if ok1 && ok2 {
				if v1 > v2 {
					t = arr1[j-1]
					arr1[j-1] = arr1[j]
					arr1[j] = t
				}
			} else if !ok1 && ok2 {
				t = arr1[j-1]
				arr1[j-1] = arr1[j]
				arr1[j] = t
			} else if !ok1 && !ok2 {
				if arr1[j-1] > arr1[j] {
					t = arr1[j-1]
					arr1[j-1] = arr1[j]
					arr1[j] = t
				}
			}
		}
	}

	return arr1
}
