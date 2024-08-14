package matrix

func RotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[j][n-1-i] = matrix[i][j]
		}
	}
	return result
}
