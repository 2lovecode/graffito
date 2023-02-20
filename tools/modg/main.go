package modg

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

type graph struct {
	root *pkg
}

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use: "modg",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, err := os.Getwd()
			goroot := os.Getenv("GOROOT")
			if err == nil {
				if goroot == "" {
					goroot = pwd
				}
				fmt.Println("****************")
				fmt.Println("Current Dir:", pwd)
				fmt.Println("GoRoot:", goroot)
				fmt.Println("****************")

				goBin := fmt.Sprintf("%s/bin/go", goroot)

				modCmd := exec.Command(goBin, "mod", "graph")
				res, e := modCmd.Output()

				if e == nil {
					reader := bytes.NewReader(res)
					//
					bufReader := bufio.NewReader(reader)
					line, _, _ := bufReader.ReadLine()
					fmt.Println(line)

				}
			}
		},
	}
}

func BuildGraph() {

}
