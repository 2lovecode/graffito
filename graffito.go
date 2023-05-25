package main

import (
	"graffito/algorithm"
	"graffito/experiment"
	"graffito/learn"
	"graffito/leetcode"
	"graffito/other"
	"graffito/pattern"
	"graffito/practice"
	"graffito/tools"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{Use: "graffito"}
	cmds := []*cobra.Command{
		tools.NewCommand(),
		algorithm.NewCommand(),
		experiment.NewCommand(),
		practice.NewCommand(),
		pattern.NewCommand(),
		leetcode.NewCommand(),
		learn.NewCommand(),
		other.NewCommand(),
	}
	rootCmd.AddCommand(cmds...)

	rootCmd.Execute()
}
