package main

import (
	"basic-authentication-with-go-fiber/config"
	"basic-authentication-with-go-fiber/routes"
	"basic-authentication-with-go-fiber/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Menghubungkan ke MongoDB
	config.ConnectDB()

	// Seed role default jika belum ada
	utils.SeedRole(config.DB)

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Menambahkan routes
	routes.UserRoutes(app)

	// Menjalankan server di port 3000
	log.Fatal(app.Listen(":3000"))
}