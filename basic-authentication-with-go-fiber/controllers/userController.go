package controllers

import (
	"basic-authentication-with-go-fiber/models"
	"basic-authentication-with-go-fiber/repositories"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllUsers returns all users (admin only)
func GetAllUsers(c *fiber.Ctx) error {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get users"})
	}
	return c.JSON(users)
}

// GetUser returns a specific user by ID (admin only)
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := repositories.GetUserByID(objID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// UpdateUser updates a user's information (admin only)
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var updateData struct {
		Username string             `json:"username"`
		Email    string             `json:"email"`
		RoleID   primitive.ObjectID `json:"role_id"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	user := models.User{
		ID:       objID,
		Username: updateData.Username,
		Email:    updateData.Email,
		RoleID:   updateData.RoleID,
	}

	if err := repositories.UpdateUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

// DeleteUser deletes a user (admin only)
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := repositories.DeleteUser(objID); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

// GetUserProfile returns the current user's profile
func GetUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// UpdateUserProfile updates the current user's profile
func UpdateUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	
	var updateData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	user := models.User{
		ID:       userID,
		Username: updateData.Username,
		Email:    updateData.Email,
	}

	if err := repositories.UpdateUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

// ChangePassword changes the current user's password
func ChangePassword(c *fiber.Ctx) error {
	userID := c.Locals("userID").(primitive.ObjectID)
	
	var passwordData struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}

	if err := c.BodyParser(&passwordData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := repositories.ChangeUserPassword(userID, passwordData.CurrentPassword, passwordData.NewPassword); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}
