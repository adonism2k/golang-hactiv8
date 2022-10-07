package models

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

func (o Order) getAllOrders() ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var orders []*Order
	result := db.WithContext(ctx).Model(&o).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (o Order) createOrder() (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Create(&o)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}

func (o Order) updateOrder() (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}

	return &o, nil
}

func (o Order) deleteOrder() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).Delete(&o)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
