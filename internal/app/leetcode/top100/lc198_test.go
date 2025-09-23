package top100

import "testing"

func TestRob(t *testing.T) {
	type td struct {
		d []int
		e int
	}
	tds := []td{
		{
			d: []int{1, 2, 3, 1},
			e: 4,
		},
		{
			d: []int{2, 7, 9, 3, 1},
			e: 12,
		},
	}
	for _, v := range tds {
		if rob(v.d) == v.e {
			t.Log("success")
		} else {
			t.Error("error")
		}
	}
}
