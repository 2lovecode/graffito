package top100

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	l := 0
	for k, _ := range m {
		if !m[k-1] {
			cN := k
			cL := 1
			for m[cN+1] {
				cN++
				cL++
			}
			if cL > l {
				l = cL
			}
		}
	}
	return l
}
