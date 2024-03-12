package cmd

import (
	"graffito/algorithm"
	"graffito/cli/practice"
	"graffito/cli/sandbox"
	"graffito/cli/search"
	"graffito/experiment"
	"graffito/learn"
	"graffito/leetcode"
	"graffito/other"
	"graffito/pattern"
	"graffito/tools"

	"github.com/spf13/cobra"
)

func NewCliCommand() *cobra.Command {
	cli := &cobra.Command{
		Use:   "cli",
		Short: "命令行工具",
	}
	cmds := []*cobra.Command{
		sandbox.NewCommand(),
		search.NewCommand(),
		practice.NewCommand(),
		tools.NewCommand(),
		algorithm.NewCommand(),
		experiment.NewCommand(),
		pattern.NewCommand(),
		leetcode.NewCommand(),
		learn.NewCommand(),
		other.NewCommand(),
	}
	cli.AddCommand(cmds...)
	return cli
}
