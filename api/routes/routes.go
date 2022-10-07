package routes

import (
	"assignment-2/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Api(h handlers.Handler) http.Handler {
	r := gin.Default()

	r.GET("/orders", h.GetOrders)
	r.POST("/orders", h.CreateOrder)
	r.PUT("/orders:orderId", h.UpdateOrder)
	r.DELETE("/orders:orderId", h.DeleteOrder)

	return r
}
