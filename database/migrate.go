package database

import (
	"Auth/models"

	"github.com/gofiber/fiber/v2"
)

func MigrateDB(c *fiber.Ctx) error {
	err := Conn.AutoMigrate(&models.User{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Migration error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Migration completed"})
}
