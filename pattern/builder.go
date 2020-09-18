package pattern

import "graffito/pattern/lib"

func main() {
	mBuilder := lib.NewMPhoneBuilder()

	mPhoneOne := mBuilder.Size(SizeBig).Color(ColorBlue).SimCard(SimYiDong).Build()
	mPhoneTwo := mBuilder.Size(SizeSmall).Color(ColorRed).SimCard(SimLianTong).Build()

	mPhoneOne.Call()
	mPhoneTwo.Call()
}
