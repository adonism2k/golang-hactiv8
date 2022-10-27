package model

import (
	"context"
	"time"

	"gorm.io/gorm/clause"
)

type Comment struct {
	ID        int       `gorm:"primarykey" json:"id" example:"1"`
	Message   string    `gorm:"not null" json:"message" example:"This is a comment"`
	UserID    int       `gorm:"not null" json:"user_id" example:"1"`
	PhotoID   int       `gorm:"not null" json:"photo_id" example:"1"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" example:"2017-07-22'T'16:28:55.444"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" example:"2017-07-22'T'16:28:55.444"`
	User      User
	Photo     Photo
}

func (c *Comment) All() ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var comments []*Comment
	result := db.WithContext(ctx).Preload("User").Preload("Photo").Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
}

func (c *Comment) Find(id int) (*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var comment Comment
	result := db.WithContext(ctx).First(&comment, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func (c *Comment) Create(comment Comment) (*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Create(&comment)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &comment, nil
}

func (c *Comment) Update(id int, comment Comment) (*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Model(&comment).Clauses(clause.Returning{}).Where("id = ?", id).Updates(comment)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &comment, nil
}

func (c *Comment) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Delete(&Comment{}, id)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
