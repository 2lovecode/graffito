package learn

import (
	"github.com/spf13/cobra"
	"graffito/internal/app/learn/algorithm"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "learn",
	}
	cmds := []*cobra.Command{
		algorithm.NewCommand(),
	}
	rootCmd.AddCommand(cmds...)
	return rootCmd
}
