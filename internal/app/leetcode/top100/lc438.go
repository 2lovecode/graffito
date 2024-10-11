package top100

// 438.找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	pf := [26]int{}

	for _, v := range p {
		pf[v-'a']++
	}

	flag := [26]int{}
	m := make(map[[26]int][]int)
	for i := 0; i < len(s); i++ {
		flag[s[i]-'a']++
		if i >= len(p) {
			flag[s[i-len(p)]-'a']--
		}
		if i >= len(p)-1 {
			m[flag] = append(m[flag], i-len(p)+1)
		}
	}
	res := make([]int, 0)
	if _, ok := m[pf]; ok {
		res = m[pf]
	}

	return res
}
