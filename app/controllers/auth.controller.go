package controllers

import (
	"fiber-go/app/models"
	"fiber-go/app/repositories"
	"fiber-go/lib"
	"fiber-go/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Login godoc
// @Summary      Login User
// @Description  Login bro
// @Tags         Login
// @Accept       json
// @Produce      json
// @Param data body models.Credentials true "The input todo struct"
// @Success      200  {object}  models.DataResponseOk
// @Failure      400  {object}  models.ResponseStatus
// @Failure      404  {object}  models.ResponseStatus
// @Failure      500  {object}  models.ResponseStatus
// @Router       /login [post]
func Login(c *fiber.Ctx) error {
	user := models.User{}
	c.BodyParser(&user)

	data, err := repositories.Login(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	password := lib.CheckPasswordHash(user.Password, data.Password)
	if !password {
		return response.Unauthorized(c)
	}

	t, err := lib.GenerateToken(data.Name, data.ID)
	if err != nil {
		return response.InternalServerError(c)
	}

	return response.Success(c, fiber.Map{"username": data.Username, "token": t})
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
