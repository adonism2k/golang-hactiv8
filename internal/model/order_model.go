package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// Order Model godoc
// @Description Order Model
type Order struct {
	ID           int            `gorm:"primarykey" swaggerignore:"true" json:"-"`
	CustomerName string         `gorm:"type:varchar(100);not null" json:"customer_name" example:"John Doe"`     // Customer Name
	OrderedAt    time.Time      `gorm:"autoCreateTime" json:"ordered_at" example:"2022-10-10T11:52:28.431369Z"` // Ordered At
	Items        []Item         `gorm:"foreignKey:OrderID" json:"items"`                                        // Items
	CreatedAt    time.Time      `gorm:"autoCreateTime" swaggerignore:"true" json:"-"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" swaggerignore:"true" json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" swaggerignore:"true" json:"-"`
} // @name OrderResponse

// OrderRequest Model godoc
// @Description OrderRequest Model
type OrderRequest struct {
	CustomerName string `json:"customer_name" example:"John Doe"` // Customer Name
	Items        []Item `json:"items"`                            // Items
} // @name OrderRequest

// Get all orders
func (o *Order) GetAll() ([]*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var orders []*Order
	result := db.WithContext(ctx).Model(&o).Preload("Items").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// Create a new order
func (o *Order) Create(order Order) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.Debug().WithContext(ctx).Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

// Update an order
func (o *Order) Update(id int, order Order) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.Debug().WithContext(ctx).Model(&o).Where("id = ?", id).Updates(order)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update the items of the order
	for _, item := range order.Items {
		item.OrderID = id
		_, err := item.Update(item)
		if err != nil {
			return nil, err
		}
	}

	return &order, nil
}

// Delete an order by id
func (o *Order) Delete(order Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.Debug().WithContext(ctx).Select("Items").Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
