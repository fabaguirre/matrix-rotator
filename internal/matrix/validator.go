package matrix

import (
	"github.com/gofiber/fiber/v2"
)

// func isSquareMatrix(matrix [][]int) bool {
// 	n := len(matrix)
// 	for _, row := range matrix {
// 		if len(row) != n {
// 			return false
// 		}
// 	}
// 	return true
// }

func ParseAndValidateMatrixRequest(c *fiber.Ctx) (*MatrixRequest, error) {
	var req MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	// if !isSquareMatrix(req.Matrix) {
	// 	return nil, fiber.NewError(fiber.StatusBadRequest, "Matrix must be square")
	// }

	return &req, nil
}
