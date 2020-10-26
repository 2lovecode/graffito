package main

import (
	"fmt"
	"graffito/pattern/lib"
)

func main() {
	mult := lib.Operation{Operator: lib.Multiplication{}}

	fmt.Println(mult.Operate(3, 4))
}
