package cli

import (
	"github.com/2lovecode/graffito/cmd/web"
	"github.com/2lovecode/graffito/internal/app/labs"
	"github.com/2lovecode/graffito/internal/app/learning"
	"github.com/2lovecode/graffito/internal/app/tools"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cliCmd := &cobra.Command{
		Use:   "cli",
		Short: "Graffito命令行工具",
		Long: `Graffito - Go开发工具箱

提供三大类功能：
  - tools  : 开发工具集（沙箱、时序图、图片处理等）
  - learn  : 学习资源集（LeetCode、算法、设计模式等）
  - labs   : 实验功能集（实验代码、实践探索等）

使用示例：
  graffito cli tools sandbox --source "package main..."
  graffito cli learn leetcode list
  graffito cli labs experiment depends
`,
	}

	// 三大类命令
	cliCmd.AddCommand(tools.NewCommand())
	cliCmd.AddCommand(learning.NewCommand())
	cliCmd.AddCommand(labs.NewCommand())
	cliCmd.AddCommand(web.NewCommand())

	return cliCmd
}
