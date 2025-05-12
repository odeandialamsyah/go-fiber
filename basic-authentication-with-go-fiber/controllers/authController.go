package controllers

import (
	"basic-authentication-with-go-fiber/models"
	"basic-authentication-with-go-fiber/repositories"
	"basic-authentication-with-go-fiber/utils"
	"basic-authentication-with-go-fiber/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Fungsi register user baru
func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(500).SendString("Failed to hash password")
	}

	user.Password = hashedPassword

	// Insert ke DB
	err = repositories.CreateUser(user)
	if err != nil {
		return c.Status(500).SendString("Failed to create user")
	}

	return c.Status(201).SendString("User registered successfully")
}

// Fungsi login user
func LoginUser(c *fiber.Ctx) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Cek apakah email ada
	user, err := repositories.GetUserByEmail(loginData.Email)
	if err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Verifikasi password
	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(500).SendString("Failed to generate JWT")
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

// Fungsi dashboard admin (contoh protected route)
func AdminDashboard(c *fiber.Ctx) error {
	return c.SendString("Welcome to the admin dashboard!")
}
