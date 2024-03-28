package router

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/markgerald/vw-order/controller"
)

func NewRouter(ordersController *controller.OrderController) *ginadapter.GinLambda {
	router := gin.Default()
	router.POST("/orders", ordersController.Create)
	router.PUT("/orders/:id", ordersController.Update)
	router.DELETE("/orders/:id", ordersController.Delete)
	router.GET("/orders", ordersController.FindAll)
	router.GET("/orders/:id", ordersController.FindByID)

	return ginadapter.New(router)
}
