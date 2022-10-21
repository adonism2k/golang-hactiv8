package handlers

import (
	"github.com/adonism2k/golang-hactiv8/internal/database"
	"github.com/adonism2k/golang-hactiv8/internal/model"
)

type Config struct {
	DB     database.Config
	Models model.Models
}

func New(db database.Config) Config {
	return Config{
		DB:     db,
		Models: model.New(db.DB),
	}
}
