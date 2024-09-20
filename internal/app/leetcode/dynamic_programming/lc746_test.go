package dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	convey.Convey("爬楼梯最小花费", t, func() {
		convey.So(minCostClimbingStairs([]int{10, 15, 20}), convey.ShouldEqual, 15)
	})
}
