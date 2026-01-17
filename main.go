package main

import (
	"github.com/2lovecode/graffito/cmd/cli"
	"github.com/2lovecode/graffito/pkg/config"
	"github.com/2lovecode/graffito/pkg/logging"
	"github.com/spf13/cobra"
)

func main() {
	// 初始化日志
	logging.Init()

	// 加载配置
	_, err := config.Load()
	if err != nil {
		logging.Warnf("配置加载失败，使用默认配置: %v", err)
	}

	rootCmd := &cobra.Command{
		Use:   "graffito",
		Short: "Graffito - Go开发工具箱",
		Long: `Graffito是一个Go开发工具箱，集成常用开发工具和学习资源。

包括：
  - 开发工具：沙箱、时序图、图片处理等
  - 学习资源：LeetCode题解、算法实现、设计模式等
  - 实验功能：实验代码和实践探索

更多信息请访问: https://github.com/2lovecode/graffito
`,
		Version: "1.0.0",
	}

	rootCmd.AddCommand(cli.NewCommand())

	if err := rootCmd.Execute(); err != nil {
		logging.Fatalf("执行命令失败: %v", err)
	}
}
