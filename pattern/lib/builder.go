package lib

import "fmt"

//定义常量
type PhoneSize string
const (
	SizeBig PhoneSize = "大尺寸"
	SizeMiddle = "中等尺寸"
	SizeSmall = "小尺寸"
)

type PhoneColor string
const (
	ColorRed PhoneColor = "红色"
	ColorGreen = "绿色"
	ColorBlue = "蓝色"
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


//实现手机接口的IPhone
type IPhone struct {
	Size PhoneSize
	Color PhoneColor
	SimCard PhoneSim
}

func (mine IPhone) Call() {
	fmt.Printf("%s,%s,%s phone calling\n", mine.Size, mine.Color, mine.SimCard)
}


//实现手机构建者接口的构建者
type IPhoneBuilder struct {
	phone IPhone
}

func NewIPhoneBuilder() PhoneBuilder {
	return IPhoneBuilder{
		phone:IPhone{},
	}
}

func (mine IPhoneBuilder) Size(size PhoneSize) PhoneBuilder {
	mine.phone.Size = size
	return mine
}

func (mine IPhoneBuilder) Color(color PhoneColor) PhoneBuilder {
	mine.phone.Color = color
	return mine
}

func (mine IPhoneBuilder) SimCard(sim PhoneSim) PhoneBuilder {
	mine.phone.SimCard = sim
	return mine
}

func (mine IPhoneBuilder) Build() PhoneInterface {
	return mine.phone
}




