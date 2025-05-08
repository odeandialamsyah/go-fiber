package main

import (
	"crud-mvc-with-go-fiber/config"
	"crud-mvc-with-go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	routes.UserRoute(app)

	app.Listen(":3000")
}