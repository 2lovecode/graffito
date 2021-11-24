package practice

import (
	"graffito/practice/slice_x"
	"graffito/practice/ts"

	"github.com/spf13/cobra"
)

func NewPracticeCommand() *cobra.Command {

	pracCmd := &cobra.Command{Use: "prac", Short: "练习代码"}

	slice1OpCmd := &cobra.Command{Use: "slice", Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			switch args[0] {
			case "1":
				slice_x.Run_1()
			case "2":
				slice_x.Run_2()
			case "3":
				slice_x.Run_3()
			case "4":
				slice_x.Run_4()
			}
		}

	}}

	pracCmd.AddCommand(slice1OpCmd)

	tsCmd := &cobra.Command{Use: "ts", Run: func(cmd *cobra.Command, args []string) {
		ts.Run()
	}}
	pracCmd.AddCommand(tsCmd)

	return pracCmd
}
