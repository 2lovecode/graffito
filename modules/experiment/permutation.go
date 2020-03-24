package experiment

import "fmt"

func PermutationRun() {
	NumList := []int{2, 3, 4, 5, 7}
	pp(NumList)
}

func pp(numList []int) [][]int{
	data := [][]int{}

	l := len(numList)
	switch l {
	case  2:
		data = [][]int{
			[]int{numList[0], numList[1]},
			[]int{numList[1], numList[0]},
		}
	default:
		for i := 0; i < l; i++ {
			tt := append(numList[:i], numList[i+1:]...)
			tmp := pp(tt)
			for _, v := range tmp {
				data = append(data, [][]int{v}...)
			}
		}
	}
	fmt.Println(data)
	return data
}
