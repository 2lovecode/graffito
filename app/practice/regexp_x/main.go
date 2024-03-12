package regexp_x

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "reg", Run: func(cmd *cobra.Command, args []string) {
		Run()
	}}
}
func Run() {
	reg, err := regexp.Compile(`\${.*?}`)

	if err == nil {
		fmt.Println(string(reg.ReplaceAll([]byte("adfaafd${aaa},cccd${ccsss}d"), []byte("1234"))))
	}

}
