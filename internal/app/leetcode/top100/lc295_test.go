package top100

import (
	"fmt"
	"testing"
)

func Test_295(t *testing.T) {
	medianFinder := Constructor295()
	medianFinder.AddNum(1)                 // arr = [1]
	medianFinder.AddNum(2)                 // arr = [1, 2]
	fmt.Println(medianFinder.FindMedian()) // 返回 1.5 ((1 + 2) / 2)
	medianFinder.AddNum(3)                 // arr[1, 2, 3]
	fmt.Println(medianFinder.FindMedian()) // return 2.0
}
