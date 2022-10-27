package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/adonism2k/golang-hactiv8/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Auth func for specify routes group with JWT authentication.
func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		log.Println("🚀 Token is empty")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	token = strings.Replace(token, "Bearer ", "", 1)

	user, err := utils.ValidateToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("user", user)
	return c.Next()
}
