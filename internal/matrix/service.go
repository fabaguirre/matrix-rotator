package matrix

import (
	"gonum.org/v1/gonum/mat"
)

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

func QRDecomposition(matrix [][]int) (
	*mat.Dense, *mat.Dense, error,
) {
	rows := len(matrix)
	cols := len(matrix[0])
	data := make([]float64, rows*cols)

	for i := range matrix {
		for j := range matrix[i] {
			data[i*cols+j] = float64(matrix[i][j])
		}
	}

	A := mat.NewDense(rows, cols, data)

	var qr mat.QR
	qr.Factorize(A)

	var Q *mat.Dense = new(mat.Dense)
	var R *mat.Dense = new(mat.Dense)

	qr.QTo(Q)
	qr.RTo(R)

	return Q, R, nil
}
