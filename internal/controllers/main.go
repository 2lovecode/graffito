package controllers

import "github.com/kataras/iris/v12"

func View(ctx iris.Context, filename string, optionalViewModel ...interface{}) {
	if err := ctx.View(filename, optionalViewModel...); err != nil {
		_, _ = ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
