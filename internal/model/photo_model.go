package model

import (
	"time"
)

type Photo struct {
	ID        int       `gorm:"primarykey"`
	Title     string    `gorm:"type:varchar(100);not null"`
	Caption   int       `gorm:"not null"`
	PhotoUrl  string    `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
