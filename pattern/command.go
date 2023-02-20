package pattern

import "github.com/spf13/cobra"

func NewCommand() *cobra.Command {
	patternCmd := &cobra.Command{Use: "pattern", Short: "设计模式"}

	return patternCmd
}
