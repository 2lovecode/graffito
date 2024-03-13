package algorithm

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCc() *cobra.Command {
	return &cobra.Command{
		Use: "cc",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World!")
		},
	}
}
