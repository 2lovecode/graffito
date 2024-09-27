package leetcode

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	leetCmd := &cobra.Command{Use: "leet", Short: "leetcode题解"}

	var count int

	cmds := []*cobra.Command{}
	// 设置 "climb_stairs" 子命令的参数
	for _, cmd := range cmds {
		switch cmd.Use {
		case "climb_stairs":
			cmd.Flags().IntVar(&count, "count", 1, "楼梯总数")
		case "fib":
			cmd.Flags().IntVar(&count, "count", 1, "n")
		case "trib":
			cmd.Flags().IntVar(&count, "count", 1, "n")
		}
	}

	leetCmd.AddCommand(cmds...)
	return leetCmd
}
