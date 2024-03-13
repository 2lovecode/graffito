package pattern

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Singleton(t *testing.T) {
	Convey("Singleton", t, func() {
		earth := NewSingleton()
		earth.Name = "地球"

		earth2 := NewSingleton()

		So(earth2.Name, ShouldEqual, "地球")
	})

}
