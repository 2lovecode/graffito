package top100

// 73. Set Matrix Zeroes
func setZeroes(matrix [][]int) {
	zeroRows := make([]int, 0)
	zeroCols := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroRows = append(zeroRows, i)
				zeroCols = append(zeroCols, j)
			}
		}
	}

	for _, row := range zeroRows {
		for j := 0; j < len(matrix[row]); j++ {
			matrix[row][j] = 0
		}
	}
	for _, col := range zeroCols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
}
