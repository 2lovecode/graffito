package pattern

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Strategy(t *testing.T) {
	Convey("Strategy", t, func() {
		mult := Operation{Operator: Multiplication{}}

		So(mult.Operate(3, 4), ShouldEqual, 12)
	})

}
