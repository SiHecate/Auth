package routes

import (
	"Auth/controllers"
	"Auth/database"
	"Auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	Database := app.Group("/db")
	Database.Get("/migrate", database.MigrateDB)

	Auth := app.Group("/Auth")
	Auth.Post("/login", controllers.Login)
	Auth.Post("/signup", controllers.Signup)
	Auth.Get("/logout", controllers.Logout)

	Lelouch := app.Group("/Lelouch")
	Lelouch.Use(middlewares.IsAuthorized())
	Lelouch.Use(middlewares.Authorization())
	Lelouch.Get("/Lelouch", func(c *fiber.Ctx) error {
		return c.SendString("Obey to Lelouch Vi Britannia")
	})

}
