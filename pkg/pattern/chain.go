package pattern

import "context"

// 责任链

type ChainHandler interface {
	Do(ctx context.Context) error
	Run(ctx context.Context) error
	SetNext(h ChainHandler) ChainHandler
}

type ChainNext struct {
	nextHandler ChainHandler
}

func (n *ChainNext) Run(ctx context.Context) error {
	if n.nextHandler != nil {
		if err := (n.nextHandler).Do(ctx); err != nil {
			return err
		}
		return (n.nextHandler).Run(ctx)
	}
	return nil
}

func (n *ChainNext) SetNext(h ChainHandler) ChainHandler {
	n.nextHandler = h
	return h
}

type NilChainHandler struct {
	ChainNext
}

func (h *NilChainHandler) Do(ctx context.Context) error {
	return nil
}
