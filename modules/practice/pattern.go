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
	case "singleton":
		patternSingleton()
	case "builder":
		patternBuilder()
	case "factory":
		patternFactory()
	default:
		patternList()
	}
}

func patternList(){
	fmt.Println("可使用的参数列表：", []string{
		"singleton",
		"builder",
		"factory",
	})
}

func patternSingleton() {
	earth := pattern.NewSingleton()
	earth.Name = "地球"

	earth2 := pattern.NewSingleton()

	fmt.Println("earth:", earth.Name)
	fmt.Println("earth2", earth2.Name)
}

func patternBuilder() {
	mBuilder := pattern.NewMPhoneBuilder()

	mPhoneOne := mBuilder.Size(pattern.SizeBig).Color(pattern.ColorBlue).SimCard(pattern.SimYiDong).Build()
	mPhoneTwo := mBuilder.Size(pattern.SizeSmall).Color(pattern.ColorRed).SimCard(pattern.SimLianTong).Build()

	mPhoneOne.Call()
	mPhoneTwo.Call()
}

func patternFactory() {
	pattern.NewPc(pattern.TypeLenovo).RunPrint()
	pattern.NewPc(pattern.TypeDell).RunPrint()
	pattern.NewPc(pattern.TypeHp).RunPrint()
}
