package practice

import (
	"graffito/practice/slice_x"

	"github.com/spf13/cobra"
)

func NewPracticeCommand() *cobra.Command {

	pracCmd := &cobra.Command{Use: "prac", Short: "练习代码"}

	slice1OpCmd := &cobra.Command{Use: "slice_1", Run: func(cmd *cobra.Command, args []string) {
		slice_x.Run_1()
	}}

	pracCmd.AddCommand(slice1OpCmd)

	return pracCmd
}
