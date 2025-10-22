package top100

// 32. 最长有效括号

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号

// 的长度。

// 左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 "(()())"。

// 示例 1：

// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"

// 示例 2：

// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"

// 示例 3：

// 输入：s = ""
// 输出：0

// 提示：

//     0 <= s.length <= 3 * 104
//     s[i] 为 '(' 或 ')'

func longestValidParentheses(s string) int {
	l := len(s)
	if l < 2 {
		return 0
	}
	dp := map[int]int{-2: 0, -1: 0, 0: 0}
	max := 0
	for i := 1; i < l; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				if (i-dp[i-1]-1) >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = 0
				}
			}
		} else {
			dp[i] = 0
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func longestValidParentheses2(s string) int {
	// 2. 栈
	stack := []int{-1}
	ms := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ms = max(ms, i-stack[len(stack)-1])
			}
		}
	}
	return ms
}

func longestValidParentheses3(s string) int {
	// 3. 双指针
	l := len(s)
	if l < 2 {
		return 0
	}
	mx := 0
	left := 0
	right := 0
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			mx = max(mx, 2*left)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := l - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right {
			mx = max(mx, 2*right)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return mx
}
