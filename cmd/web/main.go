package main

import (
	"fmt"
	"github.com/2lovecode/graffito/internal/controllers"
	web2 "github.com/2lovecode/graffito/web"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
)

func main() {
	web := iris.New()
	web.RegisterView(iris.HTML(web2.WebDir(), ".html"))

	corsMiddleware := cors.New()

	// 支持跨域
	web.UseGlobal(corsMiddleware.Handler())

	// web.OnAnyErrorCode(controllers.NewHome().Index)
	web.HandleDir("/assets", web2.AssetsDir())
	appsAPI := web.Party("/apps")
	{
		sandboxCtrl := controllers.NewSandboxController()

		appsAPI.Any("/sandbox/run", sandboxCtrl.Run)
	}

	// 前端
	web.Get("/", controllers.NewHome().Index)
	web.Get("/{apps}", controllers.NewHome().Index)

	err := web.Listen(":8081")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
