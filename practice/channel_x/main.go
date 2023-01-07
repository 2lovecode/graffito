package channel_x

import (
	"github.com/spf13/cobra"
	"graffito/practice/channel_x/impl"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "channel", Run: func(cmd *cobra.Command, args []string) {
		impl.NoBuffer()
		impl.WithBuffer()
	}}
}
