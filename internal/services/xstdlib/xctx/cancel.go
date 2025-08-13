package xctx

import (
	"sync"
	"time"
)

type CancelFunc func()

type cancelXCtx struct {
	mu       sync.Mutex
	done     chan struct{}
	children map[*cancelXCtx]struct{}
}

func (c *cancelXCtx) XDeadline() (deadline time.Time, ok bool) {
	return
}

func (c *cancelXCtx) XDone() <-chan struct{} {
	return c.done
}

func (c *cancelXCtx) XErr() error {
	return nil
}

func (c *cancelXCtx) XValue(key any) any {
	return nil
}

func (c *cancelXCtx) cancel() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.done == nil {
		return
	}
	close(c.done)
	for child, _ := range c.children {
		child.cancel()
	}
	c.children = nil
}

func WithCancel(parent XContext) (ctx XContext, cancel CancelFunc) {
	c := &cancelXCtx{
		mu:       sync.Mutex{},
		done:     make(chan struct{}),
		children: make(map[*cancelXCtx]struct{}),
	}

	if p, ok := parent.(*cancelXCtx); ok {
		c.mu.Lock()
		p.children[c] = struct{}{}
		c.mu.Unlock()
	}

	return c, func() {
		c.cancel()
	}
}
