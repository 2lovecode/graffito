package pattern

import (
	"fmt"
	"sync"
)

var singletonOnce sync.Once
var instance Earth

type Earth struct {
	Name string
}

func NewSingleton() Earth {
	singletonOnce.Do(func(){
		instance = Earth{}
	})
	return instance
}


func SingletonRun() {
	earth := NewSingleton()
	earth.Name = "地球"

	earth2 := NewSingleton()

	fmt.Println("earth:", earth.Name)
	fmt.Println("earth2", earth2.Name)
}