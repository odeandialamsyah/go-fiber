package controllers

import (
	"basic-authentication-with-go-fiber/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAdminStats returns statistics for admin dashboard
func GetAdminStats(c *fiber.Ctx) error {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get users"})
	}

	roles, err := repositories.GetAllRoles()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get roles"})
	}

	// Count active users (all users are considered active for now)
	activeUsers := len(users)

	return c.JSON(fiber.Map{
		"totalUsers": len(users),
		"activeUsers": activeUsers,
		"totalRoles": len(roles),
	})
}
