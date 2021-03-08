package tools

import (
	"graffito/tools/count"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Count(t *testing.T) {
	Convey("Count", t, func() {
		So(count.Count("我是谁呀t"), ShouldEqual, 5)
	})
}
