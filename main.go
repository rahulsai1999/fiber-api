package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rahulsai1999/fiber-api/service"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	app := service.Router()
	app.Listen(port)
}
