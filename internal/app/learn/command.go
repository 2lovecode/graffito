package learn

import (
	"github.com/2lovecode/graffito/internal/app/learn/algorithm"
	"github.com/spf13/cobra"
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
