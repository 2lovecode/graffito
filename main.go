package main

import (
	"fmt"
	"graffito/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := cobra.Command{}

	rootCmd.AddCommand(cmd.NewWebCommand())
	rootCmd.AddCommand(cmd.NewCliCommand())

	fmt.Print("")
	_ = rootCmd.Execute()
}
