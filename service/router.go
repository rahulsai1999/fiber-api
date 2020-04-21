package service

import (
	"github.com/gofiber/compression"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/rahulsai1999/fiber-api/service/api"
)

// Router -> returns app object for the server
func Router() *fiber.App {
	app := fiber.New(&fiber.Settings{
		Prefork: true,
	})

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(compression.New())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Status(200).JSON(fiber.Map{
			"message": "API Health OK",
		})
	})

	app.Get("/ping", api.Ping)
	app.Get("/blog", api.GetAllBlogs)
	app.Post("/blog", api.InsertBlog)
	app.Get("/blog/:id", api.GetBlog)
	app.Put("/blog/:id", api.UpdateBlog)
	app.Delete("/blog/:id", api.DeleteBlog)

	return app
}
