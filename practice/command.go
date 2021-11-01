package practice

import (
	"github.com/spf13/cobra"
)

func NewPracticeCommand() *cobra.Command {

	pracCmd := &cobra.Command{Use: "prac", Short: "练习代码"}

	// stringOpCmd := &cobra.Command{Use: "g3n", Run: func(cmd *cobra.Command, args []string) {
	// 	g3n.Run()
	// }, Short: "g3n游戏包示例代码"}

	// pracCmd.AddCommand(stringOpCmd)

	return pracCmd
}
