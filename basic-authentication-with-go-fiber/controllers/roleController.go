package controllers

import (
	"basic-authentication-with-go-fiber/repositories"
	"github.com/gofiber/fiber/v2"
)

// GetAllRoles returns all roles with user count
func GetAllRoles(c *fiber.Ctx) error {
	roles, err := repositories.GetAllRoles()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get roles"})
	}

	// Get user count for each role
	roleResponses := make([]fiber.Map, len(roles))
	for i, role := range roles {
		userCount, err := repositories.GetUserCountByRole(role.ID)
		if err != nil {
			userCount = 0
		}

		roleResponses[i] = fiber.Map{
			"id": role.ID.Hex(),
			"name": role.Name,
			"usersCount": userCount,
		}
	}

	return c.JSON(roleResponses)
}

// DeleteRole deletes a role if it has no users
func DeleteRole(c *fiber.Ctx) error {
	roleID := c.Params("id")
	userCount, err := repositories.GetUserCountByRoleID(roleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to check role users"})
	}

	if userCount > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot delete role with active users"})
	}

	if err := repositories.DeleteRole(roleID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete role"})
	}

	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}
