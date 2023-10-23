package middlewares

import (
	"Auth/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IsAuthorized() fiber.Handler {
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

		asd := c.Locals(claims.Role, claims.StandardClaims)

		fmt.Println(asd)
		return c.Next()
	}
}
