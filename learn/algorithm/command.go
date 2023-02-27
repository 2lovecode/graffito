package algorithm

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "alg",
	}
	cmds := []*cobra.Command{
		NewClimbStair(),
		NewIntersectLink(),
		NewReverseLink(),
	}
	rootCmd.AddCommand(cmds...)
	return rootCmd
}
