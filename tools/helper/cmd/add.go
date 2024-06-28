package cmd

import "github.com/spf13/cobra"

func NewAddCommand() *cobra.Command {

	name := ""
	desc := ""
	usage := ""

	add := &cobra.Command{Use: "add", Run: func(cmd *cobra.Command, args []string) {
		if name == "" || desc == "" || usage == "" {
			_ = cmd.Help()
			return
		}

	}, Short: "添加命令备忘", Example: "{path/to/exe} tools helper"}

	add.Flags().StringVarP(&name, "name", "n", "", "命令名称")
	add.Flags().StringVarP(&desc, "desc", "d", "", "命令描述")
	add.Flags().StringVarP(&usage, "usage", "u", "", "命令用法")
	return add
}
