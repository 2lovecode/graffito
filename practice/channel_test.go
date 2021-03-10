package practice

import (
	. "github.com/smartystreets/goconvey/convey"
	"graffito/practice/channel_x"
	"testing"
)

func Test_Channel_1(t *testing.T) {
	Convey("Channel_1", t, func() {
		x1, x2 := channel_x.Run1()

		So(x1, ShouldEqual, 1)
		So(x2, ShouldEqual, 2)
	})
}

func Test_Channel_2(t *testing.T) {
	Convey("Channel_2", t, func() {
		x1, x2 := channel_x.Run2()
		So(x1, ShouldEqual, "tick")
		So(x2, ShouldEqual, "boom")
	})
}

func Test_Channel_3(t *testing.T) {
	Convey("Channel_3", t, func() {
		So(channel_x.Run3(), ShouldEqual, "timeout")
	})
}
