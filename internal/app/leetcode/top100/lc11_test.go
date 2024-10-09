package top100

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxArea(t *testing.T) {
	convey.Convey("盛最多水的容器", t, func() {
		convey.So(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), convey.ShouldEqual, 49)
	})
}
