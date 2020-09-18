package event

import (
	"graffito/utils/params"
)

type MEvent struct {
	EMap map[string]MHandler
}

type MHandler interface {
	Execute(p params.IPayload)
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


func (mine *MEvent) Trigger(name string, p params.IPayload) {
	handler, ok := mine.EMap[name]
	if ok {
		handler.Execute(p)
	}
}


