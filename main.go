package main

import (
	"flag"
	"fmt"
	"graffito/modules"
	g_params "graffito/utils/params"
	"os"
	"strconv"
)

type params struct {
	help bool
	module string
	task string
	f string
}


var p params

func init() {

	flag.BoolVar(&p.help, "h",false, "帮助")
	flag.StringVar(&p.module,"m", "tools", "指定模块名")
	flag.StringVar(&p.task,"t", "hello", "指定选定模块下需要执行的任务名")
	flag.StringVar(&p.f,"f", "./", "文件路径")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	if p.help {
		flag.Usage()
	} else {
		paramsMap := map[string]interface{}{}
		inputParams := g_params.NewInputParams()
		for k,v := range flag.Args() {
			paramsMap[inputParams.GetInputPrefix()+strconv.Itoa(k)] = v
		}
		inputParams.FileList = []string{p.f}
		inputParams.InputNum = flag.NArg()
		inputParams.ParamsMap = paramsMap

		modules.Run(g_params.CommandParams{
			Module: p.module,
			Task:   p.task,
		}, inputParams)
	}
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `graffito version: 0.0.0
用法: go run main.go [-m module] [-t task] [-f filepath] [-h] [...]

选项:

`)
	if err != nil {
		fmt.Println("Unknown Error!")
	} else {
		flag.PrintDefaults()
	}
}


