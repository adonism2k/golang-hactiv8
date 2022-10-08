package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CustomerName string         `json:"customer_name" gorm:"type:varchar(100);not null"`
	OrderedAt    time.Time      `json:"ordered_at" gorm:"autoCreateTime"`
	Items        []Item         `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (o *Order) GetAll() ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var orders []*Order
	result := db.WithContext(ctx).Model(&o).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o *Order) Create(order Order) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (o *Order) Update(order Order) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Save(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (o *Order) Delete(order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
