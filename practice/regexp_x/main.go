package regexp_x

import (
	"fmt"
	"regexp"
)

func Run() {
	reg, err := regexp.Compile(`\${.*?}`)

	if err == nil {
		fmt.Println(string(reg.ReplaceAll([]byte("adfaafd${aaa},cccd${ccsss}d"), []byte("1234"))))
	}

}
