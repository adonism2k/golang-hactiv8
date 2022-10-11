package handlers

import (
	"assignment-2/internal/model"
	"assignment-2/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get All Orders godoc
// @Summary 	Get All Orders
// @Schemes
// @Description Fetch all orders and all of its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Success     200 {object} utils.JSONResult{data=[]model.Order,error=object} "Success"
// @Failure     400 {object} utils.JSONResult{data=object,error=utils.JSONError} "Error"
// @Router      /orders [get]
func (c *Config) GetOrders(ctx *gin.Context) {
	orders, err := c.Models.Order.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error getting orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.JSONResult{
		Success: true,
		Message: "Orders fetched successfully",
		Error:   nil,
		Data:    orders,
	})
}

// Create New Order godoc
// @Summary 	Create New Order
// @Schemes
// @Description Create a new order and its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param 		request body model.OrderRequest true "Order Request"
// @Success     201 {object} utils.JSONResult{data=model.Order,error=object} "Created"
// @Failure     400 {object} utils.JSONResult{data=object,error=utils.JSONError} "Error"
// @Router      /orders [post]
func (c *Config) CreateOrder(ctx *gin.Context) {
	// create a new order
	var newOrder model.Order
	err := ctx.BindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error creating orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	orderResult, err := c.Models.Order.Create(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error creating orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	// return a 201 created status code
	ctx.JSON(http.StatusCreated, utils.JSONResult{
		Success: true,
		Message: "Orders fetched successfully",
		Error:   nil,
		Data:    orderResult,
	})
}

// Update Order godoc
// @Summary Update Order
// @Schemes
// @Description Update an order and its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param 		request body model.OrderRequest true "Order Request"
// @Param       id path int true "Order ID"
// @Router      /orders/{id} [put]
// @Success     200 {object} utils.JSONResult{data=model.Order,error=object} "Success"
// @Failure     400 {object} utils.JSONResult{data=object,error=utils.JSONError} "Error"
func (c *Config) UpdateOrder(ctx *gin.Context) {
	// update an order
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error updating orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	var newOrder model.Order
	err = ctx.BindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error updating orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	orderResult, err := c.Models.Order.Update(id, newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error updating orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	// return a 200 ok status code
	ctx.JSON(http.StatusCreated, utils.JSONResult{
		Success: true,
		Message: "Orders updated successfully",
		Error:   nil,
		Data:    orderResult,
	})
}

// Delete Order godoc
// @Summary Delete Order
// @Schemes
// @Description Delete an order and its items
// @Tags        Order
// @Accept      json
// @Produce     json
// @Param       id path int true "Order ID"
// @Router      /orders/{id} [delete]
// @Success     204
// @Failure     400 {object} utils.JSONResult{data=object,error=utils.JSONError} "Error"
func (c *Config) DeleteOrder(ctx *gin.Context) {
	// delete an order
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error deleting orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	order := model.Order{ID: id}

	err = c.Models.Order.Delete(order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.JSONResult{
			Success: false,
			Message: "Error deleting orders",
			Error: utils.JSONError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	// return a 204 no content status code
	ctx.Status(http.StatusNoContent)
}
