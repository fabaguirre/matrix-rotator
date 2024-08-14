package matrix

import (
	"github.com/gofiber/fiber/v2"
)

func validateMatrix(matrix [][]int) error {
	length := len(matrix[0])
	if length == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Matrix cannot be empty")
	}

	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) != length {
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
