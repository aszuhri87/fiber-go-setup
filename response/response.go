package response

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseOk(c *fiber.Ctx, data any) error {
	c.SendStatus(200)

	return c.JSON(fiber.Map{
		"data":    data,
		"status":  200,
		"message": "Success",
	})
}

func ResponseBadResponse(c *fiber.Ctx, message string) error {
	c.SendStatus(400)

	return c.JSON(fiber.Map{
		"status":  400,
		"message": message,
	})
}

func ResponseInternalServerError(c *fiber.Ctx) error {
	c.SendStatus(500)

	return c.JSON(fiber.Map{
		"status":  500,
		"message": "Server Error",
	})
}

func ResponseUnauthorized(c *fiber.Ctx) error {
	c.SendStatus(401)

	return c.JSON(fiber.Map{
		"status":  401,
		"message": "Unauthorized",
	})
}
