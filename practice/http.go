package practice

import (
	"fmt"
	"graffito/utils/params"
	"io/ioutil"
	"net/http"
)

func HttpRun(p params.InputParamsInterface) {
	key := p.GetInputPrefix() + "0"
	input := p.GetString(key)
	switch input {
	case "client":
		httpClientRun()
	case "server":
		httpServerRun()
	default:
		fmt.Println("请输入参数 client/server")
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
