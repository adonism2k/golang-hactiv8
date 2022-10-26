package model

import (
	"time"

	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

var db *gorm.DB

type Models struct {
	User        User
	Photo       Photo
	Comment     Comment
	SocialMedia SocialMedia
}

func New(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		User{},
		Photo{},
		Comment{},
		SocialMedia{},
	}
}
