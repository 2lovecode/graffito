package practice

import (
	"github.com/2lovecode/graffito/cmd/cli/app/practice/xctx"
	"github.com/2lovecode/graffito/cmd/cli/app/practice/xsync"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "practice"}

	cmd.AddCommand(xsync.NewCommand(), xctx.NewCommand())
	return cmd
}
