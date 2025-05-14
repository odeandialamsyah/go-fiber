package routes

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Register routes
	api.Post("/register", controllers.RegisterUser)
	api.Post("/login", controllers.LoginUser)
}