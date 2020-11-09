package main

import "graffito/pattern/lib"

// 工厂模式
func main() {
	lib.NewPc(lib.TypeLenovo).RunPrint()
	lib.NewPc(lib.TypeDell).RunPrint()
	lib.NewPc(lib.TypeHp).RunPrint()
}
