package algorithm

import (
	"github.com/spf13/cobra"
)

func NewAlgorithmCommand() *cobra.Command {
	return &cobra.Command{Use: "alg", Short: "数据结构"}
}
