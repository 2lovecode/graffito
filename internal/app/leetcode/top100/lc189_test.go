package top100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	assert.Equal(t, []int{3, 99, -1, -100}, nums2)
}
