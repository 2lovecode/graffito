package top100

import "fmt"

func orangesRotting(grid [][]int) int {
	println("-----------")
	rows := len(grid)
	cols := len(grid[0])

	bfsQueue := make([][]int, 0)
	level := 0

	freshOrangeCnt := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				freshOrangeCnt++
			} else if grid[i][j] == 2 {
				bfsQueue = append(bfsQueue, []int{i, j})
			}
		}
	}
	if freshOrangeCnt == 0 {
		return 0
	}
	for len(bfsQueue) > 0 {
		oldL := len(bfsQueue)
		for i := 0; i < oldL; i++ {
			n, m := bfsQueue[i][0], bfsQueue[i][1]
			if grid[n][m] == 1 {
				freshOrangeCnt--
				grid[n][m] = 2
			}
			if (n+1) < rows && grid[n+1][m] == 1 {
				bfsQueue = append(bfsQueue, []int{n + 1, m})
			}
			if (n-1) >= 0 && grid[n-1][m] == 1 {
				bfsQueue = append(bfsQueue, []int{n - 1, m})
			}
			if (m+1) < cols && grid[n][m+1] == 1 {
				bfsQueue = append(bfsQueue, []int{n, m + 1})
			}
			if (m-1) >= 0 && grid[n][m-1] == 1 {
				bfsQueue = append(bfsQueue, []int{n, m - 1})
			}
		}
		fmt.Println(bfsQueue)
		level++
		bfsQueue = bfsQueue[oldL:]
	}
	if freshOrangeCnt > 0 {
		return -1
	}

	return level - 1
}
