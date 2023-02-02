package interface_x

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	interfaceCmd := &cobra.Command{Use: "interface"}
	cmds := []*cobra.Command{
		{
			Use: "compare-with-nil",
			Run: func(cmd *cobra.Command, args []string) {
				compareWithNil()
			},
		},
		{
			Use: "print-dynamic-type-and-value",
			Run: func(cmd *cobra.Command, args []string) {
				printDynamicTypeAndValue()
			},
		},
		{
			Use: "check-implement",
			Run: func(cmd *cobra.Command, args []string) {
				checkImplement()
			},
		},
	}
	interfaceCmd.AddCommand(cmds...)

	return interfaceCmd
}
