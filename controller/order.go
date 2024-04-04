package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/vw-order/data/request"
	"github.com/markgerald/vw-order/data/response"
	"github.com/markgerald/vw-order/helper"
	"github.com/markgerald/vw-order/service"
	"log"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (controller *OrderController) Create(ctx *gin.Context) {
	log.Printf("Create Order")
	createOrderRequest := request.CreateOrdersRequest{}
	err := ctx.ShouldBindJSON(&createOrderRequest)
	helper.ErrorPanic(err)

	controller.orderService.Create(createOrderRequest)
	webresponse := response.Response{
		Code:   201,
		Status: "Created",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webresponse)
}

func (controller *OrderController) Update(ctx *gin.Context) {
	log.Printf("Update Order")
	updateOrderRequest := request.UpdateOrdersRequest{}
	err := ctx.ShouldBindJSON(&updateOrderRequest)
	helper.ErrorPanic(err)

	orderId := ctx.Param("id")
	updateOrderRequest.ID = orderId
	update, err := controller.orderService.Update(updateOrderRequest)
	if err != nil {
		response.SendErrorResponse(ctx, 404, "Order Not Found")
		return
	}

	response.SendSuccessResponse(ctx, 200, "OK", update)
}

func (controller *OrderController) Delete(ctx *gin.Context) {
	log.Printf("Delete Order")
	orderId := ctx.Param("id")
	err := controller.orderService.Delete(orderId)
	if err != nil {
		response.SendErrorResponse(ctx, 404, "Order Not Found")
		return
	}

	response.SendSuccessResponse(ctx, 200, "OK", nil)
}

func (controller *OrderController) FindByID(ctx *gin.Context) {
	log.Printf("Find Order By ID")
	orderId := ctx.Param("id")
	orderResponse, err := controller.orderService.FindByID(orderId)

	if err != nil {
		response.SendErrorResponse(ctx, 404, "Order Not Found")
		return
	}

	response.SendSuccessResponse(ctx, 200, "OK", orderResponse)
}

func (controller *OrderController) FindAll(ctx *gin.Context) {
	log.Printf("Find All Orders")
	ordersResponse := controller.orderService.FindAll()

	webresponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   ordersResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webresponse)
}
