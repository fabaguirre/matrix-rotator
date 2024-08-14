package matrix

func RotateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][rows-i-1] = matrix[i][j]
		}
	}

	return result
}
