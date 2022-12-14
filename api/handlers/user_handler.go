package handlers

import (
	"net/http"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/adonism2k/golang-hactiv8/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary 	Login
// @Description Authenticates user and returns JWT token
// @Tags        User
// @Accept      json
// @Produce     json
// @Param      	request body handlers.Login.Request true "Login Request"
// @Success     200 {object} handlers.Login.Response "Success"
// @Router      /users/login [post]
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
			"error":   true,
			"message": "cannot parse json",
		})
	}

	// Validate the request
	errors := utils.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user := h.Models.User.FindByEmail(body.Email)
	if user.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Invalid email or Password"})
	}

	if err := utils.VerifyPassword(user.Password, body.Password); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Invalid email or Password"})
	}

	token, err := utils.GenerateToken(user, time.Duration(h.Env.JWTExpired), h.Env.JWTSecret)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(Response{token})
}

// Register godoc
// @Summary 	Register
// @Description Authenticates user and returns JWT token
// @Tags        User
// @Accept      json
// @Produce     json
// @Param      	request body handlers.Register.Request true "Register Request"
// @Success     200 {object} handlers.Register.Response "Success"
// @Router      /users/register [post]
func (h *Config) Register(c *fiber.Ctx) error {
	// RegisterRequest Model godoc
	// @Description RegisterRequest Model
	type Request struct {
		Age      int    `json:"age" example:"18" validate:"required,number,gte=8"`
		Username string `json:"username" example:"adnsm" validate:"required"`
		Email    string `json:"email" example:"abdianrizky11@gmail.com" validate:"required,email"`
		Password string `json:"password" example:"bcrypt hashed password" validate:"required,min=6,max=32"`
	} // @name RegisterRequest

	// RegisterResponse Model godoc
	// @Description RegisterResponse Model
	type Response struct {
		ID       int    `json:"id" example:"1"`
		Age      int    `json:"age" example:"18"`
		Username string `json:"username" example:"adnsm"`
		Email    string `json:"email" example:"abdianrizky11@gmail.com"`
	} // @name RegisterResponse

	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "cannot parse json"})
	}

	// Validate the request
	errors := utils.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Check if there is an existing user with the same email
	user := h.Models.User.FindByEmail(body.Email)
	if user.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "email already exists"})
	}

	// Check if there is an existing user with the same username
	user = h.Models.User.FindByUsername(body.Username)
	if user.ID != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "username already exists"})
	}

	// Hash the password
	password, err := utils.HashPassword(body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	// Create the user
	var newUser model.User
	newUser.Age = body.Age
	newUser.Username = body.Username
	newUser.Email = body.Email
	newUser.Password = password

	// Save the user
	user = h.Models.User.Create(newUser)

	return c.Status(http.StatusOK).JSON(Response{
		ID:       user.ID,
		Age:      user.Age,
		Username: user.Username,
		Email:    user.Email,
	})
}

// Update User godoc
// @Summary 	Update User
// @Description Update the current user data
// @Tags        User
// @Accept      json
// @Produce     json
// @Param		auth header string true "Authorization"
// @Param		id path int true "User ID"
// @Param		request body handlers.UpdateUser.Request true "Update User Request"
// @Security    ApiKeyAuth
// @Success     200 {object} handlers.UpdateUser.Response "Success"
// @Router      /users/{id} [put]
func (h *Config) UpdateUser(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	// UpdateUserRequest Model godoc
	// @Description UpdateUserRequest Model
	type Request struct {
		Username string `json:"username" example:"adnsm" validate:"required"`
		Email    string `json:"email" example:"abdianrizky11@gmail.com" validate:"required,email,min=6,max=32"`
	} // @name UpdateUserRequest

	// UpdateUserResponse Model godoc
	// @Description UpdateUserResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Username  string    `json:"username" example:"adnsm"`
		Age       int       `json:"age" example:"18"`
		Email     string    `json:"email" example:"abdianrizky11@gmail.com"`
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name UpdateUserResponse

	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "cannot parse json"})
	}

	// Validate the request
	errors := utils.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var userEdit model.User
	userEdit.Username = body.Username
	userEdit.Email = body.Email

	newUser := h.Models.User.Update(user.ID, userEdit)

	return c.Status(http.StatusOK).JSON(Response{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Age:       newUser.Age,
		Email:     newUser.Email,
		UpdatedAt: newUser.UpdatedAt,
	})
}

// Delete User godoc
// @Summary 	Delete User
// @Description Delete the current user account
// @Tags        User
// @Accept      json
// @Produce     json
// @Param 		id path int true "User ID"
// @Security    ApiKeyAuth
// @Success     200 {object} handlers.DeleteResponse "Success"
// @Router      /users/{id} [delete]
func (h *Config) DeleteUser(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	if err := h.Models.User.Delete(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": true, "message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(DeleteResponse{Message: "your account has been successfully deleted"})
}
