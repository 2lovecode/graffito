package lib

import "context"

// 责任链

type Handler interface {
	Do(ctx context.Context) error
	Run(ctx context.Context) error
	SetNext(h Handler) Handler
}

type Next struct {
	nextHandler Handler
}

func (n *Next) Run(ctx context.Context) error {
	if n.nextHandler != nil {
		if err := (n.nextHandler).Do(ctx); err != nil {
			return err
		}
		return (n.nextHandler).Run(ctx)
	}
	return nil
}

func (n *Next) SetNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

type NilHandler struct {
	Next
}

func (h *NilHandler) Do(ctx context.Context) error {
	return nil
}
