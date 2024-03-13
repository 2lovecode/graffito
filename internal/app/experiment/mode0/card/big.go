package card

import (
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"
	jsoniter "github.com/json-iterator/go"
)

type BigCard struct {
	BaseCard
	Name string `json:"name"`
}

func NewBigCard() *BigCard {
	c := &BigCard{}
	c.HandlerMap = make(map[string]handler.IHandler)
	return c
}

func (hc *BigCard) GetCardType() string {
	return "big"
}

func (hc *BigCard) Output() string {
	if nameHandler, err := hc.GetHandler("name"); err == nil {
		nameHandler.Input(hc.Ctx)
		hc.Name = nameHandler.Output().(string)
	}
	data, _ := jsoniter.MarshalToString(hc)
	return data
}
