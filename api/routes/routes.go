package routes

import (
	"assignment-2/api/handlers"
	"net/http"

	docs "assignment-2/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Api(h handlers.Config) http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		// Order
		v1.GET("/orders", h.GetOrders)
		v1.POST("/orders", h.CreateOrder)
		v1.PUT("/orders/:id", h.UpdateOrder)
		v1.DELETE("/orders/:id", h.DeleteOrder)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
