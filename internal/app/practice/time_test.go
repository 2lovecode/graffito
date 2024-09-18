package practice

import (
	"testing"

	"github.com/2lovecode/graffito/internal/app/practice/time_x"
	"github.com/smartystreets/goconvey/convey"
)

func TestTime_Run1(t *testing.T) {
	convey.Convey("Time_1", t, func() {
		convey.So(time_x.Run1(), convey.ShouldEqual, true)
	})
}
