package leetcode

import "github.com/spf13/cobra"

func NewLeetcodeCommand() *cobra.Command {
	leetCmd := &cobra.Command{Use: "leet", Short: "leetcode题解"}
	return leetCmd
}
