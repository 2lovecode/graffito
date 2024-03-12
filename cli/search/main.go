package search

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	search2 "graffito/app/search"
	"io"
	"os"
)

func NewCommand() *cobra.Command {
	file := ""
	source := ""

	cmd := &cobra.Command{Use: "search", Run: func(cmd *cobra.Command, args []string) {
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
		searchApp := search2.NewApplication()

		out, err := searchApp.Exec(context.Background(), &search2.Input{
			SourceCode: sourceCode,
		})

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		so := search2.Output{}
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
