package controllers

import (
	"fiber-go/app/models"
	"fiber-go/app/repositories"
	"fiber-go/lib"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	data, err := repositories.Login(user)
	if err != nil {
		return c.JSON(fiber.StatusInternalServerError, err.Error())
	}

	password := lib.CheckPasswordHash(user.Password, data.Password)

	if !password {
		return c.JSON(fiber.StatusUnauthorized)
	}

	// Generate encoded token and send it as response.
	t, err := lib.GenerateToken(data.Name, data.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
