package main

import (
	"api/internal/config"
	"api/internal/matrix"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()

	app.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"success": true,
			"message": "Server is running",
			"status":  200,
		})
	})

	app.Post("/api/matrix/rotate", matrix.RotateMatrixHandler)

	log.Fatal(app.Listen(cfg.Port))
}
