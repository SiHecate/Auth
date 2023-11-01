package middlewares

import (
	"Auth/database"
	"Auth/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Authorization() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		role := c.Locals("Role")
		id := c.Locals("Issuer")

		var user models.User
		database.Conn.Model(&user).Where("id = ?", id).First(&user)

		if role == "User" {
			fmt.Println("User role: " + user.Role)
			fmt.Println("Username: " + user.Name)
			fmt.Println("User email: " + user.Email)
		}

		if role == "Admin" {
			fmt.Println("Admin role: " + user.Role)
			fmt.Println("Admin username: " + user.Name)
			fmt.Println("Admin email: " + user.Email)
		}

		return c.Next()
	}
}
