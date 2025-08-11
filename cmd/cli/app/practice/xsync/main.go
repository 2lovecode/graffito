package xsync

import (
	"github.com/2lovecode/graffito/cmd/cli/app/practice/xsync/xgroup"
	"github.com/2lovecode/graffito/cmd/cli/app/practice/xsync/xmutex"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xsync"}

	cmd.AddCommand(xgroup.NewCommand())
	cmd.AddCommand(xmutex.NewCommand())
	return cmd
}
