package dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	convey.Convey("爬楼梯", t, func() {
		convey.So(climbStairs(4), convey.ShouldEqual, 5)
	})
}
