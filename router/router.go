package router

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/markgerald/vw-order/controller"
)

func NewRouter(ordersController *controller.OrderController) *ginadapter.GinLambda {
	router := gin.Default()
	router.POST("/", ordersController.Create)
	router.PUT("/:id", ordersController.Update)
	router.DELETE("/:id", ordersController.Delete)
	router.GET("/", ordersController.FindAll)
	router.GET("/:id", ordersController.FindByID)
	router.GET("/user/:userId", ordersController.FindByUserId)

	return ginadapter.New(router)
}
