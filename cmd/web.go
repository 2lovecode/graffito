package cmd

import (
	"fmt"
	"graffito/frontend"
	"graffito/internal/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/spf13/cobra"
)

func NewWebCommand() *cobra.Command {
	return &cobra.Command{Use: "web", Run: func(cmd *cobra.Command, args []string) {

		web := iris.New()
		web.RegisterView(iris.HTML(frontend.WebDir(), ".html"))

		corsMiddleware := cors.New()

		// 支持跨域
		web.UseGlobal(corsMiddleware.Handler())

		// web.OnAnyErrorCode(controllers.NewHome().Index)
		web.HandleDir("/assets", frontend.AssetsDir())
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
	},
		Short: "web工具",
	}
}
