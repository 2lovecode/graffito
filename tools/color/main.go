package color

import (
	"errors"
	"image/color"
	"regexp"
)

type ColorFormat string

const ColorFormatHEX ColorFormat = "HEX"   // #C7A68D
const ColorFormatAHEX ColorFormat = "AHEX" // #FFC7A68D
const ColorFormatHEXA ColorFormat = "HEXA" // #C7A68DFF
const ColorFormatRGB ColorFormat = "RGB"   // rgb(199, 166, 141)
const ColorFormatRGBA ColorFormat = "RGBA" // rgba(199, 166, 141, 1)
// const ColorFormatHSL ColorFormat = "HSL"   // hsl(26, 34%, 67%)
// const ColorFormatHSLA ColorFormat = "HSLA" // hsla(26, 34%, 67%, 1)

var ColorRegxs = map[ColorFormat]string{
	ColorFormatHEX:  "#[a-fA-F\\d]{6}",
	ColorFormatAHEX: "#[a-fA-F\\d]{8}",
	ColorFormatHEXA: "#[a-fA-F\\d]{8}",
	ColorFormatRGB:  "[Rr][Gg][Bb][\\(](([0-9]{1,3}),\\s*([0-9]{1,3}),\\s*([0-9]{1,3}))[\\)]",
	ColorFormatRGBA: "[Rr][Gg][Bb][Aa][\\(](([0-9]{1,3}),\\s*([0-9]{1,3}),\\s*([0-9]{1,3}),\\s*([0-9.]{1,4}))[\\)]",
	// ColorFormatHSL:  "[Hh][Ss][Ll][\\(](((([\\d]{1,3}|[\\d\\%]{2,4})[\\,]{0,1})[\\s]*){3})[\\)]",
	// ColorFormatHSLA: "[Hh][Ss][Ll][Aa][\\(](((([\\d]{1,3}|[\\d\\%]{2,4}|[\\d\\.]{1,3})[\\,]{0,1})[\\s]*){4})[\\)]",
}

var ColorHandlerMap = map[ColorFormat]ColorHandler{
	ColorFormatHEX:  &HEXColorHandler{},
	ColorFormatAHEX: &AHEXColorHandler{},
	ColorFormatHEXA: &HEXAColorHandler{},
	ColorFormatRGB:  &RGBColorHandler{},
	ColorFormatRGBA: &RGBAColorHandler{},
	// ColorFormatHSL:  &HSLColorHandler{},
	// ColorFormatHSLA: &HSLAColorHandler{},
}

func ConvertColorFormat(before string, from ColorFormat, to ColorFormat) (after string, err error) {
	after = before
	if before == "" || from == to {
		return
	}
	if ok, e := ValidateColorFormat(before, from); !ok {
		if e == nil {
			err = errors.New("颜色格式错误")
		} else {
			err = e
		}
		return
	}

	beforeColor := color.RGBA{}
	if begoreH, ok := ColorHandlerMap[from]; ok {
		beforeColor = begoreH.Decode(before)

		if afterH, okk := ColorHandlerMap[to]; okk {
			after = afterH.Encode(beforeColor)
		}
	}
	return
}

func ValidateColorFormat(color string, format ColorFormat) (flag bool, err error) {
	if reg, ok := ColorRegxs[format]; ok {
		// fmt.Println(reg, color, format)
		return regexp.MatchString(reg, color)
	}
	return false, errors.New("不支持该格颜色式！")
}
