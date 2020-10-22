package handler

import (
	"context"
)

type IHandler interface {
	GetName() string
	Input(ctx context.Context)
	Output() interface{}
}

type DataItem struct {
	Type string
	Name string
	Extra interface{}
}

type InputData struct {
	CommonData interface{}
	Data DataItem
	Extra interface{}
}

type BaseHandler struct {
	Ctx context.Context
}

func (h *BaseHandler) Input(ctx context.Context) {
	h.Ctx = ctx
}

func (h *BaseHandler) GetInputData() *InputData {
	data, ok := h.Ctx.Value("data").(*InputData)
	if ok {
		return data
	}
	return nil
}