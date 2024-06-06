package routes

import (
	"fiber-go/app/controllers"
	"fiber-go/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) *fiber.App {
	api := app.Group("/api", middleware.JwtMiddleware())

	api.Get("/user", controllers.GetUser)
	app.Post("/user", controllers.PostUser)
	app.Get("/user/:id", controllers.GetUserByID)
	app.Put("/user/:id", controllers.PutUser)
	app.Delete("/user/:id", controllers.DeleteUser)
	return app
}
