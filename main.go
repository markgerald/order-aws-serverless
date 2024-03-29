package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/go-playground/validator/v10"
	"github.com/markgerald/vw-order/controller"
	"github.com/markgerald/vw-order/db"
	"github.com/markgerald/vw-order/repository"
	"github.com/markgerald/vw-order/router"
	"github.com/markgerald/vw-order/service"
)

var ginLambda *ginadapter.GinLambda

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	DB := db.InitDb()
	validate := validator.New()
	orderRepository := repository.NewOrdersRepositoryImpl(DB)
	orderService := service.NewOrderServiceImpl(orderRepository, validate)
	orderController := controller.NewOrderController(orderService)
	router.NewRouter(orderController)
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
