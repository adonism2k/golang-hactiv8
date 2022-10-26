package model

import (
	"context"
	"time"

	"gorm.io/gorm/clause"
)

// Social Media Model godoc
// @Description Social Media Model
type SocialMedia struct {
	ID        int       `gorm:"primarykey" json:"id" example:"1"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name" example:"Facebook"`
	Url       string    `gorm:"column:social_media_url;not null" json:"social_media_url" example:"https://images.unsplash.com"`
	UserID    int       `gorm:"not null" json:"user_id" example:"1"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" example:"2017-07-22'T'16:28:55.444"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" example:"2017-07-22'T'16:28:55.444"`
	User      User
} // @name Social Media

// Social Media Request Model godoc
// @Description Social Media Request Model
type SocialMediaRequest struct {
	Name string `json:"name" example:"Facebook" validate:"required"`
	Url  string `json:"social_media_url" example:"https://images.unsplash.com" validate:"required,url"`
} // @name SocialMediaRequest

func (s *SocialMedia) All(userID int) ([]*SocialMedia, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var socialMedias []*SocialMedia
	result := db.WithContext(ctx).Preload("User").Find(&socialMedias, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return socialMedias, nil
}

func (s *SocialMedia) Find(id int) (*SocialMedia, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var socialMedia SocialMedia
	result := db.WithContext(ctx).First(&socialMedia, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &socialMedia, nil
}

func (s *SocialMedia) Create(socialMedia SocialMedia) (*SocialMedia, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Create(&socialMedia)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &socialMedia, nil
}

func (s *SocialMedia) Update(id int, socialMedia SocialMedia) (*SocialMedia, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Model(&socialMedia).Clauses(clause.Returning{}).Where("id = ?", id).Updates(socialMedia)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &socialMedia, nil
}

func (s *SocialMedia) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Delete(&SocialMedia{}, id)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
