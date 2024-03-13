package main

import (
	"github.com/2lovecode/graffito/cmd/cli/app/geo"
	"github.com/2lovecode/graffito/cmd/cli/app/media"
	"github.com/2lovecode/graffito/cmd/cli/app/plantuml"
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

	rootCmd.AddCommand(sandbox.NewCommand())
	rootCmd.AddCommand(plantuml.NewCommand())
	rootCmd.AddCommand(search.NewCommand())
	rootCmd.AddCommand(practice.NewCommand())
	rootCmd.AddCommand(tools.NewCommand())
	rootCmd.AddCommand(algorithm.NewCommand())
	rootCmd.AddCommand(experiment.NewCommand())
	rootCmd.AddCommand(pattern.NewCommand())
	rootCmd.AddCommand(leetcode.NewCommand())
	rootCmd.AddCommand(learn.NewCommand())
	rootCmd.AddCommand(other.NewCommand())
	rootCmd.AddCommand(geo.NewTransCommand())
	rootCmd.AddCommand(media.NewFileTypeCommand())

	_ = rootCmd.Execute()
}
