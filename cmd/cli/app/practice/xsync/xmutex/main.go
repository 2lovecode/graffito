package xmutex

import (
	"context"

	xmsrv "github.com/2lovecode/graffito/internal/services/xstdlib/xsync/xmutex"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "xmutex", Short: "互斥锁(Mutex)实现原理探索"}
	cmd.AddCommand(&cobra.Command{
		Use: "sema",
		Run: func(cmd *cobra.Command, args []string) {
			impl := xmsrv.NewImplSema()
			impl.Run(context.Background())
		},
		Short: "仅使用信号量实现（注意Go的信号量和Os信号量的区别）",
	})
	cmd.AddCommand(&cobra.Command{
		Use: "cas",
		Run: func(cmd *cobra.Command, args []string) {
			impl := xmsrv.NewImplCAS()
			impl.Run(context.Background())
		},
		Short: "简单自旋实现",
	})
	// cmd.AddCommand(&cobra.Command{
	// 	Use: "sm1",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		impl := xmsrv.NewSimpleMutexImplSema()
	// 		fmt.Println(impl.Run(context.Background()))
	// 	},
	// 	Short: "Mutex实现示例",
	// })
	cmd.AddCommand(&cobra.Command{
		Use: "sema-cas",
		Run: func(cmd *cobra.Command, args []string) {
			impl := xmsrv.NewImplSemaCAS()
			impl.Run(context.Background())
		},
		Short: "自旋+信号量",
	})
	cmd.AddCommand(&cobra.Command{
		Use: "sm4",
		Run: func(cmd *cobra.Command, args []string) {
			xmsrv.NewSimpleMutexImplV0()

		},
		Short: "Mutex实现示例",
	})

	return cmd
}
