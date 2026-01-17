package learning

import (
	"github.com/2lovecode/graffito/internal/app/learn"
	"github.com/2lovecode/graffito/internal/app/leetcode"
	"github.com/2lovecode/graffito/pkg/algorithm"
	"github.com/2lovecode/graffito/pkg/pattern"
	"github.com/spf13/cobra"
)

// NewCommand 创建学习资源类命令集合
func NewCommand() *cobra.Command {
	learningCmd := &cobra.Command{
		Use:   "learn",
		Short: "学习资源集",
		Long: `学习资源集合，包含算法、设计模式、LeetCode题解等。

包括以下内容：
  - LeetCode题解：Top100和动态规划等经典题目
  - 算法实现：数据结构与算法实现
  - 设计模式：常用设计模式实现
  - Go实践：Go语言特性实践代码
  - 学习笔记：Go语言学习笔记
`,
	}

	// 学习资源
	learningCmd.AddCommand(leetcode.NewCommand())
	learningCmd.AddCommand(learn.NewCommand())
	learningCmd.AddCommand(algorithm.NewCommand())
	learningCmd.AddCommand(pattern.NewCommand())
	// practice.NewCommand() 暂时注释，等实现后再启用

	return learningCmd
}
