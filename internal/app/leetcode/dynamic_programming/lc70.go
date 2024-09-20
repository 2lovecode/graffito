package dynamic_programming

//70. 爬楼梯
//简单
//相关标签
//相关企业
//提示
//
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//
//
//示例 1：
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//
//示例 2：
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
//
//
//
//提示：
//
//1 <= n <= 45

var climbStairsMap = make(map[int]int)

func climbStairs(n int) int {
	y := 0
	if x, ok := climbStairsMap[n]; ok {
		y = x
	} else {
		switch n {
		case 1:
			y = 1
		case 2:
			y = 2
		default:
			y = climbStairs(n-2) + climbStairs(n-1)
		}
		climbStairsMap[n] = y
	}

	return y
}
