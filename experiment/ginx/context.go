package ginx

import (
	"fmt"
	"graffito/experiment/ginx/render"
	"net/http"
	"time"
)

type XContext struct {
	engine *Engine

	Request *http.Request
	ResponseWriter http.ResponseWriter
}

func (c *XContext) Render(code int, r render.Render) {
	if err := r.Render(c.ResponseWriter); err != nil {
		fmt.Println("error!")
	}
}


/******** 实现 context.Context 接口 ********/

func (c *XContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *XContext) Done() <-chan struct{} {
	return nil
}

func (c *XContext) Err() error {
	return nil
}

func (c *XContext) Value(key interface{}) interface{} {
	return nil
}
