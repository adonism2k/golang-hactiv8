package handlers

import (
	"assignment-2/internal/database"
	"assignment-2/internal/models"
)

type Handler struct {
	DB     database.Config
	Models models.Models
}

func New(db database.Config, models models.Models) Handler {
	return Handler{
		DB:     db,
		Models: models,
	}
}
