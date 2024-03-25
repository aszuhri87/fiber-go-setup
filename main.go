package main

import (
	"fiber-go/app/routes"
	"fiber-go/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()
	database.Conn()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "fiber-go v1.0.1",
	})

	// cfg := swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: "./docs/swagger.json",
	// 	Path:     "swagger",
	// 	Title:    "Fiber API Docs",
	// }

	// app.Use(swagger.New(cfg))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.UserRoutes(app)

	app.Listen(":3000")
}
