package pattern

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Builder(t *testing.T) {
	Convey("Builder", t, func() {
		iBuilder := NewIPhoneBuilder()

		bbPhone := iBuilder.Size(SizeBig).Color(ColorBlue).SimCard(SimYiDong).Build()
		srPhone := iBuilder.Size(SizeSmall).Color(ColorRed).SimCard(SimLianTong).Build()

		bbN := bbPhone.Call()
		srN := srPhone.Call()
		So(bbN, ShouldEqual, "大尺寸,蓝色,移动")
		So(srN, ShouldEqual, "小尺寸,红色,联通")
	})
}
