package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSubarraySum(t *testing.T) {
	convey.Convey("数组中和为k的子数组个数", t, func() {
		convey.So(subarraySum([]int{1, 1, 1}, 2), convey.ShouldEqual, 2)
		convey.So(subarraySum([]int{1, 2, 3}, 3), convey.ShouldEqual, 2)
		convey.So(subarraySum([]int{1}, 0), convey.ShouldEqual, 0)
		convey.So(subarraySum([]int{-1, -1, 1}, 0), convey.ShouldEqual, 1)
		convey.So(subarraySum([]int{1, -1, 0}, 0), convey.ShouldEqual, 3)
		convey.So(subarraySum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0), convey.ShouldEqual, 55)
	})
}
