package leetcode

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	leetCmd := &cobra.Command{Use: "leet", Short: "leetcode题解"}
	return leetCmd
}
