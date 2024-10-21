package top100

// 48. 旋转图像
func rotate48(matrix [][]int) {
	rows := len(matrix)

	l := rows / 2
	current := l
	for current > 0 {
		for i := l - current; i < rows-1-l+current; i++ {
			matrix[l-current][i], matrix[i][rows-1-l+current] = matrix[i][rows-1-l+current], matrix[l-current][i]
			matrix[l-current][i], matrix[rows-1-l+current][rows-1-i] = matrix[rows-1-l+current][rows-1-i], matrix[l-current][i]
			matrix[l-current][i], matrix[rows-1-i][l-current] = matrix[rows-1-i][l-current], matrix[l-current][i]
		}
		current--
	}

}
