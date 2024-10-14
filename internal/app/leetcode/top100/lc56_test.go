package top100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
