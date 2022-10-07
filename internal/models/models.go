package models

import (
	"time"

	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

var db *gorm.DB

type Models struct {
	Item  Item
	Order Order
}

func New(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		Item:  Item{},
		Order: Order{},
	}
}
