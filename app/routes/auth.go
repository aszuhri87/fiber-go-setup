package routes

import (
	"fiber-go/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) *fiber.App {
	// api:= app.Group("/api")

	app.Post("/login", controllers.Login)

	return app
}
