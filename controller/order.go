package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/vw-order/data/request"
	"github.com/markgerald/vw-order/data/response"
	"github.com/markgerald/vw-order/helper"
	"github.com/markgerald/vw-order/service"
	"log"
	"strconv"
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
	if err != nil {
		response.SendErrorResponse(ctx, 400, "Validation Error: "+err.Error())
		log.Fatalf("Error binding JSON: %v", err)
		return
	}
	create, err := controller.orderService.Create(createOrderRequest)
	if err != nil {
		response.SendErrorResponse(ctx, 500, "Error creating order")
		return
	}
	response.SendSuccessResponse(ctx, 200, "OK", create)
}

func (controller *OrderController) Update(ctx *gin.Context) {
	log.Printf("Update Order")
	updateOrderRequest := request.UpdateOrdersRequest{}
	err := ctx.ShouldBindJSON(&updateOrderRequest)
	if err != nil {
		response.SendErrorResponse(ctx, 400, "Validation Error: "+err.Error())
		log.Fatalf("Error binding JSON: %v", err)
		return
	}

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
	limit := ctx.DefaultQuery("limit", "10")
	startKey := ctx.Query("startKey")
	limitInt, err := strconv.Atoi(limit)
	helper.ErrorPanic(err)
	ordersResponse, lastKey, err := controller.orderService.FindAll(limitInt, startKey)
	if err != nil {
		response.SendErrorResponse(ctx, 500, "Internal Server Error")
		return
	}
	data := map[string]interface{}{
		"orders":  ordersResponse,
		"lastKey": lastKey,
	}

	response.SendSuccessResponse(ctx, 200, "OK", data)
}

func (controller *OrderController) FindByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	limit := ctx.DefaultQuery("limit", "10")
	startKey := ctx.Query("startKey")
	ordersResponse, lastKey, err := controller.orderService.FindByUserId(userId, limit, startKey)
	if err != nil {
		response.SendErrorResponse(ctx, 500, "Internal Server Error")
		return
	}
	ctx.JSON(200, gin.H{
		"data":    ordersResponse,
		"lastKey": lastKey,
	})
}
