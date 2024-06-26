package practice

import (
	"context"
	"fmt"
	practice2 "github.com/2lovecode/graffito/internal/app/practice"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	question := ""
	list := false

	cmd := &cobra.Command{Use: "practice", Run: func(cmd *cobra.Command, args []string) {
		app := practice2.NewApplication()
		out, err := app.Exec(context.Background(), &practice2.Input{
			List:     list,
			Question: question,
		})

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		so := practice2.Output{}
		err = out.To(&so)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Output: ", so.Data)
	}}
	cmd.Flags().StringVarP(&question, "question", "q", "", "指定问题")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "列出所有的问题")
	return cmd
}
