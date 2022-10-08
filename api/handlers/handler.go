package handlers

import (
	"assignment-2/internal/database"
	"assignment-2/internal/model"
)

type Config struct {
	DB     database.Config
	Models model.Models
}
