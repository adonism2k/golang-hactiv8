package model

import (
	"time"
)

type Comment struct {
	ID        int       `gorm:"primarykey"`
	Message   string    `gorm:"not null"`
	UserID    string    `gorm:"type:varchar(100);not null"`
	PhotoID   int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
