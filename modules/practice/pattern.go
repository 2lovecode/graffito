package practice

import (
	"fmt"
	"graffito/modules/practice/pattern"
	"graffito/utils/params"
)

func PatternRun(p params.InputParamsInterface) {
	key := p.GetInputPrefix() + "0"
	input := p.GetString(key)
	switch input {
	case "factory":
		patternFactory()
	default:
		patternList()
	}
}

func patternList(){
	fmt.Println("可使用的参数列表：", []string{
		"factory",
	})
}

func patternFactory() {
	mFactory := pattern.NewMPhoneFactory()

	mPhoneOne := mFactory.Size(pattern.SizeBig).Color(pattern.ColorBlue).SimCard(pattern.SimYiDong).Build()
	mPhoneTwo := mFactory.Size(pattern.SizeSmall).Color(pattern.ColorRed).SimCard(pattern.SimLianTong).Build()

	mPhoneOne.Call()
	mPhoneTwo.Call()

}


