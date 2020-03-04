package pattern

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

//手机工厂接口
type PhoneFactory interface {
	Size(size PhoneSize) PhoneFactory
	Color(color PhoneColor) PhoneFactory
	SimCard(sim PhoneSim) PhoneFactory
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


//实现手机工厂接口的MPhone工厂
type MPhoneFactory struct {
	phone MPhone
}

func NewMPhoneFactory() MPhoneFactory {
	return MPhoneFactory{phone:MPhone{}}
}

func (mine MPhoneFactory) Size(size PhoneSize) MPhoneFactory {
	mine.phone.Size = size
	return mine
}

func (mine MPhoneFactory) Color(color PhoneColor) MPhoneFactory {
	mine.phone.Color = color
	return mine
}

func (mine MPhoneFactory) SimCard(sim PhoneSim) MPhoneFactory {
	mine.phone.SimCard = sim
	return mine
}

func (mine MPhoneFactory) Build() MPhone {
	return mine.phone
}


