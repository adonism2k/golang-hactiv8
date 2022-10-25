package model

import (
	"context"
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

func (c *Comment) All(comment Comment) ([]*Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var comments []*Comment
	result := db.WithContext(ctx).Preload("User").Preload("Photo").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return comments, nil
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

	result := tx.Model(&comment).Where("id = ?", id).Updates(comment)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	tx.Commit()

	return &comment, nil
}

func (c *Comment) Delete(comment Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	// result := tx.Unscoped().Delete(&comment)
	result := tx.Delete(&comment)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
