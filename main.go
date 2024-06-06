package main

import (
	"fiber-go/app/routes"
	"fiber-go/database"
	_ "fiber-go/docs"
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	fiberSwagger "github.com/swaggo/fiber-swagger" // fiber-swagger middleware
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample server API with fiber server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description use `Bearer <xx token xx>` to authenticate
func main() {
	// database.InitDB()
	database.Conn()

	cmd := exec.Command("/bin/sh", "-c", "make generate_docs;")
	err := cmd.Run()

	if err != nil {
		fmt.Println("command fail")
	} else {
		fmt.Println("command success")
	}

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "fiber-go v1.0.1",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		// AllowOriginsFunc: func(origin string) bool {
		// 	return os.Getenv("ENVIRONMENT") == "development"
		// },
		AllowCredentials: false,
	}))

	app.Use(helmet.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.UserRoutes(app)
	routes.AuthRoutes(app)

	app.Listen(":3000")
}
