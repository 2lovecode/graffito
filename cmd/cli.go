package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCliCommand() *cobra.Command {
	return &cobra.Command{Use: "cli", Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cli")
	}}
}
