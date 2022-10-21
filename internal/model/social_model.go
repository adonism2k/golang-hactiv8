package model

import (
	"time"
)

type SocialMedia struct {
	ID             int       `gorm:"primarykey"`
	Name           string    `gorm:"type:varchar(100);not null"`
	SocialMediaUrl int       `gorm:"not null"`
	UserID         string    `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
