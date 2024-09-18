package practice

import (
	"testing"

	"github.com/2lovecode/graffito/internal/app/practice/interface_x"
	"github.com/smartystreets/goconvey/convey"
)

func TestInterface_Run1(t *testing.T) {
	convey.Convey("Test Run1:\n", t, func() {
		interface_x.Run1()
	})
}

func TestInterface_Run2(t *testing.T) {
	convey.Convey("Test Run2:\n", t, func() {
		interface_x.Run2()
	})
}

func TestInterface_Run3(t *testing.T) {
	convey.Convey("Test Run3:\n", t, func() {
		interface_x.Run3()
	})
}
