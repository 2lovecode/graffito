package dynamic_programming

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTribonacci(t *testing.T) {
	convey.Convey("泰波契那", t, func() {
		convey.So(tribonacci(25), convey.ShouldEqual, 1389537)
	})
}
