package app

import (
	"bytes"
	"context"
	"fmt"
	"graffito/app/sandbox"
	"graffito/app/search"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewSandboxCommand() *cobra.Command {
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
	cmd.Flags().StringVarP(&file, "file", "f", "", "指定文件")
	cmd.Flags().StringVarP(&source, "source", "s", "", "指定源代码")
	return cmd
}

func NewSearchCommand() *cobra.Command {
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
		searchApp := search.NewApplication()

		out, err := searchApp.Exec(context.Background(), &search.Input{
			SourceCode: sourceCode,
		})

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		so := search.Output{}
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
