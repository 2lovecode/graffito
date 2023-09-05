package response

import "github.com/kataras/iris/v12"

func Success(ctx iris.Context, data interface{}) {
	_ = ctx.JSON(iris.Map{
		"status":       "success",
		"errorCode":    "",
		"errorMessage": "",
		"data":         data,
	})
}

func Failure(ctx iris.Context, errorCode string, errorMsg string) {
	_ = ctx.JSON(iris.Map{
		"status":       "failure",
		"errorCode":    errorCode,
		"errorMessage": errorMsg,
		"data":         nil,
	})
}
