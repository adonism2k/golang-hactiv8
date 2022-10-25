package model

import (
	"context"
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

func (s *SocialMedia) All() ([]*SocialMedia, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var socialMedias []*SocialMedia
	result := db.WithContext(ctx).Preload("User").Find(&socialMedias)
	if result.Error != nil {
		return nil, result.Error
	}

	return socialMedias, nil
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

	result := tx.Model(&socialMedia).Where("id = ?", id).Updates(socialMedia)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &socialMedia, nil
}

func (s *SocialMedia) Delete(social SocialMedia) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	// result := tx.Unscoped().Delete(&social)
	result := tx.Delete(&social)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
