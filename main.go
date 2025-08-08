package main

import (
	"github.com/2lovecode/graffito/cmd/cli"
	"github.com/2lovecode/graffito/cmd/web"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(cli.NewCommand())
	rootCmd.AddCommand(web.NewCommand())
	_ = rootCmd.Execute()
}
