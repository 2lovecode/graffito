package sandbox

import (
	"bytes"
	"context"
	"fmt"
	sandbox2 "github.com/2lovecode/graffito/internal/app/sandbox"
	"github.com/spf13/cobra"
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
	}, Short: "Go在线运行沙箱"}
	cmd.Flags().StringVarP(&file, "file", "f", "", "指定文件")
	cmd.Flags().StringVarP(&source, "source", "s", "", "指定源代码")
	return cmd
}
