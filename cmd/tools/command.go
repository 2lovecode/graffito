package tools

import (
	"github.com/2lovecode/graffito/cmd/tools/color"
	"github.com/2lovecode/graffito/cmd/tools/excel"
	"github.com/2lovecode/graffito/cmd/tools/helper"
	"github.com/2lovecode/graffito/cmd/tools/modg"
	"github.com/2lovecode/graffito/cmd/tools/modgv"
	"github.com/2lovecode/graffito/cmd/tools/onezero"
	"github.com/2lovecode/graffito/cmd/tools/redis"
	"github.com/2lovecode/graffito/cmd/tools/sql2go"
	"github.com/2lovecode/graffito/cmd/tools/string_op"
	"github.com/2lovecode/graffito/cmd/tools/subway"
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
