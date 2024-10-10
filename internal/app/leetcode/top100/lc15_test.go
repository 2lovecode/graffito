package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestThreeSum(t *testing.T) {
	convey.Convey("三数之和", t, func() {
		convey.So(threeSum([]int{0, 1, 1}), convey.ShouldResemble, [][]int{})
		convey.So(threeSum([]int{0, 0, 0}), convey.ShouldResemble, [][]int{{0, 0, 0}})
	})
}
