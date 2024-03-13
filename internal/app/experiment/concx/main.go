package concx

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "conc",
		Run: func(cmd *cobra.Command, args []string) {
			var wg conc.WaitGroup
			defer wg.Wait()

			wg.Go(func() {
				fmt.Println("Hello World!")
			})
		},
	}
}
