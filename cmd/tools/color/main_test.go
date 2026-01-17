package color

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_ConvertColorFormat(t *testing.T) {
	var ColorValueMap = map[ColorFormat]string{
		ColorFormatHEX:  "#C7A68D",
		ColorFormatAHEX: "#FFC7A68D",
		ColorFormatHEXA: "#C7A68DFF",
		ColorFormatRGB:  "rgb(199, 166, 141)",
		ColorFormatRGBA: "rgba(199, 166, 141, 1.00)",
		// ColorFormatHSL:  "hsl(26, 34%, 67%)",
		// ColorFormatHSLA: "hsla(26, 34%, 67%, 1.00)",
	}

	Convey("ConvertColorFormat\n", t, func() {

		for beforeFormat, beforeValue := range ColorValueMap {
			for afterFormat, afterValue := range ColorValueMap {
				if beforeFormat == afterFormat {
					continue
				}
				v, _ := ConvertColorFormat(beforeValue, beforeFormat, afterFormat)
				So(v, ShouldEqual, afterValue)
				msg := fmt.Sprintf("%s-%s-%s-%s-%s", beforeValue, beforeFormat, afterValue, afterFormat, v)
				t.Log(msg)
			}
		}

	})
}
