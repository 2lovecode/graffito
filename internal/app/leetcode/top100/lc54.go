package top100

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	upRows := 0
	downRows := 0
	leftCols := 0
	rightCols := 0
	rowsL := len(matrix)
	colsL := len(matrix[0])

	face := 1
	res := make([]int, 0, rowsL+colsL)
	i, j := 0, 0
	cnt := 0
	for cnt < rowsL*colsL {
		res = append(res, matrix[i][j])
		cnt++
		if face == 1 {
			if j >= colsL-rightCols-1 {
				face = 2
				i++
				upRows++
			} else {
				j++
			}
		} else if face == 2 {
			if i >= rowsL-downRows-1 {
				face = 3
				j--
				rightCols++
			} else {
				i++
			}

		} else if face == 3 {
			if j <= leftCols {
				face = 4
				i--
				downRows++
			} else {
				j--
			}
		} else if face == 4 {
			if i <= upRows {
				face = 1
				j++
				leftCols++
			} else {
				i--
			}
		}
	}
	return res
}
