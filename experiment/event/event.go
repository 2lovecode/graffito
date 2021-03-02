package event

import (
	"graffito/utils/params"
)

type MEvent struct {
	EMap map[string]MHandler
}

type MHandler interface {
	Execute(p params.IPayload) string
}

func NewMEvent() *MEvent {
	e := &MEvent{
		EMap: map[string]MHandler{},
	}
	return e
}

func (mine *MEvent) On(name string, handler MHandler) {
	mine.EMap[name] = handler
}

func (mine *MEvent) Trigger(name string, p params.IPayload) string {
	handler, ok := mine.EMap[name]
	if ok {
		return handler.Execute(p)
	}
	return ""
}

//实现handler
type MyCar struct {
}

func (mine *MyCar) Execute(p params.IPayload) string {
	data := ""
	if p.Has("name") {
		if s, ok := p.Get("name", "").(string); ok {
			data = s
		}
	}
	return "car-" + data
}

//实现handler
type MyPhone struct {
}

func (mine *MyPhone) Execute(p params.IPayload) string {
	data := ""
	if p.Has("name") {
		if s, ok := p.Get("name", "").(string); ok {
			data = s
		}
	}
	return "phone-" + data
}
