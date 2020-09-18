package main

import (
	"fmt"
	"graffito/algorithm/lib"
)

func main() {
	myA := lib.NewMyArray(5)

	testData := map[int]int{0:3, 1:4, 2:5, 3:1, 4:2, 5:6}

	for k, v := range testData {
		myA.Insert(v, k)
	}

	for k, _ := range testData {
		fmt.Printf("k : %d, v : %d\n", k, myA.Get(k))
	}
}
