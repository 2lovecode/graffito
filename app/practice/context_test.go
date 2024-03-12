package practice

import (
	"github.com/smartystreets/goconvey/convey"
	"graffito/practice/context_x"
	"testing"
)

func TestContext_Run1(t *testing.T) {
	convey.Convey("Context_1", t, func() {
		x1, x2 := context_x.Run1()

		convey.So(x1, convey.ShouldEqual, "done")
		convey.So(x2, convey.ShouldEqual, "done")
	})
}