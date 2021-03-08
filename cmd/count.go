package main

import (
	"flag"
	"fmt"
	"graffito/tools/count"
	"os"
)

type params struct {
	help  bool
	input string
}

var p params

func init() {
	flag.BoolVar(&p.help, "h", false, "帮助")
	flag.StringVar(&p.input, "i", "input", "输入值")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if p.help {
		flag.Usage()
	} else {
		fmt.Println("原字符串为 : ", p.input)
		fmt.Println("字符数为 : ", count.Count(p.input))
	}
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `count version: 0.0.0
用法: go run count.go [-i input] [-h]

`)
	if err != nil {
		fmt.Println("Unknown Error!")
	} else {
		flag.PrintDefaults()
	}
}
