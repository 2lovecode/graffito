package top100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	data := [][]int{
		{1},
		{5, 4, -1, 7, 8},
		{-1, -6},
		{-9, -6, -8},
	}
	expected := []int{1, 23, -1, -6}
	for k, _ := range data {
		assert.Equal(t, expected[k], maxSubArray(data[k]))
	}
	for k, _ := range data {
		assert.Equal(t, expected[k], maxSubArray2(data[k]))
	}
	for k, _ := range data {
		assert.Equal(t, expected[k], maxSubArray3(data[k]))
	}
}
