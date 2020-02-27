package main

import (
	"flag"
	"fmt"
	"graffito/modules"
	"os"
)

type params struct {
	help bool
	module string
	task string
}


var p params

func init() {

	flag.BoolVar(&p.help, "h",false, "帮助")
	flag.StringVar(&p.module,"m", "tools", "指定模块名")
	flag.StringVar(&p.task,"t", "hello", "指定选定模块下需要执行的任务名")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	if p.help {
		flag.Usage()
	}
	modules.Run(modules.InputParams{
		Module:p.module,
		Task:   p.task,
		Params: nil,
	})
}

func usage() {
	fmt.Fprintf(os.Stderr, `graffito version: 0.0.0
用法: go run main.go [-m module] [-t task] [-h]

选项:

`)
	flag.PrintDefaults()
}
