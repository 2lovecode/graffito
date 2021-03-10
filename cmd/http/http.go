package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type params struct {
	help bool
	client bool
}


var p params

func init() {
	flag.BoolVar(&p.help, "h",false, "帮助")
	flag.BoolVar(&p.client,"c", false, "客户端模式")
	flag.Usage = usage
}

func usage() {
	_, err := fmt.Fprintf(os.Stderr, `http version: 0.0.0
用法: go run http.go [-c client] [-h]

`)
	if err != nil {
		fmt.Println("Unknown Error!")
	} else {
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if p.help {
		flag.Usage()
	} else {
		if p.client {
			httpClientRun()
		} else {
			httpServerRun()
		}
	}
}
func httpClientRun() {
	res, err := http.Get("http://localhost.charlesproxy.com:8080/")
	if err != nil {
		fmt.Println("error")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("response error")
	}
	fmt.Println(string(body))
}

func httpServerRun() {
	http.HandleFunc("/", helloWorldHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("error")
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Fprintf(w,"Hello---World!")
}
