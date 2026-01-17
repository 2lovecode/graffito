package practice

import (
	"github.com/2lovecode/graffito/internal/app/cli-practice/xctx"
	"github.com/2lovecode/graffito/internal/app/cli-practice/xsync"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "practice"}

	cmd.AddCommand(xsync.NewCommand(), xctx.NewCommand())
	return cmd
}
