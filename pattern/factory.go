package pattern

import "graffito/pattern/lib"

func main() {
	lib.NewPc(lib.TypeLenovo).RunPrint()
	lib.NewPc(lib.TypeDell).RunPrint()
	lib.NewPc(lib.TypeHp).RunPrint()
}
