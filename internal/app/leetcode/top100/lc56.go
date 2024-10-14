package top100

import "sort"

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	idx := 0
	for _, v := range intervals {
		if len(merged) == 0 {
			merged = append(merged, v)
		} else {
			if v[0] <= merged[idx][1] {
				merged[idx][1] = max(merged[idx][1], v[1])
			} else {
				merged = append(merged, v)
				idx++
			}
		}
	}
	return merged
}
