package searchcli

import (
	"context"
	"fmt"

	search2 "github.com/2lovecode/graffito/internal/app/search"
	"github.com/2lovecode/graffito/internal/app/shared"
	"github.com/2lovecode/graffito/pkg/logging"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	file := ""
	source := ""

	base := &shared.CLIBase{}
	
	cmd := &cobra.Command{
		Use:   "search",
		Short: "搜索功能工具",
		Long: `提供文本搜索和检索功能。

支持两种输入方式：
  1. 从文件读取: --file/-f <文件路径>
  2. 直接输入内容: --source/-s <内容>

示例:
  # 从文件搜索
  graffito cli tools search --file data.txt
  
  # 直接搜索内容
  graffito cli tools search --source "搜索关键词"
`,
		Run: func(cmd *cobra.Command, args []string) {
			// 读取源代码
			sourceCode, err := base.ReadSourceCode(file, source)
			if err != nil {
				logging.Errorf("读取内容失败: %v", err)
				fmt.Fprintf(cmd.ErrOrStderr(), "错误: %v\n", err)
				fmt.Fprintf(cmd.OutOrStdout(), "使用 --help 查看帮助信息\n")
				return
			}

			// 执行搜索
			searchApp := search2.NewApplication()
			out, err := searchApp.Exec(context.Background(), &search2.Input{
				SourceCode: sourceCode,
			})

			if err != nil {
				logging.Errorf("搜索执行失败: %v", err)
				fmt.Fprintf(cmd.ErrOrStderr(), "执行错误: %v\n", err)
				return
			}

			so := search2.Output{}
			err = out.To(&so)
			if err != nil {
				logging.Errorf("结果解析失败: %v", err)
				fmt.Fprintf(cmd.ErrOrStderr(), "结果解析错误: %v\n", err)
				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), "输出:\n%s\n", so.Data)
		},
	}
	
	cmd.Flags().StringVarP(&file, "file", "f", "", "输入文件路径")
	cmd.Flags().StringVarP(&source, "source", "s", "", "输入内容")
	
	return cmd
}
