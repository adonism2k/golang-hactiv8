package handlers

import (
	"github.com/adonism2k/golang-hactiv8/internal/database"
	"github.com/adonism2k/golang-hactiv8/internal/initializers"
	"github.com/adonism2k/golang-hactiv8/internal/model"
)

type Config struct {
	DB     database.Config
	Models model.Models
	Env    initializers.Config
}

func New(db database.Config, config initializers.Config) Config {
	return Config{
		DB:     db,
		Models: model.New(db.DB),
		Env:    config,
	}
}
