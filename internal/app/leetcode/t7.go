// 7. 整数反转

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

// 示例 1:

// 输入: 123
// 输出: 321

//  示例 2:

// 输入: -123
// 输出: -321

// 示例 3:

// 输入: 120
// 输出: 21

// 注意:

// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
package leetcode

import (
	"math"
)

//func main() {
//	fmt.Println(reverse(9000000))
//}

func reverse(x int) int {

	tmp := 0
	p := 0

	for x != 0 {
		p = x % 10
		x = x / 10

		if tmp > math.MaxInt32/10 || (tmp == math.MaxInt32/10 && p > 7) {
			return 0
		}
		if tmp < math.MinInt32/10 || (tmp == math.MinInt32/10 && p < -8) {
			return 0
		}

		tmp = tmp*10 + p
	}

	return tmp
}
