package dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFib(t *testing.T) {
	convey.Convey("斐波那契", t, func() {
		convey.So(fib(4), convey.ShouldEqual, 3)
	})
}
