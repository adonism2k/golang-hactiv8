package model

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

func (i *Item) GetAll() ([]*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var items []*Item
	result := db.WithContext(ctx).Model(&i).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (i *Item) Create(item Item) (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (i *Item) Update(item Item) (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (i *Item) Delete(item Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Delete(&item)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
