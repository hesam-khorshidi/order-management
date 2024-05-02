package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-management/entity"
	"order-management/service"
)

type OrderController interface {
	UpdateOrderStatus(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func (controller *orderController) UpdateOrderStatus(ctx *gin.Context) {
	var updateRequest entity.UpdateOrderRequest
	if err := ctx.ShouldBind(&updateRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, entity.GeneralResponse{Status: 400, Message: "bad Request"})
		return
	}
	err := controller.orderService.UpdateOrderStatus(updateRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, entity.GeneralResponse{Status: 500, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entity.GeneralResponse{Status: 201, Message: "Order updated"})
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{
		orderService: orderService,
	}
}
