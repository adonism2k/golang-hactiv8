package handlers

import (
	"net/http"
	"strings"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/adonism2k/golang-hactiv8/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Config) Login(c *fiber.Ctx) error {
	// LoginRequest Model godoc
	// @Description LoginRequest Model
	type Request struct {
		Email    string `json:"email" example:"abdianrizky11@gmail.com" validate:"required,email,min=6,max=32"`
		Password string `json:"password" example:"bcrypt hashed password" validate:"required,min=6,max=32"`
	} // @name LoginRequest

	// LoginResponse Model godoc
	// @Description LoginResponse Model
	type Response struct {
		Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	} // @name LoginResponse

	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}

	// Validate the request
	errors := utils.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	var user model.User
	result := h.DB.DB.First(&user, "email = ?", strings.ToLower(body.Email))
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	if err := utils.VerifyPassword(user.Password, body.Password); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
}

func (h *Config) Register(c *fiber.Ctx) error {
	return nil
}

func (h *Config) EditUser(c *fiber.Ctx) error {
	return nil
}

func (h *Config) DeleteUser(c *fiber.Ctx) error {
	return nil
}
