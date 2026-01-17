package tools

import (
	"github.com/2lovecode/graffito/internal/app/geo"
	"github.com/2lovecode/graffito/internal/app/image2"
	"github.com/2lovecode/graffito/internal/app/media"
	"github.com/2lovecode/graffito/internal/app/plantuml"
	"github.com/2lovecode/graffito/internal/app/poi"
	"github.com/2lovecode/graffito/internal/app/sandbox-cli"
	"github.com/2lovecode/graffito/internal/app/search-cli"
	"github.com/2lovecode/graffito/cmd/tools"
	"github.com/spf13/cobra"
)

// NewCommand 创建工具类命令集合
func NewCommand() *cobra.Command {
	toolsCmd := &cobra.Command{
		Use:   "tools",
		Short: "开发工具集",
		Long: `开发工具集合，提供各种实用开发工具。

包括以下工具：
  - 沙箱环境：Go代码在线运行
  - 时序图生成：PlantUML工具
  - 图片处理：图片合成、水印等
  - POI搜索：百度地图POI搜索
  - 文件类型识别：通过文件头识别
  - 地理坐标转换：坐标系转换
  - 其他工具：模块依赖图、Redis工具等
`,
	}

	// 核心工具
	toolsCmd.AddCommand(sandboxcli.NewCommand())
	toolsCmd.AddCommand(plantuml.NewCommand())
	toolsCmd.AddCommand(searchcli.NewCommand())
	toolsCmd.AddCommand(geo.NewTransCommand())
	toolsCmd.AddCommand(media.NewFileTypeCommand())
	toolsCmd.AddCommand(poi.NewPOICommand())
	toolsCmd.AddCommand(image2.NewCommand())

	// 集成cmd/tools下的工具
	toolsSubCmd := tools.NewCommand()
	toolsSubCmd.Use = "utils"
	toolsSubCmd.Short = "辅助工具集"
	toolsCmd.AddCommand(toolsSubCmd)

	return toolsCmd
}
