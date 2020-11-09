package main

import "graffito/pattern/lib"

// 构建者模式
func main() {
	iBuilder := lib.NewIPhoneBuilder()

	bbPhone := iBuilder.Size(lib.SizeBig).Color(lib.ColorBlue).SimCard(lib.SimYiDong).Build()
	srPhone := iBuilder.Size(lib.SizeSmall).Color(lib.ColorRed).SimCard(lib.SimLianTong).Build()

	bbPhone.Call()
	srPhone.Call()
}
