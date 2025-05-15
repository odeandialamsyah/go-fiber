package routes

import (
	"basic-authentication-with-go-fiber/controllers"
	"basic-authentication-with-go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Public routes
	api.Post("/register", controllers.RegisterUser)
	api.Post("/login", controllers.LoginUser)

	// Protected routes
	protected := api.Group("/protected")
	protected.Use(middleware.AuthMiddleware)

	// Admin routes
	admin := protected.Group("/admin")
	admin.Use(middleware.RoleMiddleware("admin"))
	// Add admin endpoints
	admin.Get("/stats", controllers.GetAdminStats)
	admin.Get("/users", controllers.GetAllUsers)
	admin.Get("/users/:id", controllers.GetUser)
	admin.Put("/users/:id", controllers.UpdateUser)
	admin.Delete("/users/:id", controllers.DeleteUser)
	admin.Get("/roles", controllers.GetAllRoles)
	admin.Delete("/roles/:id", controllers.DeleteRole)

	// User routes
	user := protected.Group("/user")
	user.Use(middleware.RoleMiddleware("user"))
	// Add user endpoints
	user.Get("/profile", controllers.GetUserProfile)
	user.Put("/profile", controllers.UpdateUserProfile)
	user.Post("/change-password", controllers.ChangePassword)
}