package mode0

import (
	"context"
	"errors"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/card"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"
)

type DealerLayer struct {
	data    OutData
	cardMap map[string]card.ICard
}

func NewDealerLayer(data OutData) (dl *DealerLayer) {
	dl = &DealerLayer{
		data:    data,
		cardMap: make(map[string]card.ICard),
	}
	return dl
}

func (dl *DealerLayer) Output() []interface{} {
	var out []interface{}
	var ctx context.Context
	for _, v := range dl.data.Data {
		c, err := dl.GetCard(v.Type)
		if err == nil {
			ctx = context.WithValue(context.Background(), "data", &handler.InputData{
				Data: v,
			})
			c.Input(ctx)
			out = append(out, c.Output())
		}
	}
	return out
}

func (dl *DealerLayer) RegisterCard(card card.ICard) {
	dl.cardMap[card.GetCardType()] = card
}

func (dl *DealerLayer) GetCard(t string) (c card.ICard, e error) {
	c, ok := dl.cardMap[t]
	if !ok {
		e = errors.New("nil")
	}
	return
}
