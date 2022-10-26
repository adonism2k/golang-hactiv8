package handlers

import (
	"strconv"
	"time"

	"github.com/adonism2k/golang-hactiv8/internal/model"
	"github.com/adonism2k/golang-hactiv8/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Get All Comments godoc
// @Summary 	Get All Comments
// @Description Get the current Comments data
// @Tags        Comment
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Success     200 {array} model.Comment "Success"
// @Router      /comments/ [get]
func (h *Config) GetComments(c *fiber.Ctx) error {

	type Response struct {
		ID        int       `json:"id" example:"1"`
		Message   string    `json:"message" example:"This is a comment"`
		UserID    int       `json:"user_id" example:"1"`
		PhotoID   int       `json:"photo_id" example:"1"`
		CreateAt  time.Time `json:"create_at" example:"2022-10-10T11:52:28.431369Z"`
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"`
		User      model.User
		Photo     struct {
			ID      int    `json:"id" example:"1"`
			Title   string `json:"title" example:"First Photo"`
			Caption string `json:"caption" example:"my first photo"`
			Url     string `json:"photo_url" example:"https://images.unsplash.com"`
			UserID  int    `json:"user_id" example:"1"`
		}
	}

	com, err := h.Models.Comment.All()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	var response []Response
	for _, v := range com {
		var res Response
		res.ID = v.ID
		res.Message = v.Message
		res.UserID = v.UserID
		res.PhotoID = v.PhotoID
		res.CreateAt = v.CreatedAt
		res.UpdatedAt = v.UpdatedAt
		res.User = v.User
		res.Photo.ID = v.Photo.ID
		res.Photo.Title = v.Photo.Title
		res.Photo.Caption = v.Photo.Caption
		res.Photo.Url = v.Photo.Url
		res.Photo.UserID = v.Photo.UserID
		response = append(response, res)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Create Comment godoc
// @Summary 	Create Comment
// @Description Create a Comment
// @Tags        Comment
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Param		request body handlers.CreateComment.Request true "Create Comment Request"
// @Success     200 {object} handlers.CreateComment.Response "Success"
// @Router      /comments/ [post]
func (h *Config) CreateComment(c *fiber.Ctx) error {
	// Create Comment godoc
	// @Description Create Social Media Request
	type Request struct {
		Message string `json:"message" validate:"required,min=3,max=255"`
		PhotoID int    `json:"photo_id" validate:"required,number,min=1"`
	} // @name CreateCommentRequest

	// Create Comment Response Model godoc
	// @Description Create Comment Response Model
	type Response struct {
		ID       int       `json:"id" example:"1"`
		Message  string    `json:"message" example:"This is a comment"`
		UserID   int       `json:"user_id" example:"1"`
		PhotoID  int       `json:"photo_id" example:"1"`
		CreateAt time.Time `json:"create_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name CreateCommentResponse

	user := c.Locals("user").(model.User)

	var body Request
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

	var comment model.Comment
	comment.Message = body.Message
	comment.PhotoID = body.PhotoID
	comment.UserID = user.ID

	com, err := h.Models.Comment.Create(comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		ID:       com.ID,
		Message:  com.Message,
		UserID:   com.UserID,
		PhotoID:  com.PhotoID,
		CreateAt: com.CreatedAt,
	})
}

// Update Comment godoc
// @Summary 	Update Comment
// @Description Update the current Comment data
// @Tags        Comment
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param		auth header string true "Authorization"
// @Param		id path int true "Comment ID"
// @Param		request body handlers.UpdateComment.Request true "Update Comment Request"
// @Success     200 {object} handlers.UpdateComment.Response "Success"
// @Router      /comments/{id} [put]
func (h *Config) UpdateComment(c *fiber.Ctx) error {
	// Update Comment godoc
	// @Description Update Social Media Request
	type Request struct {
		Message string `json:"message" validate:"required"`
	} // @name UpdateCommentRequest

	// Update Comment Response Model godoc
	// @Description Update Comment Response Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Message   string    `json:"message" example:"This is a comment"`
		UserID    int       `json:"user_id" example:"1"`
		PhotoID   int       `json:"photo_id" example:"1"`
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name UpdateCommentResponse

	user := c.Locals("user").(model.User)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	sm, err := h.Models.Comment.Find(id)
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

	var body Request
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

	var comment model.Comment
	comment.Message = body.Message

	com, err := h.Models.Comment.Update(id, comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		ID:        com.ID,
		Message:   com.Message,
		UserID:    com.UserID,
		PhotoID:   com.PhotoID,
		UpdatedAt: com.UpdatedAt,
	})
}

// Delete Comment godoc
// @Summary 	Delete Comment
// @Description Delete a Comment
// @Tags        Comment
// @Accept      json
// @Produce     json
// @Param 		id path int true "User ID"
// @Security    ApiKeyAuth
// @Success     200 {object} handlers.DeleteResponse "Success"
// @Router      /comments/{id} [delete]
func (h *Config) DeleteComment(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	sm, err := h.Models.Comment.Find(id)
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

	err = h.Models.Comment.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(DeleteResponse{Message: "your comment has been successfully deleted"})
}
