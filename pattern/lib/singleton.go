package lib

import (
	"sync"
)

var singletonOnce sync.Once
var instance *Earth

type Earth struct {
	Name string
}

func NewSingleton() *Earth {
	singletonOnce.Do(func(){
		instance = &Earth{}
	})
	return instance
}