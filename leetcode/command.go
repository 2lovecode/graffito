package leetcode

import (
	"github.com/spf13/cobra"
	"graffito/leetcode/dynamic_programming/fibonacci"
)

func NewCommand() *cobra.Command {
	leetCmd := &cobra.Command{Use: "leet", Short: "leetcode题解"}

	var stairsCount int

	cmds := []*cobra.Command{
		{
			Use:   "climb_stairs",
			Short: "爬楼梯",
			Run: func(cmd *cobra.Command, args []string) {
				fibonacci.TestClimbStairs(stairsCount)
			},
		},
	}
	// 设置 "climb_stairs" 子命令的参数
	for _, cmd := range cmds {
		switch cmd.Use {
		case "climb_stairs":
			cmd.Flags().IntVar(&stairsCount, "count", 1, "楼梯总数")
		}
	}

	leetCmd.AddCommand(cmds...)
	return leetCmd
}
