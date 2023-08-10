package main

import (
	"github.com/spf13/cobra"
	"graffito/cmd"
)

func main() {
	rootCmd := cobra.Command{}

	rootCmd.AddCommand(cmd.NewWebCommand())
	rootCmd.AddCommand(cmd.NewCliCommand())

	_ = rootCmd.Execute()
}
