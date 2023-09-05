package cmd

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/spf13/cobra"
	"graffito/web/backend/controllers"
)

func NewWebCommand() *cobra.Command {
	return &cobra.Command{Use: "web", Run: func(cmd *cobra.Command, args []string) {
		web := iris.New()
		web.RegisterView(iris.HTML("./web/frontend/dist", ".html"))
		web.HandleDir("/", iris.Dir("./web/frontend/dist"))
		corsMiddleware := cors.New()

		// 使用中间件
		web.UseGlobal(corsMiddleware.Handler())

		web.Get("/", controllers.NewHome().Index)

		appsAPI := web.Party("/apps")
		{
			sandboxCtrl := controllers.NewSandboxController()

			appsAPI.Any("/sandbox/run", sandboxCtrl.Run)
		}

		err := web.Listen(":8081")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}}
}
