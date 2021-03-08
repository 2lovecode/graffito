package pattern

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 工厂模式
func Test_Factory(t *testing.T) {
	Convey("Factory", t, func() {
		So(NewPc(TypeLenovo).RunPrint(), ShouldEqual, "联想")
		So(NewPc(TypeDell).RunPrint(), ShouldEqual, "戴尔")
		So(NewPc(TypeHp).RunPrint(), ShouldEqual, "惠普")
	})

}
