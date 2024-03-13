package cache

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "cache",
		Run: func(cmd *cobra.Command, args []string) {
			Execute()
		},
	}
}
func Execute() {
	// 缓存穿透
	// 解决缓存穿透的方案

	// 缓存击穿
	// 解决缓存击穿的方案

	// 缓存雪崩
	// 解决缓存雪崩的方案
}
