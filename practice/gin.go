package practice

import (
	"github.com/gin-gonic/gin"
)

func GinRun() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
