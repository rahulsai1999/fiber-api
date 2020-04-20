package api

import "github.com/gofiber/fiber"

// Ping Custom message
func Ping(ctx *fiber.Ctx) {
	ctx.Status(200).JSON(fiber.Map{
		"message": "Pong",
	})
}
