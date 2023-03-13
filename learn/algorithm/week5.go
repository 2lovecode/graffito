package algorithm

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewLeetcode() *cobra.Command {
	return &cobra.Command{
		Use: "leetcode",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Good!")
		},
	}
}
