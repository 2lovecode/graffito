package sandboxcli

import (
	"context"
	"fmt"

	dto "github.com/2lovecode/graffito/internal/dto/sandbox"
	"github.com/2lovecode/graffito/internal/app/shared"
	"github.com/2lovecode/graffito/pkg/logging"
	srv "github.com/2lovecode/graffito/internal/services/sandbox"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	file := ""
	source := ""

	base := &shared.CLIBase{}
	
	cmd := &cobra.Command{
		Use:   "sandbox",
		Short: "Go代码在线运行沙箱",
		Long: `在沙箱环境中运行Go代码。

支持两种输入方式：
  1. 从文件读取: --file/-f <文件路径>
  2. 直接输入源代码: --source/-s <源代码>

示例:
  # 从文件运行
  graffito cli tools sandbox --file main.go
  
  # 直接运行源代码
  graffito cli tools sandbox --source "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello\")\n}"
`,
		Run: func(cmd *cobra.Command, args []string) {
			// 读取源代码
			sourceCode, err := base.ReadSourceCode(file, source)
			if err != nil {
				logging.Errorf("读取源代码失败: %v", err)
				fmt.Fprintf(cmd.ErrOrStderr(), "错误: %v\n", err)
				fmt.Fprintf(cmd.OutOrStdout(), "使用 --help 查看帮助信息\n")
				return
			}

			// 执行沙箱
			sandboxApp := srv.NewApplication()
			out, err := sandboxApp.Exec(context.Background(), dto.Input{
				SourceCode: sourceCode,
			})

			if err != nil {
				logging.Errorf("沙箱执行失败: %v", err)
				fmt.Fprintf(cmd.ErrOrStderr(), "执行错误: %v\n", err)
				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), "输出:\n%s\n", out.Data)
		},
	}
	
	cmd.Flags().StringVarP(&file, "file", "f", "", "Go源代码文件路径")
	cmd.Flags().StringVarP(&source, "source", "s", "", "Go源代码内容")
	
	return cmd
}
