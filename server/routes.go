package server

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/vw-order/controller"
)

func InitRoutes() {
	r := gin.Default()
	r.POST("/", controller.AddOrder)
	r.GET("/", controller.GetOrders)
	r.GET("/:id", controller.GetOrder)
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
