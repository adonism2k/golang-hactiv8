package model

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm/clause"
)

// PhotoResponse Model godoc
// @Description PhotoResponse Model
type Photo struct {
	ID        int       `gorm:"primarykey" json:"id" example:"1"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title" example:"First Photo"`
	Caption   string    `gorm:"not null" json:"caption" example:"my first photo"`
	Url       string    `gorm:"column:photo_url;not null" json:"photo_url" example:"https://images.unsplash.com"`
	UserID    int       `gorm:"not null" json:"user_id" example:"1"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" example:"2017-07-22'T'16:28:55.444"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" example:"2017-07-22'T'16:28:55.444"`
	User      User
} // @name PhotoResponse

// PhotoRequest Model godoc
// @Description PhotoRequest Model
type PhotoRequest struct {
	Title   string `json:"title" example:"First Photo" validate:"required"`
	Caption string `json:"caption" example:"my first photo"`
	Url     string `json:"photo_url" example:"https://images.unsplash.com" validate:"required,url"`
} // @name PhotoRequest

func (p *Photo) All() ([]*Photo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var photos []*Photo
	result := db.WithContext(ctx).Preload("User").Find(&photos)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return photos, nil
}

func (p *Photo) Create(photo Photo) (*Photo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Create(&photo)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &photo, nil
}

func (p *Photo) Update(id int, newPhoto Photo) (*Photo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	var photo Photo
	result := tx.Model(&photo).Clauses(clause.Returning{}).Where("id = ?", id).Updates(newPhoto)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &photo, nil
}

func (p *Photo) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	// result := tx.Select("Comments").Unscoped().Delete(&photo)
	result := tx.Select("Comments").Delete(&Photo{}, id)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
