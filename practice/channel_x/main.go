package channel_x

import (
	"fmt"
	"github.com/spf13/cobra"
	"graffito/practice/channel_x/impl"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "channel", Run: func(cmd *cobra.Command, args []string) {
		run1()
		impl.NoBuffer()
		impl.WithBuffer()
	}}
}

func run1() {
	c := make(chan int, 10)
	c <- 1
	close(c)

	select {
	case d, ok := <-c:
		if !ok {
			fmt.Println("no data received!")
			break
		}
		fmt.Println(d)
	}
}
