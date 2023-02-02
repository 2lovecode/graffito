package interface_x

import "fmt"

type interfaceDemo interface {
	hello()
}

// 检查某类型是否实现了某接口的方法
type implementDemo struct{}

func (impl *implementDemo) hello() {
	fmt.Println("hello")
}

func checkImplement() {
	// 检查某类型的引用是否实现了某接口
	var _ interfaceDemo = (*implementDemo)(nil)

	// 检查某类型是否实现了某接口
	//var _ interfaceDemo = implementDemo{}
}
