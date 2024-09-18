package practice

import (
	"testing"

	"github.com/2lovecode/graffito/internal/app/practice/context_x"
	"github.com/smartystreets/goconvey/convey"
)

func TestContext_Run1(t *testing.T) {
	convey.Convey("Context_1", t, func() {
		x1, x2 := context_x.Run1()

		convey.So(x1, convey.ShouldEqual, "done")
		convey.So(x2, convey.ShouldEqual, "done")
	})
}
