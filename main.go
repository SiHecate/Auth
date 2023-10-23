package main

import (
	"Auth/database"
	"Auth/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Println("Hello World")
	database.Connect()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	routes.Router(app)
	app.Listen(":8080")
}
