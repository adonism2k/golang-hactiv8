package handlers

import (
	"strconv"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/adonism2k/golang-hactiv8/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Get All Photos godoc
// @Summary 	Get All Photos
// @Description Get All Photos
// @Tags        Photo
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Success     200 {array} model.Photo "Success"
// @Router      /photos/ [get]
func (h *Config) GetPhotos(c *fiber.Ctx) error {
	photos, err := h.Models.Photo.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(photos)
}

// Create Photo godoc
// @Summary 	Create Photo
// @Description Create a Photo
// @Tags        Photo
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Param		request body model.PhotoRequest true "Create Photo Request"
// @Success     200 {object} handlers.CreatePhoto.Response "Success"
// @Router      /photos/ [post]
func (h *Config) CreatePhoto(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)

	// CreatePhotoResponse Model godoc
	// @Description CreatePhotoResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Title     string    `json:"title" example:"First Photo"`
		Caption   string    `json:"caption" example:"my first photo"`
		Url       string    `json:"photo_url" example:"https://images.unsplash.com"`
		CreatedAt time.Time `json:"created_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name CreatePhotoResponse

	var body model.PhotoRequest
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

	var newPhoto model.Photo
	newPhoto.Title = body.Title
	newPhoto.Caption = body.Caption
	newPhoto.Url = body.Url
	newPhoto.UserID = user.ID

	photo, err := h.Models.Photo.Create(newPhoto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(Response{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Url:       photo.Url,
		CreatedAt: photo.CreatedAt,
	})
}

// Update Photo godoc
// @Summary 	Update Photo
// @Description Update the current Photo data
// @Tags        Photo
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Param		id path int true "Photo ID"
// @Param		request body model.PhotoRequest true "Update Photo Request"
// @Success     200 {object} handlers.UpdatePhoto.Response "Success"
// @Router      /photos/{id} [put]
func (h *Config) UpdatePhoto(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	// UpdatePhotoResponse Model godoc
	// @Description UpdatePhotoResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Title     string    `json:"title" example:"First Photo"`
		Caption   string    `json:"caption" example:"my first photo"`
		Url       string    `json:"photo_url" example:"https://images.unsplash.com"`
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name UpdatePhotoResponse

	// check if the owner of the photo is the user
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "id must be an integer",
		})
	}

	photo, err := h.Models.Photo.Find(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if photo.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "You doesn't have access to this resources",
		})
	}

	var body model.PhotoRequest
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

	var newPhoto model.Photo
	newPhoto.Title = body.Title
	newPhoto.Caption = body.Caption
	newPhoto.Url = body.Url

	photo, err = h.Models.Photo.Update(id, newPhoto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		Url:       photo.Url,
		UpdatedAt: photo.UpdatedAt,
	})
}

// Delete Photo godoc
// @Summary 	Delete Photo
// @Description Delete a photo
// @Tags        Photo
// @Accept      json
// @Produce     json
// @Param 		id path int true "Photo ID"
// @Security    ApiKeyAuth
// @Success     200 {object} handlers.DeleteResponse "Success"
// @Router      /photos/{id} [delete]
func (h *Config) DeletePhoto(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "id must be an integer",
		})
	}

	photo, err := h.Models.Photo.Find(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if photo.UserID != user.ID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "You doesn't have access to this resources",
		})
	}

	if err := h.Models.Photo.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete photo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(DeleteResponse{Message: "Photo has been successfully deleted"})
}
