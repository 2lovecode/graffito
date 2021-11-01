package experiment

import (
	"graffito/experiment/g3n"

	"github.com/spf13/cobra"
)

func NewExperimentCommand() *cobra.Command {

	expCmd := &cobra.Command{Use: "exp", Short: "试验代码"}

	stringOpCmd := &cobra.Command{Use: "g3n", Run: func(cmd *cobra.Command, args []string) {
		g3n.Run()
	}, Short: "g3n游戏包示例代码"}

	expCmd.AddCommand(stringOpCmd)

	return expCmd
}
