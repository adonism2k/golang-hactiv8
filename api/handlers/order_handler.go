package handlers

import (
	"github.com/gin-gonic/gin"
)

// type Order struct {
// 	OrderID      uint      `json:"order_id"`
// 	OrderedAt    time.Time `json:"ordered_at"`
// 	CustomerName string    `json:"customer_name"`
// 	Items        Item      `json:"items"`
// }

// type Item struct {
// 	ItemID      uint   `json:"item_id"`
// 	Description string `json:"description"`
// 	Quantity    uint   `json:"quantity"`
// }

func (h Handler) GetOrders(ctx *gin.Context) {
	// TODO: get all orders
	// !ERROR: undefined: h.Models.Order.getAllOrders() (type models.Models has no field or method getAllOrders) why?
	orders, err := h.Models.Order.getAllOrders()

	// TODO: return a 200 ok status code
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h Handler) CreateOrder(ctx *gin.Context) {
	// TODO: create a new order

	// TODO: return a 201 created status code
}

func (h Handler) UpdateOrder(ctx *gin.Context) {
	// TODO: update an order

	// TODO: return a 200 ok status code
}

func (h Handler) DeleteOrder(ctx *gin.Context) {
	// TODO: delete an order

	// TODO: return a 204 no content status code
}
