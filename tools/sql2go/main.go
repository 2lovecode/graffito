package sql2go

import (
	"github.com/qmhball/db2gorm/gen"
	"github.com/spf13/cobra"
)

func Run() {
	dsn := "root:user237@tcp(10.10.10.237:3306)/mydb?charset=utf8&parseTime=true&loc=Local"

	//生成指定单表
	tblName := "User"
	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, tblName)
}

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "sql2go", Run: func(cmd *cobra.Command, args []string) {
		Run()
	}, Short: "sql2go", Example: "graffito tools sql2go"}
}
