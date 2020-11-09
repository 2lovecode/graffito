package main

import (
	"fmt"
	"graffito/experiment/event"
	"graffito/utils/params"
)

//实现handler
type MyCar struct {
}

func (mine *MyCar) Execute(p params.IPayload) {
	if p.Has("name") {
		fmt.Printf("Car name is : %s\n", p.Get("name", ""))
	}
}
//实现handler
type MyPhone struct {
}

func (mine *MyPhone) Execute(p params.IPayload) {
	if p.Has("name") {
		fmt.Printf("Phone name is : %s\n", p.Get("name", ""))
	}
}

func main() {
	mEvent := event.NewMEvent()

	myC := &MyCar{}
	myP := &MyPhone{}

	mEvent.On("car", myC)
	mEvent.On("phone", myP)

	p := params.NewPayload()
	p.Set("name", "vivo")

	mEvent.Trigger("car", p)
	mEvent.Trigger("phone", p)

}
