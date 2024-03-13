package card

import (
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"
	jsoniter "github.com/json-iterator/go"
)

type SmallCard struct {
	BaseCard
	Name string `json:"name"`
}

func NewSmallCard() *SmallCard {
	c := &SmallCard{}
	c.HandlerMap = make(map[string]handler.IHandler)
	return c
}

func (sm *SmallCard) GetCardType() string {
	return "small"
}

func (sm *SmallCard) Output() string {
	if nameHandler, err := sm.GetHandler("name"); err == nil {
		nameHandler.Input(sm.Ctx)
		sm.Name = nameHandler.Output().(string)
	}
	data, _ := jsoniter.MarshalToString(sm)
	return data
}
