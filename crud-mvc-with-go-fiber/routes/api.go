package routes

import (
	"crud-mvc-with-go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.CreateUser)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)
}