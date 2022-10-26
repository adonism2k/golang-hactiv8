package handlers

import (
	"strconv"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/adonism2k/golang-hactiv8/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Get All Social Media godoc
// @Summary 	Get All Social Media
// @Description Get the current Social Media data
// @Tags        Social Media
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Success     200 {array} model.SocialMedia "Success"
// @Router      /socialmedias/ [get]
func (h *Config) GetSocialMedias(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	sm, err := h.Models.SocialMedia.All(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sm)
}

// Create Social Media godoc
// @Summary 	Create Social Media
// @Description Create a Social Media
// @Tags        Social Media
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Param		request body model.SocialMediaRequest true "Create Social Media Request"
// @Success     200 {object} handlers.CreateSocialMedia.Response "Success"
// @Router      /socialmedias/ [post]
func (h *Config) CreateSocialMedia(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	// CreateSocialMediaResponse Model godoc
	// @Description CreateSocialMediaResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Name      string    `json:"name" example:"Instagram"`
		Url       string    `json:"url" example:"https://instagram.com"`
		UserID    int       `json:"user_id" example:"1"`
		CreatedAt time.Time `json:"created_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name CreateSocialMediaResponse

	var body model.SocialMediaRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var sm model.SocialMedia
	sm.Name = body.Name
	sm.Url = body.Url
	sm.UserID = user.ID

	social, err := h.Models.SocialMedia.Create(sm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		ID:        social.ID,
		Name:      social.Name,
		Url:       social.Url,
		UserID:    social.UserID,
		CreatedAt: social.CreatedAt,
	})
}

// Update Social Media godoc
// @Summary 	Update Social Media
// @Description Update the current Social Media data
// @Tags        Social Media
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Param		id path int true "Social Media ID"
// @Param		request body model.SocialMediaRequest true "Update Social Media Request"
// @Success     200 {object} handlers.UpdateSocialMedia.Response "Success"
// @Router      /socialmedias/{id} [put]
func (h *Config) UpdateSocialMedia(c *fiber.Ctx) error {
	// UpdateSocialMediaResponse Model godoc
	// @Description UpdateSocialMediaResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Name      string    `json:"name" example:"Instagram"`
		Url       string    `json:"social_media_url" example:"https://instagram.com"`
		UserID    int       `json:"user_id" example:"1"`
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name UpdateSocialMediaResponse

	user := c.Locals("user").(model.User)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	photo, err := h.Models.SocialMedia.Find(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if photo.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Forbidden",
		})
	}

	var body model.SocialMediaRequest
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var sm model.SocialMedia
	sm.Name = body.Name
	sm.Url = body.Url
	sm.UserID = user.ID

	social, err := h.Models.SocialMedia.Update(id, sm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		ID:        social.ID,
		Name:      social.Name,
		Url:       social.Url,
		UserID:    social.UserID,
		UpdatedAt: social.UpdatedAt,
	})
}

// Delete Social Media godoc
// @Summary 	Delete Social Media
// @Description Delete a Social Media
// @Tags        Social Media
// @Accept      json
// @Produce     json
// @Param 		id path int true "Social Media ID"
// @Security    ApiKeyAuth
// @Success     200 {object} handlers.DeleteResponse "Success"
// @Router      /socialmedias/{id} [delete]
func (h *Config) DeleteSocialMedia(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	sm, err := h.Models.SocialMedia.Find(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if sm.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Forbidden",
		})
	}

	err = h.Models.SocialMedia.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(DeleteResponse{Message: "Your social media has been successfully deleted"})
}
