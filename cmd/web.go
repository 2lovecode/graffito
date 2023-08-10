package cmd

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

func NewWebCommand() *cobra.Command {
	return &cobra.Command{Use: "web", Run: func(cmd *cobra.Command, args []string) {
		app := iris.New()

		booksAPI := app.Party("/books")
		{
			booksAPI.Use(iris.Compression)
		}

		app.Listen(":8080")
	}}
}
