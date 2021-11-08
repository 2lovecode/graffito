package event

import (
	"fmt"
	"graffito/utils/params"
)

type MEvent struct {
	EMap map[string][]MHandler
}

type MHandler interface {
	Execute(p params.IPayload)
}

func NewMEvent() *MEvent {
	e := &MEvent{
		EMap: map[string][]MHandler{},
	}
	return e
}

func (mine *MEvent) On(name string, handler MHandler) {
	if _, ok := mine.EMap[name]; ok {
		mine.EMap[name] = append(mine.EMap[name], handler)
	} else {
		mine.EMap[name] = []MHandler{handler}
	}
}

func (mine *MEvent) Trigger(name string, p params.IPayload) {
	handlers, ok := mine.EMap[name]
	if ok && len(handlers) > 0 {
		for _, handler := range handlers {
			handler.Execute(p)
		}
	}
}

//实现handler
type MyCar struct {
}

func (mine *MyCar) Execute(p params.IPayload) {
	data := ""
	if p.Has("name") {
		if s, ok := p.Get("name", "").(string); ok {
			data = s
		}
	}
	fmt.Println("car-" + data)
}

//实现handler
type MyPhone struct {
}

func (mine *MyPhone) Execute(p params.IPayload) {
	data := ""
	if p.Has("name") {
		if s, ok := p.Get("name", "").(string); ok {
			data = s
		}
	}
	fmt.Println("phone-" + data)
}
