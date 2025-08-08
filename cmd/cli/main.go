package cli

import (
	"github.com/2lovecode/graffito/cmd/cli/app/geo"
	"github.com/2lovecode/graffito/cmd/cli/app/image2"
	"github.com/2lovecode/graffito/cmd/cli/app/media"
	"github.com/2lovecode/graffito/cmd/cli/app/plantuml"
	"github.com/2lovecode/graffito/cmd/cli/app/poi"
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

func NewCommand() *cobra.Command {
	cliCmd := &cobra.Command{
		Use: "cli",
	}

	cliCmd.AddCommand(sandbox.NewCommand())
	cliCmd.AddCommand(plantuml.NewCommand())
	cliCmd.AddCommand(search.NewCommand())
	cliCmd.AddCommand(practice.NewCommand())
	cliCmd.AddCommand(tools.NewCommand())
	cliCmd.AddCommand(algorithm.NewCommand())
	cliCmd.AddCommand(experiment.NewCommand())
	cliCmd.AddCommand(pattern.NewCommand())
	cliCmd.AddCommand(leetcode.NewCommand())
	cliCmd.AddCommand(learn.NewCommand())
	cliCmd.AddCommand(other.NewCommand())
	cliCmd.AddCommand(geo.NewTransCommand())
	cliCmd.AddCommand(media.NewFileTypeCommand())
	cliCmd.AddCommand(poi.NewPOICommand())
	cliCmd.AddCommand(image2.NewCommand())

	return cliCmd
}
