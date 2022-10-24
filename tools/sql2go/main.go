package sql2go

import (
	"github.com/qmhball/db2gorm/gen"
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
