package routes

import(
	"basic-authentication-with-go-fiber/controllers"
	"basic-authentication-with-go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Register routes
	api.Post("/register", controllers.RegisterUser)
	api.Post("/login", controllers.LoginUser)
}