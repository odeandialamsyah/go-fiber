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

