package plan9

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "plan9", Run: func(cmd *cobra.Command, args []string) {
		GoWithPlan9()
		Plan9WithGo()
	}}
}

func output(int) (int, int, int)

func GoWithPlan9() {
	a, b, c := output(987654321)
	fmt.Println(a, b, c)
}

func output1(a, b int) int

func add(a, b int) int {
	return a + b
}

func Plan9WithGo() {
	s := output1(10, 30)
	fmt.Println(s)
}
