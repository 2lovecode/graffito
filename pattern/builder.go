package main

import "graffito/pattern/lib"

func main() {
	iBuilder := lib.NewIPhoneBuilder()

	bbPhone := iBuilder.Size(lib.SizeBig).Color(lib.ColorBlue).SimCard(lib.SimYiDong).Build()
	srPhone := iBuilder.Size(lib.SizeSmall).Color(lib.ColorRed).SimCard(lib.SimLianTong).Build()

	bbPhone.Call()
	srPhone.Call()
}
