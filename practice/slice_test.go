package practice

import (
	"github.com/smartystreets/goconvey/convey"
	"graffito/practice/slice_x"
	"testing"
)

func TestSlice_Run1(t *testing.T) {
	convey.Convey("Slice_1", t, func() {
		slice_x.Run_1()
	})
}

func TestSlice_Run2(t *testing.T) {
	convey.Convey("Slice_2", t, func() {
		slice_x.Run_2()
	})
}

func TestSlice_Run3(t *testing.T) {
	convey.Convey("Slice_3", t, func() {
		slice_x.Run_3()
	})
}
