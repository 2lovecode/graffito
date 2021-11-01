package experiment

import (
	"graffito/experiment/cache"

	"github.com/spf13/cobra"
)

func NewExperimentCommand() *cobra.Command {

	expCmd := &cobra.Command{Use: "exp", Short: "试验代码"}

	// stringOpCmd := &cobra.Command{Use: "g3n", Run: func(cmd *cobra.Command, args []string) {
	// 	g3n.Run()
	// }, Short: "g3n游戏包示例代码"}

	// expCmd.AddCommand(stringOpCmd)
	cacheCmd := &cobra.Command{Use: "cache", Run: func(cmd *cobra.Command, args []string) {
		cache.Run()
	}, Short: "3类缓存问题"}

	expCmd.AddCommand(cacheCmd)

	return expCmd
}
