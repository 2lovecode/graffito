package main

import (
	"fmt"
	"graffito/pattern/lib"
)

// 单例模式
func main() {
	earth := lib.NewSingleton()
	earth.Name = "地球"

	earth2 := lib.NewSingleton()

	fmt.Println("earth:", earth.Name)
	fmt.Println("earth2:", earth2.Name)
}
