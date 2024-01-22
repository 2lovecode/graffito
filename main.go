package main

import (
	"graffito/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{}

	rootCmd.AddCommand(cmd.NewWebCommand())
	rootCmd.AddCommand(cmd.NewCliCommand())

	_ = rootCmd.Execute()
}
