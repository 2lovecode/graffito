package xctx

import (
	"context"

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
	return cmd
}
