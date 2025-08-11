package xgroup

import (
	xgsrv "github.com/2lovecode/graffito/internal/services/xstdlib/xsync/xgroup"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xgroup", Short: "WaitGroup实现原理探索"}
	cmd.AddCommand(&cobra.Command{
		Use: "sg",
		Run: func(cmd *cobra.Command, args []string) {
			xgsrv.NewSimpleGroup()

		},
		Short: "WaitGroup实现示例",
	})
	return cmd
}
