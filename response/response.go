package response

import (
	"github.com/gofiber/fiber/v2"
)

func Success(c *fiber.Ctx, data any) error {
	c.SendStatus(200)

	return c.JSON(fiber.Map{
		"data":    data,
		"status":  200,
		"message": "Success",
	})
}

func Bad(c *fiber.Ctx, message string) error {
	c.SendStatus(400)

	return c.JSON(fiber.Map{
		"status":  400,
		"message": message,
	})
}

func InternalServerError(c *fiber.Ctx) error {
	c.SendStatus(500)

	return c.JSON(fiber.Map{
		"status":  500,
		"message": "Server Error",
	})
}

func Unauthorized(c *fiber.Ctx) error {
	c.SendStatus(401)

	return c.JSON(fiber.Map{
		"status":  401,
		"message": "Unauthorized",
	})
}
