package middleware

import (
	"strconv"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/gofiber/fiber/v2"
)

func UserOwner(c *fiber.Ctx) error {
	u := c.Locals("user").(model.User)
	params, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid id",
		})
	}

	if u.ID == params {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": "You doesn't have access to this resources",
	})
}
