package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/status", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"success": true,
			"message": "Server is running",
			"status":  200,
		})
	})

	app.Listen(":4000")
}
