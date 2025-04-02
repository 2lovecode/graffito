package top100

import "testing"

func TestLc994(t *testing.T) {
	type td struct {
		grid [][]int
		ans  int
	}

	tests := []td{
		{
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			ans: 4,
		},
		{
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			ans: -1,
		},
		{
			grid: [][]int{
				{0, 2},
			},
			ans: 0,
		},
		{
			grid: [][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}},
			ans:  1,
		},
	}

	for _, v := range tests {
		res := orangesRotting(v.grid)
		if res != v.ans {
			t.Errorf("orangesRotting(%v) = %v, want %v", v.grid, res, v.ans)
		}
	}

}
