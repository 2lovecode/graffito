package top100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	// [[1,1,1],[1,0,1],[1,1,1]]
	// [[1,0,1],[0,0,0],[1,0,1]]
	nums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(nums)
	assert.Equal(t, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, nums)
}
