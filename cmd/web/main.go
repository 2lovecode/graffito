package web

import (
	"fmt"

	"github.com/2lovecode/graffito/internal/controllers"
	"github.com/2lovecode/graffito/pkg/config"
	"github.com/2lovecode/graffito/pkg/logging"
	web2 "github.com/2lovecode/graffito/web"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	webCmd := &cobra.Command{
		Use:   "web",
		Short: "启动Web服务",
		Long:  "启动Graffito Web服务，提供Web界面和API接口",
		Run: func(cmd *cobra.Command, args []string) {
			// 加载配置
			cfg, err := config.Load()
			if err != nil {
				logging.Errorf("加载配置失败: %v", err)
				cfg = config.Get() // 使用默认配置
			}

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

			addr := fmt.Sprintf(":%s", cfg.Server.Port)
			logging.Infof("Web服务启动在 %s", addr)
			err = web.Listen(addr)
			if err != nil {
				logging.Fatalf("Web服务启动失败: %v", err)
				return
			}
		},
	}
	return webCmd
}
