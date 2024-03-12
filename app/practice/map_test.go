package practice

import (
	"github.com/smartystreets/goconvey/convey"
	"graffito/practice/map_x"
	"testing"
)

func Test_Run2(t *testing.T) {
	convey.Convey("Map_Run2:\n", t, func() {
		map_x.Run2()
	})
}
