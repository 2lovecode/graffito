package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
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
