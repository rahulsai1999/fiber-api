package main

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	app := fiber.New(&fiber.Settings{
		Prefork: true,
	})

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Hello and Welcome")
	})

	app.Listen(port)
}
