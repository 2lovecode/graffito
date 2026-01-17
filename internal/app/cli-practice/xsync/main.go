package xsync

import (
	"github.com/2lovecode/graffito/internal/app/cli-practice/xsync/xmutex"
	"github.com/2lovecode/graffito/internal/app/cli-practice/xsync/xother"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xsync"}

	cmd.AddCommand(xother.NewCommand())
	cmd.AddCommand(xmutex.NewCommand())
	return cmd
}
