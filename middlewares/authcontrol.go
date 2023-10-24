package middlewares

import (
	"Auth/utils"

	"github.com/gofiber/fiber/v2"
)

func IsAuthorized() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")

		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("Role", claims.Role)
		c.Locals("Issuer", claims.Issuer)
		c.Locals("Subject", claims.Subject)
		c.Locals("Expire", claims.ExpiresAt)

		return c.Next()
	}
}
