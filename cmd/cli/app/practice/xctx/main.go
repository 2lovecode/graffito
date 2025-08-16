package xctx

import (
	"context"
	"runtime"

	xctxsrv "github.com/2lovecode/graffito/internal/services/xstdlib/xctx"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xctx", Short: "Context实现原理探索"}
	cmd.AddCommand(&cobra.Command{
		Use: "cancel",
		Run: func(cmd *cobra.Command, args []string) {
			srv := xctxsrv.NewCancelService()
			srv.Run(context.Background())
		},
		Short: "CancelCtx实现原理",
	})
	cmd.AddCommand(&cobra.Command{
		Use: "gc",
		Run: func(cmd *cobra.Command, args []string) {

			a := data{v: 1}
			a.n = &data{v: 2}
			b := data{v: 3}
			b.n = &data{v: 4}
			t := a.n
			a.n = b.n
			runtime.GC()
			b.n = t

		},
		Short: "EmptyCtx实现原理",
	})
	return cmd
}

type data struct {
	v int
	n *data
}
