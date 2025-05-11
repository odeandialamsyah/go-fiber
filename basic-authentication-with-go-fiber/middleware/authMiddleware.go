package middleware

import (
	"basic-authentication-with-go-fiber/models"
	"go-mvc-auth/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware memverifikasi JWT token di header Authorization
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	}

	token := strings.Split(authHeader, "Bearer ")[1]
	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}

	c.Locals("userID", claims["sub"])
	return c.Next()
}
