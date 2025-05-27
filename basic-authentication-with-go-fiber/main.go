package main

import (
	"basic-authentication-with-go-fiber/config"
	"basic-authentication-with-go-fiber/repositories"
	"basic-authentication-with-go-fiber/routes"
	"basic-authentication-with-go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Menghubungkan ke MongoDB
	config.ConnectDB()

	// Initialize repositories
	repositories.InitCollections()

	// Seed role default dan admin user jika belum ada
	utils.SeedRole(config.DB)
	utils.SeedAdmin(config.DB)

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Add CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		// Handle preflight requests
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	})

	// Add API routes
	routes.UserRoutes(app)

	// Add 404 handler for undefined routes
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error": "Route not found",
		})
	})

	// Menjalankan server di port 3000	log.Fatal(app.Listen(":3000"))
}