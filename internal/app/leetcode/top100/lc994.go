package top100

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	bfsQueue := make([][]int, 0)
	round := 0

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

	for freshOrangeCnt > 0 && len(bfsQueue) > 0 {
		oldL := len(bfsQueue)
		round++
		for i := 0; i < oldL; i++ {
			n, m := bfsQueue[0][0], bfsQueue[0][1]
			bfsQueue = bfsQueue[1:]
			if (n+1) < rows && grid[n+1][m] == 1 {
				freshOrangeCnt--
				grid[n+1][m] = 2
				bfsQueue = append(bfsQueue, []int{n + 1, m})
			}
			if (n-1) >= 0 && grid[n-1][m] == 1 {
				freshOrangeCnt--
				grid[n-1][m] = 2
				bfsQueue = append(bfsQueue, []int{n - 1, m})
			}
			if (m+1) < cols && grid[n][m+1] == 1 {
				freshOrangeCnt--
				grid[n][m+1] = 2
				bfsQueue = append(bfsQueue, []int{n, m + 1})
			}
			if (m-1) >= 0 && grid[n][m-1] == 1 {
				freshOrangeCnt--
				grid[n][m-1] = 2
				bfsQueue = append(bfsQueue, []int{n, m - 1})
			}
		}
	}
	if freshOrangeCnt > 0 {
		return -1
	}

	return round
}
