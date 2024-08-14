package main

import (
	"api/internal/matrix"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"success": true,
			"message": "Server is running",
			"status":  200,
		})
	})

	app.Post("/api/matrix/rotate", matrix.RotateMatrixHandler)

	log.Fatal(app.Listen(":" + port))
}
