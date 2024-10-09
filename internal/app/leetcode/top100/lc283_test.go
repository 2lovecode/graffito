package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	convey.Convey("移动零", t, func() {
		nums := []int{0, 1, 0, 3, 12}
		moveZeroes(nums)
		convey.So(nums, convey.ShouldResemble, []int{1, 3, 12, 0, 0})
	})
}
