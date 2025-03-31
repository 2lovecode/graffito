package top100

import (
	"testing"
)

func TestLc200(t *testing.T) {
	type TestData struct {
		grid   [][]byte
		expect int
	}

	data := []*TestData{
		//{
		//	grid:   [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
		//	expect: 3,
		//},
		//{
		//	grid:   [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
		//	expect: 1,
		//},
		{
			grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
			expect: 1,
		},
	}

	for _, v := range data {
		if res := numIslands(v.grid); res != v.expect {
			t.Errorf("numIslands(%v) = %v, expect %v", v.grid, res, v.expect)
		}
	}
}
