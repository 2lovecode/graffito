package lib

import "fmt"

//定义常量
type PhoneSize string
const (
	SizeBig PhoneSize = "big"
	SizeMiddle = "middle"
	SizeSmall = "small"
)

type PhoneColor string
const (
	ColorRed PhoneColor = "red"
	ColorGreen = "green"
	ColorBlue = "blue"
)

type PhoneSim string
const (
	SimYiDong PhoneSim = "移动"
	SimLianTong = "联通"
	SimDianXin = "电信"
)

//手机接口
type PhoneInterface interface {
	Call()
}

//手机构造接口
type PhoneBuilder interface {
	Size(size PhoneSize) PhoneBuilder
	Color(color PhoneColor) PhoneBuilder
	SimCard(sim PhoneSim) PhoneBuilder
	Build() PhoneInterface
}


//实现手机接口的MPhone
type MPhone struct {
	Size PhoneSize
	Color PhoneColor
	SimCard PhoneSim
}

func (mine MPhone) Call() {
	fmt.Printf("%s,%s,%s phone calling\n", mine.Size, mine.Color, mine.SimCard)
}


//实现手机构建者接口的MPhone构建者
type MPhoneBuilder struct {
	phone MPhone
}

func NewMPhoneBuilder() MPhoneBuilder {
	return MPhoneBuilder{phone:MPhone{}}
}

func (mine MPhoneBuilder) Size(size PhoneSize) MPhoneBuilder {
	mine.phone.Size = size
	return mine
}

func (mine MPhoneBuilder) Color(color PhoneColor) MPhoneBuilder {
	mine.phone.Color = color
	return mine
}

func (mine MPhoneBuilder) SimCard(sim PhoneSim) MPhoneBuilder {
	mine.phone.SimCard = sim
	return mine
}

func (mine MPhoneBuilder) Build() MPhone {
	return mine.phone
}




