package top100

func numIslands(grid [][]byte) int {
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(grid, i, j)
			}
		}
	}
	return cnt
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	if grid[i][j] == '1' {
		grid[i][j] = '2'
		dfs(grid, i+1, j)
		dfs(grid, i, j+1)
		dfs(grid, i-1, j)
		dfs(grid, i, j-1)
	}
}
