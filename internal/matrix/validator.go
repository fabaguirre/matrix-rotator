package matrix

import (
	"github.com/gofiber/fiber/v2"
)

func validateMatrix(matrix [][]int) error {
	if len(matrix) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Matrix cannot be empty")
	}
	rowsLength := len(matrix)
	colsLength := len(matrix[0])

	if rowsLength == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Matrix rows cannot be empty")
	}

	for i := 0; i < rowsLength; i++ {
		if len(matrix[i]) != colsLength {
			return fiber.NewError(fiber.StatusBadRequest, "Matrix rows must have the same length")
		}
	}

	return nil
}

func ParseAndValidateMatrixRequest(c *fiber.Ctx) (*MatrixRequest, error) {
	var req MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := validateMatrix(req.Matrix); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return &req, nil
}
