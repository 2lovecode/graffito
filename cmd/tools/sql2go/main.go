package sql2go

import (
	"fmt"

	"github.com/qmhball/db2gorm/gen"
	"github.com/spf13/cobra"
)

func Run(dsn string) {
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
	user := ""
	pass := ""
	addr := ""
	dbname := ""
	cmd := &cobra.Command{Use: "sql2go", Run: func(cmd *cobra.Command, args []string) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pass, addr, dbname)
		Run(dsn)
	}, Short: "sql2go", Example: "graffito tools sql2go"}

	cmd.Flags().StringVarP(&user, "user", "u", "", "用户名")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "密码")
	cmd.Flags().StringVarP(&addr, "addr", "a", "", "Addr")
	cmd.Flags().StringVarP(&dbname, "dbname", "d", "", "数据库")

	return cmd
}
