package cmd

import (
	"bytes"
	"context"
	"fmt"
	"graffito/app/sandbox"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewCliCommand() *cobra.Command {
	file := ""
	source := ""

	cli := &cobra.Command{Use: "cli", Run: func(cmd *cobra.Command, args []string) {
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

		sandboxApp := sandbox.NewApplication()
		out, err := sandboxApp.Exec(context.Background(), &sandbox.Input{
			SourceCode: sourceCode,
		})

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		so := sandbox.Output{}
		err = out.To(&so)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Output: ", so.Data)
	}}
	cli.Flags().StringVarP(&file, "file", "f", "", "指定文件")
	cli.Flags().StringVarP(&source, "source", "s", "", "指定源代码")
	return cli
}
