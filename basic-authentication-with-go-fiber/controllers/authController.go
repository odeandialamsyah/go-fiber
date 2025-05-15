package controllers

import (
	"basic-authentication-with-go-fiber/models"
	"basic-authentication-with-go-fiber/repositories"
	"basic-authentication-with-go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

// Fungsi register user baru
func RegisterUser(c *fiber.Ctx) error {
	// Parse registration data with password confirmation
	var registerData struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate password length
	if len(registerData.Password) < 6 {
		return c.Status(400).JSON(fiber.Map{"error": "Password must be at least 6 characters long"})
	}

	// Validate password confirmation
	if registerData.Password != registerData.ConfirmPassword {
		return c.Status(400).JSON(fiber.Map{"error": "Passwords do not match"})
	}

	// Create user model
	user := models.User{
		Username: registerData.Username,
		Email:    registerData.Email,
		Password: registerData.Password,
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user.Password = hashedPassword

	// Set default role as 'user'
	userRole, err := repositories.GetRoleByName("user")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user role"})
	}
	user.RoleID = userRole.ID
	user.Role = userRole

	// Insert ke DB
	err = repositories.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

// Fungsi login user
func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Cek apakah username ada
	user, err := repositories.GetUserByUsername(loginData.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Verifikasi password
	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate JWT"})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"username": user.Username,
			"email": user.Email,
			"role": user.Role.Name,
		},
	})
}
