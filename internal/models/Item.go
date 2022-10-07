package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Code        string         `json:"code" gorm:"type:varchar(100);not null;unique"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	OrderID     uint           `json:"order_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (i Item) getAllItems() ([]*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var items []*Item
	result := db.WithContext(ctx).Model(&i).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (i Item) createItem() (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Create(&i)
	if result.Error != nil {
		return nil, result.Error
	}

	return &i, nil
}

func (i Item) updateItem() (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Save(&i)
	if result.Error != nil {
		return nil, result.Error
	}

	return &i, nil
}

func (i Item) deleteItem() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Delete(&i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}