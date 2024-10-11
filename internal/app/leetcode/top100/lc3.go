package top100

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	maxSubLen := 0
	i, j := 0, 0
	for ; i < len(s); i++ {
		if k, ok := m[s[i]]; ok && k >= j {
			j = k + 1
		}
		currentLen := i - j + 1
		if currentLen > maxSubLen {
			maxSubLen = currentLen
		}
		m[s[i]] = i
	}

	return maxSubLen
}

func lengthOfLongestSubstring2(s string) int {
	m := make(map[uint8]int)
	maxSubLen := 0
	i, j := 0, 0
	l := len(s)

	for ; i < l; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for j < l && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}

		currentLen := j - i

		if currentLen > maxSubLen {
			maxSubLen = currentLen
		}

	}
	return maxSubLen
}
