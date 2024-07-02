package tools

import (
	"github.com/2lovecode/graffito/tools/color"
	"github.com/2lovecode/graffito/tools/excel"
	"github.com/2lovecode/graffito/tools/helper"
	"github.com/2lovecode/graffito/tools/modg"
	"github.com/2lovecode/graffito/tools/modgv"
	"github.com/2lovecode/graffito/tools/onezero"
	"github.com/2lovecode/graffito/tools/redis"
	"github.com/2lovecode/graffito/tools/sql2go"
	"github.com/2lovecode/graffito/tools/string_op"
	"github.com/2lovecode/graffito/tools/subway"
	"github.com/spf13/cobra"
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
		modg.NewCommand(),
		modgv.NewCommand(),
		subway.NewCommand(),
		onezero.NewCommand(),
	}
	toolsCmd.AddCommand(cmds...)

	return toolsCmd
}
