package sandbox

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	sandbox2 "graffito/app/sandbox"
	"io"
	"os"
)

func NewCommand() *cobra.Command {
	file := ""
	source := ""

	cmd := &cobra.Command{Use: "sandbox", Run: func(cmd *cobra.Command, args []string) {
		sourceCode := ""
		if source != "" {
			sourceCode = source
		} else if file != "" {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			var buffer bytes.Buffer
			_, err = io.Copy(&buffer, f)
			if err != nil {
				fmt.Println(err)
			}
			sourceCode = buffer.String()
		}

		sandboxApp := sandbox2.NewApplication()
		out, err := sandboxApp.Exec(context.Background(), &sandbox2.Input{
			SourceCode: sourceCode,
		})

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		so := sandbox2.Output{}
		err = out.To(&so)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Output: ", so.Data)
	}}
	cmd.Flags().StringVarP(&file, "file", "f", "", "指定文件")
	cmd.Flags().StringVarP(&source, "source", "s", "", "指定源代码")
	return cmd
}
