package top100

import "sort"

// 字母异位词
// eat eta tea tae eta eat

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for k, v := range strs {
		str := make([]rune, 0)
		for _, vv := range v {
			str = append(str, vv)
		}
		sort.SliceStable(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		m[string(str)] = append(m[string(str)], strs[k])
	}

	mm := make([][]string, 0)
	for _, v := range m {
		mm = append(mm, v)
	}
	return mm
}

func groupAnagrams2(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for k, v := range strs {
		cnt := [26]int{}
		for _, vv := range v {
			cnt[vv-'a']++
		}
		m[cnt] = append(m[cnt], strs[k])
	}
	mm := make([][]string, 0)
	for _, v := range m {
		mm = append(mm, v)
	}
	return mm
}
