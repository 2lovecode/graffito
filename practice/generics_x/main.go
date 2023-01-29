package generics_x

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{Use: "generic", Run: func(cmd *cobra.Command, args []string) {
		Run()
	}}
}

type Ele interface {
	int | int16 | int64
}

func index[T comparable](s []T, v T) int {
	for i, each := range s {
		if each == v {
			return i
		}
	}
	return -1
}

func Run() {
	s1 := []int{11, 23, 45, 77}
	fmt.Println(index[int](s1, 77))

	s2 := []string{"a", "b", "c"}
	fmt.Println(index[string](s2, "b"))

}
