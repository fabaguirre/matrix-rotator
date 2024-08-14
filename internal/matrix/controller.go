package matrix

import (
	client "api/internal/network"

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

	rotatedMatrix := RotateMatrix(req.Matrix)

	stats, err := client.GetMatrixStatistics(rotatedMatrix)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send data to Node.js API"})
	}

	Q, R, err := QRDecomposition(req.Matrix)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to compute QR decomposition"})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data": fiber.Map{
			"qrDecomposition": &fiber.Map{
				"Q": Q.RawMatrix().Data,
				"R": R.RawMatrix().Data,
			},
			"rotatedMatrix": rotatedMatrix,
			"stats":         stats,
		},
		"status": fiber.StatusOK,
	})
}
