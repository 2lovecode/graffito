package labs

import (
	practice "github.com/2lovecode/graffito/internal/app/cli-practice"
	"github.com/2lovecode/graffito/internal/app/experiment"
	"github.com/2lovecode/graffito/internal/app/other"
	"github.com/spf13/cobra"
)

// NewCommand 创建实验功能类命令集合
func NewCommand() *cobra.Command {
	labsCmd := &cobra.Command{
		Use:   "labs",
		Short: "实验功能集",
		Long: `实验功能集合，包含实验性代码和实践探索。

包括以下内容：
  - 实验代码：各种实验性功能实现
  - CLI实践：命令行工具实践
  - 其他功能：其他实验性功能
`,
	}

	// 实验功能
	labsCmd.AddCommand(experiment.NewCommand())
	labsCmd.AddCommand(practice.NewCommand())
	labsCmd.AddCommand(other.NewCommand())

	return labsCmd
}
