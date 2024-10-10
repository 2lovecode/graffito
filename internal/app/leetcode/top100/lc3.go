package top100

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
