package algorithm

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewClimbStair() *cobra.Command {
	return &cobra.Command{
		Use: "climb-stair",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("爬楼梯问题: 爬n阶楼梯，每次只能爬1或2阶，问爬完楼梯有多少种方式？")
			exampleFloors := []int{1, 2, 3, 5, 10, 12}
			for _, eachFloor := range exampleFloors {
				fmt.Println("爬一个", eachFloor, "层的楼梯:", climb(eachFloor))
			}
		},
	}
}
func climb(n int) int {
	//
	if n == 2 {
		return 2
	}

	if n == 1 {
		return 1
	}

	return climb(n-1) + climb(n-2)

}
