package middleware

import (
	"strconv"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/gofiber/fiber/v2"
)

func UserOwner(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	params, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid id",
		})
	}

	if user.ID == params {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error":   true,
		"message": "You doesn't have access to this resources",
	})
}

func PhotoOwner(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	params, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var photos []model.Photo
	for _, v := range photos {
		if v.ID == params && v.UserID == user.ID {
			return c.Next()
		} else if v.ID == params && v.UserID != user.ID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You doesn't have access to this resources",
			})
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}

func CommentOwner(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	params, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var comments []model.Comment
	for _, v := range comments {
		if v.ID == params && v.UserID == user.ID {
			return c.Next()
		} else if v.ID == params && v.UserID != user.ID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You doesn't have access to this resources",
			})
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}

func SocialOwner(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	params, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var socials []model.SocialMedia
	for _, v := range socials {
		if v.ID == params && v.UserID == user.ID {
			return c.Next()
		} else if v.ID == params && v.UserID != user.ID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You doesn't have access to this resources",
			})
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
