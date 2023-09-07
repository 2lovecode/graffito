package cmd

import (
	"graffito/algorithm"
	"graffito/app"
	"graffito/experiment"
	"graffito/learn"
	"graffito/leetcode"
	"graffito/other"
	"graffito/pattern"
	"graffito/practice"
	"graffito/tools"

	"github.com/spf13/cobra"
)

func NewCliCommand() *cobra.Command {
	cli := &cobra.Command{
		Use: "cli",
	}
	cmds := []*cobra.Command{
		app.NewSandboxCommand(),
		tools.NewCommand(),
		algorithm.NewCommand(),
		experiment.NewCommand(),
		practice.NewCommand(),
		pattern.NewCommand(),
		leetcode.NewCommand(),
		learn.NewCommand(),
		other.NewCommand(),
	}
	cli.AddCommand(cmds...)
	return cli
}
