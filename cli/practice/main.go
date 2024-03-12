package practice

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	practice2 "graffito/app/practice"
)

func NewCommand() *cobra.Command {
	chapter := ""
	number := 0
	part := ""

	cmd := &cobra.Command{Use: "practice", Run: func(cmd *cobra.Command, args []string) {
		app := practice2.NewApplication()
		out, err := app.Exec(context.Background(), &practice2.Input{
			Chapter: chapter,
			Number:  number,
			Part:    part,
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
	cmd.Flags().StringVarP(&chapter, "chapter", "c", "", "指定章节")
	cmd.Flags().IntVarP(&number, "number", "n", 0, "第几道习题")
	cmd.Flags().StringVarP(&part, "part", "p", "all", "输出部分：q-问题，a-答案，all-所有（默认）")
	return cmd
}
