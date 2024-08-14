package matrix

import (
	"github.com/gofiber/fiber/v2"
)

func RotateMatrixHandler(c *fiber.Ctx) error {
	req, err := ParseAndValidateMatrixRequest(c)
	if err != nil {
		return c.Status(err.(*fiber.Error).Code).JSON(fiber.Map{
			"success": false,
			"message": err.(*fiber.Error).Message,
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    req.Matrix,
		"status":  fiber.StatusOK,
	})
}
