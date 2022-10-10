package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Item Model info
// @Description Item Model
type Item struct {
	ID          int            `gorm:"primarykey" swaggerignore:"true" json:"-"`
	Code        string         `gorm:"type:varchar(100);not null;unique" json:"code" example:"PD-001"`    // Item Code
	Description string         `gorm:"type:varchar(100);not null" json:"description" example:"Product 1"` // Item Description
	Quantity    int            `gorm:"type:int;not null" json:"quantity" example:"10"`                    // Quantity
	OrderID     int            `gorm:"not null;foreignkey:OrderID" swaggerignore:"true" json:"-"`
	CreatedAt   time.Time      `swaggerignore:"true" gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time      `swaggerignore:"true" gorm:"autoUpdateTime" json:"-"`
	DeletedAt   gorm.DeletedAt `swaggerignore:"true" gorm:"index" json:"-"`
} // @name ItemResponse

// Item Request info
// @Description Item Request
type ItemRequest struct {
	Code        string `json:"code" example:"PD-001"`           // Item Code
	Description string `json:"description" example:"Product 1"` // Item Description
	Quantity    int    `json:"quantity" example:"10"`           // Item Quantity
} // @name ItemRequest

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

	result := db.WithContext(ctx).Where("order_id = ?", item.OrderID).Updates(&item)
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
