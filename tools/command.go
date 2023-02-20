package tools

import (
	"github.com/spf13/cobra"
	"graffito/tools/color"
	"graffito/tools/excel"
	"graffito/tools/helper"
	"graffito/tools/modgv"
	"graffito/tools/redis"
	"graffito/tools/sql2go"
	"graffito/tools/string_op"
)

func NewCommand() *cobra.Command {

	toolsCmd := &cobra.Command{Use: "tools", Short: "工具箱"}

	cmds := []*cobra.Command{
		string_op.NewCommand(),
		helper.NewCommand(),
		redis.NewCommand(),
		color.NewCommand(),
		excel.NewCommand(),
		sql2go.NewCommand(),
		modgv.NewCommand(),
	}
	toolsCmd.AddCommand(cmds...)

	return toolsCmd
}
