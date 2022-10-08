package handlers

import (
	"assignment-2/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID           uint      `json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}

// Get All Orders godoc
// @Summary Get All Orders
// @Schemes
// @Description Fetch all orders and all of its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Success     200 {array} Order
// @Router      /orders [get]
func (c *Config) GetOrders(ctx *gin.Context) {
	orders, err := c.Models.Order.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Orders fetched successfully",
		"data":    orders,
	})
}

// Create New Order godoc
// @Summary Create New Order
// @Schemes
// @Description Create a new order and its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Success     201 {object} Order
// @Router      /orders [post]
func (c *Config) CreateOrder(ctx *gin.Context) {
	// create a new order
	var newOrder model.Order
	err := ctx.BindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	personResult, err := c.Models.Order.Create(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
			"data":    nil,
		})
	}

	// return a 201 created status code
	ctx.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "Order created successfully",
		"data":    personResult,
	})
}

// Update Order godoc
// @Summary Update Order
// @Schemes
// @Description Update an order and its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Success     200 {object} Order
// @Router      /orders/{id} [put]
// @Param       id path int true "Order ID"
func (c *Config) UpdateOrder(ctx *gin.Context) {
	// TODO: update an order

	// TODO: return a 200 ok status code
}

// Delete Order godoc
// @Summary Delete Order
// @Schemes
// @Description Delete an order and its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Success     204 {} Order
// @Router      /orders/{id} [delete]
// @Param       id path int true "Order ID"
func (c *Config) DeleteOrder(ctx *gin.Context) {
	// TODO: delete an order

	// TODO: return a 204 no content status code
}
