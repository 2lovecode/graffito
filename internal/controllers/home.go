package controllers

import "github.com/kataras/iris/v12"

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (d *Home) Index(ctx iris.Context) {
	if err := ctx.View("index.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
