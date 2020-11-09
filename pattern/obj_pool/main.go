package main

import "graffito/pattern/lib"

// 对象池模式
func main() {
	pool := lib.NewPool(10)

	for i := 0; i < 20; i++ {
		pool.Do()
	}
}
