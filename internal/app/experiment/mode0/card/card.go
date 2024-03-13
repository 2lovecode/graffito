package card

import (
	"context"
	"github.com/2lovecode/graffito/internal/app/experiment/mode0/handler"
)

type ICard interface {
	GetCardType() string
	Input(ctx context.Context)
	Output() string
	Register(handler handler.IHandler)
	GetHandler(name string) (handler handler.IHandler, err error)
}

type BaseCard struct {
	Ctx        context.Context             `json:"-"`
	HandlerMap map[string]handler.IHandler `json:"-"`
}

func (b *BaseCard) Register(handler handler.IHandler) {
	b.HandlerMap[handler.GetName()] = handler
}

func (b *BaseCard) GetHandler(name string) (handler handler.IHandler, err error) {
	return b.HandlerMap[name], nil
}

func (b *BaseCard) Input(ctx context.Context) {
	b.Ctx = ctx
}
