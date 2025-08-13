package xctx

import (
	"context"
	"fmt"
	"time"
)

type XContext interface {
	XDeadline() (deadline time.Time, ok bool)
	XDone() <-chan struct{}
	XErr() error
	XValue(key any) any
}

type backgroundCtx struct {
	emptyXCtx
}

type todoCtx struct {
	emptyXCtx
}

func Background() XContext {
	return backgroundCtx{}
}

func TODO() XContext {
	return todoCtx{}
}

type CancelService struct {
}

func NewCancelService() *CancelService {
	return &CancelService{}
}

func (ct *CancelService) Run(ctx context.Context) {
	c := Background()

	nc, cf := WithCancel(c)

	go func() {
		nc1, _ := WithCancel(nc)
		<-nc1.XDone()
		fmt.Println("cancel 1")
	}()

	go func() {
		nc2, _ := WithCancel(nc)
		<-nc2.XDone()
		fmt.Println("cancel 2")
	}()

	go func() {
		<-time.After(time.Second * 1)
		// fmt.Println("cancel")
		cf()
	}()
	<-nc.XDone()
	fmt.Println("finish")
	<-time.After(time.Second * 1)
}
