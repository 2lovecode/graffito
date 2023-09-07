package cmd

import (
	"embed"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/spf13/cobra"
	"graffito/cmd/web/backend/controllers"
	"io/fs"
)

//go:embed web/frontend/dist/*
var frontendFileSystem embed.FS

func NewWebCommand() *cobra.Command {
	return &cobra.Command{Use: "web", Run: func(cmd *cobra.Command, args []string) {

		webDir, webErr := fs.Sub(frontendFileSystem, "web")

		if webErr != nil {
			fmt.Println("web:", webErr)
			return
		}

		frontendDir, frontendErr := fs.Sub(webDir, "frontend")
		if frontendErr != nil {
			fmt.Println("frontend:", frontendErr)
			return
		}

		distDir, distErr := fs.Sub(frontendDir, "dist")
		if distErr != nil {
			fmt.Println("dist:", distErr)
			return
		}

		assetsDir, assetsErr := fs.Sub(distDir, "assets")
		if assetsErr != nil {
			fmt.Println("assets:", assetsErr)
			return
		}

		web := iris.New()
		web.RegisterView(iris.HTML(distDir, ".html"))

		corsMiddleware := cors.New()

		// 使用中间件
		web.UseGlobal(corsMiddleware.Handler())

		web.OnAnyErrorCode(controllers.NewHome().Index)
		web.HandleDir("/assets", assetsDir)
		appsAPI := web.Party("/apps")
		{
			sandboxCtrl := controllers.NewSandboxController()

			appsAPI.Any("/sandbox/run", sandboxCtrl.Run)
		}

		err := web.Listen(":8082")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}}
}
