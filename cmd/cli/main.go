package main

import (
	"github.com/2lovecode/graffito/cmd/cli/app/practice"
	"github.com/2lovecode/graffito/cmd/cli/app/sandbox"
	"github.com/2lovecode/graffito/cmd/cli/app/search"
	"github.com/2lovecode/graffito/internal/app/experiment"
	"github.com/2lovecode/graffito/internal/app/learn"
	"github.com/2lovecode/graffito/internal/app/leetcode"
	"github.com/2lovecode/graffito/internal/app/other"
	"github.com/2lovecode/graffito/pkg/algorithm"
	"github.com/2lovecode/graffito/pkg/pattern"
	"github.com/2lovecode/graffito/tools"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	list := []*cobra.Command{
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
	rootCmd.AddCommand(list...)
	_ = rootCmd.Execute()
}
