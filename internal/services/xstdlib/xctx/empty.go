package xctx

import "time"

type emptyXCtx struct{}

func (emptyXCtx) XDeadline() (deadline time.Time, ok bool) {
	return
}

func (emptyXCtx) XDone() <-chan struct{} {
	return nil
}

func (emptyXCtx) XErr() error {
	return nil
}

func (emptyXCtx) XValue(key any) any {
	return nil
}
