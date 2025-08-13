package xother

import (
	xosrv "github.com/2lovecode/graffito/internal/services/xstdlib/xsync/xother"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xother", Short: "其他同步锁实现原理探索"}
	cmd.AddCommand(&cobra.Command{
		Use: "waitgroup",
		Run: func(cmd *cobra.Command, args []string) {
			xosrv.NewWaitGroup()

		},
		Short: "其他同步锁实现原理探索",
	})
	return cmd
}
