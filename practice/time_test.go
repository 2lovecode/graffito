package practice

import (
	"github.com/smartystreets/goconvey/convey"
	"graffito/practice/time_x"
	"testing"
)

func TestTime_Run1(t *testing.T) {
	convey.Convey("Time_1", t, func() {
		convey.So(time_x.Run1(), convey.ShouldEqual, true)
	})
}
