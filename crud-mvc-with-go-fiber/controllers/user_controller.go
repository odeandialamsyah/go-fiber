package controllers

import (
	"crud-mvc-with-go-fiber/models"
	"crud-mvc-with-go-fiber/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := repositories.CreateUser(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := repositories.UpdateUser(id, user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("User updated successfully")
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repositories.DeleteUser(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString("User deleted successfully")
}
