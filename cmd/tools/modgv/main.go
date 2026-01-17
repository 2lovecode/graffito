package modgv

import (
	"github.com/spf13/cobra"
	"os"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "modgv", Run: func(cmd *cobra.Command, args []string) {
		Render(os.Stdin, os.Stdout)
	}, Short: "命令列表", Example: "graffito tools helper"}
}
