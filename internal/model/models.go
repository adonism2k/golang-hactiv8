package model

import (
	"time"

	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

var db *gorm.DB

func New(dbPool *gorm.DB) Models {
	db = dbPool

	return Models{
		Item:  Item{},
		Order: Order{},
	}
}

type Models struct {
	Item  Item
	Order Order
}
