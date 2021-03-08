// 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

// 示例 2:

// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

// 示例 3:

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring2(s))
}

// 暴力破解
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	var i, j, z int

	if sLen < 2 {
		return sLen
	}
	maxLen := 0

	for i = 0; i < sLen; i++ {
		currentMaxLen := 0
		for j = i; j < sLen; j++ {
			flag := false
			for z = i; z < j; z++ {
				if s[z] == s[j] {
					flag = true
				}
			}
			if flag {
				break
			} else {
				currentMaxLen++
			}
		}
		if currentMaxLen > maxLen {
			maxLen = currentMaxLen
		}
	}
	return maxLen
}

// 滑动窗口
func lengthOfLongestSubstring2(s string) int {
	sr := []rune(s)
	sLen := len(sr)
	hashMap := make(map[rune]int)
	var maxLen, currentLen int

	for i := 0; i < sLen; i++ {
		if idx, ok := hashMap[sr[i]]; ok {
			hashMap = make(map[rune]int)
			i = idx + 1
			currentLen = 1
		} else {
			currentLen++
		}
		hashMap[sr[i]] = i
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
