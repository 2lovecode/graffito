package xsync

import (
	"github.com/2lovecode/graffito/cmd/cli/app/practice/xsync/xmutex"
	"github.com/2lovecode/graffito/cmd/cli/app/practice/xsync/xother"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xsync"}

	cmd.AddCommand(xother.NewCommand())
	cmd.AddCommand(xmutex.NewCommand())
	return cmd
}
